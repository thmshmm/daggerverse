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
	"errors"
	"fmt"
	"path/filepath"
)

const YAMLLINT_VERSION = "1.35.1-r1"
const DEFAULT_PATH = "."
const program = "yamllint"

type Yamllint struct{}

// Lint runs yamllint on the provided directory
func (m *Yamllint) Lint(
	ctx context.Context,
	src *dagger.Directory,
	// +optional
	version string,
	// +optional
	path string,
	// +optional
	strict bool,
) (string, error) {
	if version == "" {
		version = YAMLLINT_VERSION
	}

	if path == "" {
		path = DEFAULT_PATH
	} else {
		path = filepath.Clean(path)
	}

	c := dag.Container().
		From("alpine:latest").
		WithExec([]string{"apk", "update"}).
		WithExec([]string{"apk", "add", "yamllint=" + version}).
		WithMountedDirectory("/mnt", src).
		WithWorkdir("/mnt").
		WithExec(
			command(
				program, strict, path))

	stdout, err := c.Stdout(ctx)

	if err != nil {
		var e *dagger.ExecError
		if errors.As(err, &e) {
			return stdout, fmt.Errorf("yamllint failed with exit code %d", e.ExitCode)
		}

		return stdout, err
	}

	return stdout, nil
}

func command(exec string, strict bool, path string) []string {
	var args []string

	args = append(args, exec)

	if strict {
		args = append(args, "--strict")
	}

	args = append(args, path)

	return args
}
