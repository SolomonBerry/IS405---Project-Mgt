package main

import (
	"fmt"
	"project2/student"
	"strconv"
)

func main() {
	var groupSize int = student.GroupSize()
	var numGroups int = student.NumGroups(groupSize)
	var Remainder int = student.Remainder (numGroups, groupSize)
	var Groups [][]string = student.MakeGroups(numGroups, groupSize, Remainder)

	fmt.Println(len(Groups))

	for i:=0;i < len(Groups); i++{
		fmt.Println("")
		fmt.Println("Group " + strconv.Itoa(i))
		for d:=0;d < len(Groups[i]); d++{
			fmt.Println(Groups[i][d])
		}
	}


	

}