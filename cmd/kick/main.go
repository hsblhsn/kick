package main

import (
	"log"

	"github.com/hsblhsn/kick/cli"
)

func main() {
	buildCMD := cli.NewBuildCMD()
	generateCMD := cli.NewGenerateCMD()

	rootCMD := cli.NewRootCMD()
	rootCMD.AddCommand(buildCMD, generateCMD)

	if err := rootCMD.Execute(); err != nil {
		log.Fatal(err)
	}
}
