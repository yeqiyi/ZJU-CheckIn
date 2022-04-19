package utils

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
)

type GeoInfo struct {
	Type      string   `json:"type"`
	Pos       Position `json:"position"`
	LocatType string   `json:"location_type"`
	Msg       string   `json:"message"`
	Accuracy  string   `json:"accuracy"`
	IsCvt     bool     `json:"isConverted"`
	Stat      int      `json:"status"`
	FmtAddr   string   `json:"formattedAddress"`
	AddrComp AC `json:"addressComponent"`
}

type Position struct {
	Q   float32 `json:"Q"`
	R   float32 `json:"R"`
	Lng float32 `json:"lng"`
	Lat float32 `json:"lat"`
}

type AC struct{
	Prov string `json:"province"`
	City string `json:"city"`
	Dst string `json:"district"`
}

var mp = map[string]string{
	"宁波校区": "geoInfos/cst.json",
	"玉泉校区": "geoInfos/yq.json",
}

func GetGeoInfo(campus string) (string,*GeoInfo,error) {
	if _,ok:=mp[campus];!ok{
		return "",nil,errors.New("无位置信息!(campus:"+campus+")")
	}
	log.Println("Campus is ",campus,",reading geoInfo file : ",mp[campus])
	jsonFile := mp[campus]
	raw,err:=ioutil.ReadFile(jsonFile)
	if err!=nil{
		return "",nil,err
	}
	g:=new(GeoInfo)
	err=json.Unmarshal(raw,g)
	if err!=nil{
		return "",nil,err
	}
	return string(raw),g,nil
}