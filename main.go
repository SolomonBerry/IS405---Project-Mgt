//Solomon Berry, Alexis Transfiguracion, Keanna Nabrotzky, Miles Leung, SeungEun Lee - It also prints our name in the terminal when it is done running
package main

import (
	"project2/student"
)

func main() {
	var Names []string = student.ReadNames()
	var groupSize int = student.GroupSize()
	var numGroups int = student.NumGroups(groupSize, Names)
	var Remainder int = student.Remainder (numGroups, groupSize, Names)
	var Groups [][]string = student.MakeGroups(numGroups, groupSize, Remainder, Names)


	// for i:=0;i < len(Groups); i++{
	// 	fmt.Println("")
	// 	fmt.Println("Group " + strconv.Itoa(i))
	// 	for d:=0;d < len(Groups[i]); d++{
	// 		fmt.Println(Groups[i][d])
	// 	}
	// }

	student.OutputFile(Groups, Remainder)
	student.MyName()


	

}