package app

import (
	"os"
	"path/filepath"
)

func (app *App) cloneNvim() error {
	err := app.cloneRepo(app.repos["nvim"])
	if err != nil {
		return err
	}
	return nil
}

func (app *App) setupNvim() {
	srcDir := filepath.Join(app.runningDir, "nvim-dotfiles")
	err := os.Rename(srcDir, "nvim")
	if err != nil {
		app.logger.Fatal(err.Error())
	}
	srcDir = filepath.Join(app.runningDir, "nvim")
	cfgDir := filepath.Join(app.homeDir, ".config")
	app.mvDir(srcDir, cfgDir)
}
