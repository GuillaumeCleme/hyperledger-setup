package main

import (
	"fmt"
	"docker"
)

const (
	fabricTag = ""
	version = "1.0.3"
)

//export ARCH=$(echo "$(uname -s|tr '[:upper:]' '[:lower:]'|sed 's/mingw64_nt.*/windows/')-$(uname -m | sed 's/x86_64/amd64/g')" | awk '{print tolower($0)}')
//Set MARCH variable i.e ppc64le,s390x,x86_64,i386
//MARCH=`uname -m`
var images = [...]string {"peer", "orderer", "couchdb", "ccenv", "javaenv", "kafka", "zookeeper", "tools"}

func dockerFabricPull(){
	
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