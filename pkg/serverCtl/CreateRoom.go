package serverctl

import (
	cm "ULZRoomService/common"
	cf "ULZRoomService/config"
	rd "ULZRoomService/pkg/store/redis"
	pb "ULZRoomService/proto"
	"errors"
	"log"
	"sync"
)

func (this *ULZRoomServiceBackend) CreateRoom(context.Context, *pb.RoomCreateReq) (*pb.Room, error) {
	cm.PrintReqLog(ctx, req)
	start := time.Now()
	b.mu.Lock()
	defer func() {
		b.mu.Unlock()
		elapsed := time.Since(start)
		log.Printf("Quit-Room took %s", elapsed)
	}()

	wkbox := this.searchAliveClient()
	// for loop it
	tmptime := time.Now().String() + req.HostId
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
	wkbox.Preserve(false)
	// b.Roomlist = append(b.Roomlist, &rmTmp)
	
	return &rmTmp, nil
}
