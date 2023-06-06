package util

import (
	"reflect"
	"testing"
)

func TestParseToken(t *testing.T) {
	type args struct {
		token string
	}
	tests := []struct {
		name    string
		args    args
		want    *Claims
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6ImM4MWU3MjhkOWQ0YzJmNjM2ZjA2N2Y4OWNjMTQ4NjJjIiwiZXhwIjoxNjg2MDMzNzI5LCJpc3MiOiJkZXJpIn0.fE9cwlDCQ2yHvMxuy6yida8hmtHf5n1K60qeQw--Vig",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseToken(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}
