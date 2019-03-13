package tutorial_30_days_of_code

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/zrma/1d1c/hacker_rank/common/utils"
)

var _ = Describe("https://www.hackerrank.com/challenges/30-class-vs-instance/problem", func() {
	It("문제를 풀었다", func() {
		err := utils.PrintTest(func() error {
			ClassAndInstance(-1)
			ClassAndInstance(10)
			ClassAndInstance(16)
			ClassAndInstance(18)
			return nil
		}, []string{
			"Age is not valid, setting age to 0.",
			"You are young.",
			"You are young.",
			"",
			"You are young.",
			"You are a teenager.",
			"",
			"You are a teenager.",
			"You are old.",
			"",
			"You are old.",
			"You are old.",
			"",
		})
		Expect(err).ShouldNot(HaveOccurred())
	})
})