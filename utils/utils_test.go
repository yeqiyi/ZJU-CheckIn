package utils

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestGetField(t *testing.T) {
	if raw,err:=ioutil.ReadFile("../docs/Raw");err==nil{
		if mp,err:=Getfields(raw);err==nil{
			fmt.Println(mp)
		}else{
			fmt.Println(err)
			t.Fail()
		}
	}else{
		fmt.Println(err)
		t.Fail()
	}
}