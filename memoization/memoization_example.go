package memoization

//Write a function fib() that takes an integer n and returns the nth Fibonacci number.
//Let's say our Fibonacci series is 0-indexed and starts with 0. So:
//fib(0)  # => 0
//fib(1)  # => 1
//fib(2)  # => 1
//fib(3)  # => 2
//fib(4)  # => 3

func fib(n int, fibMap map[int]int) int {
	if fibMap == nil {
		fibMap = map[int]int{}
	}
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	// check if we've calculated the term before
	val, ok := fibMap[n]
	if ok {
		return val
	}
	val = fib(n-1, fibMap) + fib(n-2, fibMap)
	fibMap[n] = val
	return val
}
