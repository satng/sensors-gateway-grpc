package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/satng/sensors-gateway-grpc/pb"
	"strings"

	//_ "github.com/satng/sensors-gateway-grpc/taosSql"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":5012"
)

type server struct {
}

const stmtSensorSql = `INSERT INTO d_%s USING sensors_data TAGS ('%s','%s','%s') VALUES`

func (s *server) DataPush(ctx context.Context, in *pb.SensorRequest) (*pb.SensorReply, error) {
	log.Printf("Received Header: %v,%v,%v", in.GetDeviceId(), in.GetRecordId(), in.GetSensorType())

	if in.GetSensorType() != "gps" {
		stmtSensorSqlHeader := fmt.Sprintf(stmtSensorSql, in.GetDeviceId(), in.GetDeviceId(), in.GetRecordId(), in.GetSensorType())
		var stmtSensorSqlContext bytes.Buffer
		for _, item := range in.GetDataStr() {
			//1565253030508;0.0000;9.8100;0.0000;57207591632172;1565253030488.63
			items := strings.Split(item, ";")
			ts_sensor := items[5]
			stmtSensorSqlContext.WriteString(fmt.Sprintf(`(%v,%v,%v,%v,%v,%v)`, ts_sensor[:strings.LastIndex(ts_sensor, ".")], items[0], items[1], items[2], items[3], items[4]))
		}

		stmtSensorSql := fmt.Sprintf(`%s %s;`, stmtSensorSqlHeader, stmtSensorSqlContext.String())
		log.Printf("Received Data: %v", stmtSensorSql)
	}

	return &pb.SensorReply{Message: "Hello " + in.DeviceId}, nil
}

func main() {

	//test taos db
	//if taosTool.Test() {
	//	return
	//}

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
