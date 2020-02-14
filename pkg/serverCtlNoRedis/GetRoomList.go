package serverCtlNoRedis

import (
	cm "ULZRoomService/pkg/common"
	pb "ULZRoomService/proto"
	"context"
	"log"
	"time"
)

// GetRoomList :
func (b *ULZRoomServiceBackend) GetRoomList(ctx context.Context, req *pb.RoomCreateReq) (res *pb.RoomListResp, err error) {
	start := time.Now()
	defer func() {
		elapsed := time.Since(start)
		log.Printf("Get-Room-List took %s", elapsed)
	}()

	var tmp []*pb.RoomSH

	for v := range b.Roomlist {
		if b.Roomlist[v].Room.CharCardNvn == req.CharCardNvn &&
			(req.CostLimitMax != 0 && req.CostLimitMax == b.Roomlist[v].Room.CostLimitMax) &&
			(req.CostLimitMax != 0 && req.CostLimitMax == b.Roomlist[v].Room.CostLimitMax) &&
			(req.CharCardLimitMax != nil && req.CharCardLimitMax == b.Roomlist[v].Room.CharCardLimitMax) &&
			(req.CharCardLimitMin != nil && req.CharCardLimitMin == b.Roomlist[v].Room.CharCardLimitMin) {

			tmp = append(tmp, cm.ToParseSH(b.Roomlist[v].Room))
		}
	}
	for v := range b.Roomlist {
		if b.Roomlist[v].Room.CharCardNvn == req.CharCardNvn &&
			(req.CostLimitMax != 0 && req.CostLimitMax == b.Roomlist[v].Room.CostLimitMax) &&
			(req.CostLimitMax != 0 && req.CostLimitMax == b.Roomlist[v].Room.CostLimitMax) {
			rtmp := false
			for k := range tmp {
				if tmp[k].Key == &b.Roomlist[v].Room.Key {
					rtmp = true
				}
			}
			if !rtmp {
				tmp = append(tmp, cm.ToParseSH(b.Roomlist[v].Room))
			}
		}
	}
	for v := range b.Roomlist {
		if b.Roomlist[v].Room.CharCardNvn == req.CharCardNvn {
			rtmp := false
			for k := range tmp {
				if tmp[k].Key == &b.Roomlist[v].Room.Key {
					rtmp = true
				}
			}
			if !rtmp {
				tmp = append(tmp, cm.ToParseSH(b.Roomlist[v].Room))
			}
		}
	}
	// log.Println("list:", tmp)
	// log.Println("typeof:", reflect.TypeOf(tmp))

	res = &pb.RoomListResp{
		Result:   tmp,
		ErrorMsg: nil,
	}
	return
}
