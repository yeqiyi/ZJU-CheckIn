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
		fmt.Println(GetUid(raw))
	}else{
		fmt.Println(err)
		t.Fail()
	}
}

func TestGetCreated(t *testing.T) {
	if raw,err:=ioutil.ReadFile(RawPath);err==nil{
		fmt.Println(GetCreated(raw))
	}else{
		fmt.Println(err)
		t.Fail()
	}
}

func TestGetId(t *testing.T) {
	if raw,err:=ioutil.ReadFile(RawPath);err==nil{
		fmt.Println(GetId(raw))
	}else{
		fmt.Println(err)
		t.Fail()
	}
}

func TestGetDate(t *testing.T) {
	//fmt.Println(GetDate())
}

func TestReadFile(t *testing.T) {
	jsonFile:="../geoinfos/cst.json"
	if raw,err:=ioutil.ReadFile(jsonFile);err==nil{
		fmt.Println(string(raw))
	}
}

func TestGetGeoInfo(t *testing.T) {
	campus:="宁波校区"
	if r,g,err:=GetGeoInfo(campus);err==nil{
		fmt.Println(r)
		fmt.Println(g)
	}else{
		fmt.Println(err)
		t.Fail()
	}
}

func TestBot(t *testing.T) {
	bot:=GetBot("../config.ini")
	err:=bot.SendMsg("Hello world 喵")
	if err!=nil{
		fmt.Println(err)
		t.Fail()
	}
}