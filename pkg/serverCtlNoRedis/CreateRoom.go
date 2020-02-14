package serverCtlNoRedis

import (
	cm "ULZRoomService/pkg/common"
	pb "ULZRoomService/proto"
	"context"
	"log"
	"time"
)

// CreateRoom :

func (this *ULZRoomServiceBackend) CreateRoom(ctx context.Context, req *pb.RoomCreateReq) (*pb.Room, error) {
	cm.PrintReqLog(ctx, req)
	start := time.Now()
	this.mu.Lock()
	defer func() {
		this.mu.Unlock()
		elapsed := time.Since(start)
		log.Printf("Quit-Room took %s", elapsed)
	}()
	// for loop it
	tmptime := time.Now().String() + req.Host.GetId()
	var f = ""
	for {
		f = cm.HashText(tmptime)
		_, xist := this.Roomlist["Rm"+f]
		if !xist {
			break
		}
	}

	rmTmp := pb.Room{
		Key:              "Rm" + f,
		Id:               f[0:5],
		Host:             req.Host,
		Dueler:           nil,
		Status:           pb.RoomStatus_ON_WAIT,
		Turns:            0,
		CharCardLimitMax: req.CharCardLimitMax,
		CharCardLimitMin: req.CharCardLimitMin,
		CharCardNvn:      req.CharCardNvn,
	}
	rmTmp1 := RoomMgr{
		Room:       rmTmp,
		clientConn: make(map[string]*pb.RoomService_ServerBroadcastServer),
	}
	this.Roomlist["Rm"+f] = &rmTmp1
	return &rmTmp, nil
}
