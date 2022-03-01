package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf  // restful 结构体
	Auth struct { // 权限校验 结构体
		AccessSecret string
		AccessExpire int64
	}
	UserRpc zrpc.RpcClientConf
}
