package qq

import (
	"github.com/Mrs4s/MiraiGo/client"
	"qq_bot/utils"
)

type QQHelper struct {
	client *client.QQClient
}

func NewQQHelper(uid int64, passwd string) *QQHelper {
	md5Passwd := utils.Md5(passwd)
	return &QQHelper{
		client: client.NewClientMd5(uid, md5Passwd),
	}
}

func (o *QQHelper) Groups() {
}
