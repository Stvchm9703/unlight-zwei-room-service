package serverCtlNoRedis

import (
	// _ "ULZRoomService"
	cm "ULZRoomService/common"
	cf "ULZRoomService/config"
	pb "ULZRoomService/proto"
	"errors"
	"fmt"
	"log"
	"sync"
	// ants "github.com/panjf2000/ants/v2"
)

var _ pb.RoomServiceServer = (*ULZRoomServiceBackend)(nil)

// Remark: the framework make consider "instant" request
//
type ULZRoomServiceBackend struct {
	// pb.ULZRoomServiceServer
	mu       *sync.Mutex
	CoreKey  string
	Roomlist []*RoomMgr
}

// New : Create new backend
func New(conf *cf.ConfTmp) *ULZRoomServiceBackend {
	ck := "RSCore" + cm.HashText(conf.APIServer.IP)

	g := ULZRoomServiceBackend{
		CoreKey: ck,
		mu:      &sync.Mutex{},
	}
	return &g
}

func (this *ULZRoomServiceBackend) Shutdown() {
	log.Println("in shtdown proc")
	/// TODO: send closing msg to all client
	for _, v := range this.Roomlist {
		log.Println("Server OS.sigKill")
		v.BroadCast("RmSvrMgr",
			&pb.RoomMsg{
				Key:    v.Key,
				FormId: "RmSvrMgr",
				ToId:   "ALL_USER",
			})
		v.ClearAll()
	}
	this.CloseDB()
	log.Println("endof shutdown proc:", this.CoreKey)
}

// PrintReqLog

// ======================================================================================================
// RoomMgr : Room Manager
type RoomMgr struct {
	pb.Room
	clientConn map[string]*pb.RoomService_ServerBroadcastServer
}

// ----------------------------------------------------------------------------------------------------
// roommgr.get_only_stream

func (rm *RoomMgr) GetGS(user_id string) *pb.RoomService_ServerBroadcastServer {
	log.Println(rm.conn_pool)
	a, ok := rm.clientConn[user_id]
	if ok {
		return a
	}
	return nil
}

func (rm *RoomMgr) AddGS(user_id string, stream *pb.RoomService_ServerBroadcastServer) (bool, error) {
	_, ok := rm.get_only_stream[user_id]
	if ok {
		return false, errors.New("StreamExist")
	}

	rm.get_only_stream[user_id] = stream
	return true, nil
}

func (rm *RoomMgr) DelGS(user_id string) (bool, error) {
	log.Println("Del Stream:", user_id)
	if rm.get_only_stream[user_id] != nil {
		*(rm.get_only_stream[user_id]) = nil
		delete(rm.get_only_stream, user_id)
		return true, nil
	}
	return false, errors.New("StreamNotExist")
}
func (rm *RoomMgr) BroadCastGS(from string, message *pb.CellStatusResp) {
	log.Println(rm.get_only_stream)

	for k, v := range rm.get_only_stream {
		if k != from {
			tmpv := *v
			tmpv.Send(message)
		}
	}

}

// ---------------------------------------------------------------------------------------------

// RoomMgr
func (rm *RoomMgr) BroadCast(from string, message *pb.CellStatusResp) {
	log.Println("BS!", message)
	rm.BroadCastGS(from, message)
}
func (rm *RoomMgr) ClearAll() {
	log.Println("ClearAll Proc")
	for k := range rm.get_only_stream {
		fmt.Println(k)
		rm.DelGS(k)
	}
}
