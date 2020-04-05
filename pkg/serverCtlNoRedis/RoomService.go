package serverCtlNoRedis

import (
	// _ "ULZRoomService"
	cm "ULZRoomService/pkg/common"
	cf "ULZRoomService/pkg/config"
	ws "ULZRoomService/pkg/websocket"
	pb "ULZRoomService/proto"
	"log"
	"sync"
	// ants "github.com/panjf2000/ants/v2"
)

var _ pb.RoomServiceServer = (*ULZRoomServiceBackend)(nil)

// Remark: the framework make consider "instant" request
//
type ULZRoomServiceBackend struct {
	// pb.ULZRoomServiceServer
	mu         *sync.Mutex
	CoreKey    string
	Roomlist   map[string]*RoomMgr
	castServer *ws.SocketHub
}

// New : Create new backend
func New(conf *cf.ConfTmp) *ULZRoomServiceBackend {
	ck := "RSCore" + cm.HashText(conf.APIServer.IP)
	// cast :=
	g := ULZRoomServiceBackend{
		CoreKey:    ck,
		mu:         &sync.Mutex{},
		Roomlist:   make(map[string]*RoomMgr),
		castServer: ws.NewHub(),
	}
	return &g
}

func (this *ULZRoomServiceBackend) Shutdown() {
	log.Println("in shtdown proc")
	/// TODO: send closing msg to all client

	log.Println("endof shutdown proc:", this.CoreKey)
}
