package townjudge

/***
In a town, there are n people labeled from 1 to n. There is a rumor that one of these people is secretly the town judge.

If the town judge exists, then:

The town judge trusts nobody.
Everybody (except for the town judge) trusts the town judge.
There is exactly one person that satisfies properties 1 and 2.
You are given an array trust where trust[i] = [ai, bi] representing that the person labeled ai trusts the person labeled bi. If a trust relationship does not exist in trust array, then such a trust relationship does not exist.

Return the label of the town judge if the town judge exists and can be identified, or return -1 otherwise.



Example 1:

Input: n = 2, trust = [[1,2]]
Output: 2
Example 2:

Input: n = 3, trust = [[1,3],[2,3]]
Output: 3
Example 3:

Input: n = 3, trust = [[1,3],[2,3],[3,1]]
Output: -1

*/

func findJudge(n int, trust [][]int) int {

	if n == 1 {
		return 1
	}

	if len(trust) == 1 {
		return trust[0][1]
	}
	// create two arrays, index represents the person
	// value represents how much it trusts or is trusted by someone else
	trusts := make([]int, n+1)
	trustedBy := make([]int, n+1)

	for _, pair := range trust {
		trusts[pair[0]]++
		trustedBy[pair[1]]++
	}

	for i := 0; i <= n; i++ {
		// judge trusts nobody, but everybody else trusts the town judge
		if trustsNoBody(trusts, i) && everybodyElseTrustsThem(trustedBy, n, i) {
			return i
		}
	}

	return -1

}

func trustsNoBody(trusts []int, person int) bool {
	return trusts[person] == 0
}

func everybodyElseTrustsThem(trustedBy []int, population int, person int) bool {
	return trustedBy[person] == population-1
}
