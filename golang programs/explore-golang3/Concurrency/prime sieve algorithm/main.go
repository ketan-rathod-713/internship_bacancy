package main

import (
	"fmt"
)

func worker(n int, i int, chPrime chan int){
	
}

func sieveOfEratosthenes(n int, ch chan int) {
	prime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		prime[i] = true
	}

	for p := 2; p*p <= n; p++ {
		if prime[p] == true {
			for i := p * p; i <= n; i += p {
				prime[i] = false
			}
		}
	}

	for p := 2; p <= n; p++ {
		if prime[p] == true {
			ch <- p
		}
	}
	close(ch)
}

func main() {
	n := 100 // Change this value as needed
	ch := make(chan int)
	go sieveOfEratosthenes(n, ch)

	var primes []int
	for prime := range ch {
		primes = append(primes, prime)
	}
	fmt.Println("Prime numbers up to", n, ":", primes)
}
