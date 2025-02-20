package serverCtl

import (
	cm "ULZRoomService/pkg/common"
	// cf "ULZRoomService/config"
	// rd "ULZRoomService/pkg/store/redis"
	pb "ULZRoomService/proto"
	"context"

	// "errors"
	"log"

	"github.com/gogo/status"
	"google.golang.org/grpc/codes"

	// "sync"
	"time"
)

func (this *ULZRoomServiceBackend) CreateRoom(ctx context.Context, req *pb.RoomCreateReq) (*pb.Room, error) {
	cm.PrintReqLog(ctx, "create-room", req)
	start := time.Now()
	this.mu.Lock()
	wkbox := this.searchAliveClient()

	defer func() {
		wkbox.Preserve(false)
		this.mu.Unlock()
		elapsed := time.Since(start)
		log.Printf("Quit-Room took %s", elapsed)
	}()
	// for loop it
	tmptime := time.Now().String() + req.Host.GetId()
	var f = ""
	for {
		f = cm.HashText(tmptime)
		l, err := (wkbox).ListRem(&f)
		if err != nil {
			log.Println(err)
			return nil, status.Errorf(codes.Internal, err.Error())
		}
		if len(*l) == 0 {
			break
		}
	}
	if req.CharCardNvn == 0 {
		req.CharCardNvn = 1
	}
	if req.CostLimitMax == 0 {
		req.CostLimitMax = 200
	}
	if req.CharCardLimitMax == nil {
		req.CharCardLimitMax = &pb.RmCharCardInfo{
			Cost: 40,
		}
	}
	rmTmp := pb.Room{
		Key:              "Rm" + f,
		Id:               f[0:5],
		Host:             req.Host,
		Dueler:           nil,
		Status:           pb.RoomStatus_ON_WAIT,
		Turns:            0,
		CostLimitMax:     req.CostLimitMax,
		CostLimitMin:     req.CostLimitMin,
		CharCardLimitMax: req.CharCardLimitMax,
		CharCardLimitMin: req.CharCardLimitMin,
		CharCardNvn:      req.CharCardNvn,
	}

	f = "Rm" + f

	// Set Para
	if _, err := wkbox.SetPara(&rmTmp.Key, rmTmp); err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	_, ok := this.roomStream[f]
	if ok {
		return nil, status.Error(codes.AlreadyExists, "ROOM_IS_EXIST")
	}

	rmStream := RoomStreamBox{
		key:        f,
		password:   req.Password,
		clientConn: make(map[string]*pb.RoomService_ServerBroadcastServer),
	}

	this.roomStream[f] = &rmStream

	return &rmTmp, nil
}
