package main

import (
	_ "embed"
	"fmt"
	"log"
	"os"

	"github.com/Omar622/secret-generator/app"
	"github.com/go-ole/go-ole"
)

func exit() {
	fmt.Println("press enter to exit.")

	var dumInput string
	fmt.Scanln(&dumInput)

	os.Exit(0)
}

func main() {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	sg := app.NewSecretGenerator()

	if err := sg.ReadInput(); err != nil {
		log.Printf("invalid input types: %v\n", err)
		exit()
	}

	if !sg.IsValid() {
		log.Println("invalid input values")
		exit()
	}

	w := app.NewWriter(sg.MatchRanges())

	if err := w.WriteReport(); err != nil {
		log.Printf("something went error while writing report: %v\n", err)
		exit()
	}

	if err := w.WriteSecretNumbers(); err != nil {
		log.Printf("something went error while writing secret numbers: %v\n", err)
		exit()
	}

	exit()
}
