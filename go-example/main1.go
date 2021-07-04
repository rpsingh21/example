package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

var logger *log.Logger

func main() {
	logFile, err := os.OpenFile("app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("Faild to open file: %v", err)
	}
	defer logFile.Close()
	logger = log.New(logFile, "", log.LstdFlags|log.Lshortfile)
	logger.Printf("Start application......")

	queue := make(chan string, 10)

	go func() {
		// for {
		// el := <-queue
		// logger.Printf("Log event %s", el)
		for v := range queue {
			logger.Printf("Log event %s", v)
		}
		time.Sleep(2 * time.Second)
		// }
	}()

	consoleReader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("-> ")
		text, _ := consoleReader.ReadString('\n')
		queue <- text
	}
}

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func main() {

// 	reader := bufio.NewReader(os.Stdin)

// 	for {
// 		fmt.Print("-> ")
// 		text, _ := reader.ReadString('\n')
// 		// convert CRLF to LF
// 		text = strings.Replace(text, "\n", "", -1)

// 		if strings.Compare("hi", text) == 0 {
// 			fmt.Println("hello, Yourself")
// 		}

// 	}

// }
