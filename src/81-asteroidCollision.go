package main

func asteroidCollision(asteroids []int) []int {
	stack := make([]int, 0, len(asteroids))
	res := []int{}
	for i := 0; i < len(asteroids); i++ {
		if asteroids[i] > 0 {
			stack = append(stack, asteroids[i])
		} else {
			for len(stack) != 0 {
				temp := stack[len(stack)-1]
				if temp > -asteroids[i] {
					break
				} else if temp < -asteroids[i] {
					stack = stack[:len(stack)-1]
				} else {
					stack = stack[:len(stack)-1]
					break
				}
			}
			if len(stack) == 0 {
				res = append(res, asteroids[i])
			}
		}
	}
	for i := 0; i < len(stack); i++ {
		res = append(res, stack[i])
	}
	return res
}
