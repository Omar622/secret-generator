package main

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/Omar622/secret-generator/app"
	"github.com/go-ole/go-ole"
)

func exit(code int) {
	fmt.Print("\npress enter to exit...")

	var dumInput string
	fmt.Scanln(&dumInput)

	os.Exit(code)
}

func main() {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	sg := app.NewSecretGenerator()

	if err := sg.ReadInput(); err != nil {
		fmt.Printf("invalid input types: %v\n", err)
		exit(1)
	}

	if !sg.IsValid() {
		fmt.Println("invalid input values")
		exit(1)
	}

	w := app.NewWriter(sg.MatchRanges())

	if err := w.WriteReport(); err != nil {
		fmt.Printf("something went error while writing report: %v\n", err)
		exit(1)
	}

	if err := w.WriteSecretNumbers(); err != nil {
		fmt.Printf("something went error while writing secret numbers: %v\n", err)
		exit(1)
	}

	exit(0)
}
