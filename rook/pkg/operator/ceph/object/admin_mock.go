/*
Copyright 2021 The Rook Authors. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package object

import "net/http"

// MockClient is the mock of the HTTP Client
// It can be used to mock HTTP request/response from the rgw admin ops API
type MockClient struct {
	// MockDo is a type that mock the Do method from the HTTP package
	MockDo MockDoType
}

// MockDoType is a custom type that allows setting the function that our Mock Do func will run instead
type MockDoType func(req *http.Request) (*http.Response, error)

// Do is the mock client's `Do` func
func (m *MockClient) Do(req *http.Request) (*http.Response, error) { return m.MockDo(req) }
