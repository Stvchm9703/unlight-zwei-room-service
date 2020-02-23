package serverCtl

import (
	cm "ULZRoomService/pkg/common"
	pb "ULZRoomService/proto"
	"encoding/json"
	"log"
	"time"

	"github.com/gogo/status"
)

// GetRoomList :
func (b *ULZRoomServiceBackend) GetRoomList(req *pb.RoomCreateReq, stream pb.RoomService_GetRoomListServer) error {
	b.mu.Lock()
	start := time.Now()
	wkbox := b.searchAliveClient()
	cm.PrintReqLog(nil, "get-room-list", req)
	defer func() {
		wkbox.Preserve(false)
		elapsed := time.Since(start)
		log.Printf("get-Room-list took %s", elapsed)
		b.mu.Unlock()
	}()
	// var tmp pb.Room
	var RmList []*pb.Room
	var strl *[]byte
	var err error
	if req.Key != "" {
		strl, err = wkbox.GetParaList(&req.Key)
	} else {
		strl, err = wkbox.GetParaList(nil)
	}

	if err != nil {
		log.Fatalln(err)
		// ignore err ()
	}
	// log.Println("strl:", string(*strl))
	err = json.Unmarshal(*strl, &RmList)
	if err != nil {
		log.Fatalln(err)
		return status.Errorf(500, err.Error())
	}

	// log.Println("list:", RmList)
	// log.Println("typeof:", reflect.TypeOf(RmList))
	var res_list []*string
	log.Println("level0")

	for v := range RmList {
		if RmList[v].CharCardNvn == req.CharCardNvn &&
			(req.CostLimitMax != 0 && req.CostLimitMax == RmList[v].CostLimitMax) &&
			(req.CostLimitMax != 0 && req.CostLimitMax == RmList[v].CostLimitMax) &&
			(req.CharCardLimitMax != nil && req.CharCardLimitMax == RmList[v].CharCardLimitMax) &&
			(req.CharCardLimitMin != nil && req.CharCardLimitMin == RmList[v].CharCardLimitMin) {

			res_list = append(res_list, (&RmList[v].Key))
			stream.Send(RmList[v])
			log.Println(RmList[v])
		}
	}
	log.Println("level1")

	for v := range RmList {
		if RmList[v].CharCardNvn == req.CharCardNvn &&
			(req.CostLimitMax != 0 && req.CostLimitMax == RmList[v].CostLimitMax) &&
			(req.CostLimitMax != 0 && req.CostLimitMax == RmList[v].CostLimitMax) {
			rtmp := false
			for k := range res_list {
				if *res_list[k] == RmList[v].Key {
					rtmp = true
				}
			}
			if !rtmp {
				res_list = append(res_list, (&RmList[v].Key))
				stream.Send(RmList[v])
				log.Println(RmList[v])
			}
		}
	}
	log.Println("level2")

	for v := range RmList {
		if RmList[v].CharCardNvn == req.CharCardNvn {
			rtmp := false
			for k := range res_list {
				if *res_list[k] == RmList[v].Key {
					rtmp = true
				}
			}
			if !rtmp {
				res_list = append(res_list, (&RmList[v].Key))
				stream.Send(RmList[v])
				log.Println(RmList[v])
			}
		}
	}
	log.Println(RmList)
	for v := range RmList {
		rtmp := false
		for k := range res_list {
			if *res_list[k] == RmList[v].Key {
				rtmp = true
			}
		}
		if !rtmp {
			res_list = append(res_list, (&RmList[v].Key))
			stream.Send(RmList[v])
			log.Println(RmList[v])
		}
	}
	return nil
}
