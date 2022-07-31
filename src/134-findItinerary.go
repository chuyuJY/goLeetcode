package main

import "sort"

type airport struct {
	target  string
	visited bool
}

type airports []*airport

func (ap airports) Len() int {
	return len(ap)
}

func (ap airports) Swap(i, j int) {
	ap[i], ap[j] = ap[j], ap[i]
}

func (ap airports) Less(i, j int) bool {
	return ap[i].target < ap[j].target
}

func findItinerary(tickets [][]string) []string {
	res := []string{"JFK"}
	// 1. 构造map
	airportsMap := map[string]airports{}
	for _, ticket := range tickets {
		if airportsMap[ticket[0]] == nil {
			airportsMap[ticket[0]] = airports{}
		}
		airportsMap[ticket[0]] = append(airportsMap[ticket[0]],
			&airport{target: ticket[1], visited: false})
	}
	// 2. map内部排序
	for src, _ := range airportsMap {
		sort.Sort(airportsMap[src])
	}
	// 3. 回溯
	var backTracking func() bool
	backTracking = func() bool {
		if len(tickets)+1 == len(res) {
			return true
		}
		for _, airport := range airportsMap[res[len(res)-1]] {
			if !airport.visited {
				res = append(res, airport.target)
				airport.visited = true
				if backTracking() {
					return true
				}
				airport.visited = false
				res = res[:len(res)-1]
			}
		}
		return false
	}
	backTracking()
	return res
}
