package serverCtl

import (
	pb "ULZRoomService/proto"
	"context"
	"errors"
	"log"
	"time"

	"github.com/gogo/status"
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


	// if tmp.Host.key == 
	

	// if tmpr.HostId == req.HostId {
	// 	// Self quit, not start play yet

	// } else {
	// 	// Dueler quit game,
	// }
	return nil, errors.New("NotImplement")

}
