
package main

import (
    "fmt"
    "reflect"
    // "strconv"
    // "math/cmplx"
)

var (
    myVariable1 int = 1
    myVariable2 string = "Hello"
    myVariable3 uint64 = 12
    myBool bool = true
    x int = 2
    y rune = 45
    z byte = 0
)

func main () {
    // fmt.Println(myVariable1,myVariable2,myVariable3, myBool,x,y,z, reflect.TypeOf(z))
    fmt.Println("In Go there are 8 types of ints, and 4 types of floats.")
    fmt.Println("\nA Rune is a unit64. See: ")
    fmt.Println(reflect.TypeOf(y))

    fmt.Println("\nA byte is a unit8. See: ")
    fmt.Println(reflect.TypeOf(z))

    fmt.Println("\nThe ints are: uint8, uint16, uint32, uint64, and their signed counterparts.")
    fmt.Println("\nThe floats include float32, float64, complex64, and complex132.")
    fmt.Println("\nIt'd be cool to know what an imaginary number looks like...")

    imaginaryNumber := complex(2,8)
    fmt.Println(imaginaryNumber)
    fmt.Println(reflect.TypeOf(imaginaryNumber))

    var floatingPoint float32 = 1.0
    fmt.Println(floatingPoint)
    fmt.Println(reflect.TypeOf(floatingPoint))

    for i := 0; i < 10; i++ {
        fmt.Println(i)
            switch i % 2 {
                case 0: fmt.Println("hi")
                case 1: fmt.Println("bye")
                default: fmt.Println("idk")
            }
        if i == 2 {
            fmt.Println("AHHHH!!!!")
        } else if i == 3 {
            fmt.Println("GAHHH!!!!")
        } else {
            fmt.Println("I like your shirt.")
        }
    }

    var myArray [5]int;

    fmt.Println(myArray)

    myArray[1] = 20
    myArray2 := [2]int{30,30}
    fmt.Println(myArray, myArray2)

    mySlice := myArray[1:3]
    mySlice2 := make([]int, 4, 5)
    mySlice3 := append(mySlice, 23,23)

    fmt.Println(mySlice, mySlice2, mySlice3)
    copy(mySlice,mySlice2)
    fmt.Println("copy(x,y) copies y to x")
    fmt.Println(mySlice, mySlice2, mySlice3)

    fmt.Println(myArray)

    for i, value := range myArray {
        fmt.Println(reflect.TypeOf(i), value)
    }

    // how do I add multiple key, values to an exisiting map?

    var myMap map[string]int
    myMap = make(map[string]int)
    myMap["poop"] = 2
    myMap["pee"] = 1

    fmt.Println(myMap)
    delete (myMap,"poop")
    fmt.Println(myMap)

    myMap2:= map[string]int{
        "hi":2,
        "hello":3,
        "goodbye":10,
    }

    for key, value := range myMap2 {
        fmt.Println(key,value)
    }

    value, ok := myMap["pee"]

    if ok {
        fmt.Println(value)
    } else {
        fmt.Println("nope")
    }

    hi := myArray2[len(myArray2)-1]

    fmt.Println(hi)

}
