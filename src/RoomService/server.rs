use std::collections::HashMap;
use std::hash::{Hash, Hasher};
use std::pin::Pin;
use std::sync::Arc;
use std::time::Instant;

use futures::{Stream, StreamExt};
use tokio::sync::mpsc;
use tonic::transport::Server;
use tonic::{Request, Response, Status};

use ulz_proto::room_service_server::{RoomService, RoomServiceServer};
use ulz_proto::{
    Empty, ErrorMsg, RmCharCardInfo, RmUserInfo, Room, RoomCreateReq, RoomListResp, RoomMsg,
    RoomReq, RoomResp, RoomSearchReq, RoomSh,
};

pub mod ulz_proto {
    tonic::include_proto!("ulz_proto");
}

// mod data;

#[derive(Debug)]
pub struct RoomServiceBackend {
    broadcast_conn: Vec<Response<ServerStreamStream>>,
}

#[tonic::async_trait]
impl RoomService for RoomServiceBackend {
    async fn create_room(&self, request: Request<RoomCreateReq>) -> Result<Response<Room>, Status> {
        println!("requset = {:?}", request);

        Ok(Response::new(Room::default()))
    }
    async fn get_room_list(
        &self,
        request: Request<RoomSearchReq>,
    ) -> Result<Response<RoomListResp>, Status> {
        Ok(Response::new(RoomListResp::default()))
    }

    async fn get_room_info(
        &self, 
        request: Request<RoomReq>
    ) -> Result<Response<RoomResp>, Status> {
        Ok(Response::new(RoomResp::default()))
    }

    async fn delete_room(
        &self, 
        request: Request<RoomReq>
    ) -> Result<Response<RoomResp>, Status> {
        Ok(Response::new(RoomResp::default()))
    }

    // type ServerStreamStream = mpsc::Receiver<Result<RoomMsg, Status>>;

    async fn server_stream(
        &self,
        request: Request<RoomReq>,
    ) -> Result<Response<Self::ServerBroadcastStream>, Status> {
        Ok( 

        )
        // Ok(Response::new(ServerStreamStream::default()))
    }
    async fn send_message(
        &self, 
        request: Request<RoomMsg>,
    ) -> Result<Response<Empty>, Status> {
        Ok(Response::new(Empty::default()))
    }

    async fn quit_room(
        &self, 
        request: Request<RoomReq>
    ) -> Result<Response<Empty>, Status> {
        Ok(Response::new(Empty::default()))
    }

    async fn quick_pair(
        &self, 
        request: Request<RoomSearchReq>
    ) -> Result<Response<Room>, Status> {
        Ok(Response::new(Room::default()))
    }
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let addr: std::net::SocketAddr = "[::1]:10000".parse().unwrap();
    println!("RouteGuideServer listening on: {}", addr);

    // let route_guide = RouteGuideService {
    //     features: Arc::new(data::load()),
    // };

    // let svc = RouteGuideServer::new(route_guide);

    // Server::builder().add_service(svc).serve(addr).await?;

    Ok(())
}
