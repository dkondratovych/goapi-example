package task

import (
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"

	"github.com/jinzhu/gorm"
	"github.com/seesawlabs/Dima-Kondravotych-Exercise/shared/config"
	"github.com/seesawlabs/Dima-Kondravotych-Exercise/shared/storages"
	"github.com/seesawlabs/Dima-Kondravotych-Exercise/shared/database"
	"github.com/seesawlabs/Dima-Kondravotych-Exercise/server"
)

type Application struct {
	DB 				*gorm.DB
	Config 			*config.Config
	Server 			*server.Server
	StorageProvider	storages.IStorageProvider
}

func(a *Application) LoadConfig(configPath string) error {
	if err := a.Config.Load(configPath); err != nil {
		return err
	}

	return nil
}

func(a *Application) Init() error {

	if err := a.InitLogger(); err != nil {
		return fmt.Errorf("Error while logger initialization %v", err)
	}

	if err := a.InitDb(); err != nil {
		return fmt.Errorf("Error while database initialization %v", err)
	}

	a.InitStorageProvider()

	return nil
}

func(a *Application) InitLogger() error {
	file, err := os.OpenFile(a.Config.Application.LogFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		return fmt.Errorf("Error openning log file %v", err)
	}

	log.SetOutput(file)
	log.SetFormatter(&log.JSONFormatter{})

	return nil
}

func(a *Application) InitDb() error {
	db, err := database.NewGormDb(a.Config.Database)
	if err != nil {
		return err
	}

	a.DB = db

	return nil
}

func(a *Application) InitStorageProvider() {
	a.StorageProvider = storages.NewStorageProvider(a.DB)
}

func (a *Application) Run() error {
	a.Server = server.NewServer(a.Config, a.StorageProvider)

	if err := a.Server.Run(); err != nil {
		return err
	}

	return nil
}

func NewApplication() *Application {
	return &Application{
		Config: config.NewConfig(),
	}
}