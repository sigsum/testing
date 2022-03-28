package main

import (
	"log"

	"git.sigsum.org/sigsum-lib-go/pkg/hex"
)

func main() {
	log.Printf("hola: %s", hex.Serialize([]byte{0,1,2,3}))
}
