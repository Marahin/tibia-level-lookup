package tibialevellookup_test

import (
	"testing"

	"github.com/marahin/tibialevellookup"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestTibiaLevelLookup(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TibiaLevelLookup Suite")
}

var experienceTableLevels = map[int]uint64{
	44:   1238400,
	501:  2070900000,
	1001: 16616800000,
	1631: 72046489000,
}

var above2k = map[int]uint64{
	2062: 145814328084,
	2059: 145183075916,
	2004: 133780473475,
}

var _ = Describe("Level to experience", func() {
	Context("level contained within the Tibia.com level table", func() {
		It("returns proper experience amount", func() {
			for level, expectedExperience := range experienceTableLevels {
				calculatedExperience := tibialevellookup.LevelToExperience(level)

				Expect(calculatedExperience).To(Equal(expectedExperience))
			}
		})
	})

	Context("level outside Tibia.com level table (calculated)", func() {
		It("returns proper experience amount", func() {
			for level, realLifeExperience := range above2k {
				Expect(realLifeExperience).To(BeNumerically(">=", tibialevellookup.LevelToExperience(level)))
				Expect(realLifeExperience).To(BeNumerically("<", tibialevellookup.LevelToExperience(level+1)))
			}
		})
	})
})

var _ = Describe("Experience to level", func() {
	It("throws an error if GenerateExperienceTable wasn't invoked", func() {
		tibialevellookup.ClearExpTable()

		_, err := tibialevellookup.ExperienceToLevel(16616800000)

		Expect(err).NotTo(BeNil())
		Expect(err.Error()).To(ContainSubstring("cannot use ExperienceToLevel() if expTable"))
	})

	It("throws an error if experience was not matched", func() {
		tibialevellookup.GenerateExperienceTable()

		_, err := tibialevellookup.ExperienceToLevel(1132933899800)

		Expect(err).NotTo(BeNil())
		Expect(err.Error()).To(ContainSubstring("not matched - perhaps pass a higher number to GenerateExperienceTable"))
	})

	Context("experience contained within the Tibia.com level table", func() {
		Context("exact number", func() {
			It("returns proper level", func() {
				tibialevellookup.GenerateExperienceTable()

				Expect(tibialevellookup.ExperienceToLevel(16616800000)).To(Equal(1001))
			})
		})

		Context("somewhere in betweeen", func() {
			It("returns proper level", func() {
				tibialevellookup.GenerateExperienceTable()

				Expect(tibialevellookup.ExperienceToLevel(16626800000)).To(Equal(1001))
			})
		})
	})

	Context("experience outside the Tibia.com level table", func() {
		It("returns a proper number", func() {
			for level, realLifeExperience := range above2k {
				tibialevellookup.GenerateExperienceTable()

				Expect(tibialevellookup.ExperienceToLevel(realLifeExperience)).To(Equal(level))
			}
		})
	})
})
