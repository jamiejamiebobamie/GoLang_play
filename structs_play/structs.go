// https://medium.com/rungo/structures-in-go-76377cc106a2
package main

import "fmt"

// uppercase structs, methods, and classes get exported with packages
// the Computer struct will get exported as will the Owner field of computer
// all other fields of Computer will be non-accessible to users of the package.
type Computer struct {
    Owner string
    color string
    works bool
}

type AnonymousFields struct {
    int
    bool
    string
}

// structs can contain metadata
// no idea how to access it ...
// they can also contain embedded structs (Computer)
type myRoom struct {
    wallpaper string `hey baby`
    Computer
}

// structs can contain anonymous functions ... which seem more like methods
// than interfaces are...

type HeyGirlType func() string

type heyGirl struct {
    HeyGirl HeyGirlType
}

func main(){

    // normal struct declaration
    heyhey := Computer{"old man", "blue", false}
    fmt.Printf("%+v\n", heyhey)

    // anonymous (nameless) struct
    animal := struct{
        body,name string
        wild bool
        Owner string
    }{
        body:"feathers",
        wild: true,
    }
    fmt.Printf("%+v\n",animal)

    // pointer to struct
    myComputer := &Computer{
        works:true,Owner:"Jamie",color:"Jet Black",
    }
    fmt.Printf("%+v\n",myComputer)

    // anonymous fields struct
    // uses the names of the datatypes as the fields' names
    // {int:7 bool:false string:Works?}
    // can only have one type of each value with this method
    newAnon := AnonymousFields{
        7,false,"Works?",
    }
    fmt.Printf("%+v\n", newAnon)

    // promoted fields of embedded structs
    bedroom := myRoom{
        wallpaper:"pink",
        Computer: Computer{Owner:"Jamie",works:true,color:"Jet Black"},
    }
    fmt.Printf("%+v\n",bedroom)

    // the embedded struct is named the same as it's containing field in myRoom
    // so myRoom has access to the fields of the Computer struct
    bedroom.Owner = "idk"
    fmt.Printf("%+v\n", bedroom)

     /*
     interfaces can be struct fields, but:

     "If we did not assign any value to salary field while creating an
     Employee struct as we did on line no. 25,
     Go will panic in a runtime error as we are trying to call a method
     on a nil value which is the default dynamic value of an interface.

     Hence, try to avoid having a struct field of an interface type."

     -https://medium.com/rungo/structures-in-go-76377cc106a2
     */

    // a ton of work:
    hello := heyGirl{
        HeyGirl: func() string {return "Hey girl."},
    }
    fmt.Println(hello.HeyGirl())


}
