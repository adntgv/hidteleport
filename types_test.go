package main

import (
	"reflect"
	"testing"

	hook "github.com/robotn/gohook"
)

func Test_eventToBytesProto(t *testing.T) {
	type args struct {
		evt hook.Event
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := eventToBytesProto(tt.args.evt)
			if (err != nil) != tt.wantErr {
				t.Errorf("eventToBytesProto() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("eventToBytesProto() = %v, want %v", got, tt.want)
			}
		})
	}
}
