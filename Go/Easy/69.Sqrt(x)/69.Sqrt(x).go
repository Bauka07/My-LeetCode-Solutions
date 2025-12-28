package main


func mySqrt(x int) int {
	if x < 2 {
			return x
	}
	l, r := 1, x
	for l <= r {
		m := l + (r-l)/2
		if m*m == x {
			return m
		}	else if m*m < x {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return r
}
