package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	MONITORING = 3
	DELAY      = 5
)

func main() {
	showIntro()
	for {
		showMenu()

		switch readCommand() {
		case 1:
			startMonitoring()
		case 2:
			printLogs()
		case 0:
			fmt.Println("Close program")
			os.Exit(0)
		default:
			fmt.Println("Command not found!")
			os.Exit(-1)
		}

		fmt.Println("")
	}

	// fmt.Scanf("%d", &command)

	// fmt.Println("The address from 'command' variable is", &command)
	// fmt.Println("The command is", command)

	// if command == 1 {
	// 	fmt.Println("Monitoring...")
	// } else if command == 2 {
	// 	fmt.Println("Show logs...")
	// } else if command == 0 {
	// 	fmt.Println("Close program")
	// } else {
	// 	fmt.Println("Command not found!")
	// }
}

func showIntro() {
	name := "Vitor"
	version := 1.1
	fmt.Println("Hello, Mr.", name)
	fmt.Println("This program is at version", version)
}

func showMenu() {
	fmt.Println("1. Init monitoring")
	fmt.Println("2. Show logs")
	fmt.Println("0. Close")
}

func readCommand() int {
	var command int
	fmt.Scan(&command)
	return command
}

func startMonitoring() {
	fmt.Println("Monitoring...")
	sites := readSitesFromFile()
	for i := 0; i < MONITORING; i++ {
		for i, site := range sites {
			fmt.Println("Test site", i, ":", site)
			testSite(site)
		}
		time.Sleep(DELAY * time.Second)
		fmt.Println("")
	}
}

func testSite(site string) {
	res, err := http.Get(site)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if res.StatusCode == 200 {
		fmt.Println("Site", site, "was loaded with successfully")
		saveLog(site, true)
	} else {
		fmt.Println("Site", site, "have a problem. Status Code", res.StatusCode)
		saveLog(site, false)
	}
}

func readSitesFromFile() []string {
	var sites []string
	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	reader := bufio.NewReader(file)
	for {
		row, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error:", err)
		}

		sites = append(sites, strings.TrimSpace(row))
	}
	file.Close()
	return sites
}

func saveLog(site string, status bool) {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")
	file.Close()
}

func printLogs() {
	fmt.Println("Showing logs...")
	file, err := ioutil.ReadFile("log.txt") // Read all file
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(file))
}
