# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [Protocol Documentation](#protocol-documentation)
  - [Table of Contents](#table-of-contents)
  - [proto/message.proto](#protomessageproto)
    - [RmCharCardInfo](#rmcharcardinfo)
    - [RmUserInfo](#rmuserinfo)
    - [Room](#room)
    - [RoomStatus](#roomstatus)
    - [RoomCreateReq](#roomcreatereq)
    - [RoomMsg](#roommsg)
    - [RoomMsg.MsgType](#roommsgmsgtype)
    - [RoomReq](#roomreq)
    - [RoomUpdateCardReq](#roomupdatecardreq)
    - [RoomUpdateCardReq.PlayerSide](#roomupdatecardreqplayerside)
  - [Scalar Value Types](#scalar-value-types)
  
    - [RoomMsg.MsgType](#ULZProto.RoomMsg.MsgType)
    - [RoomStatus](#ULZProto.RoomStatus)
    - [RoomUpdateCardReq.PlayerSide](#ULZProto.RoomUpdateCardReq.PlayerSide)
  
- [Scalar Value Types](#scalar-value-types)



<a name="proto/message.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/message.proto
import &#34;common.proto&#34;;


<a name="ULZProto.RmCharCardInfo"></a>

### RmCharCardInfo
RmCharCardInfo
     Basic Character Card Infomation


| Field   | Type            | Label    | Description                                            |
| ------- | --------------- | -------- | ------------------------------------------------------ |
| card_id | [int32](#int32) | optional | the character card's id, for example abanned card list |
| level   | [int32](#int32) | optional | the character card's level                             |
| rare    | [int32](#int32) | optional | the character card's rare                              |
| cost    | [int32](#int32) | optional | the character card's cost                              |






<a name="ULZProto.RmUserInfo"></a>

### RmUserInfo
RmUserInfo :
User basic Information for room info displaying


| Field     | Type              | Label    | Description                                           |
| --------- | ----------------- | -------- | ----------------------------------------------------- |
| id        | [string](#string) | required | user unique id                                        |
| name      | [string](#string) | required | user name                                             |
| level     | [int32](#int32)   | optional | user experience level (for future feature)            |
| avat_icon | [string](#string) | optional | user avator icon (for future feature)                 |
| title     | [string](#string) | optional | user title  (for future feature)                      |
| rank      | [int32](#int32)   | optional | user ranking in public rank list (for future feature) |






<a name="ULZProto.Room"></a>

### Room
Room :
Room Detail Information, the data structure for storing


| Field               | Type                                       | Label | Description                                                                 |
| ------------------- | ------------------------------------------ | ----- | --------------------------------------------------------------------------- |
| id                  | [string](#string)                          |       | Room-id the id                                                              |
| key                 | [string](#string)                          |       | Room-Key the simply id                                                      |
| password            | [string](#string)                          |       | Room Password                                                               |
| host                | [RmUserInfo](#ULZProto.RmUserInfo)         |       | User-info for room-host player                                              |
| dueler              | [RmUserInfo](#ULZProto.RmUserInfo)         |       | User-info for dueler-player                                                 |
| status              | [RoomStatus](#ULZProto.RoomStatus)         |       | Rooms Status                                                                |
| turns               | [int32](#int32)                            |       | Turns number of Game Duel                                                   |
| cost_limit_max      | [int32](#int32)                            |       | Limitation for Maxmum of Total Deck Cost, null mean no limitation of that   |
| cost_limit_min      | [int32](#int32)                            |       | Limitation for Minimum of Total Deck Cost, null mean no limitation of that  |
| char_card_limit_max | [RmCharCardInfo](#ULZProto.RmCharCardInfo) |       | Limitation for Maxmum Charecter Card Cost, null mean no limitation of that  |
| char_card_limit_min | [RmCharCardInfo](#ULZProto.RmCharCardInfo) |       | Limitation for Minimum Charecter Card Cost, null mean no limitation of that |
| char_card_nvn       | [int32](#int32)                            |       | 1vs1 (value:1) or 3vs3 ( value : 3)                                         |
| host_charcard_id    | [int32](#int32)                            |       | the picked character card id by host player                                 |
| host_cardset_id     | [int32](#int32)                            |       | the picked character set card id by host player                             |
| host_cardlevel      | [int32](#int32)                            |       | the picked character id id by host player                                   |
| duel_charcard_id    | [int32](#int32)                            |       | the picked character card id by dueler player                               |
| duel_cardset_id     | [int32](#int32)                            |       | the picked character set card id by dueler player                           |
| duel_cardlevel      | [int32](#int32)                            |       | the picked character id id by dueler player                                 |

<a name="ULZProto.RoomStatus"></a>

### RoomStatus

| Name       | Number | Description                          |
| ---------- | ------ | ------------------------------------ |
| ON_INIT    | 0      | initialization                       |
| ON_WAIT    | 1      | in waiting state (Room Wait Scene)   |
| ON_START   | 2      | in game start state (CardPlay Scene) |
| ON_END     | 3      | in game end state                    |
| ON_DESTROY | 4      | in destroy data state                |




<a name="ULZProto.RoomCreateReq"></a>

### RoomCreateReq 
1. For CreateRoom 
2. For GetRoomList 

| Field               | Type                                       | Label                           | Description                                                                |
| ------------------- | ------------------------------------------ | ------------------------------- | -------------------------------------------------------------------------- |
| key                 | [string](#string)                          | not required  for creating room | For getting room list to filter the room with similar room key             |
| host                | [RmUserInfo](#ULZProto.RmUserInfo)         | required                        | Host user info, which is generated from game client program                |
| password            | [string](#string)                          | optional                        | a private room                                                             |
| cost_limit_max      | [int32](#int32)                            | optional                        | Limitation for Maxmum of Total Deck Cost, null mean no limitation of that  |
| cost_limit_min      | [int32](#int32)                            | optional                        | Limitation for Minimum of Total Deck Cost, null mean no limitation of that |
| char_card_nvn       | [int32](#int32)                            | required                        | 1vs1 (value : 1) or 3vs3 (value: 3)                                        |
| char_card_limit_max | [RmCharCardInfo](#ULZProto.RmCharCardInfo) | optional                        | Limitation for Max Charecter Card Cost, null mean no limitation of that    |
| char_card_limit_min | [RmCharCardInfo](#ULZProto.RmCharCardInfo) | optional                        | Limitation for Min Charecter Card Cost, null mean no limitation of that    |






<a name="ULZProto.RoomMsg"></a>

### RoomMsg
broadcast message to pass the user message / stricker

| Field    | Type                                         | Label    | Description                                        |
| -------- | -------------------------------------------- | -------- | -------------------------------------------------- |
| key      | [string](#string)                            | required | The Room key that message sending to sepeatic room |
| from_id  | [string](#string)                            | required | The Message Sender's ID                            |
| fm_name  | [string](#string)                            | required | The Message Sender's display name                  |
| to_id    | [string](#string)                            | required | The Receiver's id that Message sending to          |
| to_name  | [string](#string)                            | required | The Receiver's name that Message sending to        |
| message  | [string](#string)                            | required | Message Text                                       |
| msg_type | [RoomMsg.MsgType](#ULZProto.RoomMsg.MsgType) | required | Message Type                                       |

<a name="ULZProto.RoomMsg.MsgType"></a>

### RoomMsg.MsgType
the enum of message type in boradcast message

| Name          | Number | Description                                      |
| ------------- | ------ | ------------------------------------------------ |
| USER_TEXT     | 0      | simple user text                                 |
| USER_STRICKER | 1      | user stricker                                    |
| SYSTEM_INFO   | 2      | system information from room service program     |
| SYSTEM_WARN   | 3      | system warning message from room service program |
| SYSTEM_ERR    | 4      | system error message from room service program   |


<a name="ULZProto.RoomReq"></a>

### RoomReq
For requesting a single room information, and for dueler get into room with password

| Field    | Type                               | Label    | Description                                                     |
| -------- | ---------------------------------- | -------- | --------------------------------------------------------------- |
| key      | [string](#string)                  | required | The requiring room                                              |
| user     | [RmUserInfo](#ULZProto.RmUserInfo) | required | The request user's information                                  |
| is_duel  | [bool](#bool)                      | required | For checking whether is dueler or watcher                       |
| password | [string](#string)                  | optional | Required for private room, able to be null if it is public open |



<a name="ULZProto.RoomUpdateCardReq"></a>

### RoomUpdateCardReq
For updating character in Room Waiting Scene. 

| Field       | Type                                                                   | Label    | Description               |
| ----------- | ---------------------------------------------------------------------- | -------- | ------------------------- |
| key         | [string](#string)                                                      | required | The requiring room        |
| side        | [RoomUpdateCardReq.PlayerSide](#ULZProto.RoomUpdateCardReq.PlayerSide) | required | The requesting player     |
| charcard_id | [int32](#int32)                                                        | required | The character card id     |
| cardset_id  | [int32](#int32)                                                        | required | The character card set id |
| level       | [int32](#int32)                                                        | required | The character card level  |


<a name="ULZProto.RoomUpdateCardReq.PlayerSide"></a>

### RoomUpdateCardReq.PlayerSide


| Name   | Number | Description                               |
| ------ | ------ | ----------------------------------------- |
| HOST   | 0      | Host Player, the first player in the room |
| DUELER | 1      | Duel Player, the first player in the room |








 





## Scalar Value Types

| .proto Type                    | Notes                                                                                                                                           | C++    | Java       | Python      | Go      | C#         | PHP            | Ruby                           |
| ------------------------------ | ----------------------------------------------------------------------------------------------------------------------------------------------- | ------ | ---------- | ----------- | ------- | ---------- | -------------- | ------------------------------ |
| <a name="double" /> double     |                                                                                                                                                 | double | double     | float       | float64 | double     | float          | Float                          |
| <a name="float" /> float       |                                                                                                                                                 | float  | float      | float       | float32 | float      | float          | Float                          |
| <a name="int32" /> int32       | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32  | int        | int         | int32   | int        | integer        | Bignum or Fixnum (as required) |
| <a name="int64" /> int64       | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64  | long       | int/long    | int64   | long       | integer/string | Bignum                         |
| <a name="uint32" /> uint32     | Uses variable-length encoding.                                                                                                                  | uint32 | int        | int/long    | uint32  | uint       | integer        | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64     | Uses variable-length encoding.                                                                                                                  | uint64 | long       | int/long    | uint64  | ulong      | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32     | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s.                            | int32  | int        | int         | int32   | int        | integer        | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64     | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s.                            | int64  | long       | int/long    | int64   | long       | integer/string | Bignum                         |
| <a name="fixed32" /> fixed32   | Always four bytes. More efficient than uint32 if values are often greater than 2^28.                                                            | uint32 | int        | int         | uint32  | uint       | integer        | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64   | Always eight bytes. More efficient than uint64 if values are often greater than 2^56.                                                           | uint64 | long       | int/long    | uint64  | ulong      | integer/string | Bignum                         |
| <a name="sfixed32" /> sfixed32 | Always four bytes.                                                                                                                              | int32  | int        | int         | int32   | int        | integer        | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes.                                                                                                                             | int64  | long       | int/long    | int64   | long       | integer/string | Bignum                         |
| <a name="bool" /> bool         |                                                                                                                                                 | bool   | boolean    | boolean     | bool    | bool       | boolean        | TrueClass/FalseClass           |
| <a name="string" /> string     | A string must always contain UTF-8 encoded or 7-bit ASCII text.                                                                                 | string | String     | str/unicode | string  | string     | string         | String (UTF-8)                 |
| <a name="bytes" /> bytes       | May contain any arbitrary sequence of bytes.                                                                                                    | string | ByteString | str         | []byte  | ByteString | string         | String (ASCII-8BIT)            |

