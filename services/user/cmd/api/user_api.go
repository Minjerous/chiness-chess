package main

import (
	"chess-common/rep"
	"chess-user/cmd/api/internal/config"
	"chess-user/cmd/api/internal/handler"
	"chess-user/cmd/api/internal/handler/backend"
	"chess-user/cmd/rpc/proto/user"
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type loginCreds struct {
	Username string
	Password string
}

func (c *loginCreds) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	return map[string]string{
		"username": c.Username,
		"password": c.Password,
	}, nil
}

func (c *loginCreds) RequireTransportSecurity() bool {
	return true
}

var Conn *grpc.ClientConn
var UserClient user.UserServiceClient

func main() {
	config.PareConfig()
	//certificate, err := tls.LoadX509KeyPair("./user.pem", "./user.key") 懒得弄CA了
	UserCfg := config.GetUserCfg()

	Conn, err := grpc.Dial(UserCfg.UserRpc.Hosts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	backend.GetRpcClis(backend.RpcClis{UserSrvCli: user.NewUserServiceClient(Conn)})
	defer Conn.Close()

	router := gin.Default()
	router.Use(rep.Recovery())
	routerEngine(router)
	//router.RunTLS(":10001", "./config/app.pem", "./config/app.key")
	router.Run(":8080")
}
func routerEngine(engine *gin.Engine) {
	new(handler.UserHandler).Router(engine)
}
