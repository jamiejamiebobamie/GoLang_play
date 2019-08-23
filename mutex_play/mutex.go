/*
https://medium.com/rungo/anatomy-of-channels-in-go-concurrency-in-go-1ec336086adb:

"From the above result, we can see that mutex helped us resolve racing conditions.
But the first rule is to avoid shared resources between goroutines."

*/

package main

import (
    "fmt"
    _"reflect"
    "sync"
    "time"
)

func doubleMessage(msg string) (doubled string) {
    doubled = msg + " " + msg
    return
}

func doubleMessageWithPointer(msg *string, wg *sync.WaitGroup, m *sync.Mutex) {
    m.Lock()
    dereference := *msg
    *msg = dereference+" "+dereference
    m.Unlock()
    wg.Done()
    fmt.Println(wg, time.Since(start))
}


// func init(){
//     start := time.Now()
// }

var start time.Time

func main() {

    start := time.Now()
    fmt.Println(start)

    msg1 := "hi!"
    msg2 := "hi!"

    var wg sync.WaitGroup
    var m sync.Mutex

    for i := 0; i < 20; i++{
        msg1 = doubleMessage(msg1)
        // go doubleMessageWithPointer(&msg)
        // doubleMessageWithPointer(&msg)
    }

    for i := 0; i < 20; i++{
        wg.Add(1)
        go doubleMessageWithPointer(&msg2, &wg, &m)
    }

    wg.Wait()
    fmt.Println(len(msg1))
    fmt.Println(len(msg2))
    fmt.Println("finished", time.Since(start))
}
