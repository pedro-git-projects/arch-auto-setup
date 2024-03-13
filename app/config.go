package app

import (
	"os"
	"os/exec"
	"path/filepath"
)

func (app *App) cloneConfig() error {
	err := app.cloneRepo(app.repos["config"])
	if err != nil {
		return err
	}
	return nil
}

func (app *App) setupWindowManager() {
	srcDir := filepath.Join(app.runningDir, "arch-dotfiles")
	cfgDir := filepath.Join(app.homeDir, "lol")

	app.mvDir(filepath.Join(srcDir, "window-manager", "dmenu"), cfgDir)
	app.mvDir(filepath.Join(srcDir, "window-manager", "dwm"), cfgDir)
	app.mvDir(filepath.Join(srcDir, "window-manager", "file-manager", "ranger"), cfgDir)
	app.mvDir(filepath.Join(srcDir, "window-manager", "slstatus"), cfgDir)
	app.mvDir(filepath.Join(srcDir, "window-manager", "st"), cfgDir)
}

func (app *App) chsh() {
	cmd := exec.Command("chsh", "-s", app.zsh, app.whoami)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	app.logger.Printf("changing shell for user %s to %s\n", app.whoami, app.zsh)
	err := cmd.Run()
	if err != nil {
		app.logger.Fatalf("failed to change shell with error: %s\n", err.Error())
	}
}

func (app *App) setupShell() {
	srcDir := filepath.Join(app.runningDir, "arch-dotfiles")
	cfgDir := filepath.Join(app.homeDir, "lol")
	app.mvDir(filepath.Join(srcDir, "shell", "zsh", ".zshrc"), cfgDir)
	app.chsh()
	//app.mvDir(filepath.Join(srcDir, "shell", "zsh", ".zshrc"), app.homeDir)
}
