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
	"testing"

	"github.com/aws/aws-sdk-go/aws"
)

func AssertAwsStringEquals(t *testing.T, prefix string, actual *string, expected string) {
	sv := aws.StringValue(actual)
	if sv != expected {
		t.Errorf("%s: Expected %s, Actual %s", prefix, sv, expected)
	}
}

func AssertAwsInt64Equals(t *testing.T, prefix string, actual *int64, expected int64) {
	sv := aws.Int64Value(actual)
	if sv != expected {
		t.Errorf("%s: Expected %d, Actual %d", prefix, sv, expected)
	}
}
