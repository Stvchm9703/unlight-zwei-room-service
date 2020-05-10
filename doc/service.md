# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [Protocol Documentation](#protocol-documentation)
  - [Table of Contents](#table-of-contents)
  - [service.proto](#serviceproto)
    - [RoomService](#roomservice)
      - [Create Room](#create-room)
      - [Get Room Info](#get-room-info)
  
- [Scalar Value Types](#scalar-value-types)



<a name="service.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## service.proto


 

 

 


<a name="ULZProto.RoomService"></a>

### RoomService


| Method Name     | Request Type                                     | Response Type                                    | Description |
| --------------- | ------------------------------------------------ | ------------------------------------------------ | ----------- |
| CreateRoom      | [RoomCreateReq](#ULZProto.RoomCreateReq)         | [Room](#ULZProto.Room)                           |             |
| GetRoomList     | [RoomCreateReq](#ULZProto.RoomCreateReq)         | [Room](#ULZProto.Room) stream                    |             |
| GetRoomInfo     | [RoomReq](#ULZProto.RoomReq)                     | [Room](#ULZProto.Room)                           |             |
| UpdateRoom      | [RoomCreateReq](#ULZProto.RoomCreateReq)         | [Room](#ULZProto.Room)                           |             |
| UpdateCard      | [RoomUpdateCardReq](#ULZProto.RoomUpdateCardReq) | [Empty](#ULZProto.Empty)                         |             |
| SendMessage     | [RoomMsg](#ULZProto.RoomMsg)                     | [Empty](#ULZProto.Empty)                         |             |
| QuitRoom        | [RoomReq](#ULZProto.RoomReq)                     | [Empty](#ULZProto.Empty)                         |             |
| QuickPair       | [RoomCreateReq](#ULZProto.RoomCreateReq)         | [Room](#ULZProto.Room)                           |             |
| JoinRoom        | [RoomReq](#ULZProto.RoomReq)                     | [Room](#ULZProto.Room)                           |             |

 
#### Create Room 
Create the Game Chatroom 

#### Get Room Info 