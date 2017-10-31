package docker

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func IsDockerInstalled() bool{
	cmd := exec.Command("docker")
	err := cmd.Run()
	if(err != nil) {
		return false
	} else {
		return true
	}
}

func DockerPull(imageName string){
	
	cmd := exec.Command("docker", "pull", imageName)
	//cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("docker pull:", out.String())
	
}

func DockerTag(imageName string){
	
	cmd := exec.Command("docker", "tag", imageName)
	//cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("docker tag:", out.String())
	
}
