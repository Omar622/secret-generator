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

	if sg.settingNumberRange.Count() >= 1000 {
		if sg.settingNumberRange.Count()%1000 != 0 || sg.secretNumberRange.Count()%1000 != 0 {
			return false
		}

		settingNumOf1000 := sg.settingNumberRange.Count() / 1000
		secretNumOf1000 := sg.secretNumberRange.Count() / 1000

		if secretNumOf1000 < 2*settingNumOf1000-1 {
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

func (sg *SecretGenerator) getRandomSecretRanges(size, count uint) []internal.Range {
	ranges := []internal.Range{}

	for i := sg.secretNumberRange.Start; i < sg.secretNumberRange.End; i += 2 * 1000 {
		ranges = append(ranges, internal.Range{Start: i, End: i + size - 1})
	}

	ranges = internal.Shuffle(ranges, sg.rnd)

	return ranges[:count]
}

func (sg *SecretGenerator) generateMatchingRanges() []internal.MatchingRange {
	// get random 1000s and split them by 100
	var secretNumberRanges100 []internal.Range

	if sg.settingNumberRange.Count() < 1000 {
		secretNumberRanges100 = internal.Split100(sg.getRandomSecretRanges(sg.settingNumberRange.Count(), 1))
	} else {
		secretNumberRanges100 = internal.Split100(sg.getRandomSecretRanges(1000, sg.settingNumberRange.Count()/1000))
	}

	// shuffle secret ranges
	internal.Shuffle(secretNumberRanges100, sg.rnd)

	// split & classify setting ranges to ranges of size 50
	settingNumberRanges100 := []internal.Range{}

	for i := sg.settingNumberRange.Start; i < sg.settingNumberRange.End; i += 100 {
		settingNumberRanges100 = append(settingNumberRanges100, internal.Range{Start: i, End: i + 100 - 1})
	}

	// match
	matchingRanges := []internal.MatchingRange{}
	for i := range settingNumberRanges100 {
		matchingRanges = append(matchingRanges, internal.MatchingRange{
			SettingNumberRange: settingNumberRanges100[i], SecretNumberRange: secretNumberRanges100[i],
		})
	}

	// sort
	matchingRanges = internal.SortMatchingRanges(matchingRanges)

	return matchingRanges
}
