package setup_test

import (
    "testing"
    "setup"
    "fmt"
)

func TestGetRuntimeArch(t *testing.T) {
	
	runtime := setup.GetRuntimeArch()
	
	if runtime == ""{
		t.Error("Could not get runtime architecture")
	}
}

func TestGetRuntimeOS(t *testing.T) {
	
	runtime := setup.GetRuntimeOS()
	
	if runtime == ""{
		t.Error("Could not get runtime OS")
	}
}

func TestGetImageArch(t *testing.T) {
	
	runtime := setup.GetImageArch()
	runtimeArch := setup.GetRuntimeArch()
	
	if runtime == ""{
		t.Error("Could not get runtime architecture")
	} else if runtimeArch == "amd64" && runtime != "x86_64" {
		t.Error("Invalid runtime architecture")
	}
}

func TestGetBinArch(t *testing.T) {
	
	binArch := setup.GetBinArch()
	
	if binArch == ""{
		t.Error("Could not get bin architecture")
	}
	
	fmt.Println("Bin Architecture: " + binArch)
}