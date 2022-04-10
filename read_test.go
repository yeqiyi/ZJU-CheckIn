package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestRead(t *testing.T) {
	raw, _ := ioutil.ReadFile("response.log")
	gr, err := gzip.NewReader(bytes.NewReader(raw))
	defer gr.Close()
	if err!=nil{
		fmt.Println(err)
		t.Fail()
	}
	raw,err=ioutil.ReadAll(gr)
	if err!=nil{
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(string(raw))
}