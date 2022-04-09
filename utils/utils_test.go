package utils

import (
	"fmt"
	"io/ioutil"
	"testing"
)

var RawPath="../docs/Raw"

func TestGetField(t *testing.T) {
	if raw,err:=ioutil.ReadFile(RawPath);err==nil{
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

func TestGetUid(t *testing.T) {
	if raw,err:=ioutil.ReadFile(RawPath);err==nil{
		if mp,err:=GetUid(raw);err==nil{
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