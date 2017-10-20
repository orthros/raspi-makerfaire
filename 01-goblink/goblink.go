package main

//import "os"
import "fmt"
import "time"
import "strconv"

//import "github.com/stianeikeland/go-rpio"

func main() {
    fmt.Println("Please enter time interval in seconds: ")
    text := ""
    fmt.Scanln(&text)
    seconds, parseErr := strconv.Atoi(text)

    for parseErr != nil {
        fmt.Println("Sorry we didnt recognize \"" + text + "\" as an integer. Please try again")
        fmt.Println("Please enter time interval in seconds: ")
        fmt.Scanln(text)
        seconds, parseErr = strconv.Atoi(text)
    }

    done := make(chan bool)

    go pinWorker(seconds, done)

    //Wait until we are Done...
    <-done

    fmt.Println("Exiting")
}

func pinWorker(seconds int, done chan<- bool) {
    fmt.Println("Starting pinWorker with interval " + strconv.Itoa(seconds))

    //err := rpio.Open()
    //if err != nil {
    //    fmt.Println(err)
    //    os.Exit(1)
    //}
    //defer rpio.Close()

    //pin := rpio.Pin(10)
    ticker := time.NewTicker(time.Second * time.Duration(seconds))
    for range ticker.C {
        //Do some work with the pins???
        fmt.Println("Hey there")
        //pin.Toggle();
    }
    done<- true
}
