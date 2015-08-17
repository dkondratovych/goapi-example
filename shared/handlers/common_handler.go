package handlers

import (
	"github.com/seesawlabs/Dima-Kondravotych-Exercise/shared/storages"
	"github.com/seesawlabs/Dima-Kondravotych-Exercise/shared/config"
)

// CommonHandler is user by all server handlers as default handler
type CommonHandler struct {
	Storage  storages.IStorageProvider
	Config *config.Config
}

func (ch *CommonHandler) SetStorage(st storages.IStorageProvider) {
	ch.Storage = st
}

func (ch *CommonHandler) SetConfig(cf *config.Config) {
	ch.Config = cf
}