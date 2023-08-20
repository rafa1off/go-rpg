package setup

import (
	"net"
)

const (
	grpcPort = "8000"
	restPort = "8080"
)

func initListener(c chan<- net.Listener, port string) {
	lis, err := net.Listen("tcp", "0.0.0.0:"+port)
	if err != nil {
		c <- nil
		return
	}
	c <- lis
}
