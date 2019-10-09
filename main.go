package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/satng/sensors-gateway-grpc/pb"
	"github.com/satng/sensors-gateway-grpc/taosTool"
	"strings"

	_ "github.com/satng/sensors-gateway-grpc/taosSql"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":5012"
)

type server struct {
}

//CREATE TABLE acc_data (ts_sensor TIMESTAMP, ts TIMESTAMP,x FLOAT,y FLOAT,z FLOAT,t TIMESTAMP) TAGS(device_id BINARY(40),record_id BINARY(20))
//CREATE TABLE gyr_data (ts_sensor TIMESTAMP, ts TIMESTAMP,x FLOAT,y FLOAT,z FLOAT,t TIMESTAMP) TAGS(device_id BINARY(40),record_id BINARY(20))
//CREATE TABLE mag_data (ts_sensor TIMESTAMP, ts TIMESTAMP,x FLOAT,y FLOAT,z FLOAT,t TIMESTAMP) TAGS(device_id BINARY(40),record_id BINARY(20))

//CREATE TABLE gps_data (time TIMESTAMP,latitude FLOAT,longitude FLOAT,altitude FLOAT,accuracy FLOAT,bearing FLOAT,speed FLOAT,flag INT) TAGS(device_id BINARY(40),record_id BINARY(20))

const stmtAccSensorSql = `INSERT INTO d_%s_acc USING acc_data TAGS ('%s','%s') VALUES`
const stmtGyrSensorSql = `INSERT INTO d_%s_gyr USING gyr_data TAGS ('%s','%s') VALUES`
const stmtMagSensorSql = `INSERT INTO d_%s_mag USING mag_data TAGS ('%s','%s') VALUES`
const stmtGpsSensorSql = `INSERT INTO d_%s_gps USING gps_data TAGS ('%s','%s') VALUES`

func (s *server) DataPush(ctx context.Context, in *pb.SensorRequest) (*pb.SensorReply, error) {
	log.Printf("Received Header: %v,%v,%v", in.GetDeviceId(), in.GetRecordId(), in.GetSensorType())
	sensorType := in.GetSensorType()
	stmtSensorSqlHeader := ""
	isGps := false
	if sensorType == "accelerometer" {
		stmtSensorSqlHeader = fmt.Sprintf(stmtAccSensorSql, in.GetDeviceId(), in.GetDeviceId(), in.GetRecordId())
	} else if sensorType == "gyroscope" {
		stmtSensorSqlHeader = fmt.Sprintf(stmtGyrSensorSql, in.GetDeviceId(), in.GetDeviceId(), in.GetRecordId())
	} else if sensorType == "magnetic" {
		stmtSensorSqlHeader = fmt.Sprintf(stmtMagSensorSql, in.GetDeviceId(), in.GetDeviceId(), in.GetRecordId())
	} else if sensorType == "gps" {
		isGps = true
		stmtSensorSqlHeader = fmt.Sprintf(stmtGpsSensorSql, in.GetDeviceId(), in.GetDeviceId(), in.GetRecordId())
	}

	if !isGps {
		i := 0
		sql := bytes.Buffer{}
		for _, item := range in.GetDataStr() {
			//1565253030508;0.0000;9.8100;0.0000;57207591632172;1565253030488.63
			items := strings.Split(item, ";")
			if len(items) != 6 {
				log.Print("strings formatter: %v", item)
				continue
			}
			tsSensor := items[5]
			tsIndex := strings.LastIndex(tsSensor, ".")
			if tsIndex == -1 {
				tsIndex = strings.LastIndex(tsSensor, ",")
			}
			if len(tsSensor) < 13 {
				continue
			}

			sql.WriteString(fmt.Sprintf(`(%v,%v,%v,%v,%v,%v)`, tsSensor[:13], items[0], items[1], items[2], items[3], items[4]))
			i++

			if i%50 == 0 {
				stmtSensorSql := fmt.Sprintf(`%s %s;`, stmtSensorSqlHeader, sql.String())
				taosTool.Insert(stmtSensorSql)
				sql.Reset()
			}
		}
		if sql.Len() > 0 {
			stmtSensorSql := fmt.Sprintf(`%s %s;`, stmtSensorSqlHeader, sql.String())
			taosTool.Insert(stmtSensorSql)
			sql.Reset()
		}

	} else {
		i := 0
		sql := bytes.Buffer{}
		for _, item := range in.GetDataStr() {
			//30.516360;114.359117;0.85;19.64;248.00;0.77;1564562033799;0
			items := strings.Split(item, ";")
			if len(items) != 8 {
				log.Print("strings formatter: %v", item)
				continue
			}
			sql.WriteString(fmt.Sprintf(`(%v,%v,%v,%v,%v,%v,%v,%v)`, items[6], items[0], items[1], items[2], items[3], items[4], items[5], items[7]))
			i++
			if i%50 == 0 {
				stmtSensorSql := fmt.Sprintf(`%s %s;`, stmtSensorSqlHeader, sql.String())
				taosTool.Insert(stmtSensorSql)
				sql.Reset()
			}

		}
		if sql.Len() > 0 {
			stmtSensorSql := fmt.Sprintf(`%s %s;`, stmtSensorSqlHeader, sql.String())
			taosTool.Insert(stmtSensorSql)
			sql.Reset()
		}
	}
	return &pb.SensorReply{Message: "Hello " + in.DeviceId}, nil
}

func main() {

	//test taos db
	taosTool.InitDB()

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSensorsServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	taosTool.CloseDB()
	log.Println("server over...")
}
