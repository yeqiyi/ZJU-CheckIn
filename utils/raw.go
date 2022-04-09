package utils

import (
	"errors"
	"regexp"
)

//获取特殊字段
func Getfields(raw []byte) (map[string]string, error) {
	mp := make(map[string]string)
	rawStr := string(raw)
	reg := regexp.MustCompile(`szgjcs.*?,(?s:.*?)\"(.*?)\"\s?:\s?\"(.*?)\",(?s:.*?)\"(.*?)\"\s?:\s?\"(.*?)\"`)
	
	res:=reg.FindAllStringSubmatch(rawStr,-1)
	//fmt.Println(res)
	for _, text := range res {
        //fmt.Println("i = ",i,",text[1] = ", text)
		for i:=1;i<len(text);i+=2{
			val:=regexp.MustCompile(`\b+?`) //验证提取的字段是否合法（仅包含数字和字母）
			if val.MatchString(text[i])&&val.MatchString(text[i+1]){
				mp[text[i]]=text[i+1]
			}else{
				return mp,errors.New("匹配错误，字段提取失败")
			}
		}
    }
	return mp,nil
}

//获取Id
func GetId(raw []byte) (string, error) {
	id:=""
	rawStr := string(raw)
	reg := regexp.MustCompile(`var.*?def.*?{.*?\"id\":\"(\d*)\".*?}`)
	
	res:=reg.FindAllStringSubmatch(rawStr,-1)
	//fmt.Println(res)
	id=res[0][1]
	return id,nil
}

//获取Uid
func GetUid(raw []byte) (string, error) {
	uid:=""
	rawStr := string(raw)
	reg := regexp.MustCompile(`var.*?def.*?{.*?\"uid\":\"(\d*)\".*?}`)
	
	res:=reg.FindAllStringSubmatch(rawStr,-1)
	//fmt.Println(res)
	uid=res[0][1]
	return uid,nil
}


//获取Date
func GetDate(raw []byte) (string, error) {
	date:=""
	rawStr := string(raw)
	reg := regexp.MustCompile(`var.*?def.*?{.*?\"date\":\"(\d*)\".*?}`)
	
	res:=reg.FindAllStringSubmatch(rawStr,-1)
	//fmt.Println(res)
	date=res[0][1]
	return date,nil
}


//获取created字段
func GetCreated(raw []byte) (string, error) {
	created:=""
	rawStr := string(raw)
	reg := regexp.MustCompile(`var.*?def.*?{.*?\"created\":\"(\d*)\".*?}`)
	
	res:=reg.FindAllStringSubmatch(rawStr,-1)
	//fmt.Println(res)
	created=res[0][1]
	return created,nil
}