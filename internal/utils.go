package internal

import (
	"fmt"
)

// ReadUInt read unsigned integer from stdin
func ReadUInt(lineToPrint string) (uint, error) {
	var input uint

	fmt.Print(lineToPrint)
	_, err := fmt.Scanln(&input)

	return input, err
}
