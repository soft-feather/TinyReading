package file

import (
	"fmt"
	"github.com/wonderivan/logger"
	"os"
)

func GetRealPath(path string) string {
	return fmt.Sprintf("%s/doc/%s", StoragePrefix, path)
}

func Init() {
	StoragePrefix, _ = os.Getwd()
	logger.Debug("file storage prefix : %s", StoragePrefix)
}
