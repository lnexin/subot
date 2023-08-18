package main

import (
	"fmt"
	"github.com/eatmoreapple/openwechat"
	"github.com/mdp/qrterminal/v3"
	"os"
	"path/filepath"
	"susu-wechat-bot/handler"
)

func main() {
	bot := openwechat.DefaultBot(openwechat.Desktop)

	// 注册消息处理函数
	//bot.MessageHandler = func(msg *openwechat.Message) {
	//	fmt.Println(msg.Content)
	//	if msg.IsText() && msg.Content == "ping" {
	//		msg.ReplyText("pong")
	//	}
	//}

	// 注册登录二维码回调
	bot.UUIDCallback = func(uuid string) {
		fmt.Println(openwechat.GetQrcodeUrl(uuid))
		qrterminal.Generate("https://login.weixin.qq.com/l/"+uuid, qrterminal.L, os.Stdout)
	}
	// 创建容器存储对象
	reloadStorage := openwechat.NewFileHotReloadStorage(filepath.Join(os.Getenv("DATA"), "storage.json"))
	defer reloadStorage.Close()

	// 热登录
	if err := bot.HotLogin(reloadStorage, openwechat.NewRetryLoginOption()); err != nil {
		// 如果登录出错并且错误信息不为nil,输出
		fmt.Print(err)
		return
	}

	//bot.PushLogin(reloadStorage, openwechat.NewRetryLoginOption())

	// 获取登录的用户
	self, err := bot.GetCurrentUser()
	if err != nil {
		fmt.Println(err)
		return
	}

	user, err := bot.GetCurrentUser()
	fmt.Println(user)
	// 获取所有的好友
	friends, err := self.Friends()
	fmt.Println(friends, err)
	//for _, f := range friends.AsMembers() {
	//	fmt.Println(f.ID())
	//	fmt.Println(f.NickName, f.UserName, f.DisplayName)
	//}
	//fmt.Println(friends, err)

	// 获取所有的群组
	groups, err := self.Groups()
	for _, g := range groups.AsMembers() {
		fmt.Println("群:", g.NickName, g.DisplayName)
	}

	//
	dispatcher := openwechat.NewMessageMatchDispatcher()
	dispatcher.SetAsync(true)

	dispatcher.OnText(handler.TextMsgHandler)
	dispatcher.OnGroup(handler.GroupTextMsgHandler)

	bot.MessageHandler = dispatcher.AsMessageHandler()
	// 阻塞主goroutine, 直到发生异常或者用户主动退出
	bot.Block()
}
