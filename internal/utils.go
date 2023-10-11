package internal

import (
	"fmt"
	"strconv"
)

// ReadUInt read unsigned integer from stdin
func ReadUInt(lineToPrint string) (uint, error) {
	var inputStr string

	fmt.Print(lineToPrint)
	if _, err := fmt.Scanln(&inputStr); err != nil {
		return 0, err
	}

	inputUInt, err := strconv.Atoi(inputStr)
	return uint(inputUInt), err
}
