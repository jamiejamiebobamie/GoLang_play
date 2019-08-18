
package main

import (
    "fmt"
    "reflect"
    "strconv"
    // "os"
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

func myFunction(poop int) (returnItem string) {
    switch poop % 2 {
        case 0: returnItem =  "poop"
        case 1: returnItem =  "pee"
    }
    return
}

func myFunction2(poop int, pee int) (returnItem1 int, returnItem2 int) {
    returnItem1 = poop + poop
    returnItem2 = pee + poop
    return
}

func myVariadicFunction(arguments ...int) (sumOfArguments int) {

    if len(arguments) > 0 {

        for _, value := range arguments {
            sumOfArguments += value

        }

        return

    } else {

        return 0

    }

}
// from the Go website:
// myClosure takes in no arguments and returns an int
func myClosure() func() int {
    // declare and intialize a variable 'i' to a unint(0)
    i := int(0)
    // return a nameless function that returns an int
    return func() (ret int) {
        // the return variable is set to the value of 'i'
        ret = i
        // 'i' is incremented by 2 (for next call of the function??)
        i+=2
        // 'ret' is returned... I am so confused.
        return
    }
}

func main() {
    defer fmt.Println("This is a defer statement.")
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

    fmt.Println(myFunction(2))

    int1,int2 := myFunction2(1,2)
    fmt.Println(strconv.Itoa(int1),strconv.Itoa(int2))

    fmt.Println(strconv.Itoa(myVariadicFunction(1,2,3,4,5,6)))

    fmt.Println("A closure is a function within a function that returns a function? \nThat stores a value between calls, so you can like increment stuff with it?")
    evenNumberGenerator := myClosure()
    fmt.Println(evenNumberGenerator())
    fmt.Println(evenNumberGenerator())
    fmt.Println(evenNumberGenerator())
    fmt.Println("The benefit of a closure is you can keep secrets inside a function.")

    fmt.Println("A defer statement is a command used inside a function to defer a line of the program to run last.")

    // fmt.Println("A panic is... ")
    //
    // _, err := os.Create("/tmp/file")
    // fmt.Println(err)
    // if err == nil {
    //  panic(err)
    // }
    fmt.Println("A pointer is a reference to an object in memory. \nIf you want to change isntances of structs / interfaces I think you have to use pointers...(Yikes)")

}
