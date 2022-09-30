package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	s := []string{"Zeno", "John", "Al", "Jenny"}
	fmt.Println("s before: ", s)
	sort.Strings(s)
	fmt.Println("s after: ", s)
	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	fmt.Println("s after reversed: ", s)

	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
	fmt.Println("n before: ", n)
	sort.Ints(n)
	fmt.Println("n after: ", n)
	sort.Sort(sort.Reverse(sort.IntSlice(n)))
	fmt.Println("n after reversed: ", n)

	studyGroup := People{"Zeno", "John", "Al", "Jenny"}
	fmt.Println("studyGroup before: ", studyGroup)
	sort.Strings(studyGroup)
	fmt.Println("studyGroup after: ", studyGroup)
	sort.Sort(sort.Reverse(sort.StringSlice(studyGroup)))
	fmt.Println("studyGroup after reversed: ", studyGroup)

	// build sorting for first name and last name
	studyGroup2 := People{"Zeno Blake", "John Alison", "Al Gore", "Jenny Myers"}
	// by default the type people will implement sort with the first name
	fmt.Println("studyGroup2 before:", studyGroup2)
	sort.Sort(studyGroup2)
	fmt.Println("studyGroup2 after sort:", studyGroup2)

	fmt.Println("studyGroup2 before:", studyGroup2)
	sort.Sort(OrderByLastName(studyGroup2))
	fmt.Println("studyGroup2 after sort last name:", studyGroup2)

}

type People []string

func (p People) Len() int {
	return len(p)
}
func (p People) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p People) Less(i, j int) bool {
	return p[i] < p[j]
}

type OrderByLastName People

func (p OrderByLastName) Len() int {
	return len(p)
}
func (p OrderByLastName) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p OrderByLastName) Less(i, j int) bool {
	ssI := strings.Split(p[i], " ")
	ssJ := strings.Split(p[j], " ")

	return ssI[1] < ssJ[1]
}
