package main

import (
    "fmt"
)

func Run() error{
    fmt.Println("Starting up app")
    return nil
}

func main(){
    fmt.Println("Go app")
    if err := Run(); err != nil {
        fmt.Println(err)
    }
}
