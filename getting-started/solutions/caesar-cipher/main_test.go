package main

import "testing"

func TestEncryption(t *testing.T) {
	type args struct {
		message string
		key     int
	}
	tests := []struct {
		name   string
		input  args
		output string
	}{
		{
			name: "encrypts an empty string",
			input: args{
				message: "",
				key:     12,
			},
			output: "",
		},
		{
			name: "encrypts 'foobarbaz'",
			input: args{
				message: "foobarbaz",
				key:     1,
			},
			output: "gppcbscba",
		},
		{
			name: "encrypts 'aaa'",
			input: args{
				message: "aaa",
				key:     1,
			},
			output: "bbb",
		},
		{
			name: "encrypts 'xxx' and wraps around",
			input: args{
				message: "xxx",
				key:     5,
			},
			output: "ccc",
		},
		{
			name: "key value of 26 gives same result",
			input: args{
				message: "golang",
				key:     26,
			},
			output: "golang",
		},
		{
			name: "encryption keeps punctuation",
			input: args{
				message: "hello world!",
				key:     1,
			},
			output: "ifmmp xpsme!",
		},
	}

	for _, tt := range tests {
		res := encrypt(tt.input.message, tt.input.key)
		if res != tt.output {
			t.Errorf("%s: got %s want %s", tt.name, res, tt.output)
		}
	}
}

func TestDecryption(t *testing.T) {
	type args struct {
		message string
		key     int
	}
	tests := []struct {
		name   string
		input  args
		output string
	}{
		{
			name: "decrypts an empty string",
			input: args{
				message: "",
				key:     12,
			},
			output: "",
		},
		{
			name: "decrypts 'gppcbscba'",
			input: args{
				message: "gppcbscba",
				key:     1,
			},
			output: "foobarbaz",
		},
		{
			name: "decrypts 'ddd'",
			input: args{
				message: "ddd",
				key:     1,
			},
			output: "ccc",
		},
		{
			name: "decrypts 'ddd' and wraps around",
			input: args{
				message: "ddd",
				key:     5,
			},
			output: "yyy",
		},
		{
			name: "key value of 26 gives same result",
			input: args{
				message: "golang",
				key:     26,
			},
			output: "golang",
		},
		{
			name: "decryption keeps punctuation",
			input: args{
				message: "ifmmp xpsme!",
				key:     1,
			},
			output: "hello world!",
		},
	}

	for _, tt := range tests {
		res := decrypt(tt.input.message, tt.input.key)
		if res != tt.output {
			t.Errorf("%s: got %s want %s", tt.name, res, tt.output)
		}
	}
}
