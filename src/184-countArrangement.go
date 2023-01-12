package main

func countArrangement(n int) int {
	res := 0
	visited := make([]bool, n+1)

	var backTrack func(index int)
	backTrack = func(index int) {
		if index == n+1 {
			res++
			return
		}
		for i := 1; i <= n; i++ {
			if !visited[i] && (index%i == 0 || i%index == 0) {
				visited[i] = true
				backTrack(index + 1)
				visited[i] = false
			}
		}
	}
	backTrack(1)
	return res
}
