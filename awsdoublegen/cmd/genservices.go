// +build doublegen
/*
 * Copyright 2020 grant@lastweekend.com.au
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"flag"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/private/model/api"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime/debug"
	"strings"
	"sync"
	"text/template"
)

// Derived from aws-sdk-go/private/model/cli/gen-api/main.go
var excludeServices = map[string]struct{}{
	"importexport": {},
}

var genPath, svcPath, svcImportPath,apisPath string

func main() {

	flag.StringVar(&genPath, "genpath", "awsdoublegen/cmd/services",
		"The path to build service generators in",
	)

	flag.StringVar(&svcPath, "path", "doubles",
		"The `path` to generate service clients in to.",
	)


	flag.StringVar(&svcImportPath, "svc-import-path",
		"github.com/aws/aws-sdk-go/service",
		"The Go `import path` to generate client to be under.",
	)

	defaultApisPath := fmt.Sprintf("%s/pkg/mod/github.com/aws/%s@v%s/models/apis",os.Getenv("GOPATH"),aws.SDKName,aws.SDKVersion)
	flag.StringVar(&apisPath, "apis", defaultApisPath, "The path to the sdk module's `models/apis` folder")
	flag.Parse()

	if len(os.Getenv("AWS_SDK_CODEGEN_DEBUG")) != 0 {
		api.LogDebug(os.Stdout)
	}

	// Make sure all paths are based on platform's pathing not Unix
	globs := flag.Args()
	for i, g := range globs {
		globs[i] = filepath.FromSlash(filepath.Join(apisPath,g))
	}
	svcPath = filepath.FromSlash(svcPath)
	genPath = filepath.FromSlash(genPath)

	modelPaths, err := api.ExpandModelGlobPath(globs...)
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to glob file pattern", err)
		os.Exit(1)
	}
	modelPaths, _ = api.TrimModelServiceVersions(modelPaths)

	loader := api.Loader{
		BaseImport:            svcImportPath,
		IgnoreUnsupportedAPIs: true,
	}

	apis, err := loader.Load(modelPaths)

	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to load API models", err)
		os.Exit(1)
	}
	if len(apis) == 0 {
		fmt.Fprintf(os.Stderr, "expected to load models, but found none")
		os.Exit(1)
	}

	if v := os.Getenv("SERVICES"); len(v) != 0 {
		svcs := strings.Split(v, ",")
		for pkgName, a := range apis {
			var found bool
			for _, include := range svcs {
				if a.PackageName() == include {
					found = true
					break
				}
			}
			if !found {
				delete(apis, pkgName)
			}
		}
	}

	var wg sync.WaitGroup
	servicePaths := map[string]struct{}{}
	for _, a := range apis {
		if _, ok := excludeServices[a.PackageName()]; ok {
			continue
		}

		// Create the output path for the model.
		pkgDir := filepath.Join(svcPath, a.PackageName())
		os.MkdirAll(pkgDir, 0775)

		if _, ok := servicePaths[pkgDir]; ok {
			fmt.Fprintf(os.Stderr,
				"attempted to generate a client into %s twice. Second model package, %v\n",
				pkgDir, a.PackageName())
			os.Exit(1)
		}
		servicePaths[pkgDir] = struct{}{}

		wg.Add(1)
		go func(a *api.API, pkgDir string) {
			defer wg.Done()
			generateService(a, pkgDir)
			generatePagers(a,pkgDir)
			generateWaiters(a,pkgDir)
		}(a,pkgDir)
	}

	wg.Wait()

	generate(svcPath)

}

type API struct {
	*api.API
}

func (a API) InterfaceName() string {
	return a.StructName() + "API"
}

func (a API) SDKVersion() string {
	return aws.SDKVersion
}

func generatePagers(a *api.API, pkgDir string) {
	t := template.Must(template.New("Doc").Parse(PagerTemplate))
	writeGoTemplate(filepath.Join(pkgDir, "pagers.go"),t,a)
}

func generateWaiters(a *api.API, pkgDir string) {
	if len(a.Waiters) > 0 {
		t := template.Must(template.New("Doc").Funcs(
			template.FuncMap{
				"titleCase": func(v string) string {
					return strings.Title(v)
				},
			}).Parse(WaitersTemplate))
		writeGoTemplate(filepath.Join(pkgDir, "waiters.go"),t,a)
	}
}
func generateService(a *api.API, pkgDir string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "Error generating %s\n%s\n%s\n", pkgDir, r, debug.Stack())
			os.Exit(1)
		}
	}()

	log.Printf("Generating %s (%s)...\n", a.PackageName(), a.Metadata.APIVersion)

	writeServiceBuilder(API{a})
	writeDoc(pkgDir,API{a})
}

func generate(svcPath string) {
	cmd := exec.Command("go","generate", "./...")
	cmd.Dir = svcPath
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	must(cmd.Run())
}

func writeServiceBuilder(a API) {
	t := template.Must(template.New("Doc").Parse(ServiceBuilderTemplate))
	writeGoTemplate(filepath.Join(genPath, a.PackageName() + ".go"),t,a)
}

func writeDoc(pkgDir string, a API) {
	t := template.Must(template.New("Doc").Parse(DocTemplate))
	writeGoTemplate(filepath.Join(pkgDir,"doc.go"),t,a)
}

func writeGoTemplate(file string,t *template.Template, data interface{}) {
	f,err := os.Create(file)
	must(err)
	defer f.Close()
	must(t.Execute(f,data))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}


const DocTemplate = `
// Code generated by go awsdoublegen; DO NOT EDIT.
// This file was generated at 2020-01-29T21:22:05+11:00
// For SDKVersion {{.SDKVersion}}
//Package {{.PackageName}}double contains test double implementation of {{.InterfacePackageName}}.{{.InterfaceName}}
package {{.PackageName}}double

//go:generate go run -tags doublegen ../../awsdoublegen/cmd/services/{{.PackageName}}.go
`

const ServiceBuilderTemplate = `
// +build doublegen
// Code generated by go awsdoublegen; DO NOT EDIT.
// This file was generated at 2020-01-29T21:22:05+11:00

package main

import (
	"github.com/aws/aws-sdk-go/service/{{.PackageName}}/{{.InterfacePackageName}}"
	"github.com/lwoggardner/awsdouble/awsdoublegen"
)

func main() {
	awsdoublegen.NewAPIInterface((*{{.InterfacePackageName}}.{{.InterfaceName}})(nil)).Generate("double.go")
}

`

const PagerTemplate = `
// Code generated by go awsdoublegen; DO NOT EDIT.
// This file was generated at 2020-01-29T21:22:05+11:00
package {{.PackageName}}double

import    "github.com/aws/aws-sdk-go/aws/request"

var paginators = map[string]*request.Paginator{
{{range .OperationList}}
		{{ if .Paginator }}"{{.Name}}": &request.Paginator{
				InputTokens: {{ .Paginator.InputTokensString }},
				OutputTokens: {{ .Paginator.OutputTokensString }},
				LimitToken: "{{ .Paginator.LimitKey }}",
				TruncationToken: "{{ .Paginator.MoreResults }}",
		},
       {{end}}
{{end}}
}
`

const WaitersTemplate = `
// Code generated by go awsdoublegen; DO NOT EDIT.
// This file was generated at 2020-01-29T21:22:05+11:00
package {{.PackageName}}double

import (
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
    "github.com/lwoggardner/awsdouble"
)

var waiters = map[string]*awsdouble.Waiter{
{{range .Waiters}}
	"{{.Name}}" : &awsdouble.Waiter{
         OperationName: "{{.OperationName}}",
         Waiter: request.Waiter{
			Name:    "WaitUntil{{ .Name }}",
			MaxAttempts: {{ .MaxAttempts }},
			Delay: request.ConstantWaiterDelay({{ .Delay }} * time.Second),
			Acceptors: []request.WaiterAcceptor{
				{{ range $_, $a := .Acceptors }}{
					State:    request.{{ titleCase .State }}WaiterState,
					Matcher:  request.{{ titleCase .Matcher }}WaiterMatch,
					{{- if .Argument }}Argument: "{{ .Argument }}",{{ end }}
					Expected: {{ .ExpectedString }},
				},
				{{ end }}
			},
         },
    },
{{end}}
}
`