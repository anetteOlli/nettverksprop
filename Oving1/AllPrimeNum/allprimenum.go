package AllPrimeNum

import (
	"fmt"
	"nettverksprog/Oving1/IsPrime"
)

func AllPrimeNum (num1 int, num2 int) [] int {
	//Create a slice (not an array but kind of a pointer to an array?, because I want to add more elements to it)
	//Zero values to a slice is a nil
	var primes []int
	//Iterates through all the numbers between the two limits
	for i := num1; i <= num2; i++ {
		//If current number is a prime
		if IsPrime.IsPrime(i) {
			//The prime number is added in the end of the slice (Go creates a new array behind the slice in the background with the new element in)
			primes = append(primes, i)
		}
	}
	return primes
}

func main() {
	num1 := 100
	num2 := 200
	primes := AllPrimeNum(num1, num2)

	//HUSK Å SORTER PRIMTALLENE
	fmt.Println("The prime numbers between ", num1, " and ", num2, " is: ", primes)
}
