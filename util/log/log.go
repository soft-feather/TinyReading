package log

import (
	"github.com/wonderivan/logger"
)

func Init() error {
	err := logger.SetLogger("./log.json")
	if err != nil {
		return err
	}
	logger.Debug("debug init : %v ", logger.LevelMap)
	return nil
}
