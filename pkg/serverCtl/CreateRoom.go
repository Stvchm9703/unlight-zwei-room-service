package serverCtl

import (
	cm "ULZRoomService/common"
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
	cm.PrintReqLog(ctx, req)
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
			return nil, status.Errorf(500, err.Error())
		}
		if len(*l) == 0 {
			break
		}
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

	if _, err := wkbox.SetPara(&rmTmp.Key, rmTmp); err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	// b.Roomlist = append(b.Roomlist, &rmTmp)

	return &rmTmp, nil
}
