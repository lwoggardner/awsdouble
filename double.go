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
	dbl "github.com/lwoggardner/godouble/godouble"
	"reflect"
	"strings"
	"time"
)

type AWSTestDouble struct {
	*dbl.TestDouble
	//Replace this with a fakeclock implementation, otherwise default is that all sleeps are 10ms
	Timewarp
}

// Called by service specific doubles
func NewDouble(t dbl.T, forInterface interface{}, configurator ...func(c *AWSTestDouble)) *AWSTestDouble {
	d := &AWSTestDouble{TestDouble: dbl.NewDouble(t, forInterface), Timewarp: DefaultTimeWarp}
	for _,c := range configurator {
		c(d)
	}
	return d
}

func (d *AWSTestDouble) SetTimewarp(tw Timewarp) {
	d.Timewarp = tw
}
// godouble uses zero values for defaults, but API methods are all ptrTo(struct) and never return nil in the real world
type defaultAWSReturnValues []reflect.Type

func (d defaultAWSReturnValues) Receive() ([]interface{}, error) {
	results := make([]interface{},len(d))
	for i,out := range d {
		if reflect.Ptr == out.Kind() && reflect.Struct == out.Elem().Kind() && strings.HasSuffix(out.String(),"Output") {
			results[i] = reflect.New(out.Elem()).Interface()
		} else {
			results[i] = reflect.Zero(out).Interface()
		}
	}
	return results, nil
}


func (d *AWSTestDouble) DefaultReturnValues(m dbl.Method) dbl.ReturnValues {
	mType := m.Reflect().Type
	results := defaultAWSReturnValues(make([]reflect.Type,mType.NumOut()))
	for i := 0; i < mType.NumOut(); i++ {
		results[i] = mType.Out(i)
	}

	if strings.HasPrefix(m.Reflect().Name,"Poll") {
		//Simulate long poll
		return dbl.RandDelayed(results,time.Duration(100) * time.Millisecond, dbl.Timewarp(d.Timewarp))
	} else {
		return results
	}
}



