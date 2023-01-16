package main

import (
	trippb "coolcar/proto/gen/go"
	"encoding/json"
	"fmt"
	"google.golang.org/protobuf/proto"
)

func main() {
	trip := trippb.Trip{
		Start:       "abc",
		End:         "def",
		DurationSec: 3600,
		FeeCent:     10000,
		StartPos: &trippb.Location{
			Lat: 30,
			Lon: 120,
		},
		EndPos: &trippb.Location{
			Lat: 31,
			Lon: 115,
		},
		PathPos: []*trippb.Location{
			{
				Lat: 32,
				Lon: 116,
			},
		},
		Status: trippb.TripStatus_FINISHED,
	}
	fmt.Println(&trip)
	marshal, err := proto.Marshal(&trip)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%X\n", marshal)

	var trip2 trippb.Trip
	err = proto.Unmarshal(marshal, &trip2)
	if err != nil {
		panic(err)
	}
	fmt.Println(&trip2)

	b, err := json.Marshal(&trip2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", b)
}
