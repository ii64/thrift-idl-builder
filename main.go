// SPDX-License-Identifier: MIT
package main

import "os"

func main() {
	if err := Main(os.Args[1:]); err != nil {
		os.Exit(1)
	}
}
