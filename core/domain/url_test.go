package domain

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewStrictUrlFromString_OK(t *testing.T) {
	type args struct {
		v string
	}
	st, _ := NewStrictUrlFromString("https://example.com")
	tests := []struct {
		name    string
		args    args
		want    StrictUrl
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				v: "https://example.com",
			},
			want:    st,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewStrictUrlFromString(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewStrictUrlFromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStrictUrlFromString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewStrictUrlFromString_Validate(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name    string
		args    args
		want    StrictUrl
		wantErr error
	}{
		{
			name:    "invalid url",
			args:    args{v: "www.google.com"},
			want:    StrictUrl{},
			wantErr: fmt.Errorf("URL must be absolute: www.google.com"),
		},
		{
			name:    "invalid schema",
			args:    args{v: "httpw://www.google.com"},
			want:    StrictUrl{},
			wantErr: fmt.Errorf("URL scheme must be http or https: httpw"),
		},
		{
			name:    "empty host",
			args:    args{v: "https://"},
			want:    StrictUrl{},
			wantErr: fmt.Errorf("URL host must not be empty"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewStrictUrlFromString(tt.args.v)
			if err == nil {
				if tt.wantErr != nil {
					t.Errorf("NewStrictUrlFromString() got = %v, want %v", got, tt.want)
					return
				}
			} else {
				if !reflect.DeepEqual(err, tt.wantErr) {
					t.Errorf("NewStrictUrlFromString() got = %v, want %v", got, tt.want)
					return
				}
			}
		})
	}
}
