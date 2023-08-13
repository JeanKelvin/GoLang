package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoring = 5
const delay = 5

func main() {

	showIntro()

	for {
		showMenu()
		command := readCommand()

		switch command {
		case 1:
			startMonitoring()
		case 2:
			printLogs()
		case 0:
			fmt.Println("Leaving the program")
			os.Exit(0)
		default:
			fmt.Println("I don't know this command")
			os.Exit(-1)
		}
	}
}

func showIntro() {
	var name string = "Jean"
	version := 1.0
	fmt.Println("Hello sr.", name)
	fmt.Println("Version", version)
	//fmt.Println("The type of the variable is", reflect.TypeOf(version))
}

func showMenu() {
	fmt.Println("1- Start processing")
	fmt.Println("2- View logs")
	fmt.Println("0- Exit the program")
}

func readCommand() int {
	var command int
	fmt.Scan(&command)
	//fmt.Println("The address of my variable is", &command)

	return command
}

func startMonitoring() {
	fmt.Println("Monitoring")

	sites := readArchiveSites()
	for i := 0; i < monitoring; i++ {
		for i, site := range sites {
			fmt.Println("Testing site", i, ":", site)
			testSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")
}

func testSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Error occurred:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "has been uploaded successfully!")
		registerLogs(site, true)
	} else {
		fmt.Println("Site:", site, "has a problem. Status code: ", resp.StatusCode)
		registerLogs(site, false)
	}
}

func readArchiveSites() []string {
	var sites []string
	file, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Error occurred:", err)
	}

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		sites = append(sites, line)

		if err == io.EOF {
			break
		}
	}

	file.Close()
	return sites
}

func registerLogs(site string, status bool) {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + "- online: " + strconv.FormatBool(status) + "\n")
	file.Close()
}

func printLogs() {
	file, err := os.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(file))
}
