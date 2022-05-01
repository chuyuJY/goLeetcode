package main

import "fmt"

// 1. 无法通过，超出时间限制，但是是最直观的方法
//func myPow(x float64, n int) float64 {
//	temp := float64(1)
//	if n > 0 {
//		for i := 0; i < n; i++ {
//			temp *= x
//		}
//		return temp
//	} else {
//		n = -n
//		x = 1 / x
//		for i := 0; i < n; i++ {
//			temp *= x
//		}
//		return temp
//	}
//}

// 2. 记忆递归
//var temp map[int]float64
//
//func myPow(x float64, n int) float64 {
//	if n < 0 {
//		temp = make(map[int]float64, 128)
//	} else {
//		temp = make(map[int]float64, 128)
//	}
//	return pow(x, n)
//}
//func pow(x float64, n int) float64 {
//	if n == 1 {
//		return x
//	}
//	if n == 0 {
//		return 1
//	} else if n > 0 {
//		if temp[n] != 0 {
//			return temp[n]
//		} else {
//			temp[n] = pow(x, n/2) * pow(x, n-n/2)
//			return temp[n]
//		}
//	} else {
//		n = -n
//		if temp[n] != 0 {
//			return temp[n]
//		} else {
//			temp[n] = pow(1/x, n/2) * pow(1/x, n-n/2)
//			return temp[n]
//		}
//	}
//}
// 3. 迭代，快速幂
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		n = -n
		x = 1 / x
	}
	// 到这里， n必>0，开始快速幂计算
	res := float64(1)
	for ; n >= 1; n >>= 1 {
		// 若这轮的二进制位为1
		if n&1 == 1 {
			// 乘相应位数的x迭代结果
			res *= x
			//n--	// 其实也不需要这个，因为向右>>移位之后，就无了，n是int
		}
		// 此时这个位应该是0
		// 0的话，其实相当于这轮的是1，所以不用写

		// 主要是为了下一轮用，对应 相应位的值  （x1,x2,x4,x8....的值）
		x *= x
	}
	return res
}

func main() {
	fmt.Println(myPow(8, 32))
}
