package main

import (
	"fmt"
	"nettverksprog/Oving1/AllPrimeNum"
	"sync"
	"sort"
	"time"
)
var wg sync.WaitGroup
var primes [] int
var mutex sync.Mutex

func RunFindPrimesUpTo(start int, end int) {
	var primeTemp = AllPrimeNum.AllPrimeNum(start,end)
	mutex.Lock()
 	primes = append(primes, primeTemp...)
 	mutex.Unlock()
	wg.Done()
}

func CreateThreads(min int, max int) int {
	var sliceSize = 1000
	var nSlices = (int) (max - min)/sliceSize
	fmt.Println("Number of threads: ", nSlices)

	wg.Add(nSlices)
	var startTimer = time.Now()
	var start = min
	var slutt = start + sliceSize -1
		for i:= 0; i < nSlices; i++  {
			go RunFindPrimesUpTo(start,slutt)
			start += sliceSize
			//second last loop adds one to get the end of the slice in order that the last number is checked if prime
			if i - 1 == nSlices{
				slutt +=sliceSize +1
			}else{
				slutt += sliceSize
			}
		}
	wg.Wait()
	fmt.Println("Elapsedtime: ", time.Since(startTimer))
	return len(primes)
}

func main()  {

	min := 0
	max := 1000000

	numOfPrimes :=CreateThreads(min, max)

	sort.Ints(primes)

	//Printing the answer
	fmt.Println("Min and max chosen:  ", min, " and ", max)
	if numOfPrimes <= 3000 {
		fmt.Println("The primenumbers is: ", primes)
	}
	fmt.Println("Number of primes: ", numOfPrimes)
}
