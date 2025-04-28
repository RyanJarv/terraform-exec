// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tfexec

import (
	"context"
	"os/exec"
)

type stateShowConfig struct {
	backup      string
	backupOut   string
	dryRun      bool
	lock        bool
	lockTimeout string
	state       string
	stateOut    string
}

var defaultStateShowOptions = stateShowConfig{
	lock:        true,
	lockTimeout: "0s",
}

// StateShowCmdOption represents options used in the Refresh method.
type StateShowCmdOption interface {
	configureStateShow(*stateShowConfig)
}

func (opt *StateOption) configureStateShow(conf *stateShowConfig) {
	conf.state = opt.path
}

// StateShow represents the terraform state rm subcommand.
func (tf *Terraform) StateShow(ctx context.Context, address string, opts ...StateShowCmdOption) ([]byte, error) {
	cmd, err := tf.stateShowCmd(ctx, address, opts...)
	if err != nil {
		return nil, err
	}
	return tf.runTerraformCmdWithOutput(ctx, cmd)
}

func (tf *Terraform) stateShowCmd(ctx context.Context, address string, opts ...StateShowCmdOption) (*exec.Cmd, error) {
	c := defaultStateShowOptions

	for _, o := range opts {
		o.configureStateShow(&c)
	}

	args := []string{"state", "show"}

	// string opts: only pass if set
	if c.state != "" {
		args = append(args, "-state="+c.state)
	}

	// positional arguments
	args = append(args, address)

	return tf.buildTerraformCmd(ctx, nil, args...), nil
}
