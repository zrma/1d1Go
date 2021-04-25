package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnagram(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/anagram/problem")

	for _, tc := range []struct {
		given string
		want  int32
	}{
		{"", -1},
		{"aaabbb", 3},
		{"ab", 1},
		{"abc", -1},
		{"mnop", 2},
		{"xyyx", 0},
		{"xaxbbbxx", 1},
		{"hhpddlnnsjfoyxpciioigvjqzfbpllssuj", 10},
		{"xulkowreuowzxgnhmiqekxhzistdocbnyozmnqthhpievvlj", 13},
		{"dnqaurlplofnrtmh", 5},
		{"aujteqimwfkjoqodgqaxbrkrwykpmuimqtgulojjwtukjiqrasqejbvfbixnchzsahpnyayutsgecwvcqngzoehrmeeqlgknnb", 26},
		{"lbafwuoawkxydlfcbjjtxpzpchzrvbtievqbpedlqbktorypcjkzzkodrpvosqzxmpad", 15},
		{"drngbjuuhmwqwxrinxccsqxkpwygwcdbtriwaesjsobrntzaqbe", -1},
		{"ubulzt", 3},
		{"vxxzsqjqsnibgydzlyynqcrayvwjurfsqfrivayopgrxewwruvemzy", 13},
		{"xtnipeqhxvafqaggqoanvwkmthtfirwhmjrbphlmeluvoa", 13},
		{"gqdvlchavotcykafyjzbbgmnlajiqlnwctrnvznspiwquxxsiwuldizqkkaawpyyisnftdzklwagv", -1},
	} {
		t.Run(tc.given, func(t *testing.T) {
			got := anagram(tc.given)
			assert.Equal(t, tc.want, got)
		})
	}
}
