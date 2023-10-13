package main

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/Omar622/secret-generator/app"
	"github.com/Omar622/secret-generator/internal"
)

func exit(code int) {
	fmt.Print("\npress enter to exit...")

	var dumInput string
	fmt.Scanln(&dumInput)

	os.Exit(code)
}

func run() ([]internal.MatchingRange, error) {
	sg := app.NewSecretGenerator()

	if err := sg.ReadInput(); err != nil {
		return []internal.MatchingRange{}, fmt.Errorf("invalid input types: %v", err)
	}

	if !sg.IsValid() {
		return []internal.MatchingRange{}, fmt.Errorf("invalid input values")
	}

	return sg.MatchRanges(), nil
}

func write(mr []internal.MatchingRange) error {
	w := app.NewWriter(mr)

	if err := w.WriteReport(); err != nil {
		return fmt.Errorf("something went error while writing report: %v", err)
	}
	fmt.Println("report has been written successfully")

	if err := w.WriteSecretNumbers(); err != nil {
		return fmt.Errorf("something went error while writing secret numbers: %v", err)
	}
	fmt.Println("secret numbers have been written successfully")

	return nil
}

func interact() {
	totalMR := []internal.MatchingRange{}
	for {
		mr, err := run()
		if err != nil {
			fmt.Println(err)
			continue
		}

		input, err := internal.ReadUInt("press enter to continue or enter 1 to exit: ")
		if input == 1 && err == nil {
			totalMR = append(totalMR, mr...)
			break
		}

		totalMR = append(totalMR, mr...)
	}

	if err := write(totalMR); err != nil {
		fmt.Println(err)
		exit(1)
	}

	exit(0)
}

func main() {
	interact()
}
