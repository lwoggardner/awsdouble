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

package awsdoublegen

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/lwoggardner/godouble/doublegen"
	"io"
	"log"
	"os"
	"reflect"
	"strings"
	"text/template"
)

type APIInterface struct {
	doublegen.Interface
}

type APIMethod struct {
	doublegen.Method
}

func NewAPIInterface(forInterface interface{}) APIInterface {

	return APIInterface{Interface: doublegen.NewGenerator(forInterface, func(iface *doublegen.Interface) {
		pkgPathParts := strings.Split(iface.Type.PkgPath(),"/")
		servicePath := pkgPathParts[len(pkgPathParts) - 2]
		ifaceParts := strings.Split(iface.Type.String(),".")
		serviceAPI := ifaceParts[1]
		service := serviceAPI[:len(serviceAPI)-3]
		pkgName := fmt.Sprintf("%sdouble",servicePath)
		typeName := fmt.Sprintf("%sDouble",service)
		iface.SetType(pkgName,typeName)
	})}
}

func (iface APIInterface) Generate(fileName string) {
	f,err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	iface.GenerateDouble(f)
	log.Printf("Created double for %v at %s",iface,fileName)
}

func (iface APIInterface) GenerateDouble(writer io.Writer) {
	t := template.New("GenerateServiceDouble")
	iface.ParseTemplates(t)
	template.Must(t.Parse(DoubleTemplate))
	template.Must(t.Parse(GenerateTemplate))
	err := t.Execute(writer, iface)
	if err != nil {
		log.Fatal(err)
	}
}

func (iface APIInterface) SDKVersion() string {
	return aws.SDKVersion
}

func (iface APIInterface) ServiceImport() string {

	pkgPathParts := strings.Split(iface.Type.PkgPath(),"/")
	return pkgPathParts[len(pkgPathParts) - 2]
}

func (iface APIInterface) Methods() []APIMethod {

	methods := iface.Interface.Methods()
	apiMethods := make([]APIMethod, len(methods))
	for i,m := range methods {
		apiMethods[i] = APIMethod{m}
	}
	return apiMethods
}

func (m APIMethod) OperationName() string {
	result := m.Name
	result = strings.TrimPrefix(result,"WaitUntil")
	result = strings.TrimSuffix(result,"WithContext")
	result = strings.TrimSuffix(result, "Pages")
	return result
}

func (m APIMethod) ServiceImport() string {
	return strings.Split(m.InputType().String(),".")[0]
}

func (m APIMethod) InputType() reflect.Type {
	argIndex := 0
	if strings.HasSuffix(m.Name,"WithContext") {
		argIndex = 1
	}
	return m.Type.In(argIndex).Elem()
}

func (m APIMethod) OutputType() reflect.Type {
	switch m.Fake() {
	case "PagesWithContext":
		return m.Type.In(2).In(0).Elem()
	case "Pages":
		return m.Type.In(1).In(0).Elem()
	default:
		return m.Type.Out(0).Elem()
	}
}

func (m APIMethod) WaiterName() string {
	return m.OperationName()
}

func (m APIMethod) Fake() string {
	return strings.Replace(m.Name,m.OperationName(),"",1)
}

const GenerateTemplate = `
{{template "PackageHeader" . }}
{{template "AWSDoubleType" . }}
{{range .Methods}}
	{{template "DoubleMethod" . }}
    {{if eq .Fake "WaitUntil"}}
        {{template "FakeWaitUntil" .}}
    {{else if eq .Fake "WaitUntilWithContext"}}
        {{template "FakeWaitUntilWithContext" .}}
    {{else if eq .Fake "PagesWithContext"}}
        {{template "FakePagesWithContext" .}}
    {{else if eq .Fake "Pages"}}
        {{template "FakePages" .}}
    {{else if eq .Fake "WithContext"}}
        {{template "FakeWithContext" .}}
	{{end}}
{{end}}
`
const DoubleTemplate = `
{{define "AWSDoubleType"}}
{{- /*gotype: github.com/lwoggardner/godouble/awsdoublegen.Interface*/ -}}
import (
    "github.com/lwoggardner/godouble/godouble"
    "github.com/lwoggardner/awsdouble"
	"context"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
    "github.com/aws/aws-sdk-go/service/{{.ServiceImport}}"
    "github.com/aws/aws-sdk-go/service/{{.ServiceImport}}/{{.ServiceImport}}iface"
)

// {{.TypeName}} is TestDouble for {{.Type}}
type {{.TypeName}} struct {
    {{.Type}}
    *awsdouble.AWSTestDouble
}

// Constructor for {{.TypeName}}
// Default configuration will ensure 
//   * API operations not explicitly stubbed will return an empty output struct pointer, and nil error.
//     To simulate long polling, "Poll" operations will return these values after a random delay of up to 100ms.
//   * WithContext methods implement a 'Fake' method that returns a cancellation error if the context is cancelled 
//      before the method is called. 
//   * Pages and PagesWithContext methods similarly implement a 'Fake' method that paginates over the underlying method.
//
//    This allows tests to only stub the simple api methods and be generally unconcerned whether the SUT is using the
//      Context or Pagination forms of the API.
//
func NewDouble(t godouble.T,configurators ...func(*awsdouble.AWSTestDouble)) *{{.TypeName}} {
    result := &{{.TypeName}}{}

    configurators = append([]func(configurator *awsdouble.AWSTestDouble){func(d *awsdouble.AWSTestDouble) {
        d.SetDefaultCall(result.defaultMethodCall)
        d.SetDefaultReturnValues(result.defaultReturnValues)
    }},configurators...)

    result.AWSTestDouble = awsdouble.NewDouble(t,(*{{.Type}})(nil), configurators...)
    return result
}

func (d *{{.TypeName}}) defaultReturnValues(m godouble.Method) godouble.ReturnValues {
    return d.DefaultReturnValues(m)
}

func (d *{{.TypeName}}) defaultMethodCall(m godouble.Method) godouble.MethodCall {
    switch m.Reflect().Name {
    {{- range .Methods -}}
    {{- if .Fake }}
    case "{{.Name}}":
        return m.Fake(d.fake{{.Name}})
    {{ end }}
    {{ end }}
    default:
        return nil
    }
}    
{{end}}

{{define "FakeWithContext"}}
    {{- /*gotype: github.com/lwoggardner/awsdouble.InterfaceMethod*/ -}}

func (d *{{.TypeName}}) fake{{.Name}}(ctx context.Context, in {{ index .Args 1}}, _ ...request.Option) ({{index .Returns 0}}, error) {
    
	select {
        case <-ctx.Done():
            return nil, awserr.New(request.CanceledErrorCode, "{{.OperationName}} cancelled", ctx.Err())
        default:
            return d.{{.OperationName}}(in)
    }
}
{{end}}


{{define "FakePages"}}
func (d *{{.TypeName}}) fake{{.Name}}(in *{{.InputType}}, pager func(*{{.OutputType}}, bool) (shouldContinue bool)) error {
	return d.Paginate("{{.OperationName}}",paginators,in,pager)
}
{{end}}

{{define "FakePagesWithContext"}}
func (d *{{.TypeName}}) fake{{.Name}}(ctx context.Context, in *{{.InputType}}, pager func(*{{.OutputType}}, bool) (shouldContinue bool), options ...request.Option) error {
	return d.PaginateWithContext("{{.OperationName}}", paginators, ctx,in,pager, options...)
}
{{end}}

{{define "FakeWaitUntil"}}
func (d *{{.TypeName}}) fake{{.Name}}(in *{{.InputType}}) error {
	return d.WaitUntil("{{.WaiterName}}",waiters,in)
}
{{end}}

{{define "FakeWaitUntilWithContext"}}
func (d *{{.TypeName}}) fake{{.Name}}(ctx context.Context, in *{{.InputType}}, waitOption ...request.WaiterOption) error {
 	return d.WaitWithContext("{{.WaiterName}}",waiters,ctx,in,waitOption...)
}
{{end}}
`