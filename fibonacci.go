package fib

func newMemoizedFib() func(uint) uint {
	cache := make(map[uint]uint)
	// NOTE: fn must be defined beforehand (Go caveat)
	var fn func(uint) uint
	fn = func(n uint) uint {
		if n < 2 {
			return n
		}
		if _, ok := cache[n]; !ok {
			cache[n] = fn(n-1) + fn(n-2)
		}
		return cache[n]
	}
	return fn
}

// Fibonacci run base sequence
func Fibonacci(n uint) uint {
	if n < 2 {
		return n
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}
