package docker

import (
    "testing"
    "fmt"
    "setup"
    "regexp"
)

func TestIsDockerInstalled(t *testing.T) {
	installed := IsDockerInstalled()
	fmt.Println("Docker Installed:", installed)
}

func TestGetDockerVersion(t *testing.T) {
	//Create version matching regex
	var versionRegEx = regexp.MustCompile(`[0-9.]*`)

	message, err := GetDockerVersion()
	
	if err != nil{
		t.Error("Error while retrieving docker version", err)
	} else if !versionRegEx.MatchString(message) {
		t.Error("Docker Version is invalid")
	} else{
		fmt.Println("Docker Version:", message)
	}
}

func TestDockerPull(t *testing.T) {
	//TODO Returns success
	message, err := DockerPull("hyperledger/fabric-peer:x86_64-" + setup.Version)
	
	if err != nil{
		t.Error("Error while running docker pull", err)
	} else{
		fmt.Println("Docker Pull:", "\n", message)
	}
}

