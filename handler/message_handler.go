package handler

import (
	"fmt"
	"github.com/eatmoreapple/openwechat"
	"golang.org/x/time/rate"
	"time"
)

var (
	friendsLimit = rate.NewLimiter(rate.Every(2*time.Second), 1)
)

func TextMsgHandler(ctx *openwechat.MessageContext) {
	// 处理好友消息

	if ctx.IsSystem() || !ctx.IsText() || ctx.IsSendByGroup() {
		fmt.Println("仅处理好友信息:TextMsgHandler")
		return
	}

	content := ctx.Content

	sender, _ := ctx.Sender()

	fmt.Println(sender)
	sendName := sender.DisplayName
	if sendName == "" {
		sendName = sender.NickName
	}
	ctx.ReplyText(sendName + " 暂无语料分析能力:" + content)

	ctx.Abort()
	ctx.AsRead()
}
