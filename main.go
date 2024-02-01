package main

func main() {
	// Code
	println("Bienvenido")
	var data = TwoNumberSumN([]int{-21, -1, 12, 4, 65, 56, 210, 11, 9, -47}, 10)
	if len(data) == 2 {
		println(data[0], data[1])
	} else {
		println("No hubo match")
	}

}
func TwoNumberSumNSquare(array []int, target int) []int {
	var pair []int
	for index, element := range array {

		for i := index + 1; i < len(array); i++ {
			if element+array[i] == target {
				return append(pair, element, array[i])
			}
		}
	}
	return []int{}
}

func TwoNumberSumN(array []int, target int) []int {
	seen := make(map[int]bool)

	for _, num := range array {
		complement := target - num

		if seen[complement] {
			return []int{num, complement}
		}

		seen[num] = true
	}

	return []int{}
}
