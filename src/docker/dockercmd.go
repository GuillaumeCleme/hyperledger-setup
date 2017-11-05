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
		return "", errors.New("Failed to run [docker version] command: " +
			 err.Error() + 
			 ". Is Docker started?")
	} else{
		
		output := out.String()
		matchResult := versionRegEx.FindStringSubmatch(output)
		
		//Make sure we have at least one match
		if len(matchResult) < 1 {
			return "", errors.New("Failed to parse Docker version, result was: " + output)
		} else{
			//Get the first match
			return matchResult[0], nil
		}
	}
}

func ExecDockerCmd(params ...string) (string, error){
	
	fmt.Println("Executing: docker", strings.Join(params," "))
	
	cmd := exec.Command("docker", params...)

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", errors.New("Failed to execute Docker command: " + 
			err.Error() +
			". Output was: \n" +
			out.String())
	} else{
		return out.String(), nil
	}	
}
