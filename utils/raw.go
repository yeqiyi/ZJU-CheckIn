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

//获取Uid
func GetUid(raw []byte) (map[string]string, error) {
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
func GetId(raw []byte) (map[string]string, error) {
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

//获取Date
func GetDate(raw []byte) (map[string]string, error) {
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


//获取created字段
func GetCreated(raw []byte) (map[string]string, error) {
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