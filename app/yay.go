package app

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func (app *App) cloneYay() error {
	err := app.cloneRepo(app.repos["yay"])
	if err != nil {
		return err
	}
	return nil
}

func (app *App) executeMakepkg(dir string) {
	err := os.Chdir(dir)
	if err != nil {
		log.Fatalf("failed to change directory: %s\n", err.Error())
	}

	cmd := exec.Command("makepkg", "-si")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	log.Printf("executing 'makepkg -si' on %s\n", dir)
	err = cmd.Run()
	if err != nil {
		log.Fatalf("failed to execute 'makepkg -si': %s\n", err.Error())
	}

	err = os.Chdir("..")
	if err != nil {
		log.Fatalf("failed to change back to original directory: %s\n", err.Error())
	}
}
func (app *App) setupYay() {
	app.executeMakepkg(filepath.Join(app.runningDir, "yay"))
}
