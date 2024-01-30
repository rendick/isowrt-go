package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

var (
	// color
	Red   = "\033[31m"
	Reset = "\033[0m"
)

var SelectFile int
var SelectUsb string
var Select string
var files []string

func main() {
	root := "/home/"

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			fmt.Println(err)
			return nil
		}

		if !info.IsDir() && filepath.Ext(path) == ".iso" {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	for num, file := range files {
		fmt.Printf(Red+"%d"+Reset+": %s\n", num, file)
	}

	fmt.Print("Select a file: ")
	fmt.Scan(&SelectFile)

	selectedIso := files[SelectFile]
	fmt.Printf("You selected: %s\n\n", selectedIso)

	usbCmd, err := exec.Command("sh", "-c", "lsblk -d | awk '/ 8:/'").Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(usbCmd))

	fmt.Print("Enter you USB device: ")
	fmt.Scan(&SelectUsb)
	fmt.Printf("Your flash drive: /dev/%s (y or n): ", SelectUsb)
	fmt.Scan(&Select)

	if Select == "n" || Select == "no" {
		fmt.Println("Exiting program.")
		os.Exit(0)
	} else if Select == "y" || Select == "yes" {
		fmt.Println("yes")
	} else if Select != "y" || Select != "yes" {
		fmt.Println("Invalid option!")
		return
	}

	time.Sleep(2 * time.Second)

	ddCmd := fmt.Sprintf("sudo dd if='%s' of='/dev/%s' bs=4M status=progress  && sync", selectedIso, SelectUsb)
	fmt.Println("Wait...")
	cmd := exec.Command("sh", "-c", ddCmd)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()

}
