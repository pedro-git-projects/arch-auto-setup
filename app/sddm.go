package app

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func (app *App) chownDisplayManager() {
	cmd := exec.Command("sudo", "chown", "-R", fmt.Sprintf("%s:%s", app.whoami, app.whoami), app.sddmDir)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	app.logger.Printf("taking ownership of %s\n", app.sddmDir)
	err := cmd.Run()
	if err != nil {
		app.logger.Fatalf("failed take ownership with error: %s\n", err.Error())
	}

	cmd = exec.Command("sudo", "chown", "-R", fmt.Sprintf("%s:%s", app.whoami, app.whoami), app.sddmThemeDir)
	app.logger.Printf("taking ownership of %s\n", app.sddmDir)
	err = cmd.Run()
	if err != nil {
		app.logger.Fatalf("failed take ownership with error: %s\n", err.Error())
	}
}

func (app *App) enableSddm() {
	cmd := exec.Command("systemctl", "enable", "sddm.service")
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		app.logger.Fatalf("failed to enable sddm")
	}
}

func (app *App) setupDisplayManager() {
	app.chownDisplayManager()
	srcDir := filepath.Join(app.runningDir, "arch-dotfiles")
	app.mvDir(filepath.Join(srcDir, "display-manager", "sddm", "usr-lib-sddm-sddm.conf.d", "sddm.conf"), app.sddmDir)
	app.mvDir(filepath.Join(srcDir, "display-manager", "sddm", "usr-share-sddm-themes", "pedro-default"), app.sddmThemeDir)
	app.enableSddm()
}
