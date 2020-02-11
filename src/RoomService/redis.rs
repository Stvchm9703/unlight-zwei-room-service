use ulz_proto::{
    Empty, ErrorMsg, 
    RmCharCardInfo, RmUserInfo, 
    Room, RoomListResp, 
    RoomMsg, RoomReq, RoomResp,
    RoomSearchReq, RoomSh, RoomCreateReq
};
use std::vec ;
use futures::prelude::*;
use redis::AsyncCommands;
use serde::{Deserialize, Serialize};
use serde_json::Result;

fn init(){

}

pub struct RedisCliBox {
    conn : Arc<redis::Client>
}

impl redisFunc for RedisCliBox{

    async fn insert_room(room: Room)->Result<Room, Status> {

    }

    fn fetch_room_list(req: RoomSearchReq) -> Vec<Room>  {
        // if (req.key != "") {
        //     let result = redis::cmd("SSCAN")
        //         .arg("");

        //     let dat = req.parse
        // }
    }
}

fn main(){
    let y = RedisCliBox{
        conn : Arc::new(redis::Client::open()),
    };


}

