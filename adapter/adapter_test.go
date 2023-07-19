package adapter

import (
	"fmt"
	"net/url"
	"reflect"
	"testing"
)

func TestCheckUrlStrictly_OK(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name    string
		args    args
		want    *url.URL
		wantErr bool
	}{
		{
			name:    "valid url",
			args:    args{v: "https://www.google.com"},
			want:    &url.URL{Scheme: "https", Host: "www.google.com"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CheckUrlStrictly(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckUrlStrictly() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CheckUrlStrictly() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckUrlStrictly_NG(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name    string
		args    args
		want    *url.URL
		wantErr error
	}{
		{
			name:    "invalid url",
			args:    args{v: "www.google.com"},
			want:    nil,
			wantErr: fmt.Errorf("URL must be absolute: www.google.com"),
		},
		{
			name:    "invalid schema",
			args:    args{v: "httpw://www.google.com"},
			want:    nil,
			wantErr: fmt.Errorf("URL scheme must be http or https: httpw"),
		},
		{
			name:    "empty host",
			args:    args{v: "https://"},
			want:    nil,
			wantErr: fmt.Errorf("URL host must not be empty"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CheckUrlStrictly(tt.args.v)
			if err == nil {
				if tt.wantErr != nil {
					t.Errorf("TestCheckUrlStrictly() got = %v, want %v", got, tt.want)
					return
				}
			} else {
				if !reflect.DeepEqual(err, tt.wantErr) {
					t.Errorf("TestCheckUrlStrictly() got = %v, want %v", got, tt.want)
					return
				}
			}
		})
	}
}
