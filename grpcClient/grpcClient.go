package main

import (
	"log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	proto "code.nurture.farm/BloodBankSystemService/zerotouch/golang/proto/BloodBankSystemService/BloodBankSystemService"
)

const (
	LOCAL_URL = ":6000"
)

var ENV = LOCAL_URL

func TestFindPassword() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(ENV, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewBloodBankSystemServiceClient(conn)

	request := &proto.FindPasswordRequest{
	    //Set your request here
	}

	response, err := c.ExecuteFindPassword(context.Background(), request)
	if err != nil {
		log.Fatalf("Error when calling FindPasswordRequest: %s", err)
	}
	log.Println(response)
}

func TestAddUser() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(ENV, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewBloodBankSystemServiceClient(conn)

	request := &proto.AddUserRequest{
	    //Set your request here
	}

	response, err := c.ExecuteAddUser(context.Background(), request)
	if err != nil {
		log.Fatalf("Error when calling AddUserRequest: %s", err)
	}
	log.Println(response)
}

func TestAddUserBulk() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(ENV, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewBloodBankSystemServiceClient(conn)

	request := &proto.BulkAddUserRequest{
	    //Set your request here
	}

	response, err := c.ExecuteAddUserBulk(context.Background(), request)
	if err != nil {
		log.Fatalf("Error when calling BulkAddUserRequest: %s", err)
	}
	log.Println(response)
}

func TestFindBlood() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(ENV, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewBloodBankSystemServiceClient(conn)

	request := &proto.FindBloodRequest{
	    //Set your request here
	}

	response, err := c.ExecuteFindBlood(context.Background(), request)
	if err != nil {
		log.Fatalf("Error when calling FindBloodRequest: %s", err)
	}
	log.Println(response)
}

func TestAddBlood() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(ENV, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewBloodBankSystemServiceClient(conn)

	request := &proto.AddBloodRequest{
	    //Set your request here
	}

	response, err := c.ExecuteAddBlood(context.Background(), request)
	if err != nil {
		log.Fatalf("Error when calling AddBloodRequest: %s", err)
	}
	log.Println(response)
}

func TestAddBloodBulk() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(ENV, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewBloodBankSystemServiceClient(conn)

	request := &proto.BulkAddBloodRequest{
	    //Set your request here
	}

	response, err := c.ExecuteAddBloodBulk(context.Background(), request)
	if err != nil {
		log.Fatalf("Error when calling BulkAddBloodRequest: %s", err)
	}
	log.Println(response)
}



