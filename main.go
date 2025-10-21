package main

import "fmt"

func main()  {
	// Step 1
	scoresStep1 := []int{50, 75, 66, 20, 32, 90}
	editedScores := []int{}
	editedScores = append(editedScores, scoresStep1[:3]...)
	editedScores = append(editedScores, 88)
	editedScores = append(editedScores, scoresStep1[4:]...)

	fmt.Println("Result Step 1")
	fmt.Println(editedScores)

	// Step 2
	scoresStep2 := []int{50, 66, 75, 20, 32, 90}
	newScores := []int{}
	for x := range len(scoresStep2) {
		if (scoresStep2[x] == 66) {
			newScores = append(newScores, scoresStep2[:x]...)
			newScores = append(newScores, 88)
			newScores = append(newScores, scoresStep2[x:]...)
		}
	}

	fmt.Println("Result Step 2")
	fmt.Println(newScores)
}