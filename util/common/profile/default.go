package profile

import "os"

var (
	etcProfile *EtcProfile
)

type EtcProfile struct {
	CWD string
}

func Init() error {
	var err error
	etcProfile.CWD, err = os.Getwd()
	return err
}

func GetEtcProfile() *EtcProfile {
	return etcProfile
}
