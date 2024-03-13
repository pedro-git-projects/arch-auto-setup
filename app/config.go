package app

func (app *App) cloneConfig() error {
	err := app.cloneRepo(app.repos["config"])
	if err != nil {
		return err
	}
	return nil
}
