package integer

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

// Divide two integers without using multiplication, division and mod operator.
// If it is overflow, return MAX_INT.
func Divide(dividend int, divisor int) int {
	if divisor == 0 {
		panic("divisor is zero")
	}
	if dividend == MinInt && divisor == -1 {
		return MaxInt
	}

	flag := true
	if dividend < 0 && divisor > 0 || dividend > 0 && divisor < 0 {
		flag = false
	}
	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}

	ans := 0
	for dividend >= divisor {
		x := divisor
		y := 1
		for dividend >= (x << 1) {
			x <<= 1
			y <<= 1
		}
		dividend -= x
		ans += y
	}
	if flag {
		return ans
	} else {
		return -ans
	}
}
