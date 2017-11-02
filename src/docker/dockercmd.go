package docker

import (
	"bytes"
	"os/exec"
	"regexp"
	"errors"
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

func GetDockerVersion() (string, error){
	//Create version parsing regex
	var versionRegEx = regexp.MustCompile(`[0-9.]+`)
	
	cmd := exec.Command("docker", "version")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", errors.New("Failed to run [docker version] command: " + err.Error())
	} else{
		
		output := out.String()
		matchResult := versionRegEx.FindStringSubmatch(output)
		
		//Make sure we have at least one match
		if len(matchResult) < 1 {
			return "", errors.New("Failed to parse docker version, result was: " + output)
		} else{
			//Get the first match
			return matchResult[0], nil
		}
	}
}

func DockerPull(imageName string) (string, error){
	
	cmd := exec.Command("docker", "pull", imageName)

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", errors.New("Failed to run [docker pull] command: " + err.Error())
	} else{
		return out.String(), nil
	}	
}

func DockerTag(imageName string) (string, error){
	
	cmd := exec.Command("docker", "tag", imageName)

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", errors.New("Failed to run [docker tag] command: " + err.Error())
	} else{
		return out.String(), nil
	}
}
