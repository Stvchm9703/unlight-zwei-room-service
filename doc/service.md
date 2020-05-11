# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [Protocol Documentation](#protocol-documentation)
  - [Table of Contents](#table-of-contents)
  - [service.proto](#serviceproto)
    - [RoomService](#roomservice)
      - [Create Room](#create-room)
        - [Work Flow](#work-flow)
      - [Get Room List](#get-room-list)
        - [Work Flow](#work-flow-1)
      - [Get Room Info](#get-room-info)
      - [Workflow](#workflow)
      - [Update Room](#update-room)
      - [Workflow](#workflow-1)
      - [Update Card](#update-card)
      - [Workflow](#workflow-2)
      - [Send Message](#send-message)
        - [WorkFlow](#workflow-3)
      - [QuitRoom](#quitroom)
        - [WorkFlow](#workflow-4)
      - [JoinRoom](#joinroom)
  
- [Scalar Value Types](#scalar-value-types)



<a name="service.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## service.proto

<a name="ULZProto.RoomService"></a>

### RoomService


| Method Name | Request Type                                     | Response Type                 | Description |
| ----------- | ------------------------------------------------ | ----------------------------- | ----------- |
| CreateRoom  | [RoomCreateReq](#ULZProto.RoomCreateReq)         | [Room](#ULZProto.Room)        |             |
| GetRoomList | [RoomCreateReq](#ULZProto.RoomCreateReq)         | [Room](#ULZProto.Room) stream |             |
| GetRoomInfo | [RoomReq](#ULZProto.RoomReq)                     | [Room](#ULZProto.Room)        |             |
| UpdateRoom  | [RoomCreateReq](#ULZProto.RoomCreateReq)         | [Room](#ULZProto.Room)        |             |
| UpdateCard  | [RoomUpdateCardReq](#ULZProto.RoomUpdateCardReq) | [Empty](#ULZProto.Empty)      |             |
| SendMessage | [RoomMsg](#ULZProto.RoomMsg)                     | [Empty](#ULZProto.Empty)      |             |
| QuitRoom    | [RoomReq](#ULZProto.RoomReq)                     | [Empty](#ULZProto.Empty)      |             |
| QuickPair   | [RoomCreateReq](#ULZProto.RoomCreateReq)         | [Room](#ULZProto.Room)        |             |
| JoinRoom    | [RoomReq](#ULZProto.RoomReq)                     | [Room](#ULZProto.Room)        |             |

 
#### Create Room 
Create the Game Room , including the limitation.

##### Work Flow
```mermaid 
sequenceDiagram
participant GC as Game Client
participant RS as Room Service
participant Rds as Redis

  GC -->> RS: CreatRoom( RoomCreateReq )

  loop generate hash key
    RS -->> Rds:  Check Room is Not exist
    alt is Exist 
      RS -->> RS: hash new room key
    else Not Exist 
      RS -->> Rds:  Set Room Data
    end
  end

  RS -->> GC: RoomInfo ( Room )

  opt Redis execution error
    RS -->> GC: error (code.internal) 
  end

```


#### Get Room List 
Fetch the Game Room with searching parameter.

##### Work Flow
```mermaid 
sequenceDiagram
participant GC as Game Client
participant RS as Room Service
participant Rds as Redis

  GC -->> RS: GetRoomList ( RoomCreateReq )

  RS -->> Rds: Get Room list
  Rds -->> RS: return full room list 

  loop room in room list 
    alt is similar 
      RS -->> GC: return room 
    end
  end

  opt Redis execution error
    RS -->> GC: error (code.internal) 
  end
```


#### Get Room Info
get game room detail information.

#### Workflow
```mermaid 
sequenceDiagram
participant GC as Game Client
participant RS as Room Service
participant Rds as Redis

  GC -->> RS: GetRoom ( RoomReq )

  RS -->> Rds: Get Room by Room key
  Rds -->> RS: return room 

  opt Redis execution error
    RS -->> GC: return Error (code = Not Found) 
  end
   
  alt password is vaild
    RS -->> GC: return Room-info ( Room )
  else password invild 
    RS -->> GC: return Error (permission denied)
  else public open room 
    RS -->> GC: return Room-info ( Room )
  end
  
```

#### Update Room
For host player to update the limitation.

#### Workflow
```mermaid 
sequenceDiagram 
participant OGC as Other Game Client
participant GC as Game Client
participant RS as Room Service
participant Rds as Redis

  GC -->> RS: UpdateRoom ( RoomCreateReq )

  RS -->> Rds: Get Room by Room key
  Rds -->> RS: return room 
  opt Redis execution error
    RS -->> GC: return Error (code = Not Found) 
  end

  opt password invild 
    RS -->> GC: return Error (permission denied)
  end

  Note over GC,Rds : password is vaild or public open room 

  RS -->> RS: Set Parameter from RoomCreateReq to Fetched Room
  
  RS -->> Rds: Set Room by Room key
  Rds -->> RS: Complete Update
  opt Redis execution error
    Rds -->> RS: Fail Update
    RS -->> GC: return Error (code = Internal) 
  end
  
  par return the result to request-client (Host player)
    RS -->> GC: return Room-info ( Room )  
  and broadcast to other client
    RS ->> OGC: broadcast via NAT-message system
  end
```

#### Update Card 
For both player to update the Character Card.
#### Workflow
```mermaid 
sequenceDiagram 
participant OGC as Other Game Client
participant GC as Game Client
participant RS as Room Service
participant Rds as Redis

  GC -->> RS: UpdateRoom ( RoomUpdateCardReq )

  RS -->> Rds: Get Room by Room key
  Rds -->> RS: return room 
  opt Redis execution error
    RS -->> GC: return Error (code = Not Found) 
  end

  opt password invild 
    RS -->> GC: return Error (permission denied)
  end

  Note over GC,Rds : password is vaild or public open room 

  RS -->> RS: Set Parameter from RoomUpdateCardReq to Fetched Room
  
  RS -->> Rds: Set Room by Room key

  Rds -->> RS: Complete Update
  opt Redis execution error
    Rds -->> RS: Fail Update
    RS ->> GC: return Error (code = Internal) 
  end
  
  par return the complete status to request-client (Host player)
    RS -->> GC: return Empty ( Empty )  
  and broadcast to all client
    RS ->> OGC: broadcast via NAT-message system
    RS ->> OGC: broadcast via NAT-message system
  end
```

#### Send Message 
sending the command / broadcast message

##### WorkFlow 
```mermaid 
sequenceDiagram 
participant OGC as Other Game Client
participant GC as Game Client
participant RS as Room Service

  GC -->> RS: SendMessage ( RoomUpdateCardReq )

  RS ->> OGC: broadcast message via NAT-message system
  RS ->> GC: broadcast message via NAT-message system
  RS ->> GC: return complete status (empty)

```

#### QuitRoom 
Player / Watcher quiting the game room.
- broadcast who leaving the room, and unregister from NAT-message system(if it is needed)
- if it is host player, the game room is set as \[ON_DESTROY\], so unable other player to join this room.
- if it is dueler player, the game room is set as \[ON_WAIT\], so other player is able to join the game as dueler player.

##### WorkFlow 
```mermaid 
sequenceDiagram 
participant OGC as Other Game Client
participant GC as Game Client
participant RS as Room Service
participant Rds as Redis

  GC -->> RS: QuitRoom ( RoomReq )

  RS -->> Rds: Get Room by Room key
  Rds -->> RS: return room 

  opt Redis execution error
    RS -->> GC: return Error (code = Not Found) 
  end

  Note over GC,Rds : check passwrod is vaild 
  opt password invild 
    RS -->> GC: return Error (permission denied)
  end

  Note over GC,Rds : password is vaild or public open room 

  RS ->> OGC: broadcast "player quiting" system message via NAT-message system

  alt the player is Host player 
    RS -->> OGC: broadcast "Host Quit Room" system message via NAT-message system
  
    RS -->> RS: Set Room's Status as [ON_DESTROY] 

    RS -->> Rds: Set Room into Redis
    Rds -->> RS: Complete Status 
    opt Redis execution error
        Rds -->> RS: Fail Update
        RS ->> GC: return Error (code = Internal) 
      end
  else the player is Dueler player
    RS -->> RS: Set Room's Dueler as [Null]  
    RS -->> RS: Set Room's Status as [ON_Wait] 

    RS -->> Rds: Set Room into Redis
    Rds -->> RS: Complete Status 
    opt Redis execution error
      Rds -->> RS: Fail Update
      RS ->> GC: return Error (code = Internal) 
    end
  end

  RS ->> GC: return Complete status  

```



#### JoinRoom
To allow player to join to game room , \
if the game room have dueler player, other player is rejected to join as dueler.

```mermaid 
sequenceDiagram 
participant OGC as Other Game Client
participant GC as Game Client
participant RS as Room Service
participant Rds as Redis

  GC -->> RS: QuitRoom ( RoomReq )

  RS -->> Rds: Get Room by Room key
  Rds -->> RS: return room 
  opt Redis execution error
    RS -->> GC: return Error (code = Not Found) 
  end

  Note over GC,Rds : check passwrod is vaild 
  opt password invild 
    RS -->> GC: return Error (permission denied)
  end

  Note over GC,Rds : password is vaild or public open room 

  alt the player is allowed to be duler player 
    
    RS -->> RS: Set Room's Dueler as income player 

    RS -->> Rds: Set Room into Redis
    Rds -->> RS: Complete Status 
    opt Redis execution error
        Rds -->> RS: Fail Update
        RS ->> GC: return Error (code = Internal) 
    end

    RS -->> OGC: broadcast "DuelerJoinRoom" system message via NAT-message system
  

  else the player is Watcher
     RS -->> OGC: broadcast "WatcherJoinRoom" system message via NAT-message system
  end

  RS ->> GC: return Complete status  

```
