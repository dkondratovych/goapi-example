package handlers

import (
	"github.com/seesawlabs/Dima-Kondravotych-Exercise/shared/storages"
)

// CommonHandler is user by all server handlers as default handler
type CommonHandler struct {
	Storage  storages.IStorageProvider
}

func (c *CommonHandler) SetStorage(st storages.IStorageProvider) {
	c.Storage = st
}