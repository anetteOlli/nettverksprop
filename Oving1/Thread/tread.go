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
	var primeTemp [] int = AllPrimeNum.AllPrimeNum(start,end)
	mutex.Lock()
 	primes = append(primes, primeTemp...)
 	mutex.Unlock()
	wg.Done()
}

func CreateThreads(min int, max int){
	var sliceSize = 1000
	var nSlices int = (int) (max - min)/sliceSize
	fmt.Println(nSlices)

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
	fmt.Println("elapsedtime: ", time.Since(startTimer))

}

func main()  {


	//wg.Add(3)
	//	go RunFindPrimesUpTo(1,1000)
	//	go RunFindPrimesUpTo(1001,2000)
	//	go RunFindPrimesUpTo(2001,3000)
	//wg.Wait()

	CreateThreads(0, 10000000)

	sort.Ints(primes)

	fmt.Println("The prime numbers between ", 1, " and ", 1000, " is: ", ", len(primes): ", len(primes))
}
