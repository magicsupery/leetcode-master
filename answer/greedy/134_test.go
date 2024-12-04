package greedy

func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	delta := make([]int, n, n)

	total := 0
	for i := 0; i < n; i++ {
		delta[i] = gas[i] - cost[i]
		total += delta[i]
	}
	if total < 0 {
		return -1
	}

	// find maxSubArray

	startIndex := 0
	deltaNum := 0
	for i := 0; i < n; i++ {
		deltaNum += delta[i]
		if deltaNum < 0 {
			deltaNum = 0
			startIndex = i + 1
		}
	}

	return startIndex
}
