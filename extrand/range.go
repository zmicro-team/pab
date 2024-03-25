// Copyright [2022] [thinkgos] thinkgo@aliyun.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package extrand

// Intx returns, as an int, a non-negative pseudo-random number in the half-open interval [min,max).
// It panics if max < min.
func Intx(min, max int) int {
	if min > max {
		panic("extrand: invalid argument to Int")
	}
	if min == max {
		return min
	}
	return Intn(max-min) + min
}

// Int31x returns, as an int32, a non-negative pseudo-random number in the half-open interval [min,max).
// It panics if max < min.
func Int31x(min, max int32) int32 {
	if min > max {
		panic("extrand: invalid argument to Int31")
	}
	if min == max {
		return min
	}
	return Int31n(max-min) + min
}

// Int63x returns, as an int64, a non-negative pseudo-random number in the half-open interval [min,max).
// It panics if max < min.
func Int63x(min, max int64) int64 {
	if min > max {
		panic("extrand: invalid argument to Int63")
	}
	if min == max {
		return min
	}
	return Int63n(max-min) + min
}

// Float64x returns, as an float64, a non-negative pseudo-random number in the half-open interval [min,max).
// It panics if max < min.
func Float64x(min, max float64) float64 {
	if min > max {
		panic("extrand: invalid argument to Float64")
	}
	return min + (max-min)*Float64()
}
