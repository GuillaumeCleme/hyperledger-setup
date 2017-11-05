package configuration

import (
	"runtime"
	"log"
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


func GetBinArch() string{
	
	if(contains(goOs, SUPPORTED_OS[:]) && contains(goArch, SUPPORTED_ARCHS[:])){
		return goOs + "-" + goArch
	} else{
		log.Fatal("Retrieved incompatible package architecture")
		return ""
	}
}

func contains(str string, arr []string) bool {
   for _, a := range arr {
      if a == str {
         return true
      }
   }
   return false
}
