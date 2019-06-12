package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"go.zenithar.org/nest/cmd/nest-snake/cmd"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	command := cmd.NewDefaultSnakeCommand()

	if err := command.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
