package setup

import (
	app "go-rpg/app/services"
	"go-rpg/proto"
	"net"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const defaultPort int = 8000

type Server struct {
	engine *grpc.Server
}

func NewServer() *Server {
	return &Server{}
}

func initListener(c chan<- net.Listener) {
	lis, err := net.Listen("tcp", "0.0.0.0:"+strconv.Itoa(defaultPort))
	if err != nil {
		c <- nil
		return
	}
	c <- lis
}

func (s *Server) Run() {
	lis := make(chan net.Listener)
	go initListener(lis)

	srv := grpc.NewServer()

	proto.RegisterCharactersServer(srv, app.NewCharServer())
	proto.RegisterRaidsServer(srv, nil)

	reflection.Register(srv)

	s.engine = srv

	go Logger.Info("server listening on: http://localhost:" + strconv.Itoa(defaultPort))
	Logger.Sugar().Fatal(s.engine.Serve(<-lis))
}
