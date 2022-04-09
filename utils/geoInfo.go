package utils

import (
	"encoding/json"
	"io/ioutil"
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
	"宁波校区": "../geoInfos/cst.json",
}

func GetGeoInfo(campus string) (string,*GeoInfo,error) {
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