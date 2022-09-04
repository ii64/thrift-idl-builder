// SPDX-License-Identifier: MIT
package main

import (
	"fmt"
	"os"
	"strings"
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
		"-gen", fmt.Sprintf(`go:%s`, strings.Join([]string{
			"thrift_import_path=github.com/apache/thrift/lib/go/thrift",
			"package_prefix=github.com/ii64/thrift-idl-builder/internal/test/gen/",
		}, ",")),
	})
	if err != nil {
		t.Fail()
	}
}
