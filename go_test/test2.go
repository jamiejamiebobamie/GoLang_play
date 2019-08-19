package main

import (
    "fmt"
    // "strconv"
)

// function "ping" takes in a channel that only recieves strings, and takes in a string
// the function sends the input string along the channel
func ping(pings chan<- string, msg string) {
    pings <- msg
}

// the function takes in a channel that only sends, and only sends strings
// the function also takes in a channel that recieves strings

// <-chan == a channel that only sends stuff
// chan<- == a channel that only recieves stuff


// the function takes the message from the sending channel
// and sends the message through the recieving channel
func pong(pings <-chan string, pongs chan<- string) {
    msg := <-pings
    pongs <- msg
}


func main() {
    pings := make(chan string, 2)
    pongs := make(chan string, 2)
    go ping(pings, "hi")
    go ping(pings, "goodbye")



    for i:=0; i < 10; i++ {
        // fmt.Println(i)
        fmt.Println(len(pings))
        fmt.Println(len(pongs))
    }


    go pong(pings, pongs)
    go pong(pings, pongs)


}
