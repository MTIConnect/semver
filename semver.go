package main

import (
	"fmt"

	"github.com/MTIConnect/semver/args"
)

func main() {
	args, err := args.ParseArgs()
	if err != nil {
		panic(err)
	}
	if err := args.Process(); err != nil {
		panic(err)
	}
	fmt.Print(args.Version.String())
}
