package main
/*
 This will contain an implementation of map reduce and some other similar functionalities

*/

/*
##### SORRY BUT MORE THAN HALF OF THIS FILE IS GONNA BE DATA STRUCTURES ##########

RDDs:
what I think it should be:
	- Should be an immputable read-only create only data structure. 
	- Could be a 2-d array. Maybe I think this cus Idk what else to do
	- should just be that. No idea how to do cache or persist

#### STEPS ######
* create s stuct with 2-d rows. 

## AIM should be atleast to be able to implement one of the classic spark questions ###
* Word count seems the easiest

	*/

import (
	"fmt"
	"strings"
	"github.com/jinzhu/copier"
	"github.com/davecgh/go-spew/spew"
)


// Example mapper functions 
func upper(str interface{}) interface{}{
	return strings.ToUpper(str.(string))
}

func customSpacer(str interface{}) interface{}{
	return strings.Join(strings.Split(str.(string), ""), " ")
}

// The base RDD will accept any type of data structure
type BaseRDD struct{
	file interface{}
}

/*
* ALL TRANSFORM OPERATIONS ON THE RDD
* ------ right now I'm assuming that everything is a string ---
*/


// MAPPER
// first class ANON func for transformations
type transformer func(interface{}) interface{}

func (base BaseRDD) mapper(t transformer, index int) BaseRDD{
	var returnRDD = BaseRDD{}
	// Create a copy
	copier.Copy(&returnRDD.file, &base.file)
	for iter_index, val := range base.file.([][]string)[index]{
		returnRDD.file.([][]string)[index][iter_index] = t(val).(string)
	}
	return returnRDD
}

// SHOW
func (base BaseRDD) show(){
	rddStr := spew.Sdump(base)
	fmt.Println(rddStr+"\n")
}


// Limitation is that I have to depend on a single signature for passing data
func main() {
	var RDD1 = BaseRDD{file:[][]string{{"aBCdEF", "consume"}, {"vaporwave"}}}
	RDD1.show()
	RDD2 := RDD1.mapper(upper, 0)
	fmt.Println("After transformation 1:\n")
	RDD2.show()
	fmt.Println("After tranformation 2:\n")
	RDD3 := RDD2.mapper(customSpacer, 1)
	RDD3.show()
}
