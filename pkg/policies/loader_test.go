//
// Copyright 2024 The Chainloop Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package policies

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEnsureScheme(t *testing.T) {
	tests := []struct {
		name     string
		ref      string
		expected []string
		wantPath string
		wantErr  bool
	}{
		{
			name:     "we want ta file scheme and received one",
			ref:      "file:///tmp/policy.json",
			expected: []string{fileScheme},
			wantPath: "/tmp/policy.json",
		},
		{
			name:     "we want a file scheme and received a chainloop scheme",
			ref:      "chainloop:///policy.json",
			expected: []string{fileScheme},
			wantErr:  true,
		},
		{
			name:     "we want a https scheme",
			ref:      "https://example.com/policy.json",
			expected: []string{httpsScheme},
			wantPath: "example.com/policy.json",
		},
		{
			name:     "it works with both http and https",
			ref:      "http://example.com/policy.json",
			expected: []string{httpsScheme, httpScheme},
			wantPath: "example.com/policy.json",
		},
		{
			name:     "doest not supports default schema",
			ref:      "built-in-policy",
			expected: []string{chainloopScheme},
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotScheme, err := ensureScheme(tt.ref, tt.expected...)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.wantPath, gotScheme)
		})
	}
}
