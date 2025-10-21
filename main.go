package main

import "fmt"

func main()  {
	scores := []int{50, 75, 66, 20, 32, 90}
	editedScored := []int{}
	editedScored = append(editedScored, scores[:3]...)
	editedScored = append(editedScored, 88)
	editedScored = append(editedScored, scores[4:]...)

	for x:= range len(editedScored) {
		fmt.Println(editedScored[x])
	}
	
}