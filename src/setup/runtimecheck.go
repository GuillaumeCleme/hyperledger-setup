package setup

import (
	"runtime"
)

const(
	goArch string = runtime.GOARCH
	goOs string = runtime.GOOS
)

func GetRuntimeArch() string{
	return goArch
}

func GetRuntimeOS() string{
	return goOs
}
