package app

import (
	"log"
	"os"
	"os/user"
	"path/filepath"
)

type App struct {
	logger     *log.Logger
	repos      map[string]string
	homeDir    string
	runningDir string
	whoami     string
	zsh        string
}

func New() *App {
	repos := map[string]string{
		"config": "https://github.com/pedro-git-projects/arch-dotfiles",
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
		logger:     log.New(os.Stdout, "==> ", log.Ldate|log.Ltime),
		repos:      repos,
		homeDir:    os.Getenv("HOME"),
		runningDir: runningDir,
		whoami:     whoami,
		zsh:        "/usr/bin/zsh",
	}
}

func (app *App) Run() {
	app.cloneConfig()
	app.setupWindowManager()
	app.setupShell()
}