package botsfwmodels

import (
	"reflect"
	"testing"
)

func TestChatKey_String(t *testing.T) {
	type fields struct {
		BotID  string
		ChatID string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "empty",
			want: ":",
		},
		{
			name: "botID",
			fields: fields{
				BotID:  "bot1",
				ChatID: "chat1",
			},
			want: "bot1:chat1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := ChatKey{
				BotID:  tt.fields.BotID,
				ChatID: tt.fields.ChatID,
			}
			if got := k.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChatKey_Validate(t *testing.T) {
	type fields struct {
		BotID  string
		ChatID string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "valid",
			fields: fields{
				BotID:  "bot1",
				ChatID: "chat1",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := ChatKey{
				BotID:  tt.fields.BotID,
				ChatID: tt.fields.ChatID,
			}
			if err := k.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewChatKey(t *testing.T) {
	type args struct {
		botID  string
		chatID string
	}
	tests := []struct {
		name string
		args args
		want ChatKey
	}{
		{
			name: "valid",
			args: args{
				botID:  "bot1",
				chatID: "chat1",
			},
			want: ChatKey{
				BotID:  "bot1",
				ChatID: "chat1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewChatKey(tt.args.botID, tt.args.chatID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewChatKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
