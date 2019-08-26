package main

import (
    "fmt"
)

/*
"Syntax to define a slice is pretty similar to that of an array
but without specifying the elements count."

- https://medium.com/rungo/the-anatomy-of-slices-in-go-6450e3bb2b94
*/

func main() {

    var myArray [3]int = [3]int{1,2,3}
    // slice is just a reference to an array.
    // and without a reference, the zero value is a 'nil'
    var mySlice []int

    fmt.Println(myArray,mySlice)

    // array[includes:doesNotInclude]
    // array[from:toTheElementBeforeThisOne]
    mySlice = myArray[0:3]
    fmt.Println(myArray,mySlice)

    // a slice is a reference to an array
    // changing an element of a slice
    // changes the referenced array
    mySlice[0] = 11
    fmt.Println(myArray)

    // you can grow a slice beyond the capacity of its underlying array
    // len() == show the current # of elements,
    // cap() == show how much the array could hold
    fmt.Println(len(mySlice), cap(mySlice))

    /*
    "append function takes a slice as the first argument,
    one/many elements as further arguments to append to the slice
    and returns a new slice of the same data type."
    */

    newSlice := append(mySlice, 11,12,14)
    fmt.Println(newSlice)
    fmt.Println(len(newSlice), cap(newSlice))

    // "append does not mutate original slice"
    // unless you want it to...
    mySlice = append(mySlice, newSlice...)
    fmt.Println("h",mySlice)

    // but it does mutate the underlying array...
    fmt.Println(myArray)

    // "slices are no easy business.
    // Use append only to self assign the new slice like s = append(s, ...)
    // which is more manageable."

    newerSlice := []int{1,2}
    // copy returns an int, the # of items copied.
    n1 := copy(newerSlice, mySlice)
    fmt.Println(n1,newerSlice)

    // slices as function parameters::
    // "slices are still passed by value to the function but since
    // they reference the same underneath the array,
    // it looks like that they are passed by reference."

    // comparing two slices:
    // "You will get error invalid operation: s1 == s2
    // (slice can only be compared to nil) which means that slices
    // can be only checked for the condition of nil or not.
    // If you really need to compare two slices, use for range loop
    // to match each element of the two slices or use DeepEquel
    // function of reflect package."

    // anonymous slice:
    // can't access underlying array that slice references
    anonSlice := []int{1,2,3}

    


}
