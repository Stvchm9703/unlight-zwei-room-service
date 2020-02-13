package serverCtl

import (
	pb "ULZRoomService/proto"
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/gogo/status"
)

// GetRoomList :
func (b *ULZRoomServiceBackend) GetRoomList(ctx context.Context, req *pb.RoomSearchReq) (*pb.RoomListResp, error) {
	start := time.Now()
	wkbox := b.searchAliveClient()
	defer func() {
		wkbox.Preserve(false)
		elapsed := time.Since(start)
		log.Printf("Quit-Room took %s", elapsed)
	}()
	// var tmp pb.Room
	var RmList []*pb.Room
	var strl *[]byte
	var err error
	if req.Key != "" {
		strl, err = wkbox.GetParaList(&req.Key)
	}

	if err != nil {
		log.Fatalln(err)
		// ignore err ()
	}
	// log.Println("strl:", string(*strl))
	err = json.Unmarshal(*strl, &RmList)
	if err != nil {
		log.Fatalln(err)
		return nil, status.Errorf(500, err.Error())
	}

	// log.Println("list:", RmList)
	// log.Println("typeof:", reflect.TypeOf(RmList))
	var res_list []*pb.RoomSH

	for v := range RmList {
		res_list = append(res_list, ToParseSH(RmList[v]))
	}

	res := &pb.RoomListResp{
		Result:   res_list,
		ErrorMsg: nil,
	}
	return res, nil
}

func ToParseSH(this *pb.Room) *pb.RoomSH {
	return &pb.RoomSH{
		Key:        this.Key,
		HostName:   this.Host.Name,
		HostLv:     this.Host.Level,
		DuelerName: this.Dueler.Name,
		DuelerLv:   this.Dueler.Level,
		Status:     this.Status,
		Turns:      this.Turns,
	}
}
