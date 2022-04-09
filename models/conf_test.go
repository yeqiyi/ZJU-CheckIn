package models

import (
	"fmt"
	"testing"
)

func TestLoadConf(t *testing.T) {
	if conf,err:=LoadConf("../config.ini");err==nil{
		fmt.Println(conf)
	}else{
		fmt.Println(err)
		t.Fail()
	}
}