package main

func main() {
	count := 0
	const N = 100

	// 1 1/2 1/4 1/8 => 2 ( S - S/2 = 1 )
	//
	// 1 to inf sum (1/2)^n => 2
	//
	// 1..10 => 10*11/2 (n n+1 /2)
	//
	// O(N)
	for i := N; i > 0; i /= 2 {
		for j = 0; j < i; j++ {
			count += 1
		}
	}
}
