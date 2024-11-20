package main

import "sort"

func bagOfTokensScore(tokens []int, power int) int {
	maxScore, score := 0, 0
	sort.Ints(tokens)

	l, r := 0, len(tokens)-1

	for l <= r {
		if power >= tokens[l] {
			// play face up
			power -= tokens[l]
			l++
			score++
			maxScore = max(maxScore, score)
		} else if score > 0 {
			// play face down
			power += tokens[r]
			score--
			r--
			maxScore = max(maxScore, score)
		} else {
			return maxScore
		}
	}

	return maxScore
}
