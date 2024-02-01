package main

func main() {
	// Code
	println("Bienvenido")
	var data = TwoNumberSum([]int{-21, 301, 12, 4, 65, 56, 210, 356, 9, -47}, 164)
	println(data)

}
func TwoNumberSum(array []int, target int) []int {
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
