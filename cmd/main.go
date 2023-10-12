package main

import (
	_ "embed"
	"fmt"

	"github.com/Omar622/secret-generator/app"
	"github.com/go-ole/go-ole"
)

func main() {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	for {
		fmt.Println("start new session:")

		sg := app.NewSecretGenerator()

		if err := sg.ReadInput(); err != nil {
			fmt.Printf("invalid input types: %v\n\n", err)
			continue
		}

		if !sg.IsValid() {
			fmt.Print("invalid input values\n\n")
			continue
		}

		w := app.NewWriter(sg.MatchRanges())

		if err := w.WriteReport(); err != nil {
			fmt.Printf("something went error while writing report: %v\n\n", err)
			continue
		}

		if err := w.WriteSecretNumbers(); err != nil {
			fmt.Printf("something went error while writing secret numbers: %v\n\n", err)
			continue
		}
	}

}
