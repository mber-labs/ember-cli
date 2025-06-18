package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		printHelp()
		return
	}

	cmd := os.Args[1]
	nodeURL, err := readNodeURL()
	if err != nil {
		fmt.Println("❌", err)
		return
	}

	switch cmd {
	case "help":
		printHelp()
	case "register":
		callEndpoint(nodeURL + "/register")
	case "getaddress":
		callAndPrint(nodeURL + "/eth-address")
	default:
		fmt.Println("❌ Unknown command:", cmd)
		printHelp()
	}
}

func readNodeURL() (string, error) {
	data, err := os.ReadFile(".ember-node")
	if err != nil {
		return "", fmt.Errorf("could not read .ember-node file")
	}
	return strings.TrimSpace(string(data)), nil
}

func callEndpoint(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("❌ Failed to call %s: %v\n", url, err)
		return
	}
	defer resp.Body.Close()
	fmt.Printf("📡 Register Operator called!")
}

func callAndPrint(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("❌ Failed to call %s: %v\n", url, err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("📬 %s\n", string(body))
}

func printHelp() {
	fmt.Println("📘 ember-cli commands:")
	fmt.Println("   help         Show available commands")
	fmt.Println("   register     Register the node as an operator")
	fmt.Println("   getaddress   Get the Ethereum address of the node")
}
