package tibialevellookup

import (
	"errors"
	"fmt"
	"math"
)

var expTable []int

// Calculates approximate experience based on level.
// The actual math formula has been taken from: https://tibia.fandom.com/wiki/Experience_Table
// It is worth to mention that Tibia.com's experience table does not exceed level 2000: https://www.tibia.com/library/?subtopic=experiencetable
// It is unsure how precise the formula is itself, but when checking highscores for Bobeek and Goraca (the highest levels at the time of writing this)
// it seems accurate.
func formula(level float64) float64 {
	return math.Floor(((50.0 * math.Pow(level, 3)) / 3.0) - (100 * math.Pow(level, 2)) + ((850 * level) / 3) - 200)
}

// Returns base required experience for a level.
// Uses pregenerated expTable if there's one (see GenerateExperienceTable()),
// otherwise falls back to the formula().
func LevelToExperience(level int) int {
	if len(expTable) >= level {
		return expTable[level]
	} else {
		return int(math.Floor(formula(float64(level))))
	}
}

// Returns approximate level based on experience - basing off of math formula for calculating experience for a level.
// Returns an error if level is out of the table or expTable is nonexistent - see GenerateExperienceTable.
func ExperienceToLevel(experience int) (int, error) {
	if len(expTable) == 0 {
		return 0, errors.New("cannot use ExperienceToLevel() if expTable was not pregenerated - make sure GenerateExperienceTable() is called before calling ExperienceToLevel!")
	}

	for id, _ := range expTable {
		if expTable[id] <= experience && expTable[id+1] > experience {
			return id, nil
		}
	}

	return 0, fmt.Errorf("experience %d not matched - perhaps pass a higher number to GenerateExperienceTable()", experience)
}

// Generates experience table up to the first passed integer argument. Default is 2500.
// Required for ExperienceToLevel function to work.
func GenerateExperienceTable(params ...int) {
	var targetLevel int

	if len(params) == 0 {
		targetLevel = 2500
	} else {
		targetLevel = params[0]
	}

	if len(expTable) >= targetLevel {
		return
	}

	tmpExpTable := make([]int, targetLevel+1)

	for k := 1; k < targetLevel+1; k++ {
		tmpExpTable[k] = LevelToExperience(k)
	}

	expTable = tmpExpTable
}

// Replaces expTable with an empty integer array
func ClearExpTable() {
	expTable = make([]int, 0)
}
