package tests

import (
	"crypto/cipher"
	"sendios/internal"
	"testing"
)

func TestEncrypt_Decrypt(t *testing.T) {
	type fields struct {
		block cipher.Block
	}
	type args struct {
		encryptedData string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encrypt, err := internal.MakeEncrypt()
			if (err != nil) != tt.wantErr {
				t.Errorf("MakeEncrypt error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			got, err := encrypt.Decrypt(tt.args.encryptedData)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Decrypt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncrypt_EncryptData(t *testing.T) {
	type fields struct {
		block cipher.Block
	}
	type args struct {
		dataToEncrypt []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encrypt, err := internal.MakeEncrypt()
			if (err != nil) != tt.wantErr {
				t.Errorf("MakeEncrypt error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			got, err := encrypt.EncryptData(tt.args.dataToEncrypt)
			if (err != nil) != tt.wantErr {
				t.Errorf("EncryptData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("EncryptData() got = %v, want %v", got, tt.want)
			}
		})
	}
}
