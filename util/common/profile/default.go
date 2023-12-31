package profile

import "os"

var (
	etcProfile *EtcProfile
)

type StorageProfile struct {
	SqliteDBFileName string
}

type EtcProfile struct {
	CWD string
	StorageProfile
}

func Init() error {
	var (
		err error
		cwd string
	)
	if cwd, err = os.Getwd(); err != nil {
		return err
	}
	etcProfile = &EtcProfile{
		CWD: cwd,
		StorageProfile: StorageProfile{
			SqliteDBFileName: "storage.db",
		},
	}

	return err
}

func GetEtcProfile() *EtcProfile {
	return etcProfile
}
