package storage

import (
	"github.com/soft-feather/TinyReading/util/storage/file"
	"github.com/soft-feather/TinyReading/util/storage/sqlite"
)

func Init() error {
	file.Init()
	// json_storage.Init()
	err := sqlite.Init()
	if err != nil {
		return err
	}
	return nil
}
