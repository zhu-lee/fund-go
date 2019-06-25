package app

import "flag"

var (
	profileFlag = flag.String("p", "", `application active profiles`)
	signalFlag  = flag.String("s", "", `send signal to the daemon
		quit - graceful shutdown
		stop - fast shutdown
		reload - reloading the configuration file`)
)

func GetProfileFlag()*string  {
	parseFlag()
	return profileFlag
}

func GetSignalFlag()*string  {
	parseFlag()
	return signalFlag
}

func parseFlag()  {
	if !flag.Parsed() {
		flag.Parse()
	}
}