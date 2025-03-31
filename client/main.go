package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/abhijeet-singh800/grpc-test/grpcPkg/calc"
	"github.com/abhijeet-singh800/grpc-test/grpcPkg/crud"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Cat Establish the connection Error: %v", err)
	}
	defer conn.Close()

	calcClient := calc.NewCalcServiceClient(conn)
	crudClient := crud.NewCrudServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Calls for Calculator Service

	req := &calc.Two{First: 32, Sec: 22}
	// Add Request
	fmt.Printf("%d, %d", 32, 22)
	resp, err := calcClient.Add(ctx, req)
	if err != nil {
		log.Fatalf("Cant Add Error %v", err)
	}
	fmt.Printf("\nSum %d", resp.Num)

	// Product Request
	resp, err = calcClient.Mul(ctx, req)
	if err != nil {
		log.Fatalf("Cant Multiply Err %v", err)
	}
	fmt.Printf("\nProduct %d", resp.Num)

	// Subtract Request
	resp, err = calcClient.Sub(ctx, req)
	if err != nil {
		log.Fatalf("Cant Subtract Error %v ", err)
	}
	fmt.Printf("\nSubtract %d", resp.Num)

	//  Divide Request
	resp, err = calcClient.Div(ctx, req)
	if err != nil {
		log.Fatalf("Cant Divide Error %v", err)
	}
	fmt.Printf("\nDivided : %d", resp.Num)

	// Sum Up to Request
	reqNum := &calc.NumList{Nums: []int32{1, 2, 3, 4, 5, 6}}
	resp, err = calcClient.SumUpTo(ctx, reqNum)
	if err != nil {
		log.Fatalf("Cant Sum Error %v", err)
	}
	fmt.Printf("\nSumed %d \n", resp.Num)

	// Count Up to
	reqOne := &calc.One{Num: 20}
	rsp, err := calcClient.CountUpTo(ctx, reqOne)
	if err != nil {
		log.Fatalf("Cant Count Up to Error: %v", err)
	}
	for _, num := range rsp.Nums {
		fmt.Printf("  %v", num)
	}

	// Calls for Crud Service

	// Get User
	req_id := &crud.Id{Id: 32}
	user, err := crudClient.GetUser(ctx, req_id)
	if err != nil {
		log.Fatalf("Cant Get User Error: %v", err)
	}
	fmt.Printf("\n Get User\n Name: %v \n Emp Id: %v \n Gender: %v \n Membership: %v", user.Name, user.EmpId, user.Gender, user.Premium)

	// New User
	new_usr := &crud.UserSpec{
		Name:    "Navjeet Singh",
		Id:      55,
		EmpId:   3,
		Gender:  crud.Gender_Male,
		Premium: crud.Membership_Premium,
	}
	new_id, err := crudClient.NewUser(ctx, new_usr)
	if err != nil {
		log.Fatalf("Cant Create User Error: %v", err)
	}
	fmt.Printf("\n Created User \n New User Id : %v", new_id.Id)

	// Delete User
	del_id := &crud.Id{
		Id: 52,
	}
	deleted_usr, err := crudClient.DeleteUser(ctx, del_id)
	if err != nil {
		log.Fatalf("Cant delete user Error: %v", err)
	}
	fmt.Printf("\n Deleted User \n Name: %v\n Emp Id: %v\n Gender: %v\n Membership: %v", deleted_usr.Name, deleted_usr.EmpId, deleted_usr.Gender, deleted_usr.Premium)

	// Set User
	setUsr := &crud.UserSpec{
		Id:      1,
		Name:    "Abhijeet Singh",
		EmpId:   11,
		Gender:  crud.Gender_Male,
		Premium: crud.Membership_Premium,
	}
	updatedUsr, err := crudClient.SetUser(ctx, setUsr)
	if err != nil {
		log.Fatalf("Cant update the User Error : %v ", err)
	}
	fmt.Printf("\n Updated User \n Name: %v\n Emp Id: %v\n Gender: %v\n Membership: %v", updatedUsr.Name, updatedUsr.EmpId, updatedUsr.Gender, updatedUsr.Premium)

	// Streaming the Data
	nil_req := &crud.Nil{}
	stream, err := crudClient.GetAll(context.Background(), nil_req)
	if err != nil {
		log.Fatalf("Cant Start Streaming Error: %v", err)
	}

	fmt.Printf("\n Straming Started")
	for {
		resp_str, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Cant continue streaming Error: %v", err)
		}
		fmt.Printf("\n Id: %v \n Name: %v\n Emp Id: %v\n Gender: %v\n Membership: %v", resp_str.Id, resp_str.Name, resp_str.EmpId, resp_str.Gender, resp_str.Premium)
	}
	fmt.Printf("\n Streaming Ended")
}
