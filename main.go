package main

import "fmt"

func main()  {
	scores := []int{50, 75, 20, 32, 66, 90}
	editedScores := []int{}
	editedScores = append(editedScores, scores[:3]...)
	editedScores = append(editedScores, 88)
	editedScores = append(editedScores, scores[4:]...)

	// Next Step
	newScores := []int{}
	for x := range len(scores) {
		if (scores[x] == 66) {
			newScores = append(newScores, scores[:x]...)
			newScores = append(newScores, 88)
			newScores = append(newScores, scores[x:]...)
		}
	}

	fmt.Println("Result Step 1")
	for x:= range len(editedScores) {
		fmt.Println(editedScores[x])
	}

	fmt.Println("Result Step 2")
	for x:= range len(newScores) {
		fmt.Println(newScores[x])
	}
}