package dal

import (
	"github.com/DavidWisdom/tiktok-backend/app/user/biz/dal/mysql"
	"github.com/DavidWisdom/tiktok-backend/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
