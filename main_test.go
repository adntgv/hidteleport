package main

import (
	"reflect"
	"testing"

	hook "github.com/robotn/gohook"
)

func Test_eventsFromBytes(t *testing.T) {
	type args struct {
		bz []byte
	}
	tests := []struct {
		name    string
		args    args
		wantEvt hook.Event
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				bz: []byte{},
			},
			wantEvt: hook.Event{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotEvt, err := eventsFromBytesProto(tt.args.bz)
			if (err != nil) != tt.wantErr {
				t.Errorf("eventsFromBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotEvt, tt.wantEvt) {
				t.Errorf("eventsFromBytes() = %v, want %v", gotEvt, tt.wantEvt)
			}
		})
	}
}
