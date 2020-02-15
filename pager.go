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

package awsdouble

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/request"
	"reflect"
)

func (d AWSTestDouble) PaginateWithContext(operationName string, paginators map[string]*request.Paginator, ctx context.Context, in interface{}, pager interface{}, options ...request.Option) error {

	d.TestDouble.T().Helper()
	paginator := paginators[operationName]

	if in == nil {
		return errors.New("input arg must not be nil")
	}
	if pager == nil {
		return errors.New("pager arg must not be nil")
	}

	next := in
	pagerV := reflect.ValueOf(pager)
	inV := reflect.ValueOf(in).Elem()
	for {
		returns := d.Invoke(operationName + "WithContext", ctx,next,options)
		out := returns[0]
		err,_ := returns[1].(error)
		if err != nil {
			return err
		} else {
			nextPageTokens := nextPageTokens(paginator,out)
			last := nextPageTokens == nil
			cont := pagerV.Call([]reflect.Value{reflect.ValueOf(out),reflect.ValueOf(last)})[0].Bool()
			if !cont || last {
				return nil
			}
			nextV := reflect.New(inV.Type())
			nextV.Elem().Set(inV) //Copy
			next = nextV.Interface()
			for i,tokenPath := range paginator.InputTokens {
				awsutil.SetValueAtPath(next,tokenPath, nextPageTokens[i])
			}
		}
	}

}

func (d AWSTestDouble) Paginate(operationName string, paginators map[string]*request.Paginator, in interface{}, pager interface{}) error {
	d.TestDouble.T().Helper()
	return d.PaginateWithContext(operationName,paginators, context.Background(), in, pager)
}

//derived from SDK request based code
func nextPageTokens(p *request.Paginator, out interface{}) []interface{}{
	if p.TruncationToken != "" {
		tr, _ := awsutil.ValuesAtPath(out, p.TruncationToken)
		if len(tr) == 0 {
			return nil
		}

		switch v := tr[0].(type) {
		case *bool:
			if !aws.BoolValue(v) {
				return nil
			}
		case bool:
			if v == false {
				return nil
			}
		}
	}

	tokens := []interface{}{}
	tokenAdded := false
	for _, outToken := range p.OutputTokens {
		vs, _ := awsutil.ValuesAtPath(out, outToken)
		if len(vs) == 0 {
			tokens = append(tokens, nil)
			continue
		}
		v := vs[0]

		switch tv := v.(type) {
		case *string:
			if len(aws.StringValue(tv)) == 0 {
				tokens = append(tokens, nil)
				continue
			}
		case string:
			if len(tv) == 0 {
				tokens = append(tokens, nil)
				continue
			}
		}

		tokenAdded = true
		tokens = append(tokens, v)
	}

	if tokenAdded {
		return tokens
	} else {
		return nil
	}
}
