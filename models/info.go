package models

type Info struct{
	Sfymqjczrj string //家人是否14天入境或拟入境，为0
	Zjdfgj string //14天到过的国家地区，为空

	Sfyrjjh string //未来14天是否有入境规划（有：1；没有：0）
	Cfgj string //没入境规划为空
	Tjgj string //没入境规划为空
	Nrjrq string //没入境规划为0
	Rjka string //没入境规划为空
	Jnmddsheng string //没入境规划为空
	Jnmddshi string //没入境规划为空
	Jnmddqu string //没入境规划为空
	Jnmddxiangxi string //没入境规划为空
	Rjjtfs string //没入境规划为空
	Rjjtfs1 string //没入境规划为空
	Rjjtgjbc string //没入境规划为空
	Jnjtfs string //没入境规划为空
	Jnjtfs1 string //没入境规划为空
	Jnjtgjbc string //没入境规划为空

	Sfqrxxss string //提交前确认信息是否准确（ok：1）
	Sfqtyyqjwdg string //是否因发热外其他原因未到岗（是：1；否：0）
	Sffrqjwdg string //是否因发热请假未到岗（是：1；否：0）

	Sfhsjc string //含义不明，为空


	Zgfx14rfh string //14日是否入境，拟入境，为0
	Zgfx14rfhdd string //与字段Zgfx14rfh相关联(若Zgfx14rfh为0,此字段为空)
	

	Sfyxjzxgym string //是否意向接种（有：1；没有：0）
	Sfbyjzrq string //是否不宜接种人群（1：对疫苗成分过敏；2：患急性疾病者；3：处于慢性疾病的急性发作期者；4：妊娠期妇女；5：否）
	Jzxgymqk string //当前接种情况（1：第一针；4：第二针（已满6个月）；5：第二针（未满6个月）；6：第三针；3：未接种）

	Campus string //所在校区：宁波校区

	Id string
	Uid string
	Date string
	Tw string //是否发热
	Sfcxtz string
	Sfyyjc string //与sfcxtz字段相关联(若sfcxtz为0，此字段值为0)
	Jcjgqr string  //与Sfyyjc字段相关联（若Sfyyjc为0，此字段值为0）
	Jcjg string //与Sfyyjc字段相关联（若Sfyyjc为0，此字段值为空）

	Sfjcbh string //是否接触病患
	Sfcxzysx string //是否有任何与疫情相关的，值得注意的情况（1：有；0：没有）

	Remark string  //其他信息，与Sfcxzysx字段相关联（若Sfcxzysx为0，此字段为空）

	Address string //就是formattedAddress
	Area string // "浙江省 宁波市 鄞州区"
	Province string //"浙江省"
	City string //"宁波市"
	GeoApiInfo string //具体定位信息，从config/getinfo.json中读取
	Created string //上次打卡时间，时间戳（从Raw中的def截取）
	Qksm string  //情况说明，与字段Sfcxzysx相关联（若Sfcxzysx为0，此字段为空）

	Sfzx string //今日是否在校（在：1；不在：0）
	Sfjcwhry string //含义不明，一般为0

	Sfcyglq string //是否居家观察 为0
	Gllx string //隔离场所,为空
	Glksrq string //隔离观察开始时间，为空

	Jcbhlx string //与Sfjcbh字段关联（若Sfjcbh为0，此字段为空）
	Jcbhrq string //与Sfjcbh字段关联（若Sfjcbh为0，此字段为空）

	Sftjwh string //含义不明，一般为0
	Sftjhb string //含义不明，一般为0

	Fxyy string //返校原因 不管写没写传到服务端都为空
	Bztcyy string   //不同城原因
					//	"1" : 其他
					//  "2" : 探亲
					//  "3" : 旅游
					//  "4" : 回家
					//  "5" : 返校
					//  "6" : 出差


	fjsj string //含义不明，为0
	Sfjchbry string //含义不明，为0
	Jrsfqzys string //含义不明，为空
	Jrsfqzfy string //含义不明，为空

	Sfyqjzgc string //是否被当地管理部门要求集中观察，为空
	Jrdqjcqk string //含义不明，为空
	Sfjcqz string //废弃字段，为空
	Jcqzrq string //含义不明，为空
	Jcwhryfs string //含义不明，为空
	Jchbryfs string //含义不明，为空
	Xjzd string //含义不明，为空
	Szgj string //所在国家，为空

	Sfsfbh string //含义不明，为0
	Jhfjrq string //含义不明，为空
	Jhfjjtgj string //含义不明，为空
	Jhfjhbcc string //含义不明，为空
	Jhfjsftjwh string //含义不明，为0
	Jhfjsftjhb string //含义不明，为0
	Szsqsfybl int //含义不明，为0

	Sfsqhzjkk string //是否申领所在地健康码（已领：1；未领：0）
	Sqhzjkkys string //健康码颜色（1:绿；2：红；3：黄；4：橙）

	Sfygtjzzfj int //含义不明，为0
	Gtjzzfjsj string //含义不明，为空

	Gwszgz string //自选打卡地点（已废弃），为空
	Gwszgzcs string //自选打卡地点（已废弃），为空
	Gwszdd string //自选打卡地点（已废弃），为空

	Szgjcs string //所在国家城市，在国内为空

	SpecFiled map[string]string //服务器签发特殊字段

	Ismoved string //打卡地址是否和前一天相同。相同为0，不同为1

	Zgfx14rfhsj string //与字段Zgfx14rfh相关联(若Zgfx14rfh为0,此字段为空)

}