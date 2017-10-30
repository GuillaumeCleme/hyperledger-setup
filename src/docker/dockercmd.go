package docker

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)


func DockerPull(imageName string){
	
	cmd := exec.Command("docker", "pull", imageName)
	//cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("docker pull: %q\n", out.String())
	
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
	fmt.Printf("docker tag: %q\n", out.String())
	
}
