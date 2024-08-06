package bootstrap

import (
	"github.com/aimerneige/openaibot/internal/config"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/driver"

	_ "github.com/aimerneige/openaibot/internal/openai"
)

func StartBot() {
	common := config.Conf.Common
	zero.RunAndBlock(&zero.Config{
		NickName:      common.NickName,
		CommandPrefix: common.CommandPrefix,
		SuperUsers:    common.SuperUsers,
		Driver: []zero.Driver{
			// 正向 WS
			driver.NewWebSocketClient(common.WSServer, common.WSToken),
		},
	}, nil)
}
