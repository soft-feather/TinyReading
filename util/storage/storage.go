package storage

import (
	"github.com/soft-feather/TinyReading/util/storage/file"
	"github.com/soft-feather/TinyReading/util/storage/sqlite"
)

func Init() {
	file.Init()
	// json_storage.Init()
	sqlite.Init()
}
