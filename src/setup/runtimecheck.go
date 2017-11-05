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

func GetImageArch() string{
	
	switch arch := goArch; arch {
		case "s390x":
			return arch
		case "ppc64le":
			return arch
		default:
			return "x86_64"
	}
}
