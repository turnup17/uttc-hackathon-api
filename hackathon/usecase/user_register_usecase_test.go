package usecase

import (
	"main/model"
	"testing"
)

func Test_check_input(t *testing.T) {
	type args struct {
		user_info model.UserResForHTTPPost
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{model.UserResForHTTPPost{Name: "John", Age: 30}},
			want: false,
		},
		{
			name: "test2",
			args: args{model.UserResForHTTPPost{Name: "Bobby", Age: 85}},
			want: true,
		},
		{
			name: "test3",
			args: args{model.UserResForHTTPPost{Name: "BobbyBobbyBobbyBobbyBobbyBobbyBobbyBobbyBobbyBobbyBobbyBobbyBobbyBobbyBobbyBobbyBobbyBobbyBobbyBobbyBobbyBobby", Age: 40}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Check_input(tt.args.user_info); got != tt.want {
				t.Errorf("Check_input() = %v, want %v", got, tt.want)
			}
		})
	}
}
