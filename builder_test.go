// SPDX-License-Identifier: MIT
package main

import (
	"os"
	"testing"
)

func TestBuilder(t *testing.T) {
	os.RemoveAll("./internal/test/gen")
	err := Main([]string{
		"-errors",
		"-wrk", "10",
		"-source-dir", "./idl",
		"-o", "./internal/test/gen",
		"-bin", "thriftgo",
		"-gen", "go",
	})
	if err != nil {
		t.Fail()
	}
}
