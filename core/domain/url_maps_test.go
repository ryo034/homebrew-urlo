package domain

import (
	"reflect"
	"testing"
)

func TestUrlMaps_GetTitles_OK(t *testing.T) {
	type fields struct {
		values []UrlMap
	}
	v, _ := NewStrictUrlFromString("https://example.com")
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "OK",
			fields: fields{
				values: []UrlMap{
					NewUrlMap("title1", v),
				},
			},
			want: []string{"title1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			us := &UrlMaps{
				values: tt.fields.values,
			}
			if got := us.GetTitles(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTitles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUrlMaps_GetItemFromLabel_OK(t *testing.T) {
	type fields struct {
		values []UrlMap
	}
	type args struct {
		label Title
	}
	v, _ := NewStrictUrlFromString("https://example.com")
	v2, _ := NewStrictUrlFromString("https://example2.com")
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    UrlMap
		want1   int
		wantErr bool
	}{
		{
			name:    "OK",
			fields:  fields{values: []UrlMap{NewUrlMap("title1", v), NewUrlMap("title2", v2)}},
			args:    args{label: "title1"},
			want:    NewUrlMap("title1", v),
			want1:   0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			us := &UrlMaps{
				values: tt.fields.values,
			}
			got, got1, err := us.GetItemFromLabel(tt.args.label)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetItemFromLabel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetItemFromLabel() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetItemFromLabel() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestUrlMaps_TitleMaxLen_OK(t *testing.T) {
	type fields struct {
		values []UrlMap
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "OK",
			fields: fields{
				values: []UrlMap{
					NewUrlMap("title1", StrictUrl{}),
					NewUrlMap("title2", StrictUrl{}),
				},
			},
			want: 6,
		},
		{
			name: "Japanese OK",
			fields: fields{
				values: []UrlMap{
					NewUrlMap("タイトル1", StrictUrl{}),
					NewUrlMap("タイトル", StrictUrl{}),
				},
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			us := &UrlMaps{
				values: tt.fields.values,
			}
			if got := us.TitleMaxLen(); got != tt.want {
				t.Errorf("TitleMaxLen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUrlMaps_FilterByRegex_OK(t *testing.T) {
	type fields struct {
		values []UrlMap
	}
	type args struct {
		query string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *UrlMaps
		wantErr bool
	}{
		{
			name: "OK",
			fields: fields{
				values: []UrlMap{
					NewUrlMap("title1", StrictUrl{}),
					NewUrlMap("title2", StrictUrl{}),
				},
			},
			args:    args{query: "1"},
			want:    &UrlMaps{values: []UrlMap{NewUrlMap("title1", StrictUrl{})}},
			wantErr: false,
		},
		{
			name: "No match",
			fields: fields{
				values: []UrlMap{
					NewUrlMap("title1", StrictUrl{}),
					NewUrlMap("title2", StrictUrl{}),
				},
			},
			args:    args{query: "3"},
			want:    &UrlMaps{values: []UrlMap{}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			us := &UrlMaps{
				values: tt.fields.values,
			}
			got, err := us.FilterByRegex(tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("FilterByRegex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterByRegex() got = %v, want %v", got, tt.want)
			}
		})
	}
}
