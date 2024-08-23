package marshal

import (
	"reflect"
	"testing"
)

func Test_customJsonUnmarshal(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "",
			args:    args{b: []byte(`{"Message":"hello"}`)},
			want:    "hello",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := JsonDecoderUnmarshal(tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("JsonDecoderUnmarshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JsonDecoderUnmarshal() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_defaultJsonUnmarshal(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "",
			args:    args{b: []byte(`{"Message":"hello"}`)},
			want:    "hello",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := defaultJsonUnmarshal(tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("defaultJsonUnmarshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("defaultJsonUnmarshal() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_CustomJsonUnmarshal(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "",
			args:    args{b: []byte(`{"Message":"hello"}`)},
			want:    "hello world!",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := JsonCustomUnmarshal(tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("defaultJsonUnmarshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("defaultJsonUnmarshal() got = %v, want %v", got, tt.want)
			}
		})
	}
}
