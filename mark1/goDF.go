// Another fuck all implementation of GoDF
package main

import (
	"fmt"
	"time"
)

// DF STRUCT LOOKS SOMETHING LIKE THIS

type Col struct{
	data []string
	name string
	length int

}

type DF struct{
	columns []string
	data map[string] Col // Data will be stored as slices
}


// parse columns
func (df *DF) makeDataframe(cList []Col){
	/*
		This function will make a dataframe from cColumns.
	*/

	// initialize the map by allocating some space
	df.data = make(map[string] Col)
	for _, element := range cList{
		df.data[element.name] = element
		// Try to make changes to the original element
		
	}
}



func main(){
	// path := "/home/shiv/projects/golang/in.csv"

	// Two sample columns that will make up the cols
	var col1 = []string{"1", "2", "3", "4", "5"}
	var col2 = []string{"A", "B", "C", "D", "E"}

	// Push them into the Col struct
	C1 := Col{data:col1, name:"Sr. No"}
	C2 := Col{data:col2, name: "Letters"}




	start := time.Now()
	fmt.Println("----- TEST TO SEE HOW LONG IT TAKES TO LOAD A CSV INTO MEM -----------------")
	fmt.Println(time.Since(start))

	// An attempt at pushing data into the DF
	df := DF{columns:[]string{"A", "Col2"}}

	df.makeDataframe([]Col{C1, C2})


	fmt.Println(df)
}