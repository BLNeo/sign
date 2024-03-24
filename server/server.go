package server

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"net"
	"sign/conf"
	"sign/models"
	signPb "sign/proto/sign"
	"sign/router"
	"sign/service/sign_rpc_service"
	"sign/tool/log"
	"sign/tool/mysql"
	"sign/tool/redis"
	"sign/tool/util"
)

type Server struct {
	gin  *gin.Engine  // 路由服务
	grpc *grpc.Server // grpc服务
	db   *gorm.DB     // db数据库服务
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Init(conf *conf.Config) error {
	var err error
	log.Init(conf.App.ServerName)

	s.db, err = mysql.InitEngine(conf.Mysql)
	if err != nil {
		return err
	}

	// 赋值到models包的db
	models.InitDb(s.db)
	// 建表同步
	models.CreateTable()

	// redis
	err = redis.InitClient(conf.Redis)
	if err != nil {
		return err
	}

	// jwt
	util.InitJwtSecret()
	// gin
	gin.SetMode(conf.Http.Mode)
	s.gin = gin.Default()
	router.InitRouter(s.gin)
	// validator // 入参校验翻译器
	err = util.InitTrans()
	if err != nil {
		return err
	}
	// grpc
	s.grpc = grpc.NewServer()

	log.Logger.Info("server init success")
	return nil
}

func (s *Server) GinRun() error {
	return s.gin.Run(conf.Conf.Http.Port)
}

func (s *Server) GrpcRun() {
	signPb.RegisterSignServer(s.grpc, sign_rpc_service.NewRealSignRpc())
	listen, err := net.Listen("tcp", ":"+conf.Conf.Grpc.Port)
	if err != nil {
		log.Logger.Error("GrpcRun Listen TCP err:" + err.Error())
	}

	err = s.grpc.Serve(listen)
	if err != nil {
		log.Logger.Error("GrpcRun Serve listen err:" + err.Error())
	}

}
