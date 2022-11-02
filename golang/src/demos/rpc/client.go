package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc/jsonrpc"
	"time"
)

func main() {

	conn, err := net.DialTimeout("tcp", "localhost:8080", 30*time.Second)
	if err != nil {
		fmt.Println("err msg:", err)
	}

	defer conn.Close()

	client := jsonrpc.NewClient(conn)
	var item Items
	client.Call("ServiceHandler.GetName", 1, &item)
	log.Printf("ServiceHandler.GetName 返回结：%v\n", item)

	//var resp Response
	item = Items{2, "学员均"}
	log.Println(item)
	log.Println(&item)
	client.Call("ServiceHandler.SaveName", "haha", &item)
	log.Printf("ServiceHandler.GetName 返回结：%v\n", item)

	//var serverAddress="localhost"
	//client,err:=rpc.DialHTTP("tcp",serverAddress+":8080")
	//if err!=nil{
	//	log.Fatal("异常：",err)
	//}
	//
	//args := new(utils.Args)
	//args.A=10
	//args.B=10
	//
	////args := &utils.Args{10,10}
	//var reply int
	//err=client.Call("MathService.Multiply",args,&reply)
	//if err!=nil{
	//	log.Fatal("异常：",err)
	//}
	//fmt.Printf("%d*%d=%d\n",args.A,args.B,reply)
	//
	//divideCall:=client.Go("MathService.Divide",args,&reply,nil)
	//for{
	//	select {
	//	case <-divideCall.Done:
	//		fmt.Printf("%d/%d=%d\n",args.A,args.B,reply)
	//		return
	//	}
	//}

}
