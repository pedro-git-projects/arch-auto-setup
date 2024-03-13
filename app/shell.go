package app

import (
	"os"
	"os/exec"
	"path/filepath"
)

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
	app.chsh()
	app.mvDir(filepath.Join(srcDir, "shell", "zsh", ".zshrc"), app.homeDir)
}
