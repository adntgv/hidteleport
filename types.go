package main

import (
	"encoding/json"

	pb "github.com/adntgv/hidteleport/proto"
	hook "github.com/robotn/gohook"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func eventToBytesProto(evt hook.Event) ([]byte, error) {
	pbEvt := &pb.Event{
		Kind:     uint32(evt.Kind),
		When:     timestamppb.New(evt.When),
		Mask:     uint32(evt.Mask),
		Reserved: uint32(evt.Reserved),

		Keycode: uint32(evt.Keycode),
		Rawcode: uint32(evt.Rawcode),
		Keychar: evt.Keychar,

		Button: uint32(evt.Button),
		Clicks: uint32(evt.Clicks),

		X: int32(evt.X),
		Y: int32(evt.Y),

		Amount:    uint32(evt.Amount),
		Rotation:  evt.Rotation,
		Direction: uint32(evt.Direction),
	}
	return proto.Marshal(pbEvt)
}

func eventsFromBytesProto(bz []byte) (hook.Event, error) {
	pbEvt := &pb.Event{}
	err := proto.Unmarshal(bz, pbEvt)
	if err != nil {
		return hook.Event{}, nil
	}

	evt := hook.Event{
		Kind:     uint8(pbEvt.Kind),
		When:     pbEvt.When.AsTime(),
		Mask:     uint16(pbEvt.Mask),
		Reserved: uint16(pbEvt.Reserved),

		Keycode: uint16(pbEvt.Keycode),
		Rawcode: uint16(pbEvt.Rawcode),
		Keychar: pbEvt.Keychar,

		Button: uint16(pbEvt.Button),
		Clicks: uint16(pbEvt.Clicks),

		X: int16(pbEvt.X),
		Y: int16(pbEvt.Y),

		Amount:    uint16(pbEvt.Amount),
		Rotation:  pbEvt.Rotation,
		Direction: uint8(pbEvt.Direction),
	}

	return evt, nil
}

func eventToBytesJSON(evt hook.Event) ([]byte, error) {
	return json.Marshal(evt)
}

func eventsFromBytesJSON(bz []byte) (hook.Event, error) {
	evt := hook.Event{}
	err := json.Unmarshal(bz, &evt)
	return evt, err
}

type screenSize struct {
	width  float32
	height float32
}

type scaledPosition struct {
	X float32
	Y float32
}

type eventWrapper struct {
	hook.Event
	ScaledMousePosition scaledPosition
}

func (ew eventWrapper) String() string {
	bz, _ := json.Marshal(ew)
	return ew.Event.String() + " " + string(bz)
}
