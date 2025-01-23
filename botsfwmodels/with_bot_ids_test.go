package botsfwmodels

import (
	"github.com/strongo/validation"
	"testing"
)

func TestWithBotIDs_Validate(t *testing.T) {
	type fields struct {
		BotIDs []string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Should pass for good records",
			fields: fields{
				BotIDs: []string{"bot1", "bot2", "bot3"},
			},
			wantErr: false,
		},
		{
			name: "Should pass for nil",
			fields: fields{
				BotIDs: []string{"bot1"},
			},
			wantErr: false,
		},
		{
			name: "Should pass for empty",
			fields: fields{
				BotIDs: []string{"bot1"},
			},
			wantErr: false,
		},
		{
			name: "Should fail for non trimmed",
			fields: fields{
				BotIDs: []string{" bot1"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := WithBotIDs{
				BotIDs: tt.fields.BotIDs,
			}
			if err := v.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWithRequiredBotIDs_Validate(t *testing.T) {
	tests := []struct {
		name    string
		v       WithRequiredBotIDs
		wantErr bool
	}{
		{
			name: "Should pass",
			v: WithRequiredBotIDs{
				BotIDs: []string{"bot1"},
			},
			wantErr: false,
		},
		{
			name:    "Should fail on nil",
			v:       WithRequiredBotIDs{BotIDs: []string{}},
			wantErr: true,
		},
		{
			name:    "Should fail on empty",
			v:       WithRequiredBotIDs{BotIDs: []string{}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.v.Validate()
			if err == nil && tt.wantErr || err != nil && !tt.wantErr || tt.wantErr && !validation.IsBadFieldValueError(err) {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
