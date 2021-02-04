package main

import "fmt"

func maxDivide(a int, b int) int {
	if a%b != 0 {
		return a
	}
	return maxDivide(a / b, b)
}

func isUgly(no int) int {
	a := maxDivide(no, 2)
	b := maxDivide(a, 3)
	c := maxDivide(b, 5)

	if c == 1 {
		return 1
	}
	return 0
}

func getNthUglyNo(n int, i int, count int) int {
	if n < count {
		return i-1
	}
	if isUgly(i) == 1 {
		return getNthUglyNo(n, i+1, count+1)
	} else {
		return getNthUglyNo(n, i+1, count)
	}
}

func main() {
	i := 10
	fmt.Printf("[INPUT] %d\n", i)
	o := getNthUglyNo(i, 1, 1)
	fmt.Printf("[OUTPUT] %d\n", o)
}
