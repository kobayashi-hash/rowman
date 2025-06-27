package main

import (
	"reflect"
	"testing"
	"strings"
)

func TestHeadRows(t *testing.T) {
	data := [][]string{ // 仮のcsvデータ
		{"a", "b"},
		{"c", "d"},
		{"e", "f"},
	}

	got := headRows(data, 2) // 2行出力
	want := [][]string{
		{"a", "b"},
		{"c", "d"},
	}

	if !reflect.DeepEqual(got, want) { // 比較
		t.Errorf("headRows() = %v, want %v", got, want)
	}
}

func TestFilter(t *testing.T) {
	data := [][]string{
		{"apple", "red"},
		{"strawberry", "red"},
		{"grape", "purple"},
	}

	got := filter(data, "red")
	want := [][]string{
		{"apple", "red"},
        {"strawberry", "red"},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("filter() = %v, want %v", got, want)
	}
}

func TestHeadCols(t *testing.T) {
	data := [][]string{
		{"a", "b", "c"},
		{"d", "e", "f"},
	}

	got := headCols(data, 2)
	want := [][]string{
		{"a", "b"},
		{"d", "e"},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("headCols() = %v, want %v", got, want)
	}
}

func TestReadCSVFromStdin(t *testing.T) { // 標準入力から受け取ったcsvを出力できるか
	input := "a,b,c\n1,2,3\n"
	r := strings.NewReader(input)
	records, err := readCSVFromStdin(r)
	if err != nil {
		t.Fatalf("readCSVFromStdin error: %v", err)
	}

	want := [][]string{
		{"a", "b", "c"},
		{"1", "2", "3"},
	}

	if !reflect.DeepEqual(records, want) {
		t.Errorf("readCSVFromStdin = %v, want %v", records, want)
	}
}

func TestValidateOpts(t *testing.T) {
	tests := []struct {
		name  string
		opts  options
		wantErr bool
	}{
		{
			name: "only rowNum set",
			opts: options{rowNum: 5},
			wantErr: false,
		},
		{
			name: "only colNum set",
			opts: options{colNum: 3},
			wantErr: false,
		},
		{
			name: "only filterText set",
			opts: options{filterText: "foo"},
			wantErr: false,
		},
		{
			name: "no options set",
			opts: options{},
			wantErr: true,
		},
		{
			name: "multiple options set",
			opts: options{rowNum: 1, colNum: 2},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateOpts(&tt.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateOpts() error = %v, wantErr = %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateArgs(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{
			name: "no args (stdin)",
			args: []string{},
			wantErr: false,
		},
		{
			name: "one file arg",
			args: []string{"file.csv"},
			wantErr: false,
		},
		{
			name: "two file args",
			args: []string{"file1.csv", "file2.csv"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateArgs(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateArgs() error = %v, wantErr = %v", err, tt.wantErr)
			}
		})
	}
}

