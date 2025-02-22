package macos

import (
	"go.uber.org/zap"
)

type OsManager struct {
	logger *zap.Logger
}

func NewOsManager(logger *zap.Logger) *OsManager {
	return &OsManager{logger: logger}
}
