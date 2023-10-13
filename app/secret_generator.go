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

	if sg.settingNumberRange.Start, err = internal.ReadUInt("enter the start setting number: "); err != nil {
		return err
	}
	if sg.settingNumberRange.End, err = internal.ReadUInt("enter the last setting number: "); err != nil {
		return err
	}
	if sg.secretNumberRange.Start, err = internal.ReadUInt("enter the start secret number: "); err != nil {
		return err
	}
	if sg.secretNumberRange.End, err = internal.ReadUInt("enter the last secret number: "); err != nil {
		return err
	}

	return nil
}

// IsValid check input validity to generate matching
func (sg *SecretGenerator) IsValid() bool {
	if sg.settingNumberRange.Start%1000 != 1 || sg.settingNumberRange.End%100 != 0 ||
		sg.settingNumberRange.Start > sg.settingNumberRange.End {
		return false
	}

	if sg.secretNumberRange.Start%1000 != 1 || sg.secretNumberRange.End%100 != 0 ||
		sg.secretNumberRange.Start > sg.secretNumberRange.End {
		return false
	}

	if sg.settingNumberRange.Count() >= 500 {
		if sg.settingNumberRange.Count()%500 != 0 || sg.secretNumberRange.Count()%500 != 0 {
			return false
		}

		settingNumOf500 := sg.settingNumberRange.Count() / 500
		secretNumOf500 := sg.secretNumberRange.Count() / 500

		if secretNumOf500 < 2*settingNumOf500-1 {
			return false
		}
	} else if sg.secretNumberRange.Count() < sg.settingNumberRange.Count() {
		return false
	}

	return true
}

// MatchRanges match setting number ranges with secret number ranges randomly
func (sg *SecretGenerator) MatchRanges() []internal.MatchingRange {
	var matchingRanges []internal.MatchingRange

	for {
		matchingRanges = sg.generateMatchingRanges()

		internal.PrintMatchingRanges(matchingRanges)

		flag, err := internal.ReadUInt("enter 1 to write data into access files or enter anything else to generate another result: ")
		if flag == 1 && err == nil {
			break
		}
	}

	return matchingRanges
}

func (sg *SecretGenerator) getRandomRanges(size, count uint) []internal.Range {
	ranges := []internal.Range{}

	for i := sg.secretNumberRange.Start; i < sg.secretNumberRange.End; i += 1000 {
		ranges = append(ranges, internal.Range{Start: i, End: i + size - 1})
	}

	ranges = internal.Shuffle(ranges, sg.rnd)

	return ranges[:count]
}

func (sg *SecretGenerator) split50(ranges []internal.Range) []internal.Range {
	ranges50 := []internal.Range{}

	for _, r := range ranges {
		for i := r.Start; i < r.End; i += 50 {
			ranges50 = append(ranges50, internal.Range{Start: i, End: i + 49})
		}
	}

	return ranges50
}

func (sg *SecretGenerator) generateMatchingRanges() []internal.MatchingRange {
	// get random 500s and split them by 50
	var ranges50 []internal.Range

	if sg.settingNumberRange.Count() < 500 {
		ranges50 = sg.split50(sg.getRandomRanges(sg.settingNumberRange.Count(), 1))
	} else {
		ranges50 = sg.split50(sg.getRandomRanges(500, sg.settingNumberRange.Count()/500))
	}

	// classify secret ranges
	secretNumberRanges01 := []internal.Range{}
	secretNumberRanges51 := []internal.Range{}

	for _, r := range ranges50 {
		if r.Start%100 == 1 {
			secretNumberRanges01 = append(secretNumberRanges01, r)
		} else {
			secretNumberRanges51 = append(secretNumberRanges51, r)
		}
	}

	// shuffle secret ranges
	internal.Shuffle(secretNumberRanges01, sg.rnd)
	internal.Shuffle(secretNumberRanges51, sg.rnd)

	// split & classify setting ranges to ranges of size 50
	settingNumberRanges01 := []internal.Range{}
	settingNumberRanges51 := []internal.Range{}

	for i := sg.settingNumberRange.Start; i < sg.settingNumberRange.End; i += 50 {
		if i%100 == 1 {
			settingNumberRanges01 = append(settingNumberRanges01, internal.Range{Start: i, End: i + 49})
		} else {
			settingNumberRanges51 = append(settingNumberRanges51, internal.Range{Start: i, End: i + 49})
		}
	}

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
