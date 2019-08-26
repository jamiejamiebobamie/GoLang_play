package main


import (
    "fmt"
    "reflect"

)


/*

    extensible (?) functions take in an interface and return a struct

    interfaces are powerful because of type assertion which changes the
    type of a struct by changing the interface

    idk if this is a good metaphor but structs are like characters of Mario:
        mario, luigi, wario, bowser, yoshi

    and interfaces are the powerups:
        the feather, the mushroom, the fire flower

    interfaces are like clothes that characters wear that power them up:
        giving them superficial differences as well as new abilities

    but imagine if you couldn't be accepted to level 2 if you weren't
        wearing the racoon tail (the feather powerup)

    polymorphism is a struct wearing different interfaces to be accepted by
        different functions, as well as employing different methods?

    AKA:
    a function can accept many different types
    of structs if they are wearing
    the guise of a specific interface?

    let me try it...
        ... regrettable choice of structs...

*/

type Waste struct {
    difficultToPass bool
}

type poop struct {
    color string
    consistency string
    Waste
}

type pee struct {
    flow string
    color string
    Waste
}

type evacuation interface {
    remove()
}
// methods that implement the interface "evacuation"
func (p poop) remove() {
    fmt.Println("I pooped!")
}
func (p pee) remove() {
    fmt.Println("I peed!")
}

func goToTheBathroom (e evacuation) (w Waste) {
    // switch answer := fmt.Scanln("Was it tough? y/n"){
    // case 'y': w.difficultToPass = true
    // case 'n': w.difficultToPass = false
    // default:
    // e.(poop)
    w.difficultToPass = true
    // }
    return
}

func main() {

    peepee := pee{}
    peepee.remove()
    // peepee.(evacuate)
    fmt.Println(reflect.TypeOf(peepee))

    var pooper evacuation = poop{consistency:"stringy"}
    // pooper.consistency = "stringy"

    fmt.Println(reflect.TypeOf(pooper))

    _, ok := pooper.(evacuation)
    fmt.Println(ok,reflect.TypeOf(pooper), pooper)

    // i don't understand this...
    // pooper.difficultToPass = goToTheBathroom(pooper)
    fmt.Println(reflect.TypeOf(pooper),pooper)


}
