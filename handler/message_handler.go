package handler

import (
	"fmt"
	"github.com/eatmoreapple/openwechat"
	"golang.org/x/time/rate"
	"regexp"
	"time"
)

var (
	friendsLimit = rate.NewLimiter(rate.Every(2*time.Second), 1)
)

func TextMsgHandler(ctx *openwechat.MessageContext) {
	// 处理好友消息

	if ctx.IsSystem() || !ctx.IsText() || ctx.IsSendByGroup() {
		//fmt.Println("仅处理好友信息:TextMsgHandler")
		return
	}

	content := ctx.Content
	sender, _ := ctx.Sender()

	fmt.Println(sender)
	sendName := sender.DisplayName
	if sendName == "" {
		sendName = sender.NickName
	}

	r, err := regexp.MatchString("淘宝", content)
	if err != nil {
		fmt.Println(err)
	}

	if r {
		// 转链功能待添加
		ctx.ReplyText(sendName + " 淘宝转链功能待添加")
	}

	//ctx.Abort()
	//ctx.AsRead()
}
