package main

//go:generate go run directives_generate.go

import (
	"github.com/threefoldtech/coredns/coremain"

	// Plug in CoreDNS
	_ "github.com/threefoldtech/coredns/core/plugin"
)

func main() {
	coremain.Run()
}
