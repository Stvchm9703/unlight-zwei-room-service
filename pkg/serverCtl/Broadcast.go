package serverCtl

import (
	cm "ULZRoomService/pkg/common"
	cf "ULZRoomService/pkg/config"
	ws "ULZRoomService/pkg/websocket"
	pb "ULZRoomService/proto"

	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gogo/status"

	"google.golang.org/grpc/codes"
	// "time"
)

func (this *ULZRoomServiceBackend) ServerBroadcast(rReq *pb.RoomReq, stream pb.RoomService_ServerBroadcastServer) error {
	cm.PrintReqLog(nil, "server-broadcast", rReq)
	return status.Error(codes.Internal, "SkipStream")
	// _, err := this.AddStream(&rReq.Key, &rReq.User.Id, &stream)
	// if err != nil {
	// 	return status.Error(codes.NotFound, err.Error())
	// }

	// go func() {
	// 	<-stream.Context().Done()
	// 	log.Println("close done")
	// 	_, err := this.DelStream(&rReq.Key, &rReq.User.Id)
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// 	this.BroadCast(&rReq.Key, &rReq.User.Id,
	// 		cm.MsgUserQuitRoom(&rReq.Key, &rReq.User.Id, &rReq.User.Name))
	// }()
	// for {
	// }
}

func (this *ULZRoomServiceBackend) SendMessage(ctx context.Context, msg *pb.RoomMsg) (*pb.Empty, error) {
	cm.PrintReqLog(ctx, "send-message", msg)
	this.BroadCast(msg)
	return &pb.Empty{}, nil
}

func (this *ULZRoomServiceBackend) BroadCast(msg *pb.RoomMsg) {
	// this.
}

func (rsb *ULZRoomServiceBackend) RunWebSocketServer(config cf.CfAPIServer) error {
	hub := ws.NewHub()
	go hub.Run()
	router := gin.New()
	router.GET("/:roomId", Wrapfunc(rsb, hub))
	rsb.castServer = hub
	return router.Run(config.IP + ":" + strconv.Itoa(config.PollingPort))
}

// wraper to gin handler
func Wrapfunc(rsb *ULZRoomServiceBackend, hub *ws.SocketHub) gin.HandlerFunc {
	return func(c *gin.Context) {
		serveWs(rsb, hub, c)
	}
}

// serveWs handles websocket requests from the peer.
func serveWs(rsb *ULZRoomServiceBackend, hub *ws.SocketHub, c *gin.Context) {
	conn, err := ws.Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(c.Param("roomId"))
	reqKey := c.Param("roomId")

	var tmp pb.Room
	wkbox := rsb.searchAliveClient()
	if _, err := wkbox.GetPara(&reqKey, &tmp); err != nil {
		c.AbortWithStatus(412)
	}
	client := ws.NewClient(c.Param("roomId"), hub, conn)
	go client.WritePump()
	go client.ReadPump()
}
