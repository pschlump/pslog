package pslog

import (
	"fmt"
	"os"
)

func Info(x ...interface{}) {
	Infof("%s\n", x...)
}

func Infof(fmt string, x ...interface{}) {
	msg("Info:"+fmt, x...)
}

func Warn(x ...interface{}) {
	Warnf("%s\n", x...)
}

func Warnf(fmt string, x ...interface{}) {
	msg("Warning:"+fmt, x...)
}

func Error(x ...interface{}) {
	Errorf("%s\n", x...)
}

func Errorf(fmt string, x ...interface{}) {
	msg("Error:"+fmt, x...)
}

func Fatal(x ...interface{}) {
	Fatalf("%s\n", x...)
}

func Fatalf(fmt string, x ...interface{}) {
	msg("Fatal:"+fmt, x...)
	os.Exit(1)
}

func msg(ff string, x ...interface{}) {
	// TODO - log date time - find the config stuff in func()init - commented out and implement it.
	// TODO - Destination for logging
	// TODO - Location called from
	fmt.Printf(ff, x...)
}
