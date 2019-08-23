package main

import (
    "fmt"
    "reflect"
)

// function type definition
type myFunc func(string) string

// function that takes a function as a parameter
func myFunc2(a string, f myFunc) int {
    return len(f(a))
}

// function that is being passed in as a myFunc type
func myFunc3(a string) string{
    return a
}

// anonymous function
var anon = func() string {
    return "hello"
}

// Immediately-invoked function
var hey = func (hey string) string {
    return hey
} ("hey")

/*
    "If a method receives a pointer receiver, you don’t need to using (*e)
    syntax to deference pointer or get the value of the pointer.
    You can use simple e which will be the address of the value that
    pointer points to but Go will understand that you are trying to perform
    an operation on the value itself and under the hood, it will make e to
    (*e)."
*/
type Greeting struct {
    salutation string
}
func (g *Greeting) sayGreeting() string {
    return g.salutation
}

// writing custom methods for non-struct types
type float float64
func (f float) double() float {
    return f+f
}

func main(){
    // why would i ever use this?
    fmt.Println(myFunc2("hi", myFunc3))

    fmt.Println(anon())

    fmt.Println(hey)

    // METHODS

    /*
        "A method is defined with different syntax than normal function.
        It required an additional parameter known as a receiver which is
        a type to which the method belongs.
        A method can access properties of the receiver it belongs to."

        "One major difference between function and method is many
        methods can have the same name while no two functions with the
        same name can be defined in a package."

        -https://medium.com/rungo/anatomy-of-methods-in-go-f552aaa8ac4a
    */


    /*
        "Also, you don’t need to call a method on the pointer if the
        method receives a pointer receiver. Instead, you can use value as it is,
        e and call a method on it. Go will pass the pointer of e under the hood
        if the method expects pointer receiver."
    */
    g := Greeting{"hello?"}
    fmt.Println(g.sayGreeting())

    a := float(3.4)
    b := float(5.9)
    c := a.double()
    fmt.Println(reflect.TypeOf(c),a,b,c)

}
