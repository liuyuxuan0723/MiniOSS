package repo

import (
	"awesomeProject/common"
	"awesomeProject/config"
	"fmt"
	"testing"
)

func TestQueryFileByCond(t *testing.T) {
	type args struct {
		name     string
		location string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"test",
			args{name: "123", location: "12"},
			false,
		},
	}
	config.LoadConfig()
	common.MysqlInit()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := QueryFileByCond(tt.args.name, tt.args.location)
			fmt.Println(got[0])
		})
	}
}
