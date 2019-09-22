package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {

	// get server urls slice and time
	links, t := getInput()

	// create channel of type string that will be used to communicate among go routines
	c := make(chan string)

	// loop through each url & create each new go routine to hit respective url
	for _, link := range links {
		go checkLink(link, c)
	}

	// listen for some value from other routine & loop through each value
	for l := range c {
		// function literal running on separte routine to again hit url after 't' seconds
		go func(link string) {
			fmt.Println()
			time.Sleep(time.Second * time.Duration(t))
			checkLink(link, c)
		}(l)
	}
}

func getInput() ([]string, int64) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to server health check program.\nEnter the number of server urls that you want to perform health check for.")
	var n int
	_, err := fmt.Scanf("%d\n", &n)
	if err != nil {
		fmt.Println("Wrong Input")
		os.Exit(0)
	}
	var links []string
	for i := 0; i < n; i++ {
		fmt.Printf("Enter your %d server url\n", i+1)
		url, _ := reader.ReadString('\n')
		temp := strings.Split(url, "\n")
		links = append(links, temp[0])
	}
	fmt.Println("How frequently you want to hit servers? Please Enter time in number of seconds.")
	var t int64
	_, err = fmt.Scanf("%d\n", &t)
	if err != nil {
		fmt.Println("Wrong Input")
		os.Exit(0)
	}
	fmt.Println("\nServer Urls : ", links)
	fmt.Printf("Time : %d seconds\n\n", t)
	return links, t
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is down")
		c <- link
		return
	}
	c <- link
	fmt.Println(link, "is up")
}
