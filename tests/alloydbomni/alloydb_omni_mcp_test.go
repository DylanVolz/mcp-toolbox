// Copyright 2026 Google LLC
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

package alloydbomni

import (
	"context"
	"testing"
	"time"

	"github.com/googleapis/mcp-toolbox/tests"
)

func TestAlloyDBOmniMCPCallTool(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	pool, uniqueID := setupAlloyDBOmniTest(t, ctx, true, false)
	runAlloyDBOmniTests(t, ctx, pool, uniqueID, tests.WithMCPExec())
}

func TestAlloyDBOmniMCPColumnar(t *testing.T) {
	skipUnlessColumnar(t)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	setupAlloyDBOmniTest(t, ctx, true, true)
	runAlloyDBOmniColumnarTests(t, tests.WithMCPExec())
}
