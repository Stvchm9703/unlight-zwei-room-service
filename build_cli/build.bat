TITLE BUILD-BAT

ECHO GO BUILD CLI : GO-MOD GET PACKAGE TO VENDOR 
go mod vendor

ECHO GO BUILD CLI : ROOM-STATUS SERVICE PROGRAM
go build -o bin\room_status.exe build_cli\room_status.go

