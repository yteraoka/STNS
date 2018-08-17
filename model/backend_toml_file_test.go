package model

import (
	"reflect"
	"testing"
)

func Test_tomlFileFindByID(t *testing.T) {
	type args struct {
		id   int
		list map[string]UserGroup
	}
	tests := []struct {
		name string
		args args
		want map[string]UserGroup
	}{
		{
			name: "ok",
			args: args{
				id: 1,
				list: map[string]UserGroup{
					"test": &User{
						Base: Base{
							ID:   1,
							Name: "User1",
						},
					},
				},
			},
			want: map[string]UserGroup{
				"test": &User{
					Base: Base{
						ID:   1,
						Name: "User1",
					},
				},
			},
		},
		{
			name: "not found",
			args: args{
				id: 2,
				list: map[string]UserGroup{
					"test": &User{
						Base: Base{
							ID:   1,
							Name: "User1",
						},
					},
				},
			},
			want: map[string]UserGroup{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tomlFileFindByID(tt.args.id, tt.args.list); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tomlFileFindByID() = %v, want %v", got, tt.want)
			}
		})
	}
}
