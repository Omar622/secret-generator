package main

import (
	_ "embed"
	"fmt"

	"github.com/Omar622/secret-generator/app"
	"github.com/go-ole/go-ole"
)

func run() {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	fmt.Print("start new session:\n\n")

	sg := app.NewSecretGenerator()

	if err := sg.ReadInput(); err != nil {
		fmt.Printf("invalid input types: %v\n\n", err)
		return
	}

	if !sg.IsValid() {
		fmt.Print("invalid input values\n\n")
		return
	}

	w := app.NewWriter(sg.MatchRanges())

	if err := w.WriteReport(); err != nil {
		fmt.Printf("something went error while writing report: %v\n\n", err)
		return
	}

	if err := w.WriteSecretNumbers(); err != nil {
		fmt.Printf("something went error while writing secret numbers: %v\n\n", err)
		return
	}
}

func main() {
	for {
		run()
	}
}
