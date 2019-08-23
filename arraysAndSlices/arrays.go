package main

import (
    "fmt"
    "reflect"
)

/*
    "mixing different values belonging to different data types in an array
    is not allowed... In Go, an array has a fixed length.
    ...a composite data type
    ...composed of primitive data types

    "
    - https://medium.com/rungo/the-anatomy-of-arrays-in-go-24429e4491b7
*/

var myArray [10]int

var myArray2 [3]string = [3]string{"laughter", "joy", "love"}

func doesNotChange(array [3]int) {
    for index, value := range array {
        array[index] = value*3
    }
    fmt.Println(array)
}

// good for intializing global variables before main() runs.
func init(){
    for i := 0; i < 10; i++ {
        myArray[i] = i+1
    }
}

func yellowJello() {
    myArray := [2]string{"yellow", "jello",}
    fmt.Println(myArray)
}

func main() {
    fmt.Println(myArray)
    fmt.Println(myArray2)

    // the simplest way to create an array:
    // := can only be used inside function bodies.
    myArray3 := [...]int{1,2,3}
    fmt.Println(myArray3)
    yellowJello()

    /*
    "For an array to be the equal or the same as the second array,
    both array should be of the same type, must have the same elements in them
    and all elements must be in the same order."
    */
    fmt.Println(len(myArray3), reflect.TypeOf(myArray3),
    reflect.TypeOf(myArray2),
    reflect.TypeOf(myArray3) == reflect.TypeOf(myArray2))

    // iterating over an array
    for index, value := range myArray2 {
        fmt.Println(index, value)
    }
    for _, value := range myArray2 {
        fmt.Println(value)
    }

    /*
    multi-dimensional arrays:
    "an array is a collection of same data types and array is a type in itself,
    a multi-dimensional array must have arrays
    that belong to the same data type."
    */

    // declaring a multidimensional array
    var threeNumbers [2][3]int = [2][3]int{
        [3]int{1,2,3},
        [3]int{4,5,6},
    }
    fmt.Println(threeNumbers)

    // shorter syntax:
    // the interior elements of a multi-dimensional array must be the same
    // here they are [2]string
    multi := [...][2]string{ {"hi","hey"}, {"hi","hey"} }
    fmt.Println(multi)

    // iterating over a multi-dimensional array:
    for _, interiorArray := range multi {
        for _, item := range interiorArray {
            fmt.Println(item)
        }
    }

    /*
    "When you pass an array to a function,
    they are passed by value like int or string data type.
    The function receives only a copy of it.
    Hence, when you make changes to an array inside a function,
    it won’t be reflected in the original array."
    */

    // Example:
    fmt.Println(myArray3)
    doesNotChange(myArray3)
    fmt.Println(myArray3)

    // slices are pass by reference...

}
