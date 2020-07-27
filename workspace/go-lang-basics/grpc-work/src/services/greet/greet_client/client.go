package main

import (
	"../greetpb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func main () {
	fmt.Println("Hi i'm in client... ")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Sorry client cannot talk to server: %v: ", err)
	}
	defer conn.Close()

	c := greetpb.NewGreetServiceClient(conn)

	callGreet(c)
	callFullName(c)
}

func  callGreet(c greetpb.GreetServiceClient)  {
	fmt.Println("in Call Greet Function... ")

	req:= &greetpb.GreetRequest {
		Greeting: &greetpb.Greeting {
			FistName :"roh",
			LastName:"bha",
		},
	}

	res, err := c.Greet(context.Background(), req)

	if err!= nil {
		log.Fatalf("Error While calleding Greet : %v ", err)
	}

	log.Println("Response From the Greet ", res.Result)
}

func callFullName(c greetpb.GreetServiceClient) {
	gr := &greetpb.Greeting{
		FistName:"fstn",
		LastName:"lastn",
	}

	req := &greetpb.GreetRequest{
		Greeting:gr,
	}

	resp, err := c.GreetFullName(context.Background(), req)
	if err!=nil {
		log.Fatalf("some error %v", err)
	}

	log.Println("response rerce ", resp.Greeting)
}












