package serverCtl

import (
	cm "ULZRoomService/pkg/common"
	pb "ULZRoomService/proto"
	"context"
	"log"
	"time"

	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
)

// QuitRoom : Handle
func (b *ULZRoomServiceBackend) QuitRoom(ctx context.Context, req *pb.RoomReq) (*pb.Empty, error) {
	start := time.Now()
	b.mu.Lock()
	wkbox := b.searchAliveClient()
	cm.PrintReqLog(ctx, "quit-room", req)
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
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	// _, err := b.DelStream(&req.Key, &req.User.Id)
	// if err != nil {
	// 	log.Println(err)
	// }

	// broadcast to room
	// b.BroadCast(&req.Key, &b.CoreKey,
	// 	cm.MsgUserQuitRoom(&req.Key, &req.User.Id, &req.User.Name))
	b.BroadCast(
		cm.MsgUserQuitRoom(&req.Key, &req.User.Id, &req.User.Name))

	//
	// edit room
	if tmp.Host.Id == req.User.Id {
		tmp.Host = nil
		// remove room stream
		// b.BroadCast(&req.Key, &b.CoreKey,
		// 	cm.MsgHostQuitRoom(&tmp.Key, &tmp.Id))
		b.BroadCast(
			cm.MsgHostQuitRoom(&tmp.Key, &tmp.Id))

		// a, _ := b.roomStream[req.Key]
		// a.ClearAll()
		// b.roomStream[req.Key] = nil
		// delete(b.roomStream, req.Key)

		// clean room memory
		tmp.Status = pb.RoomStatus_ON_DESTROY
		// it will wait watcher to remove
	} else if tmp.Dueler.Id == req.User.Id {
		tmp.Dueler = nil
		tmp.Status = pb.RoomStatus_ON_WAIT
		// available for new player

	}

	wkbox.UpdatePara(&req.Key, tmp)
	// return nil, errors.New("NotImplement")
	return nil, nil

}
