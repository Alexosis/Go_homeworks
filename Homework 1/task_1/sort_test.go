package main

import (
	"reflect"
	"testing"
)

func TestSorter(t *testing.T) {
	tests := []struct {
		name     string
		args     flags
		filename string
		want     []string
		wantErr  bool
	}{
		{
			name:     "simple sort",
			args:     flags{},
			filename: "data.txt",
			want:     []string{"Apple", "BOOK", "Book", "Go", "Hauptbahnhof", "January", "January", "Napkin"},
			wantErr:  false,
		},
		{
			name: "sort without register",
			args: flags{
				flagF: true,
			},
			filename: "data.txt",
			want:     []string{"Apple", "Book", "BOOK", "Go", "Hauptbahnhof", "January", "January", "Napkin"},
			wantErr:  false,
		},
		{
			name: "sort in reverse order",
			args: flags{
				flagR: true,
			},
			filename: "data.txt",
			want:     []string{"Napkin", "January", "January", "Hauptbahnhof", "Go", "Book", "BOOK", "Apple"},
			wantErr:  true,
		},
		{
			name: "sort with unique bad rows",
			args: flags{
				flagU: true,
				flagK: -1,
			},
			filename: "data.txt",
			want:     nil,
			wantErr:  true,
		},
		{
			name: "sort numbers",
			args: flags{
				flagN: true,
			},
			filename: "numbers.txt",
			want:     []string{"1", "2", "3", "4", "4", "5", "5", "6", "7", "8", "9", "10"},
			wantErr:  false,
		},
		{
			name: "sort with unique",
			args: flags{
				flagU: true,
			},
			filename: "data.txt",
			want:     []string{"Apple", "BOOK", "Book", "Go", "Hauptbahnhof", "January", "Napkin"},
			wantErr:  false,
		},
		{
			name: "sort with unique and without register",
			args: flags{
				flagU: true,
				flagF: true,
			},
			filename: "data.txt",
			want:     []string{"Apple", "BOOK", "Go", "Hauptbahnhof", "January", "Napkin"},
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := Sorter(tt.filename, tt.args); !reflect.DeepEqual(got, tt.want) && (err != nil) != tt.wantErr {
				t.Errorf("UnixSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
