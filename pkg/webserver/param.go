package webserver

import "flag"

type WebserverConfig struct {
	Address string
	Port    string

	DryRun  bool
	Profile string
}

var (
	WebConfig              = flag.NewFlagSet("web", flag.ExitOnError)
	DefaultWebserverConfig *WebserverConfig
)

func init() {
	DefaultWebserverConfig = &WebserverConfig{}
	WebConfig.String(DefaultWebserverConfig.Address, "default", "this is b")
}
