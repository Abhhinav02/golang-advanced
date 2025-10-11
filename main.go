package main

import (
	"fmt"
	"sort"
)

// SORTING

type Person struct{
	Name string
	Age int
}

type By func(p1,p2 *Person)bool

type PersonSorter struct{
	people []Person
	by func(p1,p2 *Person) bool
}

func (s *PersonSorter)Len()int{
return len(s.people)
}

func (s *PersonSorter) Less(i,j int) bool{
	return s.by(&s.people[i],&s.people[j])
}

func (s *PersonSorter) Swap (i,j int){
	s.people[i], s.people[j] = s.people[j],s.people[i]
}

func (by By) Sort(people []Person){
ps:= &PersonSorter{
	people: people,
	by: by,
}
sort.Sort(ps)
}

func main() {

	people:=[]Person{
		{"Skyy",30},
		{"Bob",25},
		{"Anna",35},
		{"Max",32},
	}

	age:= func (p1,p2 *Person)bool{
		return p1.Age<p2.Age
	}

	By(age).Sort(people)

	fmt.Println("Sorted by age:",people)

}

/* ------------------------------------------------------------------
*/

// type ByAge []Person
// type ByName []Person

// func (ba ByAge) Len()int{
// 	return len(ba)
// }

// func (ba ByName) LenByName()int{
// 	return len(ba)
// }


// func (ba ByAge) Less(i,j int)bool{
// 	return ba[i].Age<ba[j].Age
// }

// func (ba ByName) LessName(i,j int)bool{
// 	return ba[i].Name<ba[j].Name
// }
// func (ba ByAge) Swap(i,j int){
// 	ba[i], ba[j] = ba[j] , ba[i]
// }

// func (ba ByName) SwapName(i,j int){
// 	ba[i], ba[j] = ba[j] , ba[i]
// }
