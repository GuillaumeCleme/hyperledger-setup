package docker

import (
    "testing"
    "fmt"
    "setup"
)

func TestIsDockerInstalled(t *testing.T) {
	//TODO Returns success
	installed := IsDockerInstalled()
	fmt.Println("Docker Installed:", installed)
}

func TestDockerPull(t *testing.T) {
	//TODO Returns success
	DockerPull("hyperledger/fabric-peer:x86_64-" + setup.Version)
}

