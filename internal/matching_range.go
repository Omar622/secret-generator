package internal

import (
	"fmt"
	"sort"
)

// MatchingRange is a MatchingRange struct
type MatchingRange struct {
	SettingNumberRange Range
	SecretNumberRange  Range
}

// Compare compare two Ranges
func (g *MatchingRange) Compare(other MatchingRange) int {
	settingNumberRangeStatus := g.SettingNumberRange.Compare(other.SettingNumberRange)
	if settingNumberRangeStatus != 0 {
		return settingNumberRangeStatus
	}

	return g.SecretNumberRange.Compare(other.SecretNumberRange)
}

// SortMatchingRanges sort MatchingRanges
func SortMatchingRanges(matchingRanges []MatchingRange) []MatchingRange {
	sort.Slice(matchingRanges, func(i, j int) bool {
		return matchingRanges[i].Compare(matchingRanges[j]) == -1
	})

	return matchingRanges
}

// PrintMatchingRanges print MatchingRanges in stdout
func PrintMatchingRanges(MatchingRanges []MatchingRange) {
	fmt.Println("generated MatchingRanges:")

	for _, g := range MatchingRanges {
		fmt.Printf("%v -> %v\n", g.SettingNumberRange, g.SecretNumberRange)
	}
}
