package botsfwmodels

import "testing"

func TestAwaitingReplyToPath(t *testing.T) {
	type args struct {
		awaitingReplyTo string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "empty", args: args{awaitingReplyTo: ""}, want: ""},
		{name: "just_path", args: args{awaitingReplyTo: "test_path_1"}, want: "test_path_1"},
		{name: "with_query", args: args{awaitingReplyTo: "test_path_2?param=1&param=2"}, want: "test_path_2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AwaitingReplyToPath(tt.args.awaitingReplyTo); got != tt.want {
				t.Errorf("AwaitingReplyToPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAwaitingReplyToQuery(t *testing.T) {
	type args struct {
		awaitingReplyTo string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "empty", args: args{awaitingReplyTo: ""}, want: ""},
		{name: "just_path", args: args{awaitingReplyTo: "test_path_1"}, want: ""},
		{name: "with_query", args: args{awaitingReplyTo: "test_path_2?param=1&param=2"}, want: "param=1&param=2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AwaitingReplyToQuery(tt.args.awaitingReplyTo); got != tt.want {
				t.Errorf("AwaitingReplyToQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}
