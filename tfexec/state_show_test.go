// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tfexec

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-exec/tfexec/internal/testutil"
)

func TestStateShowCmd(t *testing.T) {
	td := t.TempDir()

	tf, err := NewTerraform(td, tfVersion(t, testutil.Latest_v1))
	if err != nil {
		t.Fatal(err)
	}

	// empty env, to avoid environ mismatch in testing
	tf.SetEnv(map[string]string{})

	t.Run("defaults", func(t *testing.T) {
		stateShowCmd, err := tf.stateShowCmd(context.Background(), "testAddress")
		if err != nil {
			t.Fatal(err)
		}

		assertCmd(t, []string{
			"state",
			"show",
			"testAddress",
		}, nil, stateShowCmd)
	})

	t.Run("override all defaults", func(t *testing.T) {
		stateShowCmd, err := tf.stateShowCmd(context.Background(), "testAddress", State("teststate"))
		if err != nil {
			t.Fatal(err)
		}

		assertCmd(t, []string{
			"state",
			"show",
			"-state=teststate",
			"testAddress",
		}, nil, stateShowCmd)
	})
}
