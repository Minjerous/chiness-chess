package backend

import "chess-user/cmd/rpc/proto/user"

type Client user.UserServiceClient

type RpcClis struct {
	UserSrvCli user.UserServiceClient
}

func GetRpcClis(clis RpcClis) {
	UserClient = clis.UserSrvCli
}

var UserClient user.UserServiceClient

func GetUserClient() user.UserServiceClient {
	return UserClient
}
