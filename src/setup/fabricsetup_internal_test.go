package main

import (
    "testing"
    "fmt"
	"strconv"
	"strings"
)

const(
	bin = "../../bin"
)

func TestGetBinaries(t *testing.T) {
	
	//Void call
	getBinaries(bin)	
}

func TestDownloadFromUrl(t *testing.T) {
	
	location, fileName := downloadFromUrl("https://github.com/hyperledger/fabric/blob/release/README.md", bin)
	
	fmt.Println("Fetched file: " + fileName)
	
	if fileName != "README.md" {
		t.Error("Filename did not match actual file")
	} else if !strings.HasSuffix(location, "/"){
		t.Error("Invalid path returned from downloadFromUrl() function")
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