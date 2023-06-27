package tests

import (
	"testing"

	"github.com/quarks-tech/sendios-go-sdk/internal"
)

func TestBase64Encoder(t *testing.T) {
	type args struct {
		toEncode string
	}
	var tests []struct {
		name string
		args args
		want string
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := internal.Base64Encoder(tt.args.toEncode); got != tt.want {
				t.Errorf("Base64Encoder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSha1Encoder(t *testing.T) {
	type args struct {
		toEncode string
	}
	var tests []struct {
		name string
		args args
		want string
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := internal.Sha1Encoder(tt.args.toEncode); got != tt.want {
				t.Errorf("Sha1Encoder() = %v, want %v", got, tt.want)
			}
		})
	}
}
