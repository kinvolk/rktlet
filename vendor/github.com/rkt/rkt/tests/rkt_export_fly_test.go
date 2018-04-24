// Copyright 2016 The rkt Authors
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

// +build fly

package main

import (
	"testing"

	"github.com/rkt/rkt/common"
	"github.com/rkt/rkt/tests/testutils"
)

func TestExport(t *testing.T) {
	if err := common.SupportsOverlay(); err != nil {
		t.Skipf("Overlay fs not supported: %v", err)
	}
	ctx := testutils.NewRktRunCtx()
	defer ctx.Cleanup()

	// TODO(iaguis): we need a new function to unmount the fly pod so we can also test
	// overlaySimulateReboot
	exportTestCases["overlaySimpleTest"].Execute(t, ctx)
}
