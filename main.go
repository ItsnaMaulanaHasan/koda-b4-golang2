package main

import "fmt"

func main()  {
	scores := []int{50, 75, 66, 20, 32, 90}
	editedScores := []int{}
	editedScores = append(editedScores, scores[:3]...)
	editedScores = append(editedScores, 88)
	editedScores = append(editedScores, scores[4:]...)

	for x:= range len(editedScores) {
		fmt.Println(editedScores[x])
	}
	
}