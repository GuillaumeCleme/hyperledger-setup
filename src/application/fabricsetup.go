package main

import (
	"fmt"
	"log"
	"docker"
	"setup"
	"strconv"
	"strings"
	"errors"
)

//export ARCH=$(echo "$(uname -s|tr '[:upper:]' '[:lower:]'|sed 's/mingw64_nt.*/windows/')-$(uname -m | sed 's/x86_64/amd64/g')" | awk '{print tolower($0)}')
//Set MARCH variable i.e ppc64le,s390x,x86_64,i386
//MARCH=`uname -m`

func dockerFabricPull(){
	
	images := setup.IMAGES
	
	for i := 0; i < len(images); i++ {
		fmt.Println(images[i])
		
		docker.DockerPull(images[i])
	}
	
}

func dockerCaPull(){
	
}

func checkDockerReq() (bool, error){
	
	if docker.IsDockerInstalled() {
		
		version, err := docker.GetDockerVersion()
		
		checkErr("Error while getting Docker version", err)
			
		verInt, err := strconv.Atoi(strings.Replace(version, ".", "", -1))
		checkErr("Error while parsing Docker version", err)
		
		minVerInt, err := strconv.Atoi(strings.Replace(setup.MIN_DOCKER_VER, ".", "", -1))
		checkErr("Error while parsing minimum Docker version", err)
		
		if verInt < minVerInt {
			//Minimum version not satisfied
			return false, errors.New("Minimum Docker version [" + setup.MIN_DOCKER_VER + "] not satisfied")
		} else{
			return true, nil
		}
		
	} else{
		//Error, docker not installed
		return false, errors.New("Docker not installed")
	}
	
}

func checkErr(message string, err error) {
    if err != nil {
        log.Fatal("ERROR: ", message, ": ", err)
    }
}

func main(){
	//dockerFabricPull()
	fmt.Println(setup.GetRuntimeArch())
	fmt.Println(setup.GetRuntimeOS())
	
	dockerReq, err := checkDockerReq()
	checkErr("Docker requirements unsatisfied", err)
	
	fmt.Println("Docker requirements satisfied: " + strconv.FormatBool(dockerReq))
}