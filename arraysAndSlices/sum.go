package arraysAndSlices

func Sum(arrayOfNumbers []int) int {
	var sum int
	for _, number := range arrayOfNumbers {
		sum += number
	}
	return sum
}

func SumAll(slicesToSum ...[]int) []int {
	var slicesOfSum []int
	for _, numbers := range slicesToSum {
		slicesOfSum = append(slicesOfSum, Sum(numbers))
	}
	return slicesOfSum
}

func SumAllTails(slicesToSum ...[]int) []int {
	var slicesOfSum []int
	for _, tail := range slicesToSum {
		if len(tail) == 0 {
			slicesOfSum = append(slicesOfSum, 0)
		} else {
			tailOfSlices := tail[1:]
			slicesOfSum = append(slicesOfSum, Sum(tailOfSlices))
		}
	}
	return slicesOfSum
}
