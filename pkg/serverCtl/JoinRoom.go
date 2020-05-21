package serverCtl

import (
	cm "ULZRoomService/pkg/common"
	pb "ULZRoomService/proto"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
)

// JoinRoom :
func (b *ULZRoomServiceBackend) JoinRoom(ctx context.Context, req *pb.RoomReq) (*pb.Room, error) {
	start := time.Now()
	b.mu.Lock()
	wkbox := b.searchAliveClient()
	cm.PrintReqLog(ctx, "join-room", req)
	defer func() {
		(wkbox).Preserve(false)
		b.mu.Unlock()
		elapsed := time.Since(start)
		log.Printf("join-Room took %s", elapsed)
	}()

	var tmp pb.Room
	if _, err := wkbox.GetPara(&req.Key, &tmp); err != nil {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}
	if tmp.Dueler == nil && req.IsDuel {
		tmp.Dueler = req.User
		wkbox.UpdatePara(&tmp.Key, tmp)
	}
	side := ""
	if req.IsDuel {
		side = "Dueler"
	} else {
		side = "Watcher"
	}

	tmpmsg := pb.RoomMsg{
		Key:     req.Key,
		FromId:  req.User.Id,
		FmName:  req.User.Name,
		ToId:    "ALL",
		ToName:  "ALL",
		MsgType: pb.RoomMsg_SYSTEM_INFO,
		Message: fmt.Sprintf("%s is joined to this room", side),
	}

	b.BroadCast(&tmpmsg)

	return &tmp, nil
}
