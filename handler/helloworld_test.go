/*
   @Author:huolun
   @Date:2021/5/13
   @Description
*/
package handler

import (
	"context"
	"fmt"
	pb "helloworld/proto"
	"testing"
	"time"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/store"
)

func TestHelloworld_Call(t *testing.T) {
	svr := service.New(service.Name("helloworld"))
	client := pb.NewHelloworldService("helloworld", svr.Client())

	rsp, err := client.Call(context.Background(), &pb.Request{
		Name: "liuzhe",
	})
	if err != nil {
		fmt.Println("Error calling helloworld: ", err)
		return
	}

	// 打印响应内容
	fmt.Println("Response: ", rsp.Msg)
	// let's delay the process for exiting for reasons you'll see below
	time.Sleep(time.Second * 1)
	svr.Init()
	records, err := store.Read("mykey")
	if err != nil {
		fmt.Println("Error reading from store: ", err)
	}

	if len(records) == 0 {
		fmt.Println("No records")
	}
	for _, record := range records {
		fmt.Printf("key: %v, value: %v\n", record.Key, string(record.Value))
	}

}
