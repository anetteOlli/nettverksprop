package main

import (
	"fmt"
	"nettverksprog/Oving1/AllPrimeNum"
	"sync"
	"sort"
	"time"
	"math"
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

//simplified abc-formula for calculating end of slice
func calculateEnd(startByCost int, partCost int) int{
	var c = float64( 1 + 4*(startByCost-1)*(startByCost) + 8* partCost)
	return int((-1 + math.Sqrt(c))/2)
}

func CreateThreads(min int, max int, numThreads int) int {
	primes = nil
	//var nSlices =max - min
	//var sliceSize = 1

	//removes negative numbers
	if min < 0{
		min = 0
	}
	//quick check if there's no primes...
	if max <= 1{
		return 0
	}

	//user can set threading higher than the number of numbers needed to be checked
	var threading = max - min
	if numThreads < threading{
		threading = numThreads
	}

	//calculates totalCost as a the sum of the harmonic-sequence
	//known issue with this method when the number of threads is equal to the numbers checked, the last couple of
	//ends gets negative. This is fixed during the last loop
	var totalCost = (int) (max *(max +1))/2 - ((min-1)*min)/2
	var partCost = totalCost/threading

	var points [] int = nil

	fmt.Println("number of threads:" , threading)

	wg.Add(threading)
		var startTimer = time.Now()
		var startByCost = min
		var endByCost = calculateEnd(startByCost, partCost)

		for i:=0; i < threading; i++{
			fmt.Println("startpoint thread ", i, ":", startByCost, ", endpoint: ", endByCost, "i -1 == threading", i ==threading -2 )
			points = append(points, startByCost, endByCost )


			go RunFindPrimesUpTo(startByCost, endByCost)

			startByCost = endByCost +1


			if i  == threading - 2{
				endByCost = max
			}else{
				endByCost = calculateEnd(startByCost, partCost)
			}
		}

	wg.Wait()
	fmt.Println("Elapsedtime: ", time.Since(startTimer))
	fmt.Println("end & startpoints: ", points)
	return len(primes)

}

func main()  {
	min := 0
	max := 1000000
	numThreads := 6

	numOfPrimes := CreateThreads(min, max, numThreads)

	sort.Ints(primes)

	//Printing the answer
	fmt.Println("Min and max chosen:  ", min, " and ", max)
	if numOfPrimes <= 3000 {
		fmt.Println("The primenumbers is: ", primes)
	}
	fmt.Println("Number of primes: ", numOfPrimes)
}
