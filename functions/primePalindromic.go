package functions

func IsPrime(number int) bool {
	numberOfFactor := GetFactorSize(number)
	if numberOfFactor != 2 {
		return false
	}
	return true
}

func IsPalindromic(number int) bool {
	if number <= 0 {
		return false
	}
	input_num := number
	var remainder int
	res := 0
	for number > 0 {
		remainder = number % 10
		res = (res * 10) + remainder
		number = number / 10
	}
	return input_num == res
}

func GetFactorSize(number int) int {
	var count int
	for counter := 1; counter <= number; counter++ {
		if number%counter == 0 {
			count++
		}
	}
	return count
}

func findIsPrimeIsPalindromic(results chan<- int) {
	num := 2
	for {
		if IsPrime(num) && IsPalindromic(num) {
			results <- num
		}
		num++
	}
}

func sumUp(results <-chan int, done chan<- struct{}, N int) int {
	sum := 0
	count := 0

	for num := range results {
		sum += num
		count++
		if count == N {
			close(done)
			return sum
		}
	}
	return sum
}

func FindAndSumUpPrimePalindromic(numb int) int {
	results := make(chan int)
	done := make(chan struct{})

	go findIsPrimeIsPalindromic(results)

	var sum int
	go func() {
		sum = sumUp(results, done, numb)
	}()

	<-done
	close(results)

	return sum
}
