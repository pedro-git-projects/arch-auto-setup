package app

import (
	"log"
	"os"
	"os/user"
	"path/filepath"
)

type App struct {
	logger       *log.Logger
	repos        map[string]string
	homeDir      string
	runningDir   string
	sddmDir      string
	sddmThemeDir string
	whoami       string
	zsh          string
}

func New() *App {
	repos := map[string]string{
		"config": "https://github.com/pedro-git-projects/arch-dotfiles",
		"yay":    "https://aur.archlinux.org/yay.git",
		"nvim":   "https://github.com/pedro-git-projects/nvim-dotfiles",
	}

	p, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	runningDir := filepath.Dir(p)

	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	whoami := user.Username

	return &App{
		logger:       log.New(os.Stdout, "==> ", log.Ldate|log.Ltime),
		repos:        repos,
		homeDir:      os.Getenv("HOME"),
		runningDir:   runningDir,
		sddmDir:      "/usr/lib/sddm/sddm.conf.d",
		sddmThemeDir: "/usr/share/sddm/themes",
		whoami:       whoami,
		zsh:          "/usr/bin/zsh",
	}
}

func (app *App) Run() {
	err := app.cloneConfig()
	if err != nil {
		app.logger.Fatal(err.Error())
	}
	err = app.cloneYay()
	if err != nil {
		app.logger.Fatal(err.Error())
	}
	err = app.cloneNvim()
	if err != nil {
		app.logger.Fatal(err.Error())
	}
	app.setupNvim()
	app.setupYay()
	app.setupWindowManager()
	app.setupShell()
	app.setupDisplayManager()
}
