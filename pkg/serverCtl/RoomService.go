package serverCtl

import (
	// _ "ULZRoomService"
	cm "ULZRoomService/pkg/common"
	cf "ULZRoomService/pkg/config"
	rd "ULZRoomService/pkg/store/redis"
	ws "ULZRoomService/pkg/websocket"
	pb "ULZRoomService/proto"
	"log"
	"strconv"
	"sync"
	"time"
	// ants "github.com/panjf2000/ants/v2"
)

var _ pb.RoomServiceServer = (*ULZRoomServiceBackend)(nil)

// Remark: the framework make consider "instant" request
//
type ULZRoomServiceBackend struct {
	// pb.ULZRoomServiceServer
	mu      *sync.Mutex
	CoreKey string
	redhdlr []*rd.RdsCliBox
	// roomStream map[string](*RoomStreamBox)
	castServer *ws.SocketHub
}

// New : Create new backend
func New(conf *cf.ConfTmp) *ULZRoomServiceBackend {
	ck := "RSCore" + cm.HashText(conf.APIServer.IP)
	rdfl := []*rd.RdsCliBox{}
	for i := 0; i < conf.CacheDb.WorkerNode; i++ {
		rdf := rd.New(ck, "wKU"+cm.HashText("num"+strconv.Itoa(i)))

		if cm.Mode == "prod" || cm.Mode == "Debug" {
			rdf.MarshalMethods = "proto"
		}

		if _, err := rdf.Connect(&conf.CacheDb); err == nil {
			rdfl = append(rdfl, rdf)
		}
	}

	// wsWorker := ws.New()

	g := ULZRoomServiceBackend{
		CoreKey: ck,
		mu:      &sync.Mutex{},
		redhdlr: rdfl,
		// roomStream: make(map[string](*RoomStreamBox)),
		castServer: ws.NewHub(),
	}
	// g.InitDB(&conf.Database)
	return &g
}

func (this *ULZRoomServiceBackend) Shutdown() {
	/// TODO: send closing msg to all client
	// for _, v := range this.roomStream {
	// 	log.Println("Server OS.sigKill")
	// 	v.ClearAll()
	// }
	log.Println("in shutdown proc")
	for _, v := range this.redhdlr {
		if _, err := v.CleanRem(); err != nil {
			log.Println(err)
		}
		if _, e := v.Disconn(); e != nil {
			log.Println(e)
		}
	}
	// this.CloseDB()
	log.Println("endof shutdown proc:", this.CoreKey)
}

// PrintReqLog

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
