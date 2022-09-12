package student

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

//var studentNames = [...]string{"Behling, John", "Bell, Dan", "Benson, Cole", "Brockbank, Stephen", "Carlock, Cody", "Choi, Brittany", "Choi, Ye Jin", "Clark, Madison", "Ferolino, Alexis Jane", "Guevara Alvarenga, Stefany", "Hansen, Nathan", "Hoather, Jeff", "Horlacher, Ethan", "Hunt, Brandon", "Jensen, David", "Jung, Euigun", "Kimball, Logan", "Ladle, Dallin", "Lee, SeungEun", "Lei, Sheng", "Leung, Miles", "Lo, Shaun", "Marks, Greg", "Marquis, Caden", "McConkie, Liberty", "McCord, Matthew", "McMillan, Zac", "Monson, Bailey", "Nelson, Sloan", "Peterson, James", "Piscione, Michael", "Prettyman, Samantha", "Ridd, Hayden", "Salvesen, Connor", "Shipley, David", "Stanley, Madison", "Sweeten, Daniela", "Tempest, Jordan", "Trammell,  Mark", "Andelin, Kyle", "Anderson, Taylor", "Baker, Nathan", "Barton, Zachary", "Berry, Solomon", "Bullock, Taylor", "Busco, Brian", "Davis, Michael", "Egbert, Seth", "Glazier, Tanner", "Goulding, Matt", "Jackson, Spencer", "Jensen, Emily", "Karras, Caden", "King, Spencer", "Lund, Thomas", "Luper, Abbie", "Maxfield, Chase", "Miller, Brinley", "Moody, Josh", "Moulton, McKay", "Nabrotzky, Keanna", "Nelson, Hunter", "Nielsen, Dustin", "Prock, Kamryn", "Sanderson, Ian", "Schow, Jackson", "Scorse, Brett", "Smith, Ali", "Sorensen, Stephen", "Souter, Kaden", "Spencer, Jessie", "Taylor, Chandler", "Washburn, Jackson", "Williams, Brennan"}


func ReadNames() []string {
	var Names []string

	f, err := os.Open("C:/Users/saber/Documents/IS405 - Project Mgt/Week2/student/Names.txt")

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    scanner := bufio.NewScanner(f)

    for scanner.Scan() {

		Names = append(Names, scanner.Text())
        //fmt.Println(scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

	return Names
}
func GroupSize() int {

	var groupSize int
	fmt.Println("How many students do you want in each group?")

	fmt.Scanln(&groupSize)

	return groupSize


} 

func NumGroups(groupSize int, Names []string) int {
	//var studentNames = [...]string{"Behling, John", "Bell, Dan", "Benson, Cole", "Brockbank, Stephen", "Carlock, Cody", "Choi, Brittany", "Choi, Ye Jin", "Clark, Madison", "Ferolino, Alexis Jane", "Guevara Alvarenga, Stefany", "Hansen, Nathan", "Hoather, Jeff", "Horlacher, Ethan", "Hunt, Brandon", "Jensen, David", "Jung, Euigun", "Kimball, Logan", "Ladle, Dallin", "Lee, SeungEun", "Lei, Sheng", "Leung, Miles", "Lo, Shaun", "Marks, Greg", "Marquis, Caden", "McConkie, Liberty", "McCord, Matthew", "McMillan, Zac", "Monson, Bailey", "Nelson, Sloan", "Peterson, James", "Piscione, Michael", "Prettyman, Samantha", "Ridd, Hayden", "Salvesen, Connor", "Shipley, David", "Stanley, Madison", "Sweeten, Daniela", "Tempest, Jordan", "Trammell,  Mark", "Andelin, Kyle", "Anderson, Taylor", "Baker, Nathan", "Barton, Zachary", "Berry, Solomon", "Bullock, Taylor", "Busco, Brian", "Davis, Michael", "Egbert, Seth", "Glazier, Tanner", "Goulding, Matt", "Jackson, Spencer", "Jensen, Emily", "Karras, Caden", "King, Spencer", "Lund, Thomas", "Luper, Abbie", "Maxfield, Chase", "Miller, Brinley", "Moody, Josh", "Moulton, McKay", "Nabrotzky, Keanna", "Nelson, Hunter", "Nielsen, Dustin", "Prock, Kamryn", "Sanderson, Ian", "Schow, Jackson", "Scorse, Brett", "Smith, Ali", "Sorensen, Stephen", "Souter, Kaden", "Spencer, Jessie", "Taylor, Chandler", "Washburn, Jackson", "Williams, Brennan"}

	//Randomizing the array before assigning the students into groups
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(Names), func(i, j int) {
		Names[i], Names[j] = Names[j], Names[i]
	});

	var numGroups int = len(Names)/groupSize

	return numGroups
}

func Remainder(numGroups int, groupSize int, Names []string) int {
	
	var Remainder int = len(Names) - (numGroups * groupSize)

	return Remainder
}

var counter int = 0
func MakeGroups (numGroups int, groupSize int, remainder int, Names []string) [][]string {
	//Declare slice of slices

	//var studentNames = [...]string{"Behling, John", "Bell, Dan", "Benson, Cole", "Brockbank, Stephen", "Carlock, Cody", "Choi, Brittany", "Choi, Ye Jin", "Clark, Madison", "Ferolino, Alexis Jane", "Guevara Alvarenga, Stefany", "Hansen, Nathan", "Hoather, Jeff", "Horlacher, Ethan", "Hunt, Brandon", "Jensen, David", "Jung, Euigun", "Kimball, Logan", "Ladle, Dallin", "Lee, SeungEun", "Lei, Sheng", "Leung, Miles", "Lo, Shaun", "Marks, Greg", "Marquis, Caden", "McConkie, Liberty", "McCord, Matthew", "McMillan, Zac", "Monson, Bailey", "Nelson, Sloan", "Peterson, James", "Piscione, Michael", "Prettyman, Samantha", "Ridd, Hayden", "Salvesen, Connor", "Shipley, David", "Stanley, Madison", "Sweeten, Daniela", "Tempest, Jordan", "Trammell,  Mark", "Andelin, Kyle", "Anderson, Taylor", "Baker, Nathan", "Barton, Zachary", "Berry, Solomon", "Bullock, Taylor", "Busco, Brian", "Davis, Michael", "Egbert, Seth", "Glazier, Tanner", "Goulding, Matt", "Jackson, Spencer", "Jensen, Emily", "Karras, Caden", "King, Spencer", "Lund, Thomas", "Luper, Abbie", "Maxfield, Chase", "Miller, Brinley", "Moody, Josh", "Moulton, McKay", "Nabrotzky, Keanna", "Nelson, Hunter", "Nielsen, Dustin", "Prock, Kamryn", "Sanderson, Ian", "Schow, Jackson", "Scorse, Brett", "Smith, Ali", "Sorensen, Stephen", "Souter, Kaden", "Spencer, Jessie", "Taylor, Chandler", "Washburn, Jackson", "Williams, Brennan"}

	//Randomizing the array before assigning the students into groups
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(Names), func(i, j int) {
		Names[i], Names[j] = Names[j], Names[i]
	});
	var Groups [][]string
	
	for i:=0;i<numGroups;i++{
		var Group []string
		if (i < remainder){
			for d:=0;d < groupSize + 1; d++{
				
				Group = append(Group, Names[counter])
				counter++
			}
		} else{

			for d:=0;d < groupSize; d++{
				Group = append(Group, Names[counter])
				counter++
			}
			
		}
		Groups = append(Groups, Group)
	}

	return Groups
}
func OutputFile(Groups [][]string, Remainder int) {
    f, err := os.Create("C:/Users/saber/Documents/IS405 - Project Mgt/Week2/studentdata.txt")

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

	for i:=0;i < len(Groups); i++{
		//fmt.Println("Group " + strconv.Itoa(i))
		if(i < 9){
			f.WriteString("Group " + strconv.Itoa(i + 1) + ": ")	
		} else{
			f.WriteString("Group " + strconv.Itoa(i + 1) + ": ")
		}

		
		for d:=0;d < len(Groups[i]); d++{
			//fmt.Println(Groups[i][d])
			if (d < 1){
				if (i < 9){
					f.WriteString(" " + Groups[i][d])
					f.WriteString("\n")
				} else{
					f.WriteString(Groups[i][d])
					f.WriteString("\n")
				}
			} else{
				f.WriteString("          ")
				f.WriteString(Groups[i][d])
				f.WriteString("\n")
			}
			
			
		}
		f.WriteString("\n")
	}
    _, err2 := f.WriteString("\n")

    if err2 != nil {
        log.Fatal(err2)
    }

    fmt.Println("Done, you have " + strconv.Itoa(Remainder) + " groups with an extra member")

}

func MyName() string {
	return "Solomon Berry"
}

