package serverCtl

import (
	cm "ULZRoomService/common"
	// cf "ULZRoomService/config"
	// rd "ULZRoomService/pkg/store/redis"
	pb "ULZRoomService/proto"
	"context"
	// "errors"
	"github.com/gogo/status"
	"log"
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
		Key: "Rm" + f,
		// HostId:     req.HostId,
		// DuelerId:   "",
		// Status:     0,
		// Round:      0,
		// Cell:       -1,
		// CellStatus: nil,
	}
	if _, err := wkbox.SetPara(&rmTmp.Key, rmTmp); err != nil {
		log.Println(err)
		return nil, status.Errorf(500, err.Error())
	}

	// b.Roomlist = append(b.Roomlist, &rmTmp)

	return &rmTmp, nil
}
