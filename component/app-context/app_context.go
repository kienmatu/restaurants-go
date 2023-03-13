package appContext

import (
	"github.com/kienmatu/restaurants-go/utils"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
}

func NewAppContext(db *gorm.DB, cfg *utils.Configuration) AppContext {
	return &appCtx{
		db:     db,
		config: cfg,
	}
}

type appCtx struct {
	db     *gorm.DB
	config *utils.Configuration
}

func (a *appCtx) GetMainDBConnection() *gorm.DB {
	return a.db
}
