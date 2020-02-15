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
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/lwoggardner/awsdouble"
	ec2double "github.com/lwoggardner/awsdouble/doubles/ec2"
	"github.com/lwoggardner/godouble/godouble"
	"strings"
	"testing"
	"time"
)

func TestWaitEC2Stopped(t *testing.T) {
	var fakeTime time.Time

	ec2c := ec2double.NewDouble(t, func(d *awsdouble.AWSTestDouble) {
		d.SetTimewarp(func(d time.Duration) <-chan time.Time {
			fakeTime = fakeTime.Add(d)
			result := make(chan time.Time,1)
			result <- fakeTime
			close(result)
			return result
		})
	})

	defer ec2c.Verify()

	mock := ec2c.Mock("DescribeInstances")
	mock.Returning(&ec2.DescribeInstancesOutput{
		Reservations: []*ec2.Reservation{{
			Instances: []*ec2.Instance{{
				InstanceId: aws.String("i-one"),
				State: &ec2.InstanceState{
					Name: aws.String(ec2.InstanceStateNameStopping),
				},

			}},
		}},
	},nil)
	mock.Returning(&ec2.DescribeInstancesOutput{
		Reservations: []*ec2.Reservation{{
			Instances: []*ec2.Instance{{
				InstanceId: aws.String("i-one"),
				State: &ec2.InstanceState{
					Name: aws.String(ec2.InstanceStateNameStopped),
				},

			}},
		}},
	},nil)

	mock.Expect(godouble.Twice())
	err := ec2c.WaitUntilInstanceStopped(&ec2.DescribeInstancesInput{})

	if err != nil {
		t.Errorf("Unexpected error in waiter %v",err)
	}

	if fakeTime.Second() != 15 {
		t.Errorf("Expected to sleep 15 seconds, got %d",fakeTime.Second())
	}
}

func TestWaitEC2StoppedWithOptions(t *testing.T) {
	var fakeTime time.Time

	ec2c := ec2double.NewDouble(t, func(d *awsdouble.AWSTestDouble) {
		d.SetTimewarp(func(d time.Duration) <-chan time.Time {
			fakeTime = fakeTime.Add(d)
			result := make(chan time.Time,1)
			result <- fakeTime
			close(result)
			return result
		})
	})

	defer ec2c.Verify()

	mock := ec2c.Mock("DescribeInstances")
	mock.Returning(&ec2.DescribeInstancesOutput{
		Reservations: []*ec2.Reservation{{
			Instances: []*ec2.Instance{{
				InstanceId: aws.String("i-one"),
				State: &ec2.InstanceState{
					Name: aws.String(ec2.InstanceStateNameStopping),
				},

			}},
		}},
	},nil)

	mock.Expect(godouble.Twice())
	err := ec2c.WaitUntilInstanceStoppedWithContext(context.Background(),&ec2.DescribeInstancesInput{}, func(waiter *request.Waiter) {
		waiter.Delay = func(_ int) time.Duration {
			return time.Duration(40) * time.Second
		}
		waiter.MaxAttempts = 2
	})


	if err == nil {
		t.Errorf("Expected error in waiter, but received nil")
	} else if !strings.Contains(err.Error(),"attempts") {
		t.Errorf("Expected error %s to contain 'attempts'",err.Error())
	}

	if fakeTime.Second() != 40 {
		t.Errorf("Expected to sleep 40 seconds, got %d",fakeTime.Second())
	}
}

