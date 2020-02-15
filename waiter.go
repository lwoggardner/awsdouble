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
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/request"
)

type Waiter struct {
	request.Waiter
	OperationName string
}

func (d *AWSTestDouble) WaitUntil(waiterName string, waiters map[string]*Waiter, in interface{}) error {
	return d.WaitWithContext(waiterName,waiters, context.Background(), in)
}

func (d *AWSTestDouble) WaitWithContext(waiterName string, waiters map[string]*Waiter, ctx context.Context, in interface{}, wOptions ...request.WaiterOption) error {
	w := waiters[waiterName]
	w.ApplyOptions(wOptions...)

	for attempt := 1; ; attempt++ {
		returns := d.Invoke(w.OperationName + "WithContext",ctx,in,w.RequestOptions)
		out := returns[0]
		err,_ := returns[1].(error)
		// See if any of the acceptors match the request's response, or error
		for _, a := range w.Acceptors {
			if matched, matchErr := match(a, out, err); matched {
				return matchErr
			}
		}

		// The Waiter should only check the resource state MaxAttempts times
		// This is here instead of in the for loop above to prevent delaying
		// unnecessary when the waiter will not retry.
		if attempt == w.MaxAttempts {
			break
		}

		// Delay to wait (with timewarp) before inspecting the resource again
		delay := w.Delay(attempt)
		sleepCtxFn := w.SleepWithContext
		if sleepCtxFn == nil {
			tw := timeWarpFromContext(ctx)
			if tw == nil {
				tw = d.Timewarp
			}
			sleepCtxFn = tw.SleepWithContext
		}

		if err := sleepCtxFn(ctx, delay); err != nil {
			return awserr.New(request.CanceledErrorCode, "waiter context canceled", err)
		}
	}

	return awserr.New(request.WaiterResourceNotReadyErrorCode, "exceeded wait attempts", nil)
}

//Unfortunately not exported in a usable form
// match returns if the acceptor found a match with the passed in request
// or error. True is returned if the acceptor made a match, error is returned
// if there was an error attempting to perform the match.
func  match(a request.WaiterAcceptor, out interface{}, err error) (bool, error) {
	result := false
	var vals []interface{}

	switch a.Matcher {
	case request.PathAllWaiterMatch, request.PathWaiterMatch:
		// Require all matches to be equal for result to match
		vals, _ = awsutil.ValuesAtPath(out, a.Argument)
		if len(vals) == 0 {
			break
		}
		result = true
		for _, val := range vals {
			if !awsutil.DeepEqual(val, a.Expected) {
				result = false
				break
			}
		}
	case request.PathAnyWaiterMatch:
		// Only a single match needs to equal for the result to match
		vals, _ = awsutil.ValuesAtPath(out, a.Argument)
		for _, val := range vals {
			if awsutil.DeepEqual(val, a.Expected) {
				result = true
				break
			}
		}
	case request.PathListWaiterMatch:
		// ignored matcher
	case request.StatusWaiterMatch:
		//TODO: Can we extract status code from err.(awserr.Error) - or just set something up to simulate it
		//s := a.Expected.(int)
		//result = s == req.HTTPResponse.StatusCode
	case request.ErrorWaiterMatch:
		if aerr, ok := err.(awserr.Error); ok {
			result = aerr.Code() == a.Expected.(string)
		}
	default:
		//TODO: unknown waiter - sdk just logs
	}

	if !result {
		// If there was no matching result found there is nothing more to do
		// for this response, retry the request.
		return false, nil
	}

	switch a.State {
	case request.SuccessWaiterState:
		// waiter completed
		return true, nil
	case request.FailureWaiterState:
		// Waiter failure state triggered
		return true, awserr.New(request.WaiterResourceNotReadyErrorCode,
			"failed waiting for successful resource state", err)
	case request.RetryWaiterState:
		// clear the error and retry the operation
		return false, nil
	default:
		//TODO: Unknown state
		return false, nil
	}
}


