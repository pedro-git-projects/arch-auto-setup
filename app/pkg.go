package app

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func (app *App) installPackages() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <file>")
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("Error opening file: %s\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var packages []string

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		if len(parts) >= 2 && parts[0] != "" {
			packages = append(packages, parts[0])
		}
	}

	if len(packages) == 0 {
		fmt.Println("No packages found in the file")
		return
	}

	cmd := exec.Command("yay", append([]string{"-S"}, packages...)...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	fmt.Printf("Installing %d packages...\n", len(packages))
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Error running pacman -S: %s\n", err)
		return
	}

	fmt.Println("Packages installed successfully")
}
