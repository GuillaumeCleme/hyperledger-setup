package docker

import (
	"bytes"
	"os/exec"
	"regexp"
	"errors"
	"fmt"
	"strings"
)

func IsDockerInstalled() bool{
	if _, err := ExecDockerCmd(); err != nil {
		return false
	} else {
		return true
	}
}

func GetDockerVersion() (string, error){
	//Create version parsing regex
	var versionRegEx = regexp.MustCompile(`[0-9.]+`)
	
	
	message, err := ExecDockerCmd("version")
	
	if err != nil {
		return "", errors.New("Failed to run [docker version] command, is Docker started?: " +
			 err.Error())
	} else{
		
		matchResult := versionRegEx.FindStringSubmatch(message)
		
		//Make sure we have at least one match
		if len(matchResult) < 1 {
			return "", errors.New("Failed to parse Docker version, result was: " + message)
		} else{
			//Get the first match
			return matchResult[0], nil
		}
	}
}

func ExecDockerCmd(params ...string) (string, error){
	
	fmt.Println("Executing: docker", strings.Join(params," "))
	
	cmd := exec.Command("docker", params...)

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return "", errors.New("Failed to execute Docker command: " + 
			err.Error() +
			". Output was: \n" +
			stderr.String())
	} else{
		return stdout.String(), nil
	}	
}
