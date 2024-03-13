package app

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func (app *App) symlinkScript(scriptPath, symlinkName string) {
	cmd := exec.Command("sudo", "ln", "-s", scriptPath, fmt.Sprintf("/usr/local/bin/%s", symlinkName))
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		app.logger.Fatalf("failed to symlink script %s\n", scriptPath)
	}
}

func (app *App) createScriptsFolderIfNotExists(path string) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err := os.Mkdir(path, 0755)
		if err != nil {
			app.logger.Fatalf("failed to create folder: %s", path)
		}
	}
}

func (app *App) makeExecutable(dir, filename string) {
	err := os.Chdir(dir)
	if err != nil {
		log.Fatalf("failed to change directory: %s\n", err.Error())
	}

	cmd := exec.Command("chmod", "+x", filename)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	log.Printf("making script executable on %s\n", dir)
	err = cmd.Run()
	if err != nil {
		log.Fatalf("failed to execute make script executable: %s\n", err.Error())
	}

	err = os.Chdir("..")
	if err != nil {
		log.Fatalf("failed to change back to original directory: %s\n", err.Error())
	}
}

func (app *App) setupScripts() {
	srcDir := filepath.Join(app.runningDir, "arch-dotfiles")

	target := filepath.Join(app.homeDir, ".scripts")
	app.createScriptsFolderIfNotExists(target)

	app.mvDir(filepath.Join(srcDir, "scripts", "kbd.sh"), target)
	app.makeExecutable(target, "kbd.sh")

	app.mvDir(filepath.Join(srcDir, "scripts", "restore_wallpaper.sh"), target)
	app.makeExecutable(target, "restore_wallpaper.sh")
	app.symlinkScript(filepath.Join(target, "restore_wallpaper.sh"), "restore_wallpaper")
}
