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
	"time"
)

//Replacement for time.After - eg if you are using a fake clock
type Timewarp func(d time.Duration) <-chan time.Time

//This makes every duration wait for 10 ms
func DefaultTimeWarp(_ time.Duration) <- chan time.Time {
	return time.After(time.Duration(10) * time.Millisecond)
}

type timewarpContextKey string
const timeWarpKey = "timeWarpContextKey"

//Inject a Timewarp function into context
func WithTimeWarp(ctx context.Context, warp Timewarp) context.Context {
	return context.WithValue(ctx, timewarpContextKey(timeWarpKey), warp)
}

func timeWarpFromContext(ctx context.Context) Timewarp {
	tw, _ := ctx.Value(timewarpContextKey(timeWarpKey)).(Timewarp)
	return tw
}

func (tw Timewarp) SleepWithContext(ctx context.Context, duration time.Duration) error {
	select {
	case <-tw(duration):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
