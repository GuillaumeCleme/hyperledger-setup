package main

import (
	"fmt"
	"log"
	"docker"
	"setup"
	"strconv"
	"strings"
	"errors"
	"github.com/mholt/archiver"
	"os"
	"net/http"
	"io"
)

//export ARCH=$(echo "$(uname -s|tr '[:upper:]' '[:lower:]'|sed 's/mingw64_nt.*/windows/')-$(uname -m | sed 's/x86_64/amd64/g')" | awk '{print tolower($0)}')
//Set MARCH variable i.e ppc64le,s390x,x86_64,i386
//MARCH=`uname -m`

func dockerFabricPull(){
	
	arch := setup.GetImageArch()
	
	for i := 0; i < len(setup.IMAGES); i++ {
		
		imageName := setup.DOCKER_IMG_PREFIX + setup.IMAGES[i] + ":" + arch + "-" + setup.VERSION
		
		fmt.Println("Pulling Docker image: " + imageName)
		
		message, err := docker.ExecDockerCmd("pull", imageName)
		
		//Will stop here if an error is encountered
		checkErr("Error while running docker pull", err)

		fmt.Println("Docker Pull:", "\n", message)
	}
}

func getBinaries(){

	arch := setup.GetBinArch()
	
	if arch == ""{
		log.Fatal("Retrieved incompatible package architecture")
	}
	
	dlName := "hyperledger-fabric-" + arch + "-" + setup.VERSION + ".tar.gz"
	
	fmt.Println("Downloading file: " + dlName)
	
	tarFile := downloadFromUrl(setup.BIN_DL_ROOT + arch + "-" + setup.VERSION + "/" + dlName) //TODO
	
	err := archiver.TarGz.Open(tarFile, setup.BIN)
	
	//Will stop here if an error is encountered
	checkErr("Error while extracting TAR file: " + tarFile, err)	
}

func downloadFromUrl(url string) string {
	tokens := strings.Split(url, "/")
	fileName := tokens[len(tokens)-1]
	fmt.Println("Downloading", url, "to", fileName)

	output, err := os.Create(fileName)	
	defer output.Close()
	checkErr("Error while creating: " + fileName, err)

	response, err := http.Get(url)
	defer response.Body.Close()
	checkErr("Error while downloading from: " + url, err)

	n, err := io.Copy(output, response.Body)
	checkErr("Error while downloading from: " + url, err)

	fmt.Println(n, "bytes downloaded.")
	
	return fileName
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
	
	//Check Docker Requirements
	dockerReq, err := checkDockerReq()
	checkErr("Docker requirements unsatisfied", err)
	
	fmt.Println("Docker requirements satisfied: " + strconv.FormatBool(dockerReq))
	
	//Pull down images
	dockerFabricPull()
	
	//Get binaries
	getBinaries()
}