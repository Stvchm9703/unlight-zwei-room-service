package serverCtlNoRedis

import (
	cm "ULZRoomService/common"
	pb "ULZRoomService/proto"
	"context"
	"errors"
	"log"
	"sync"
	"time"
)

// CreateRoom :

func (this *ULZRoomServiceBackend) CreateRoom(ctx context.Context, req *pb.RoomCreateReq) (*pb.Room, error) {
	cm.PrintReqLog(ctx, req)
	start := time.Now()
	this.mu.Lock()
	// wkbox := this.searchAliveClient()

	defer func() {
		// wkbox.Preserve(false)
		this.mu.Unlock()
		elapsed := time.Since(start)
		log.Printf("Quit-Room took %s", elapsed)
	}()
	// for loop it
	tmptime := time.Now().String() + req.Host.GetId()
	var f = ""
	// for {
	// 	f = cm.HashText(tmptime)
	// 	l, err := (wkbox).ListRem(&f)
	// 	if err != nil {
	// 		log.Println(err)
	// 		return nil, status.Errorf(codes.Internal, err.Error())
	// 	}
	// 	if len(*l) == 0 {
	// 		break
	// 	}
	// }
	for k, v := range this.Roomlist {

	}

	rmTmp := pb.Room{
		Key:              "Rm" + f,
		Host:             req.Host,
		Dueler:           nil,
		Status:           pb.RoomStatus_ON_WAIT,
		Turns:            0,
		CharCardLimitMax: req.CharCardLimitMax,
		CharCardLimitMin: req.CharCardLimitMin,
		CharCardNvn:      req.CharCardNvn,
	}
	f = "Rm" + f

	// Set Para
	// if _, err := wkbox.SetPara(&rmTmp.Key, rmTmp); err != nil {
	// 	log.Println(err)
	// 	return nil, status.Errorf(codes.Internal, err.Error())
	// }

	_, ok := this.roomStream[f]
	if !ok {
		return nil, status.Error(codes.AlreadyExists, "ROOM_IS_EXIST")
	}

	rmStream := RoomStreamBox{
		key:      f,
		password: req.Password,
	}

	this.roomStream[f] = &rmStream

	return &rmTmp, nil
}
