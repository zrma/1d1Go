package tutorial_30_days_of_code

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/zrma/1d1c/hacker_rank/common/utils"
)

var _ = Describe("https://www.hackerrank.com/challenges/30-data-types/problem", func() {
	It("문제를 풀었다", func() {
		err := utils.PrintTest(func() error {
			dataType(
				4, 12,
				4.0, 4.0,
				"HackerRank ", "is the best place to learn and practice coding!",
			)
			return nil
		}, []string{
			"16",
			"8.0",
			"HackerRank is the best place to learn and practice coding!",
		})
		Expect(err).ShouldNot(HaveOccurred())
	})
})