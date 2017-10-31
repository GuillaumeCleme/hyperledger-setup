package main

import (
	"fmt"
	"docker"
	"setup"
)

//export ARCH=$(echo "$(uname -s|tr '[:upper:]' '[:lower:]'|sed 's/mingw64_nt.*/windows/')-$(uname -m | sed 's/x86_64/amd64/g')" | awk '{print tolower($0)}')
//Set MARCH variable i.e ppc64le,s390x,x86_64,i386
//MARCH=`uname -m`

func dockerFabricPull(){
	
	images := setup.Images
	
	for i := 0; i < len(images); i++ {
		fmt.Println(images[i])
		
		docker.DockerPull(images[i])
	}
	
}

func dockerCaPull(){
	
}



func main(){
	
	dockerFabricPull()
}