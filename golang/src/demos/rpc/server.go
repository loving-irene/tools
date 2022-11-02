package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"rpc/utils"
)

type MathService struct {
}

func (m *MathService) Multiply(args *utils.Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (m *MathService) Divide(args *utils.Args, reply *int) error {
	if args.B == 0 {
		return errors.New("除数不能为0")
	}
	*reply = args.A / args.B
	return nil
}

type ServiceHandler struct {
}

func (serviceHandler *ServiceHandler) GetName(id int, item *Items) error {
	fmt.Printf("receive GetName Call,id:%d\n", id)
	item.Id = id
	item.Name = "lovan"
	return nil
}

func (serviceHandler *ServiceHandler) SaveName(name string, item *Items) error {
	log.Println(item)
	log.Println(&item)
	fmt.Printf("receive SaveName Call,name:%s\n", name)
	item.Name = name
	return nil
}

func main() {
	server := rpc.NewServer()
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("err msg:", err)
	}
	defer listener.Close()

	serviceHandler := &ServiceHandler{}
	err = server.Register(serviceHandler)
	if err != nil {
		fmt.Println("err msg:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("err msg:", err)
		}
		go server.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
	//math:=new(MathService)
	//
	//rpc.Register(math)
	//rpc.HandleHTTP()
	//
	//listener,err :=net.Listen("tcp",":8080")
	//if err!=nil{
	//	log.Fatal("启动失败:",err)
	//}
	//err=http.Serve(listener,nil)
	//if err!=nil{
	//	log.Fatal("启动http服务失败",err)
	//}
}
