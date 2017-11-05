package setup

import (
    "testing"
)

func TestGetRuntimeArch(t *testing.T) {
	
	runtime := GetRuntimeArch()
	
	if runtime == ""{
		t.Error("Could not get runtime architecture")
	}
}

func TestGetRuntimeOS(t *testing.T) {
	
	runtime := GetRuntimeOS()
	
	if runtime == ""{
		t.Error("Could not get runtime OS")
	}
}

func TestGetImageArch(t *testing.T) {
	
	runtime := GetImageArch()
	runtimeArch := GetRuntimeArch()
	
	if runtime == ""{
		t.Error("Could not get runtime architecture")
	} else if runtimeArch == "amd64" && runtime != "x86_64" {
		t.Error("Invalid runtime architecture")
	}
}