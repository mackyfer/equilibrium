package main

func main() {}
func solution(A []int) int {
	equilibrium := -1

	if A == nil {
		panic("A is empty")
	} else if len(A) == 0 {
		panic("A cannot have 0 elements")
	} else {

		sumLeft := 0
		sumRight := 0
		for i := 0; i < len(A); i++ {
			sumRight += A[i]
		}

		for i := 0; i < len(A); i++ {
			tempRight := sumRight - A[i]
			if sumLeft == tempRight {
				equilibrium = i
				break
			} else {
				sumLeft += A[i]
				sumRight = tempRight
			}
		}
	}
	return equilibrium
}
