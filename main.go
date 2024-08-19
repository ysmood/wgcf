package main

import (
	"log"

	"github.com/ysmood/wgcf/v2/cmd"
	"github.com/ysmood/wgcf/v2/util"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(util.GetErrorMessage(err))
	}
}
