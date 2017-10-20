package main

import "os"
import "os/signal"
import "syscall"
import "fmt"
import "time"
import "strconv"

import "github.com/stianeikeland/go-rpio"

func main() {
    fmt.Println("Please enter time interval in seconds: ")
    text := ""
    fmt.Scanln(&text)
    seconds, parseErr := strconv.Atoi(text)

    for parseErr != nil {
        fmt.Println("Sorry we didnt recognize \"" + text + "\" as an integer. Please try again")
        fmt.Println("Please enter time interval in seconds: ")
        fmt.Scanln(&text)
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

    err := rpio.Open()
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    defer rpio.Close()

    //Physical pin 19
    pin := rpio.Pin(10)
    pin.Output()

    sigc := make(chan os.Signal, 1)
    signal.Notify(sigc,
                  syscall.SIGINT,
                  syscall.SIGTERM,
                  syscall.SIGQUIT)
    go func() {
            <-sigc
            pin.Low()
            rpio.Close()
            done<- true
    }()

    ticker := time.NewTicker(time.Second * time.Duration(seconds))
    for range ticker.C {
        //fmt.Println("Tick")
        pin.Toggle();
    }
    done<- true
}
