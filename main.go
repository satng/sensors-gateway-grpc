package main

import (
	"context"
	"github.com/satng/sensors-gateway-grpc/pb"
	_ "github.com/satng/sensors-gateway-grpc/taosSql"
	"github.com/satng/sensors-gateway-grpc/taosTool"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":5012"
)

type server struct {
}

func (s *server) DataPush(ctx context.Context, in *pb.SensorRequest) (*pb.SensorReply, error) {
	log.Printf("Received Header: %v,%v,%v", in.GetDeviceId(), in.GetRecordId(), in.GetSensorType())
	//log.Printf("Received Data: %v", in.GetDataStr())
	return &pb.SensorReply{Message: "Hello " + in.DeviceId}, nil
}

func main() {

	//test taos db
	if taosTool.Test() {
		return
	}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSensorsServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	log.Println("server over...")
}
