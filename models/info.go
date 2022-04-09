package models

import (
	"gopkg.in/ini.v1"
	"net/url"
)


type Config struct{
	Campus string `ini:"Campus"`
	Ismoved bool `ini:"Ismoved"`
}

func LoadConf(configPath string)(*Config,error){
	conf:=new(Config)
	err:=ini.MapTo(conf,configPath)
	if err!=nil{
		return nil,err
	}
	return conf,nil
}

func(c Config)GetReqBody() url.Values {
	urlVals := url.Values{}

	urlVals.Add("sfymqjczrj","0") //家人是否14天入境或拟入境，为0
	urlVals.Add("zjdfgj","") //14天到过的国家地区，为空
	urlVals.Add("sfyrjjh","0") //未来14天是否有入境规划（有：1；没有：0）
	urlVals.Add("cfgj","") //没入境规划为空
	urlVals.Add("tjgj","") //没入境规划为空
	urlVals.Add("nrjrq","0") //没入境规划为0
	urlVals.Add("rjka","") //没入境规划为空
	urlVals.Add("jnmddsheng","") //没入境规划为空
	urlVals.Add("jnmddshi","") //没入境规划为空
	urlVals.Add("jnmddqu","") //没入境规划为空
	urlVals.Add("jnmddxiangxi","") //没入境规划为空
	urlVals.Add("rjjtfs","") //没入境规划为空
	urlVals.Add("rjjtfs1","") //没入境规划为空
	urlVals.Add("rjjtgjbc","") //没入境规划为空
	urlVals.Add("jnjtfs","") //没入境规划为空
	urlVals.Add("jnjtfs1","") //没入境规划为空
	urlVals.Add("jnjtgjbc","") //没入境规划为空
	urlVals.Add("sfqrxxss","1")  //提交前确认信息是否准确（ok：1）
	urlVals.Add("sfqtyyqjwdg","0") //是否因发热外其他原因未到岗（是：1；否：0）
	urlVals.Add("sffrqjwdg","0") //是否因发热请假未到岗（是：1；否：0）
	urlVals.Add("sfhsjc","") //含义不明，为空
	urlVals.Add("zgfx14rfh","0")  //14日是否入境或拟入境，为0
	urlVals.Add("zgfx14rfhdd","") //与字段zgfx14rfh相关联(若Zgfx14rfh为0,此字段为空)
	urlVals.Add("sfyxjzxgym","1") //是否意向接种（有：1；没有：0）
	urlVals.Add("sfbyjzrq","5") //是否不宜接种人群:否
	urlVals.Add("jzxgymqk","6") //当前接种情况:已接种第三针
	urlVals.Add("campus","宁波校区") //所在校区
	urlVals.Add("id","")




	return urlVals
}