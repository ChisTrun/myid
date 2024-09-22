package myid

import (
	myid "github.com/ChisTrun/myid/api"
)

func NewServer() myid.MyIdServer {
	return &myidServer{}
}

type myidServer struct {
	myid.UnimplementedMyIdServer
}
