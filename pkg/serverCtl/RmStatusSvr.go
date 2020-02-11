package serverctl

import (
	// _ "ULZRoomService"
	cm "ULZRoomService/common"
	cf "ULZRoomService/config"
	rd "ULZRoomService/pkg/store/redis"
	pb "ULZRoomService/proto"
	"errors"
	"log"
	"sync"
	// ants "github.com/panjf2000/ants/v2"
)

var _ pb.RoomServiceServer = (*ULZRoomServiceBackend)(nil)

// Remark: the framework make consider "instant" request
//
type ULZRoomServiceBackend struct {
	// pb.ULZRoomServiceServer
	mu        *sync.Mutex
	CoreKey   string
	redhdlr   []*rd.RdsCliBox
	bc_stream map[string]*pb.RoomService_ServerBroadcastServer
}

// New : Create new backend
func New(conf *cf.ConfTmp) *ULZRoomServiceBackend {
	ck := "RSCore" + cm.HashText(conf.APIServer.IP)

	g := ULZRoomServiceBackend{
		CoreKey: ck,
		mu:      &sync.Mutex{},
	}
	// g.InitDB(&conf.Database)
	return &g
}

func (this *ULZRoomServiceBackend) Shutdown() {
	log.Println("in shtdown proc")
	/// TODO: send closing msg to all client
	for _, v := range this.bc_stream {
		log.Println("Server OS.sigKill")
		// v.BroadCast("RmSvrMgr",&pb.RoomMsg{})
	}
	// v.ClearAll()
	// this.CloseDB()
	log.Println("endof shutdown proc:", this.CoreKey)
}

// PrintReqLog

// ----------------------------------------------------------------------------------------------------
// roommgr.get_only_stream

func (rm *ULZRoomServiceBackend) GetGS(user_id string) *pb.RoomService_ServerBroadcastServer {
	a, ok := rm.bc_stream[user_id]
	if ok {
		return a
	}
	return nil
}

func (rm *ULZRoomServiceBackend) AddGS(user_id string, stream *pb.RoomService_ServerBroadcastServer) (bool, error) {
	_, ok := rm.bc_stream[user_id]
	if ok {
		return false, errors.New("StreamExist")
	}

	rm.bc_stream[user_id] = stream
	return true, nil
}

func (rm *ULZRoomServiceBackend) DelGS(user_id string) (bool, error) {
	log.Println("Del Stream:", user_id)
	if rm.bc_stream[user_id] != nil {
		*(rm.bc_stream[user_id]) = nil
		delete(rm.bc_stream, user_id)
		return true, nil
	}
	return false, errors.New("StreamNotExist")
}
func (rm *ULZRoomServiceBackend) BroadCastGS(from string, message *pb.RoomMsg) {
	log.Println(rm.bc_stream)

	for k, v := range rm.bc_stream {
		if k != from {
			tmpv := *v
			tmpv.Send(message)
		}
	}

}

// ---------------------------------------------------------------------------------------------

// RoomMgr
func (rm *ULZRoomServiceBackend) BroadCast(from string, message *pb.RoomMsg) {
	log.Println("BS!", message)
	rm.BroadCastGS(from, message)
}
func (rm *ULZRoomServiceBackend) ClearAll() {
	log.Println("ClearAll Proc")
	// for k := range rm.bc_stream {
	// 	fmt.Println(k)
	// 	rm.DelGS(k)
	// }
}

// -------------------------------------------------------------------------

func (b *ULZRoomServiceBackend) searchAliveClient() *rd.RdsCliBox {
	for {
		wk := b.checkAliveClient()
		if wk == nil {
			// log.Println("busy at " + time.Now().String())
			time.Sleep(500)
		} else {
			wk.Preserve(true)
			return wk
		}
	}
}

// checkAliveClient
func (b *ULZRoomServiceBackend) checkAliveClient() *rd.RdsCliBox {
	for _, v := range b.redhdlr {
		if !*v.IsRunning() {
			return v
		}
	}
	return nil
}

/// <<<=== Worker Goroutine function
