package main

import (
	"reflect"
	"testing"
)

func TestCalc(t *testing.T) {
	tests := []struct {
		name    string
		args    string
		want    string
		wantErr bool
	}{
		{
			name:    "simple calculation",
			args:    "2+3-1",
			want:    "4",
			wantErr: false,
		},
		{
			name:    "calculation with different priorities",
			args:    "2+2*2",
			want:    "6",
			wantErr: false,
		},
		{
			name:    "calculation with brackets",
			args:    "(2+2)*2",
			want:    "8",
			wantErr: false,
		},
		{
			name:    "complex calculation",
			args:    "(4*20)/(100-80)+40",
			want:    "44",
			wantErr: false,
		},
		{
			name:    "calc with wrong count of brackets 1",
			args:    "(2+2*2",
			want:    errWithClosing,
			wantErr: true,
		},
		{
			name:    "calc with wrong count of brackets 2",
			args:    "2+2)*2",
			want:    errWithOpening,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := calc(tt.args); !reflect.DeepEqual(got, tt.want) && (err != nil) != tt.wantErr {
				t.Errorf("TestCalc() = %v, want %v", got, tt.want)
			}
		})
	}
}
