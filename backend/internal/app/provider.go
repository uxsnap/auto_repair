package app

import "github.com/uxsnap/auto_repair/backend/internal/config"

func (a *App) GetConfigDB() *config.ConfigDB {
	if a.configDB == nil {
		a.configDB = config.NewConfigDB()
	}
	return a.configDB
}

func (a *App) GetConfigHTTP() *config.ConfigHTTP {
	if a.configHTTP == nil {
		a.configHTTP = config.NewConfigHttp()
	}
	return a.configHTTP
}
