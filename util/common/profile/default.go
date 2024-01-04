package profile

import "os"

var (
	etcProfile *EtcProfile
)

type EtcProfile struct {
	CWD string
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
	}

	return err
}

func GetEtcProfile() *EtcProfile {
	return etcProfile
}
