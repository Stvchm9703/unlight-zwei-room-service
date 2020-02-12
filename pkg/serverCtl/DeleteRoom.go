package serverCtl

import (
	pb "ULZRoomService/proto"
	"context"
	"log"
	"time"

	// types "github.com/gogo/protobuf/types"
	"github.com/gogo/status"
)

// DeleteRoom :
func (b *ULZRoomServiceBackend) DeleteRoom(ctx context.Context, req *pb.RoomReq) (*pb.RoomResp, error) {
	start := time.Now()
	b.mu.Lock()
	wkbox := b.searchAliveClient()
	defer func() {
		wkbox.Preserve(false)
		b.mu.Unlock()
		elapsed := time.Since(start)
		log.Printf("Quit-Room took %s", elapsed)
	}()

	if _, err := wkbox.RemovePara(&req.RoomKey); err != nil {
		log.Fatalln(err)
		return nil, status.Errorf(500, err.Error())
	}

	done := false
	// clear broadcast
	if !done {
		return nil, status.Errorf(500, "RoomNotExist")
	}

	// log.Println("b.RoomList", b.Roomlist)
	return nil, nil
}
