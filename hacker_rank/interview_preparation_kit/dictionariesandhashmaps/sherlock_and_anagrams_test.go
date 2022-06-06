package dictionariesandhashmaps

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSherlockAndAnagrams(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/sherlock-and-anagrams/problem")

	//goland:noinspection SpellCheckingInspection
	for _, tt := range []struct {
		s    string
		want int32
	}{
		{"abba", 4},
		{"abcd", 0},
		{"ifailuhkqq", 3},
		{"kkkk", 10},
		{"cdcd", 5},
	} {
		t.Run(tt.s, func(t *testing.T) {
			got := sherlockAndAnagrams(tt.s)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestSherlockAndAnagramsPerformance(t *testing.T) {
	//goland:noinspection SpellCheckingInspection
	for _, tt := range []struct {
		s    string
		want int32
	}{
		{"dbcfibibcheigfccacfegicigcefieeeeegcghggdheichgafhdigffgifidfbeaccadabecbdcgieaffbigffcecahafcafhcdg", 1464},
		{"dfcaabeaeeabfffcdbbfaffadcacdeeabcadabfdefcfcbbacadaeafcfceeedacbafdebbffcecdbfebdbfdbdecbfbadddbcec", 2452},
		{"gjjkaaakklheghidclhaaeggllagkmblhdlmihmgkjhkkfcjaekckaafgabfclmgahmdebjekaedhaiikdjmfbmfbdlcafamjbfe", 873},
		{"fdbdidhaiqbggqkhdmqhmemgljaphocpaacdohnokfqmlpmiioedpnjhphmjjnjlpihmpodgkmookedkehfaceklbljcjglncfal", 585},
		{"bcgdehhbcefeeadchgaheddegbiihehcbbdffiiiifgibhfbchffcaiabbbcceabehhiffggghbafabbaaebgediafabafdicdhg", 1305},
		{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", 166650},
		{"mhmgmbbccbbaffhbncgndbffkjbhmkfncmihhdhcebmchnfacdigflhhbekhfejblegakjjiejeenibemfmkfjbkkmlichlkbnhc", 840},
		{"fdacbaeacbdbaaacafdfbbdcefadgfcagdfcgbgeafbfbggdedfbdefdbgbefcgdababafgffedbefdecbaabdaafggceffbacgb", 2134},
		{"bahdcafcdadbdgagdddcidaaicggcfdbfeeeghiibbdhabdhffddhffcdccfdddhgiceciffhgdibfdacbidgagdadhdceibbbcc", 1571},
		{"dichcagakdajjhhdhegiifiiggjebejejciaabbifkcbdeigajhgfcfdgekfajbcdifikafkgjjjfefkdbeicgiccgkjheeiefje", 1042},
	} {
		t.Run(tt.s, func(t *testing.T) {
			assert.Eventually(t, func() bool {
				got := sherlockAndAnagrams(tt.s)
				return assert.Equal(t, tt.want, got)
			}, time.Second, time.Millisecond*100, "시간 초과")

		})
	}
}
