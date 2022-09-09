package student

import (
	"fmt"
	"math/rand"
	"time"
)

var studentNames = [...]string{"Behling, John", "Bell, Dan", "Benson, Cole", "Brockbank, Stephen", "Carlock, Cody", "Choi, Brittany", "Choi, Ye Jin", "Clark, Madison", "Ferolino, Alexis Jane", "Guevara Alvarenga, Stefany", "Hansen, Nathan", "Hoather, Jeff", "Horlacher, Ethan", "Hunt, Brandon", "Jensen, David", "Jung, Euigun", "Kimball, Logan", "Ladle, Dallin", "Lee, SeungEun", "Lei, Sheng", "Leung, Miles", "Lo, Shaun", "Marks, Greg", "Marquis, Caden", "McConkie, Liberty", "McCord, Matthew", "McMillan, Zac", "Monson, Bailey", "Nelson, Sloan", "Peterson, James", "Piscione, Michael", "Prettyman, Samantha", "Ridd, Hayden", "Salvesen, Connor", "Shipley, David", "Stanley, Madison", "Sweeten, Daniela", "Tempest, Jordan", "Trammell,  Mark", "Andelin, Kyle", "Anderson, Taylor", "Baker, Nathan", "Barton, Zachary", "Berry, Solomon", "Bullock, Taylor", "Busco, Brian", "Davis, Michael", "Egbert, Seth", "Glazier, Tanner", "Goulding, Matt", "Jackson, Spencer", "Jensen, Emily", "Karras, Caden", "King, Spencer", "Lund, Thomas", "Luper, Abbie", "Maxfield, Chase", "Miller, Brinley", "Moody, Josh", "Moulton, McKay", "Nabrotzky, Keanna", "Nelson, Hunter", "Nielsen, Dustin", "Prock, Kamryn", "Sanderson, Ian", "Schow, Jackson", "Scorse, Brett", "Smith, Ali", "Sorensen, Stephen", "Souter, Kaden", "Spencer, Jessie", "Taylor, Chandler", "Washburn, Jackson", "Williams, Brennan"}


func GroupSize() int {

	var groupSize int
	fmt.Println("How many students do you want in each group?")

	fmt.Scanln(&groupSize)

	return groupSize


} 

func NumGroups(groupSize int) int {
	var studentNames = [...]string{"Behling, John", "Bell, Dan", "Benson, Cole", "Brockbank, Stephen", "Carlock, Cody", "Choi, Brittany", "Choi, Ye Jin", "Clark, Madison", "Ferolino, Alexis Jane", "Guevara Alvarenga, Stefany", "Hansen, Nathan", "Hoather, Jeff", "Horlacher, Ethan", "Hunt, Brandon", "Jensen, David", "Jung, Euigun", "Kimball, Logan", "Ladle, Dallin", "Lee, SeungEun", "Lei, Sheng", "Leung, Miles", "Lo, Shaun", "Marks, Greg", "Marquis, Caden", "McConkie, Liberty", "McCord, Matthew", "McMillan, Zac", "Monson, Bailey", "Nelson, Sloan", "Peterson, James", "Piscione, Michael", "Prettyman, Samantha", "Ridd, Hayden", "Salvesen, Connor", "Shipley, David", "Stanley, Madison", "Sweeten, Daniela", "Tempest, Jordan", "Trammell,  Mark", "Andelin, Kyle", "Anderson, Taylor", "Baker, Nathan", "Barton, Zachary", "Berry, Solomon", "Bullock, Taylor", "Busco, Brian", "Davis, Michael", "Egbert, Seth", "Glazier, Tanner", "Goulding, Matt", "Jackson, Spencer", "Jensen, Emily", "Karras, Caden", "King, Spencer", "Lund, Thomas", "Luper, Abbie", "Maxfield, Chase", "Miller, Brinley", "Moody, Josh", "Moulton, McKay", "Nabrotzky, Keanna", "Nelson, Hunter", "Nielsen, Dustin", "Prock, Kamryn", "Sanderson, Ian", "Schow, Jackson", "Scorse, Brett", "Smith, Ali", "Sorensen, Stephen", "Souter, Kaden", "Spencer, Jessie", "Taylor, Chandler", "Washburn, Jackson", "Williams, Brennan"}

	//Randomizing the array before assigning the students into groups
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(studentNames), func(i, j int) {
		studentNames[i], studentNames[j] = studentNames[j], studentNames[i]
	});

	var numGroups int = len(studentNames)/groupSize

	return numGroups
}

func Remainder(numGroups int, groupSize int) int {
	
	var Remainder int = len(studentNames) - (numGroups * groupSize)

	return Remainder
}

var counter int = 0
func MakeGroups (numGroups int, groupSize int, remainder int) [][]string {
	//Declare slice of slices

	var studentNames = [...]string{"Behling, John", "Bell, Dan", "Benson, Cole", "Brockbank, Stephen", "Carlock, Cody", "Choi, Brittany", "Choi, Ye Jin", "Clark, Madison", "Ferolino, Alexis Jane", "Guevara Alvarenga, Stefany", "Hansen, Nathan", "Hoather, Jeff", "Horlacher, Ethan", "Hunt, Brandon", "Jensen, David", "Jung, Euigun", "Kimball, Logan", "Ladle, Dallin", "Lee, SeungEun", "Lei, Sheng", "Leung, Miles", "Lo, Shaun", "Marks, Greg", "Marquis, Caden", "McConkie, Liberty", "McCord, Matthew", "McMillan, Zac", "Monson, Bailey", "Nelson, Sloan", "Peterson, James", "Piscione, Michael", "Prettyman, Samantha", "Ridd, Hayden", "Salvesen, Connor", "Shipley, David", "Stanley, Madison", "Sweeten, Daniela", "Tempest, Jordan", "Trammell,  Mark", "Andelin, Kyle", "Anderson, Taylor", "Baker, Nathan", "Barton, Zachary", "Berry, Solomon", "Bullock, Taylor", "Busco, Brian", "Davis, Michael", "Egbert, Seth", "Glazier, Tanner", "Goulding, Matt", "Jackson, Spencer", "Jensen, Emily", "Karras, Caden", "King, Spencer", "Lund, Thomas", "Luper, Abbie", "Maxfield, Chase", "Miller, Brinley", "Moody, Josh", "Moulton, McKay", "Nabrotzky, Keanna", "Nelson, Hunter", "Nielsen, Dustin", "Prock, Kamryn", "Sanderson, Ian", "Schow, Jackson", "Scorse, Brett", "Smith, Ali", "Sorensen, Stephen", "Souter, Kaden", "Spencer, Jessie", "Taylor, Chandler", "Washburn, Jackson", "Williams, Brennan"}

	//Randomizing the array before assigning the students into groups
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(studentNames), func(i, j int) {
		studentNames[i], studentNames[j] = studentNames[j], studentNames[i]
	});
	var Groups [][]string
	
	for i:=0;i<numGroups;i++{
		var Group []string
		if (i < remainder){
			for d:=0;d < groupSize + 1; d++{
				
				Group = append(Group, studentNames[counter])
				counter++
			}
		} else{

			for d:=0;d < groupSize; d++{
				Group = append(Group, studentNames[counter])
				counter++
			}
			
		}
		Groups = append(Groups, Group)
	}

	return Groups
}

func MyName() string {
	return "Solomon Berry"
}

