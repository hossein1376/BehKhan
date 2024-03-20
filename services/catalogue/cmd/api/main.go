package main

import (
	"fmt"
	"os"

	"github.com/hossein1376/BehKhan/catalogue/cmd/api/command"
)

func main() {
	err := command.Run()
	if err != nil {
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
