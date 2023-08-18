package handler

import (
	"fmt"
	"github.com/eatmoreapple/openwechat"
	"golang.org/x/time/rate"
	"strings"
	"time"
)

var (
	groupLimit = rate.NewLimiter(rate.Every(2*time.Second), 1)
)

func GroupTextMsgHandler(ctx *openwechat.MessageContext) {
	fmt.Println(ctx.MsgType.String())
	// 仅处理群组消息
	if ctx.IsSystem() || !ctx.IsText() || !ctx.IsAt() || !ctx.IsSendByGroup() {
		//fmt.Println("消息非群组内消息,不支持此 GroupTextMsgHandler 处理")
		return
	}
	sender, _ := ctx.Sender()
	receiver := sender.MemberList.SearchByUserName(1, ctx.ToUserName)

	if receiver != nil {
		if !groupLimit.Allow() {
			ctx.Abort()
			return
		}

		msgContent := strings.TrimSpace(openwechat.FormatEmoji(ctx.Content))

		// 群组内消息需要去除@+接收者名字
		var atName string
		receiverName := receiver.First().DisplayName
		if receiverName == "" {
			receiverName = receiver.First().NickName
		}
		recName := openwechat.FormatEmoji(receiverName)

		if strings.Contains(msgContent, "\u2005") {
			atName = "@" + recName + "\u2005"
		} else {
			atName = "@" + recName
		}

		// 去除 @+name 的内容
		content := strings.TrimSpace(strings.TrimPrefix(msgContent, atName))

		sendByG, _ := ctx.SenderInGroup()

		sendName := sendByG.DisplayName
		if sendName == "" {
			sendName = sendByG.NickName
		}
		ctx.ReplyText("@" + sendName + " received message: [" + content + "]")

		ctx.Abort()
		ctx.AsRead()
	}
}
