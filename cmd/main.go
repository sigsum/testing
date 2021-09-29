package main

//
// $ go run .
// main.go:9:2: no required module provides package golang.sigsum.org/sigsum-log-go/pkg/types; to add it:
//    go get golang.sigsum.org/sigsum-log-go/pkg/types
// $ go get golang.sigsum.org/sigsum-log-go/pkg/types
// go get: unrecognized import path "golang.sigsum.org/sigsum-log-go/pkg/types": https fetch: Get "https://golang.sigsum.org/sigsum-log-go/pkg/types?go-get=1": dial tcp: lookup golang.sigsum.org: no such host
//
// I.e., the above `go get` command should work with proper setup on our end.
//
// (If this is per-repo configuration in cgit -> need for sigsum-{lib,log}-go.)
//
// Edit: seems to be working now w/ git.sigsum.org

import (
	"fmt"
	"git.sigsum.org/sigsum-log-go/pkg/types"
)

func main() {
	fmt.Printf("types.HashSize: %v\n", types.HashSize)
}
