package checkIn

import (
	"checkIn/client"
	"checkIn/models"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type CheckIn struct {
	client *http.Client
	cookie string 
}

var cookiePath="../.cookie"
var confPath="../config.ini"
var url1="https://healthreport.zju.edu.cn/ncov/wap/default/index"
var url2="https://healthreport.zju.edu.cn/ncov/wap/default/index/ncov/wap/default/save"

func SetUp()(CheckIn,error){
	c:=CheckIn{}
	c.client=client.NewClient()
	err:=c.getCookie()
	if err!=nil{
		return c,err 
	}
	return c,nil
}

func(c *CheckIn)getCookie() error{
	raw,err:=ioutil.ReadFile(cookiePath)
	if err!=nil{
		return err
	}
	c.cookie=string(raw)
	return nil
}

func(c CheckIn)SignIn()error{
	conf,err:=models.LoadConf(confPath)
	if err!=nil{
		return err
	}
	req1,err:=http.NewRequest(http.MethodGet,url1,nil)
	req1.Header.Add("cookie",c.cookie)
	req1.Header.Add("user-agent","Mozilla/5.0 (iPhone; CPU iPhone OS 15_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/19E258 Ariver/1.1.0 AliApp(AP/10.2.59.2500) Nebula WK RVKType(1) AlipayDefined(nt:WIFI,ws:390|780|3.0) AlipayClient/10.2.59.2500 Language/en Region/CN NebulaX/1.0.0")
	req1.Header.Add("accept","text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req1.Header.Add("connection","keep-alive")
	if err!=nil{
		return err
	}
	resp1,err:=c.client.Do(req1)
	if err!=nil{
		return err
	}
    defer resp1.Body.Close()
	raw,err:=ioutil.ReadAll(resp1.Body)
	urlVals,err:=conf.GetReqBody(raw)
	if err!=nil{
		return err
	}
	req2Body:=urlVals.Encode()
	//签到
	req2,err:=http.NewRequest(http.MethodPost,url2,strings.NewReader(req2Body))
	req2.Header.Add("origin","https://healthreport.zju.edu.cn")
	req2.Header.Add("user-agent","Mozilla/5.0 (iPhone; CPU iPhone OS 15_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/19E258 Ariver/1.1.0 AliApp(AP/10.2.59.2500) Nebula WK RVKType(1) AlipayDefined(nt:WIFI,ws:390|780|3.0) AlipayClient/10.2.59.2500 Language/en Region/CN NebulaX/1.0.0")
	req2.Header.Add("content-type","application/x-www-form-urlencoded; charset=UTF-8")
	req2.Header.Add("Accept-Language","en-US")
	req2.Header.Add("Accept-Encoding","gzip, deflate, br")
	req2.Header.Add("X-Requested-With","XMLHttpRequest") //请求是否由ajax发起
	req2.Header.Add("referer","https://healthreport.zju.edu.cn/ncov/wap/default/index")
	if err!=nil{
		return err
	}
	resp2,err:=c.client.Do(req2)
	if err!=nil{
		return err
	}
	defer resp2.Body.Close()
	rawMsg,err:=ioutil.ReadAll(resp2.Body)
	if err!=nil{
		return err
	}
	fmt.Println(string(rawMsg))
	return nil
}