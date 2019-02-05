package IsPrime

import "math"

//Check if the input number is a prime, if yes returns true if not returns false
func IsPrime (num int) bool {

	for i := 2; i <= int(math.Floor(math.Sqrt(float64(num)))); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

/*
In a worst case let’s say we want to find if 100 is a prime numberand we loop 100 times.
This is improved by dividing by 2 to the halfway point giving us 50 loop iterations. However,
if we were to use the square root, this would leave us with only 10 loop iterations which is much better.

This is possible because we’re only looking for the smallest possible divisor when checking if a number is prime or not.
The largest possible divisor we can use is when the two divisors equal each other.
More information on this can be found in the trial division section of the Khan Academy.

Source: https://www.thepolyglotdeveloper.com/2016/12/determine-number-prime-using-golang/
*/

