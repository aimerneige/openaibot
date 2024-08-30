package openai

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/aimerneige/openaibot/internal/openai/query"
	"github.com/aimerneige/openaibot/internal/pkg/common"
	"github.com/sirupsen/logrus"
	zero "github.com/wdvxdr1123/ZeroBot"
	"gopkg.in/yaml.v3"
)

const configFilePath = "./config/openai.yaml"

const (
	QueryEnvOnline = "online"
	QueryEnvTest   = "test"
)

type OpenaiConfig struct {
	Secret OpenaiSecret `yaml:"secret"`
}

type OpenaiSecret struct {
	Account string `yaml:"account"`
	AppId   string `yaml:"appid"`
	Token   string `yaml:"token"`
	AesKey  string `yaml:"aeskey"`
}

var accessToken string

func init() {
	confData, err := os.ReadFile(configFilePath)
	if err != nil {
		logrus.Errorln("[openai]", "Fail to read config file", err)
		return
	}
	var config OpenaiConfig
	if err := yaml.Unmarshal(confData, &config); err != nil {
		logrus.Errorln("[openai]", "Fail to unmarshal config data", err)
		return
	}
	secret := config.Secret
	go func() {
		for {
			accessToken = query.GetToken(secret.Account, secret.AppId, secret.Token)
			time.Sleep(time.Hour)
		}
	}()

	engine := zero.New()
	engine.OnMessage(zero.OnlyToMe).
		SetPriority(2).
		SetBlock(true).Handle(func(ctx *zero.Ctx) {
		userQuery := strings.TrimSpace(ctx.MessageString())
		response := query.SendQueryRequest(query.ApiReq{
			Query:    userQuery,
			Env:      QueryEnvTest,
			UserName: ctx.Event.Sender.NickName,
			Avatar:   fmt.Sprintf("https://q2.qlogo.cn/headimg_dl?dst_uin=%d&spec=100", ctx.Event.Sender.ID),
			Userid:   fmt.Sprintf("qq_%d", ctx.Event.Sender.ID),
		}, accessToken, secret.Token, secret.AesKey)
		if response == "" {
			return
		}
		response = strings.Replace(response, "LINE_BREAK", "\n", -1)
		ctx.Send(response)
	})
	engine.UseMidHandler(common.DefaultSpeedLimit)
}
