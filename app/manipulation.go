package app

import (
	"os"
	"os/exec"
)

func (app *App) cloneRepo(repo string) error {
	cmd := exec.Command("git", "clone", repo)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	app.logger.Printf("cloning repo: %s\n", repo)
	return cmd.Run()
}

func (app *App) mvDir(source, target string) {
	app.logger.Printf("moving %s to %s\n", source, target)
	cmd := exec.Command("mv", source, target)
	err := cmd.Run()
	if err != nil {
		app.logger.Fatalf("failed to move directories with error %v\n", err)
	}
}
