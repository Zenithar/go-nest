package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"go.zenithar.org/nest/cmd/nest/cmd"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	command := cmd.NewDefaultNestCommand()

	if err := command.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
