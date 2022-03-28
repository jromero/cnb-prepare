package types_test

import (
	"testing"

	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"

	"github.com/jromero/cnb-prepare/pkg/project/types"
	"github.com/jromero/cnb-prepare/pkg/testhelpers"
)

func TestAPIVersion(t *testing.T) {
	spec.Run(t, "APIVersion", testAPIVersion, spec.Parallel(), spec.Report(report.Terminal{}))
}

func testAPIVersion(t *testing.T, when spec.G, it spec.S) {
	when("#Equal", func() {
		it("is equal to comparison", func() {
			subject := types.ParseVersion("0.2")
			comparison := types.ParseVersion("0.2")

			testhelpers.Equals(t, subject.Equal(comparison), true)
		})

		it("is not equal to comparison", func() {
			subject := types.ParseVersion("0.2")
			comparison := types.ParseVersion("0.3")

			testhelpers.Equals(t, subject.Equal(comparison), false)
		})
	})

	when("IsSupersetOf", func() {
		when("0.x", func() {
			it("matching Minor value", func() {
				v := types.ParseVersion("0.2")
				target := types.ParseVersion("0.2")

				testhelpers.Equals(t, v.IsSupersetOf(target), true)
			})

			it("Minor > target Minor", func() {
				v := types.ParseVersion("0.2")
				target := types.ParseVersion("0.1")

				testhelpers.Equals(t, v.IsSupersetOf(target), false)
			})

			it("Minor < target Minor", func() {
				v := types.ParseVersion("0.1")
				target := types.ParseVersion("0.2")

				testhelpers.Equals(t, v.IsSupersetOf(target), false)
			})
		})

		when("1.x", func() {
			it("matching Major and Minor", func() {
				v := types.ParseVersion("1.2")
				target := types.ParseVersion("1.2")

				testhelpers.Equals(t, v.IsSupersetOf(target), true)
			})

			it("matching Major but Minor > target Minor", func() {
				v := types.ParseVersion("1.2")
				target := types.ParseVersion("1.1")

				testhelpers.Equals(t, v.IsSupersetOf(target), true)
			})

			it("matching Major but Minor < target Minor", func() {
				v := types.ParseVersion("1.1")
				target := types.ParseVersion("1.2")

				testhelpers.Equals(t, v.IsSupersetOf(target), false)
			})

			it("Major < target Major", func() {
				v := types.ParseVersion("1.0")
				target := types.ParseVersion("2.0")

				testhelpers.Equals(t, v.IsSupersetOf(target), false)
			})

			it("Major > target Major", func() {
				v := types.ParseVersion("2.0")
				target := types.ParseVersion("1.0")

				testhelpers.Equals(t, v.IsSupersetOf(target), false)
			})
		})
	})

	when("#LessThan", func() {
		var subject = types.ParseVersion("0.3")
		var toTest = map[string]bool{
			"0.2": false,
			"0.3": false,
			"0.4": true,
		}
		it("returns the expected value", func() {
			for comparison, expected := range toTest {
				testhelpers.Equals(t, subject.LessThan(comparison), expected)
			}
		})
	})

	when("#AtLeast", func() {
		var subject = types.ParseVersion("0.3")
		var toTest = map[string]bool{
			"0.2": true,
			"0.3": true,
			"0.4": false,
		}
		it("returns the expected value", func() {
			for comparison, expected := range toTest {
				testhelpers.Equals(t, subject.AtLeast(comparison), expected)
			}
		})
	})
}
