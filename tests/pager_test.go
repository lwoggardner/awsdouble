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

package tests

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/swf"
	swfdouble "github.com/lwoggardner/awsdouble/doubles/swf"
	. "github.com/lwoggardner/godouble/godouble"
	"testing"
)

func TestAWSTestDouble_Paginate(t *testing.T) {
	swfc := swfdouble.NewDouble(t)

	m := swfc.Mock("GetWorkflowExecutionHistory")
	m.Returning(&swf.GetWorkflowExecutionHistoryOutput{NextPageToken: aws.String("page0",)},nil)
	m.Returning(&swf.GetWorkflowExecutionHistoryOutput{NextPageToken: aws.String("page1",)},nil)
	m.Returning(&swf.GetWorkflowExecutionHistoryOutput{NextPageToken: nil}, nil)
	m.Expect(Exactly(3))

	count := 0
	err := swfc.GetWorkflowExecutionHistoryPages(&swf.GetWorkflowExecutionHistoryInput{}, func(output *swf.GetWorkflowExecutionHistoryOutput, lastPage bool) bool {
		count++
		if count == 3 && !lastPage {
			t.Errorf("On 3rd page we expect last page to be true, but was false")
		}
		if count < 3 && lastPage {
			t.Errorf("On first 2 pages we expect last page to be false, but was true")
		}
		return true
	})

	if err != nil {
		t.Errorf("Unexpected error during paging %v",err)
	}
	if count != 3 {
		t.Errorf("Expected pager 3 times, got %d",count)
	}

	s := swfc.Spy("GetWorkflowExecutionHistoryWithContext")
	s.Expect(Exactly(3))
}


func TestAWSTestDouble_PaginateRespectsPagerShouldContinue(t *testing.T) {
	swfc := swfdouble.NewDouble(t)

	m := swfc.Mock("GetWorkflowExecutionHistory")
	m.Returning(&swf.GetWorkflowExecutionHistoryOutput{NextPageToken: aws.String("page0",)},nil)
	m.Returning(&swf.GetWorkflowExecutionHistoryOutput{NextPageToken: aws.String("page1",)},nil)
	m.Returning(&swf.GetWorkflowExecutionHistoryOutput{NextPageToken: nil}, nil)
	m.Expect(Exactly(2))

	count := 0
	err := swfc.GetWorkflowExecutionHistoryPages(&swf.GetWorkflowExecutionHistoryInput{}, func(out *swf.GetWorkflowExecutionHistoryOutput, b bool) bool {
		count++
		if count == 2 && aws.StringValue(out.NextPageToken) == "page1"{
			return false //stop paging early
		}
		return true
	})

	if err != nil {
		t.Errorf("Unexpected error during paging %v",err)
	}
	if count != 2 {
		t.Errorf("Expected pager 2 times, got %d",count)
	}

	s := swfc.Spy("GetWorkflowExecutionHistoryWithContext")
	s.Expect(Exactly(2))
}