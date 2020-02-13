package serverCtl

import (
	pb "ULZRoomService/proto"
	"context"
	"log"
	"time"

	"github.com/gogo/status"
)

var (
	systemID = "SYSTEM"
)

func (b *ULZRoomServiceBackend) QuitRoom(ctx context.Context, req *pb.RoomReq) (*pb.Empty, error) {
	start := time.Now()
	b.mu.Lock()
	wkbox := b.searchAliveClient()
	defer func() {
		wkbox.Preserve(false)
		b.mu.Unlock()
		elapsed := time.Since(start)
		log.Printf("Quit-Room took %s", elapsed)
	}()

	var tmp pb.Room

	// read room
	if _, err := wkbox.GetPara(&req.Key, &tmp); err != nil {
		log.Fatalln(err)
		return nil, status.Errorf(500, err.Error())
	}

	_, err := b.DelStream(&req.Key, &req.UserId)
	if err != nil {
		log.Println(err)
	}

	// broadcast to room
	b.BroadCast(&req.Key, &systemID, &pb.RoomMsg{
		Key:     tmp.Key,
		FormId:  systemID,
		ToId:    "ALL_USER",
		Message: req.UserId + " is quited room",
		MsgType: pb.RoomMsg_SYSTEM_INFO,
	})

	//
	// edit room
	if tmp.Host.Id == req.UserId {
		tmp.Host = nil
		// remove room stream
		b.BroadCast(&req.Key, &systemID, &pb.RoomMsg{
			Key:     tmp.Key,
			FormId:  systemID,
			ToId:    "ALL_USER",
			Message: "host is quited room, this room is going abort",
			MsgType: pb.RoomMsg_SYSTEM_INFO,
		})
		a, _ := b.roomStream[req.Key]
		a.ClearAll()
		b.roomStream[req.Key] = nil
		delete(b.roomStream, req.Key)

		// clean room memory
		tmp.Status = pb.RoomStatus_ON_DESTROY
		// it will wait watcher to remove
	} else if tmp.Dueler.Id == req.UserId {
		tmp.Dueler = nil
		tmp.Status = pb.RoomStatus_ON_WAIT
		// available for new player

	}

	wkbox.UpdatePara(&req.Key, tmp)

	// return nil, errors.New("NotImplement")
	return nil, nil

}
