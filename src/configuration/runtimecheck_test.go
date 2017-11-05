package configuration_test

import (
    "testing"
    "configuration"
    "fmt"
)

func TestGetRuntimeArch(t *testing.T) {
	
	runtime := configuration.GetRuntimeArch()
	
	if runtime == ""{
		t.Error("Could not get runtime architecture")
	}
}

func TestGetRuntimeOS(t *testing.T) {
	
	runtime := configuration.GetRuntimeOS()
	
	if runtime == ""{
		t.Error("Could not get runtime OS")
	}
}

func TestGetImageArch(t *testing.T) {
	
	runtime := configuration.GetImageArch()
	runtimeArch := configuration.GetRuntimeArch()
	
	if runtime == ""{
		t.Error("Could not get runtime architecture")
	} else if runtimeArch == "amd64" && runtime != "x86_64" {
		t.Error("Invalid runtime architecture")
	}
}

func TestGetBinArch(t *testing.T) {
	
	binArch := configuration.GetBinArch()
	
	if binArch == ""{
		t.Error("Could not get bin architecture")
	}
	
	fmt.Println("Bin Architecture: " + binArch)
}