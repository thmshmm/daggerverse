// A module for yamllint functions
//
// This module provides functions for linting YAML files using yamllint. The
// module can be used to lint a directory of YAML files and return the output.
// The yamllint version can be specified as an optional argument, otherwise the
// default version will be used.

package main

import (
	"context"
	"dagger/yamllint/internal/dagger"
)

const YAMLLINT_VERSION = "1.35.1-r1"

type Yamllint struct{}

// Lint runs yamllint on the provided directory
func (m *Yamllint) Lint(
	ctx context.Context,
	src *dagger.Directory,
	// +optional
	version string,
) (string, error) {
	if version == "" {
		version = YAMLLINT_VERSION
	}

	return dag.Container().
		From("alpine:latest").
		WithExec([]string{"apk", "update"}).
		WithExec([]string{"apk", "add", "yamllint=" + version}).
		WithMountedDirectory("/mnt", src).
		WithWorkdir("/mnt").
		WithExec([]string{"yamllint", "."}).
		Stdout(ctx)
}
