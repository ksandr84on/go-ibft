package main

import (
	"fmt"
	"io"

	"github.com/ksandr84on/go-ibft/core"
)

func main() {
	b := core.NewIBFT(nil, nil, nil)

	// prevent golang compiler from removing the whole function
	fmt.Sprint(io.Discard, b)
}
