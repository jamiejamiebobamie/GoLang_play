package main

import (
    "fmt"
    "reflect"
)

type Person struct {
    job string
    age int
    Thing
}

type Chair struct {
    purpose string
    material string
    Thing
}

type Thing struct {
    living bool
}

type Being interface {
    existing() string
    breathing() string
}

type Human interface {
    pooping() string
}

func (t Thing) breathing() string {
    if t.living {
        return "I am alive and breathing."
    } else {
        return "I am not alive. I don't breathe."
    }
}

func (t Thing) existing() string {
    if t.living {
        return "I exist and I know it!"
    } else {
        return "I exist and I know it?"
    }
}

func (p Person) pooping() string {
    return "I'm on the toilet!"
}

func main() {

/*  https://medium.com/rungo/interfaces-in-go-ab1601159b3a:
    Now I guess you can understand why the type and value of the interface are dynamic.
    Like what we saw in slices lesson that a slice hold reference to an array,
    we can say that interface is also works in a similar way by dynamically
    holding a reference to the underlying type. */


 Anne := Person{job:"bus driver", age: 5}
 Anne.living = true

 fmt.Println(Anne.breathing())
 fmt.Println(Anne.existing())

 var recliner Chair = Chair{material:"poop"}
 recliner.purpose = "sitting in the den"
 recliner.living = false

 fmt.Println(recliner.breathing())
 fmt.Println(recliner.existing())

 fmt.Println(Anne, recliner)
 fmt.Printf("%+v\n", Anne)

 var inter Being = Thing{true}

 /*
 Type assertion is not only used to check if an interface has a concrete
 value of some given type but to convert a given variable of an interface type
 to a different interface type as well.
 */
 _, ok1 := inter.(Thing) // type assertion variableName.(Type)

 if ok1 {
     fmt.Println(inter.breathing())
     fmt.Println(inter.existing())
 }

 _, ok2 := inter.(Human) // type assertion variableName.(Type)

 fmt.Println(ok2, reflect.TypeOf(inter))

}
