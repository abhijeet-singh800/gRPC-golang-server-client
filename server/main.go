package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	"github.com/abhijeet-singh800/grpc-test/grpcPkg/calc"
	"github.com/abhijeet-singh800/grpc-test/grpcPkg/crud"
	"github.com/abhijeet-singh800/grpc-test/sql"
)

// Calculator Service
type calcServer struct {
	calc.UnimplementedCalcServiceServer
}

func (s *calcServer) Add(ctx context.Context, req *calc.Two) (*calc.One, error) {
	res := req.First + req.Sec
	return &calc.One{Num: res}, nil
}

func (s *calcServer) Sub(ctx context.Context, req *calc.Two) (*calc.One, error) {
	res := req.First - req.Sec
	return &calc.One{Num: res}, nil
}

func (s *calcServer) Mul(ctx context.Context, req *calc.Two) (*calc.One, error) {
	res := req.First * req.Sec
	return &calc.One{Num: res}, nil
}

func (s *calcServer) Div(ctx context.Context, req *calc.Two) (*calc.One, error) {
	if req.Sec == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Second number cant be zero")
	}
	res := req.First / req.Sec
	return &calc.One{Num: res}, nil
}

func (s *calcServer) SumUpTo(ctx context.Context, req *calc.NumList) (*calc.One, error) {
	var sum int32
	for _, num := range req.Nums {
		sum = sum + num
	}
	return &calc.One{Num: sum}, nil
}

func (s *calcServer) CountUpTo(ctx context.Context, req *calc.One) (*calc.NumList, error) {
	num := []int32{}
	for i := 0; i < int(req.Num); i++ {
		num = append(num, int32(i))
	}
	return &calc.NumList{Nums: num}, nil
}

// CRUD Service
type crudServer struct {
	crud.UnimplementedCrudServiceServer
}

func (s *crudServer) GetUser(ctx context.Context, req *crud.Id) (*crud.UserSpec, error) {
	usr, err := sql.Get(int(req.Id))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Cant get users form Database Error: %v", err)
	}

	resp := &crud.UserSpec{
		Id:      int32(usr.Id),
		Name:    usr.Name,
		EmpId:   int32(usr.EmpID),
		Gender:  toGenderEnum(usr.Gender),
		Premium: toMembershipEnum(usr.Premium),
	}

	return resp, nil
}

func (s *crudServer) DeleteUser(ctx context.Context, req *crud.Id) (*crud.UserSpec, error) {
	usr, err := sql.Get(int(req.Id))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Cant get users form Database Error: %v", err)
	}
	resp := &crud.UserSpec{
		Id:      int32(usr.Id),
		Name:    usr.Name,
		EmpId:   int32(usr.EmpID),
		Gender:  toGenderEnum(usr.Gender),
		Premium: toMembershipEnum(usr.Premium),
	}

	err = sql.Delete(int(req.Id))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Cant delete Error: %v", err)
	}

	return resp, nil
}

func (s *crudServer) NewUser(ctx context.Context, req *crud.UserSpec) (*crud.Id, error) {
	usr := &sql.User{
		Id:      int(req.Id),
		Name:    req.Name,
		EmpID:   int(req.EmpId),
		Gender:  toGender(req.Gender),
		Premium: int(req.Premium),
	}
	id, err := sql.Add(usr)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Cant create user in Database Error: %v", err)
	}
	resp := &crud.Id{
		Id: int32(id),
	}
	return resp, nil
}

func (s *crudServer) SetUser(ctx context.Context, req *crud.UserSpec) (*crud.UserSpec, error) {
	usr := sql.User{
		Id:      int(req.Id),
		Name:    req.Name,
		EmpID:   int(req.EmpId),
		Gender:  toGender(req.Gender),
		Premium: int(req.Premium),
	}

	err := sql.Set(&usr)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Cant Update user Error %v", err)
	}

	updated_usr, err := sql.Get(int(req.Id))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Cant get users form Database Error: %v", err)
	}
	resp := &crud.UserSpec{
		Id:      int32(updated_usr.Id),
		Name:    updated_usr.Name,
		EmpId:   int32(updated_usr.EmpID),
		Gender:  toGenderEnum(updated_usr.Gender),
		Premium: toMembershipEnum(updated_usr.Premium),
	}

	return resp, nil

}

func (s *crudServer) GetAll(req *crud.Nil, stream crud.CrudService_GetAllServer) error {
	sqlData, err := sql.ReadAll()
	if err != nil {
		return status.Errorf(codes.Internal, "Cant Get Users form Database Error %v", err)
	}
	for _, usr := range sqlData {
		res := &crud.UserSpec{
			Id:      int32(usr.Id),
			EmpId:   int32(usr.EmpID),
			Name:    usr.Name,
			Gender:  toGenderEnum(usr.Gender),
			Premium: toMembershipEnum(usr.Premium),
		}

		if err = stream.Send(res); err != nil {
			return status.Errorf(codes.Internal, "Cant Stream Error: %v", err)
		}
		time.Sleep(time.Second / 2)
	}
	return nil
}

func main() {
	listner, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Cant Listen to Port 50051 Error : %v", err)
	}

	server := grpc.NewServer()
	calc.RegisterCalcServiceServer(server, &calcServer{})
	crud.RegisterCrudServiceServer(server, &crudServer{})

	reflection.Register(server)

	fmt.Printf("Serving to Port :50051")
	if err := server.Serve(listner); err != nil {
		log.Fatalf("Cant serve to port 50051 Error: %v", err)
	}
}

// Helper Functions

func toGenderEnum(gender string) crud.Gender {
	if gender == "Male" {
		return crud.Gender_Male
	} else {
		return crud.Gender_Female
	}
}

func toGender(enum crud.Gender) string {
	if enum == crud.Gender_Male {
		return "Male"
	} else {
		return "Female"
	}
}

func toMembershipEnum(num int) crud.Membership {
	if num == 1 {
		return crud.Membership_Premium
	} else {
		return crud.Membership_Free
	}
}
