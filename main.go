package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// stop/cont.

func main() {
	pid:= os.Getpid()
	fmt.Println("🔵 Process ID:",pid)
	sigs:= make(chan os.Signal,1)
	done:= make(chan bool, 1)

	// Notify channel on interrupt or terminate signals
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig:= <-sigs
		fmt.Println("We recievd signal:",sig)
		done<-true
	}()

	go func() {

		for {
			select{
			case <-done:
				fmt.Println("Stopping work, due to signal.")
				return
			default:
				fmt.Println("🟢 Working..")	
				time.Sleep(time.Second)
			}
		}

	}()
	for {
		time.Sleep(time.Second)
	}

	// 💡 O/P:
	// $ go run main.go
	
}