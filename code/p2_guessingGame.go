package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "strconv"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())
    gameOver := false
    var target int = rand.Intn(100) + 1
    var reader *bufio.Reader = bufio.NewReader(os.Stdin)
    fmt.Println("\n\nWelcome to guess my number!\n")
    fmt.Println("I chose a number between 1 and 100...")
    fmt.Println("Now it is up to you to find out what it is!\n")
    for !gameOver{
        guess := getInput(reader)
        if guess > target {
            fmt.Println("Lower!")
        }else if guess < target {
            fmt.Println("Higher!")
        }else{
            fmt.Println("Congrats! You Won!")
            fmt.Println("Target Was:", target)
            gameOver = true
        }
    }
}

func getInput(reader *bufio.Reader) int{
    fmt.Println("Guess a number between 1 and 100: ")
    numRead := false
    var num int
    var err error
    for !numRead {
        inp, _ := reader.ReadString('\n')
        num, err = strconv.Atoi(strings.TrimSpace(inp))
        if (err != nil || num > 100 || num < 1) {
            fmt.Println("Bad Input.  Need num between 1 and 100")
        }else{
            numRead = true
        }
    }
    return num
}