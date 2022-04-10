package models

import (
	"checkIn/utils"
	"net/url"
	"gopkg.in/ini.v1"
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

func(c Config)GetReqBody(raw []byte) (url.Values,error) {
	urlVals := url.Values{}
	rawGeoInfo,g,err:=utils.GetGeoInfo(c.Campus)
	if err!=nil{
		return urlVals,err
	}

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
	urlVals.Add("id",utils.GetId(raw))
	urlVals.Add("uid",utils.GetUid(raw))
	urlVals.Add("date",utils.GetDate())
	urlVals.Add("tw","0") //是否发热
	urlVals.Add("sfcxtz","0")
	urlVals.Add("sfyyjc","0") //与sfcxtz字段相关联(若sfcxtz为0，此字段值为0)
	urlVals.Add("jcjgqr","0") //与Sfyyjc字段相关联（若Sfyyjc为0，此字段值为0）
	urlVals.Add("jcjg","") //与Sfyyjc字段相关联（若Sfyyjc为0，此字段值为空）
	urlVals.Add("sfjcbh","0") //是否接触病患
	urlVals.Add("sfcxzysx","0") //是否有任何与疫情相关的，值得注意的情况（1：有；0：没有）
	urlVals.Add("remark","") //其他信息，与Sfcxzysx字段相关联（若Sfcxzysx为0，此字段为空）
	urlVals.Add("address",g.FmtAddr) //就是formattedAddress
	urlVals.Add("area",g.AddrComp.Prov+" "+g.AddrComp.City+" "+g.AddrComp.Dst) 
	urlVals.Add("province",g.AddrComp.Prov) 
	urlVals.Add("city",g.AddrComp.City)
	urlVals.Add("geo_api_info",rawGeoInfo) ///具体定位信息，从geoinfos/*.json中读取
	urlVals.Add("created",utils.GetCreated(raw)) //上次打卡时间，时间戳（从Raw中的def截取）
	urlVals.Add("qksm","") //情况说明，与字段Sfcxzysx相关联（若Sfcxzysx为0，此字段为空）
	urlVals.Add("sfzx","1") //今日是否在校（在：1；不在：0）
	urlVals.Add("sfjcwhry","0") 
	urlVals.Add("sfcyglq","0") //是否居家观察 为0
	urlVals.Add("gllx","") //隔离场所,为空
	urlVals.Add("glksrq","") //隔离观察开始时间，为空
	urlVals.Add("jcbhlx","") //与Sfjcbh字段关联（若Sfjcbh为0，此字段为空）
	urlVals.Add("jcbhrq","") //与Sfjcbh字段关联（若Sfjcbh为0，此字段为空）
	urlVals.Add("sftjwh","0")
	urlVals.Add("sftjhb","0")
	urlVals.Add("fxyy","") //返校原因
	urlVals.Add("bztcyy","5") //不同城原因,返校为5
	urlVals.Add("fjsj","0")
	urlVals.Add("sfjchbry","0")
	urlVals.Add("jrsfqzys","")
	urlVals.Add("jrsfqzfy","")
	urlVals.Add("sfyqjzgc","") //是否被当地管理部门要求集中观察，为空
	urlVals.Add("jrdqjcqk","")
	urlVals.Add("sfjcqz","")
	urlVals.Add("jcqzrq","")
	urlVals.Add("jcwhryfs","")
	urlVals.Add("jchbryfs","")
	urlVals.Add("xjzd","")
	urlVals.Add("szgj","")
	urlVals.Add("sfsfbh","0")
	urlVals.Add("jhfjrq","")
	urlVals.Add("jhfjjtgj","")
	urlVals.Add("jhfjhbcc","")
	urlVals.Add("jhfjsftjwh","0")
	urlVals.Add("jhfjsftjhb","0")
	urlVals.Add("szsqsfybl","0")
	urlVals.Add("sfsqhzjkk","1") //是否申领所在地健康码（已领：1；未领：0）
	urlVals.Add("sqhzjkkys","1") //健康码颜色（1:绿；2：红；3：黄；4：橙）
	urlVals.Add("sfygtjzzfj","0")
	urlVals.Add("gtjzzfjsj","")
	urlVals.Add("gwszgz","")
	urlVals.Add("gwszgzcs","")
	urlVals.Add("gwszdd","")
	urlVals.Add("szgjcs","")
	specFiled,err:=utils.Getfields(raw)
	if err!=nil{
		return urlVals,err
	}
	for k,v:=range specFiled{
		urlVals.Add(k,v)
	}
	if c.Ismoved{
		urlVals.Add("ismoved","1")
	}else{
		urlVals.Add("ismoved","0")
	}
	urlVals.Add("zgfx14rfhsj","")

	return urlVals,nil
}