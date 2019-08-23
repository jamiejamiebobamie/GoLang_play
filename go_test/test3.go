package main


import (
    "fmt"
)

func pingPonger(channel chan<- string, msg string) {
        fmt.Println(msg)
        channel <- msg
}


func main() {
    channel := make(chan string,5)
    fmt.Println(len(channel))

    for i:= 0; i < 10; i++{
        go pingPonger(channel, string(i))
        // message := <- channel
        // fmt.Println(message)
    }
    fmt.Scanln()
}
