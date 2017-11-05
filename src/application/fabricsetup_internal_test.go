package main

import (
    "testing"
    "fmt"
	"strconv"
//	"setup"
//	"log"
)

func TestGetBinaries(t *testing.T) {
	
	getBinaries()
	
//	fmt.Println("Fetched file: " + fileName)
//	
//	if fileName != "README.md" {
//		t.Error("Filename did not match actual file")
//	}
	
}

func TestDownloadFromUrl(t *testing.T) {
	
	fileName := downloadFromUrl("https://github.com/hyperledger/fabric/blob/release/README.md")
	
	fmt.Println("Fetched file: " + fileName)
	
	if fileName != "README.md" {
		t.Error("Filename did not match actual file")
	}
	
}

func TestCheckDockerReq(t *testing.T) {
	dockerReq, err := checkDockerReq()
	checkErr("Docker requirements unsatisfied", err)
	
	fmt.Println("Docker requirements satisfied: " + strconv.FormatBool(dockerReq))
}

func TestCheckErr(t *testing.T) {
	
	checkErr("Test w/o err", nil)
	
	//TODO Monkey patch
	//checkErr("Test w/ err", errors.New("Simulated error"))
	
	
}