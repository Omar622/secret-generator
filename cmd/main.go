package main

import (
	_ "embed"
	"log"

	"github.com/Omar622/secret-generator/app"
	"github.com/go-ole/go-ole"
)

func main() {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	sg := app.NewSecretGenerator()

	if err := sg.ReadInput(); err != nil {
		log.Fatalf("invalid input types: %v\n", err)
	}

	if !sg.IsValid() {
		log.Fatal("invalid input values\n")
	}

	w := app.NewWriter(sg.MatchRanges())

	if err := w.WriteReport(); err != nil {
		log.Fatalf("something went error while writing report: %v\n", err)
	}

	if err := w.WriteSecretNumbers(); err != nil {
		log.Fatalf("something went error while writing secret numbers: %v\n", err)
	}
}
