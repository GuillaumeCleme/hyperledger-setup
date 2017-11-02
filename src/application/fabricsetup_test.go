package main

import (
    "testing"
    "fmt"
	"strconv"
)

func TestCheckDockerReq(t *testing.T) {
	dockerReq, err := checkDockerReq()
	checkErr("Docker requirements unsatisfied", err)
	
	fmt.Println("Docker requirements satisfied: " + strconv.FormatBool(dockerReq))
}

