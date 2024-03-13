package app

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func (app *App) setupWindowManager() {
	srcDir := filepath.Join(app.runningDir, "arch-dotfiles")
	cfgDir := filepath.Join(app.homeDir, ".config")

	app.mvDir(filepath.Join(srcDir, "window-manager", "dmenu"), cfgDir)
	app.executeMakeCleanInstall(filepath.Join(cfgDir, "dmenu"))

	app.mvDir(filepath.Join(srcDir, "window-manager", "dwm"), cfgDir)
	app.executeMakeCleanInstall(filepath.Join(cfgDir, "dwm"))

	app.mvDir(filepath.Join(srcDir, "window-manager", "slstatus"), cfgDir)
	app.executeMakeCleanInstall(filepath.Join(cfgDir, "slstatus"))

	app.mvDir(filepath.Join(srcDir, "window-manager", "st"), cfgDir)
	app.executeMakeCleanInstall(filepath.Join(cfgDir, "st"))

	app.mvDir(filepath.Join(srcDir, "window-manager", "file-manager", "ranger"), cfgDir)
	app.mvDir(filepath.Join(srcDir, "terminal-emulators", "alacritty"), cfgDir)
}

func (app *App) executeMakeCleanInstall(dir string) {
	err := os.Chdir(dir)
	if err != nil {
		log.Fatalf("failed to change directory: %s\n", err.Error())
	}

	cmd := exec.Command("sudo", "make", "clean", "install")
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	log.Printf("executing 'sudo make clean install' on %s\n", dir)
	err = cmd.Run()
	if err != nil {
		log.Fatalf("failed to execute 'sudo make clean install': %s\n", err.Error())
	}

	err = os.Chdir("..")
	if err != nil {
		log.Fatalf("failed to change back to original directory: %s\n", err.Error())
	}
}
