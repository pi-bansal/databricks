// Copyright (c) 2025 ADBC Drivers Contributors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//         http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package databricks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithArrowNativeDecimalDSN(t *testing.T) {
	tests := []struct {
		name string
		dsn  string
		want string
	}{
		{
			name: "no query",
			dsn:  "token:abc@host:443/sql/1.0/warehouses/w",
			want: "token:abc@host:443/sql/1.0/warehouses/w?useArrowNativeDecimal=true",
		},
		{
			name: "existing query",
			dsn:  "host:443/sql/1.0/warehouses/w?authType=OAuthM2M&clientID=id",
			want: "host:443/sql/1.0/warehouses/w?authType=OAuthM2M&clientID=id&useArrowNativeDecimal=true",
		},
		{
			name: "empty query",
			dsn:  "host:443/path?",
			want: "host:443/path?useArrowNativeDecimal=true",
		},
		{
			name: "trailing ampersand",
			dsn:  "host:443/path?maxRows=10&",
			want: "host:443/path?maxRows=10&useArrowNativeDecimal=true",
		},
		{
			name: "explicitly disabled",
			dsn:  "host:443/path?useArrowNativeDecimal=false",
			want: "host:443/path?useArrowNativeDecimal=false",
		},
		{
			name: "explicit setting is case insensitive",
			dsn:  "host:443/path?UseArrowNativeDecimal=false",
			want: "host:443/path?UseArrowNativeDecimal=false",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, withArrowNativeDecimalDSN(tt.dsn))
		})
	}
}
