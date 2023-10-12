package app

import (
	"math/rand"
	"time"

	"github.com/Omar622/secret-generator/internal"
)

// NewSecretGenerator is a factory function for SecretGenerator
func NewSecretGenerator() SecretGenerator {
	return SecretGenerator{rnd: rand.New(rand.NewSource(time.Now().Unix()))}
}

// SecretGenerator is a SecretGenerator struct
type SecretGenerator struct {
	settingNumberRange internal.Range
	secretNumberRange  internal.Range

	rnd *rand.Rand
}

// ReadInput read SecretGenerator ranges from stdin
func (sg *SecretGenerator) ReadInput() error {
	var err error

	if sg.settingNumberRange.Start, err = internal.ReadUInt("enter the Start setting number: "); err != nil {
		return err
	}
	if sg.settingNumberRange.End, err = internal.ReadUInt("enter the last setting number: "); err != nil {
		return err
	}
	if sg.secretNumberRange.Start, err = internal.ReadUInt("enter the Start secret number: "); err != nil {
		return err
	}
	if sg.secretNumberRange.End, err = internal.ReadUInt("enter the last secret number: "); err != nil {
		return err
	}

	return nil
}

// IsValid check input validity to generate matching
func (sg *SecretGenerator) IsValid() bool {
	if sg.settingNumberRange.Start%50 != 1 || sg.settingNumberRange.End%50 != 0 ||
		sg.settingNumberRange.Start > sg.settingNumberRange.End {
		return false
	}

	if sg.secretNumberRange.Start%50 != 1 || sg.secretNumberRange.End%50 != 0 ||
		sg.secretNumberRange.Start > sg.secretNumberRange.End {
		return false
	}

	settingMatchingRangesCount := (sg.settingNumberRange.End - sg.settingNumberRange.Start + 1) / 50
	secretMatchingRangesCount := (sg.secretNumberRange.End - sg.secretNumberRange.Start + 1) / 50

	if secretMatchingRangesCount < settingMatchingRangesCount {
		return false
	}

	if secretMatchingRangesCount == settingMatchingRangesCount &&
		sg.settingNumberRange.Start%100 != sg.secretNumberRange.Start%100 {
		return false
	}

	return true
}

// MatchRanges match setting number ranges with secret number ranges randomly
func (sg *SecretGenerator) MatchRanges() []internal.MatchingRange {
	matchingRanges := []internal.MatchingRange{}

	generateAgain := true
	for generateAgain {
		matchingRanges = sg.generateMatchingRanges()

		internal.PrintMatchingRanges(matchingRanges)

		flag, err := internal.ReadUInt("enter 1 to write data into access files or enter anything else to generate another result: ")
		if flag == 1 && err == nil {
			generateAgain = false
		}
	}

	return matchingRanges
}

func (sg *SecretGenerator) generateMatchingRanges() []internal.MatchingRange {
	// split ranges to ranges of size 50
	settingNumberRanges01 := []internal.Range{}
	settingNumberRanges51 := []internal.Range{}

	for i := sg.settingNumberRange.Start; i < sg.settingNumberRange.End; i += 50 {
		if i%100 == 1 {
			settingNumberRanges01 = append(settingNumberRanges01, internal.Range{Start: i, End: i + 49})
		} else {
			settingNumberRanges51 = append(settingNumberRanges51, internal.Range{Start: i, End: i + 49})
		}
	}

	secretNumberRanges01 := []internal.Range{}
	secretNumberRanges51 := []internal.Range{}

	for i := sg.secretNumberRange.Start; i < sg.secretNumberRange.End; i += 50 {
		if i%100 == 1 {
			secretNumberRanges01 = append(secretNumberRanges01, internal.Range{Start: i, End: i + 49})
		} else {
			secretNumberRanges51 = append(secretNumberRanges51, internal.Range{Start: i, End: i + 49})
		}
	}

	// shuffle
	internal.Shuffle(secretNumberRanges01, sg.rnd)
	internal.Shuffle(secretNumberRanges51, sg.rnd)

	// match
	matchingRanges := []internal.MatchingRange{}
	for i := range settingNumberRanges01 {
		matchingRanges = append(matchingRanges, internal.MatchingRange{
			SettingNumberRange: settingNumberRanges01[i], SecretNumberRange: secretNumberRanges01[i],
		})
	}

	for i := range settingNumberRanges51 {
		matchingRanges = append(matchingRanges, internal.MatchingRange{
			SettingNumberRange: settingNumberRanges51[i], SecretNumberRange: secretNumberRanges51[i],
		})
	}

	// sort
	matchingRanges = internal.SortMatchingRanges(matchingRanges)

	return matchingRanges
}
