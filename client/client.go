package main

import (
	"flag"
//	"time"

	ui "hybridseatreservation/client/user_interface"
//	pb "hybridseatreservation/reservation"

)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	flag.Parse()
        ui.InitConnection(*addr)
        defer ui.ConnectionClose()
	ui.Init()
        select{}
}




