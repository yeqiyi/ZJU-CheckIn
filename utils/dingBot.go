package utils

import (
	"gopkg.in/ini.v1"
	"fmt"

	"github.com/blinkbean/dingtalk"
)

type Bot struct{
	bot *dingtalk.DingTalk
}

func GetBot(configPath string)*Bot{
	cfg,_:=ini.Load(configPath)

	token:=cfg.Section("DingBot").Key("Token").String()
	secret:=cfg.Section("DingBot").Key("Secret").String()
	fmt.Println(token,secret)
	return &Bot{
		dingtalk.InitDingTalkWithSecret(token,secret),
	}
}

func(b Bot)SendMsg(content string)error{
	err:=b.bot.SendTextMessage(content)
	return err
}