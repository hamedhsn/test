package main

import (
    "fmt"
    "log"

    "github.com/user/stringutil"
)

func init()  {
    log.Println("enter to init")
    fmt.Println("Init!!") //print line
    log.Println("leave from init")//print
}

func main() {
    fmt.Printf(stringutil.Reverse("!oG ,olleH"))
    //fmt.Printf(stringutil.Multiplier(2,10))
}