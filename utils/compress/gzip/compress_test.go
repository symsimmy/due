package gzip

import (
	"reflect"
	"testing"
)

func BenchmarkGzipEncode(b *testing.B) {
	input := make([]byte, 1024)
	for i := 0; i < 1024; i++ {
		input[i] = 1
	}

	for i := 0; i < b.N; i++ {
		_, err := Encode(input)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkGzipDecode(b *testing.B) {
	input := make([]byte, 1024)
	for i := 0; i < 1024; i++ {
		input[i] = 1
	}

	got, _ := Encode(input)

	for i := 0; i < b.N; i++ {
		_, err := Decode(got)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func TestGzipDecode(t *testing.T) {
	type args struct {
		input []byte
	}

	decodeBytes := make([]byte, 1024)
	for i := 0; i < 1024; i++ {
		decodeBytes[i] = 1
	}

	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
		{"decode", args{input: []byte{31, 139, 8, 0, 0, 0, 0, 0, 0, 255, 98, 28, 5, 163, 96, 20, 140, 88, 0, 8, 0, 0, 255, 255, 29, 36, 130, 250, 0, 4, 0, 0}}, decodeBytes, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Decode(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("GzipDecode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GzipDecode() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGzipEncode(t *testing.T) {
	type args struct {
		input []byte
	}
	inputBytes := make([]byte, 1024)
	for i := 0; i < 1024; i++ {
		inputBytes[i] = 1
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
		{"encode", args{input: inputBytes}, []byte{31, 139, 8, 0, 0, 0, 0, 0, 0, 255, 98, 28, 5, 163, 96, 20, 140, 88, 0, 8, 0, 0, 255, 255, 29, 36, 130, 250, 0, 4, 0, 0}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Encode(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("GzipEncode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GzipEncode() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncodeRatio(t *testing.T) {
	input := []byte{4, 0, 79, 29, 0, 18, 174, 1, 10, 155, 1, 8, 134, 192, 144, 128, 128, 40, 18, 19, 109, 105, 95, 50, 48, 50, 51, 49, 48, 49, 49, 48, 52, 55, 48, 57, 48, 55, 57, 48, 7, 64, 2, 106, 14, 13, 0, 0, 128, 64, 21, 84, 109, 125, 64, 24, 1, 32, 1, 114, 4, 40, 1, 48, 1, 130, 1, 98, 10, 12, 50, 55, 52, 56, 57, 50, 48, 50, 51, 50, 55, 48, 10, 12, 50, 55, 52, 56, 57, 50, 48, 50, 51, 50, 56, 53, 10, 12, 50, 55, 52, 56, 57, 49, 57, 53, 54, 50, 55, 55, 10, 12, 50, 55, 52, 56, 57, 49, 56, 55, 48, 51, 49, 56, 10, 12, 50, 55, 52, 56, 57, 49, 56, 55, 48, 51, 50, 52, 10, 12, 50, 55, 52, 56, 57, 49, 56, 55, 48, 51, 49, 54, 1, 0, 12, 50, 55, 52, 56, 57, 49, 56, 55, 48, 51, 50, 50, 16, 179, 18, 24, 16, 32, 232, 252, 255, 255, 255, 255, 255, 255, 255, 1, 26, 15, 10, 13, 8, 128, 160, 128, 214, 154, 234, 136, 242, 16, 1, 6, 233, 7, 26, 247, 18, 10, 16, 8, 128, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 96, 56, 1, 10, 16, 8, 129, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 10, 56, 1, 10, 17, 8, 130, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 145, 1, 56, 1, 10, 17, 8, 131, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 188, 1, 56, 1, 10, 17, 8, 132, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 136, 2, 56, 1, 10, 16, 8, 133, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 112, 56, 1, 10, 17, 8, 134, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 171, 1, 56, 1, 10, 17, 8, 135, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 227, 1, 56, 1, 10, 16, 8, 136, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 44, 56, 1, 10, 16, 8, 137, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 105, 56, 1, 10, 16, 8, 138, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 23, 56, 1, 10, 16, 8, 139, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 19, 56, 1, 10, 16, 8, 140, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 114, 56, 1, 10, 16, 8, 141, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 102, 56, 1, 10, 17, 8, 142, 160, 128, 218, 1, 54, 234, 136, 242, 16, 16, 1, 24, 135, 2, 56, 1, 10, 17, 8, 143, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 133, 2, 56, 1, 10, 17, 8, 144, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 2, 4, 243, 1, 56, 1, 10, 17, 8, 145, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 237, 1, 56, 1, 10, 17, 8, 146, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 247, 1, 56, 1, 10, 17, 8, 147, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 138, 1, 56, 1, 10, 17, 8, 148, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 181, 1, 56, 1, 10, 16, 8, 149, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 120, 56, 1, 10, 17, 8, 150, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 167, 1, 56, 1, 10, 16, 8, 151, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 46, 56, 1, 10, 16, 8, 152, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 122, 56, 1, 10, 17, 8, 153, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 131, 2, 56, 1, 10, 16, 8, 154, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 22, 56, 1, 10, 17, 8, 155, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 179, 1, 56, 1, 10, 16, 8, 156, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 31, 5, 6, 1, 10, 17, 8, 157, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 175, 1, 56, 1, 10, 17, 8, 158, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 248, 1, 56, 1, 10, 16, 8, 159, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 58, 56, 1, 10, 16, 8, 160, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 7, 56, 1, 10, 16, 8, 161, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 18, 56, 1, 10, 16, 8, 162, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 37, 56, 1, 10, 16, 8, 163, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 27, 56, 1, 10, 16, 8, 164, 160, 128, 21, 8, 154, 234, 136, 242, 16, 16, 1, 24, 100, 56, 1, 10, 16, 8, 165, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 55, 56, 1, 10, 17, 8, 166, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 132, 2, 56, 1, 10, 17, 8, 167, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 240, 1, 56, 1, 10, 16, 8, 168, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 103, 56, 1, 10, 16, 8, 169, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 61, 56, 1, 10, 16, 8, 170, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 126, 56, 1, 10, 17, 8, 171, 160, 128, 218, 154, 234, 136, 242, 16, 1, 6, 1, 24, 162, 1, 56, 1, 10, 16, 8, 172, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 62, 56, 1, 10, 16, 8, 173, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 94, 56, 1, 10, 16, 8, 174, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 71, 56, 1, 10, 17, 8, 175, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 139, 1, 56, 1, 10, 17, 8, 176, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 226, 1, 56, 1, 10, 16, 8, 177, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 21, 56, 1, 10, 17, 8, 178, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 225, 1, 56, 1, 10, 1, 6, 8, 179, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 56, 56, 1, 10, 16, 8, 180, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 42, 56, 1, 10, 17, 8, 181, 160, 128, 218, 154, 234, 13, 6, 242, 16, 16, 1, 24, 246, 1, 56, 1, 10, 17, 8, 182, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 234, 1, 56, 1, 10, 17, 8, 183, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 140, 1, 5, 6, 1, 10, 16, 8, 184, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 45, 56, 1, 10, 16, 8, 185, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 41, 56, 1, 10, 16, 8, 186, 160, 128, 218, 15, 4, 234, 136, 242, 16, 16, 1, 24, 1, 56, 1, 10, 16, 8, 187, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 110, 56, 1, 10, 17, 8, 188, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 194, 1, 56, 1, 10, 17, 8, 189, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 250, 1, 56, 1, 10, 16, 8, 190, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 91, 56, 1, 10, 16, 8, 191, 160, 128, 2, 18, 154, 234, 136, 242, 16, 16, 1, 24, 69, 56, 1, 10, 16, 8, 192, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 72, 56, 1, 10, 17, 8, 193, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 190, 1, 56, 1, 10, 16, 8, 194, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 121, 56, 1, 10, 17, 8, 195, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 174, 1, 56, 1, 10, 17, 8, 196, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 232, 1, 56, 1, 10, 17, 8, 197, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 159, 1, 56, 1, 10, 16, 8, 198, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 65, 56, 1, 10, 16, 8, 199, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 95, 56, 1, 10, 16, 8, 200, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 74, 56, 1, 10, 17, 8, 20, 1, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 146, 1, 56, 1, 10, 17, 8, 202, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 142, 1, 56, 1, 10, 17, 8, 203, 160, 128, 218, 154, 234, 13, 6, 242, 16, 16, 1, 24, 242, 1, 56, 1, 10, 17, 8, 204, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 129, 1, 56, 1, 10, 17, 8, 205, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 165, 1, 5, 6, 1, 10, 17, 8, 206, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 198, 1, 56, 1, 10, 17, 8, 207, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 160, 1, 56, 1, 10, 17, 8, 208, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 170, 1, 56, 1, 10, 17, 8, 209, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 161, 1, 56, 1, 10, 17, 8, 210, 160, 128, 218, 154, 234, 136, 242, 16, 1, 6, 1, 24, 183, 1, 56, 1, 10, 16, 8, 211, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 79, 56, 1, 10, 16, 8, 212, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 14, 56, 1, 10, 17, 8, 213, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 222, 1, 56, 1, 10, 17, 8, 214, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 135, 1, 56, 1, 10, 17, 8, 215, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 210, 1, 56, 1, 10, 16, 8, 216, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 30, 56, 1, 10, 17, 8, 217, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 148, 1, 56, 1, 1, 0, 17, 8, 218, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 133, 1, 56, 1, 10, 17, 8, 219, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 176, 1, 56, 1, 10, 17, 8, 220, 160, 128, 218, 1, 54, 234, 136, 242, 16, 16, 1, 24, 168, 1, 56, 1, 10, 16, 8, 221, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 108, 56, 1, 10, 17, 8, 222, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 158, 1, 56, 1, 10, 17, 8, 223, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 251, 1, 56, 1, 10, 16, 8, 224, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 106, 56, 1, 10, 16, 8, 225, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 68, 56, 1, 10, 16, 8, 226, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 43, 56, 1, 10, 17, 8, 227, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 235, 1, 56, 1, 10, 16, 8, 228, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 3, 56, 1, 10, 17, 8, 229, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 193, 1, 56, 1, 10, 17, 8, 230, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 204, 1, 56, 1, 10, 17, 8, 231, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 238, 1, 56, 1, 10, 16, 8, 232, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 101, 56, 1, 10, 17, 8, 233, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 143, 1, 56, 1, 10, 16, 8, 234, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 86, 56, 1, 10, 17, 8, 235, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 195, 1, 56, 1, 10, 17, 8, 236, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 180, 1, 56, 1, 10, 16, 8, 237, 160, 128, 218, 15, 4, 234, 136, 242, 16, 16, 1, 24, 70, 56, 1, 10, 16, 8, 238, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 123, 56, 1, 10, 17, 8, 239, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 191, 1, 56, 1, 10, 16, 8, 240, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 16, 56, 1, 10, 17, 8, 241, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 202, 1, 56, 1, 10, 16, 8, 242, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 83, 56, 1, 10, 17, 8, 243, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 252, 1, 56, 1, 10, 17, 8, 244, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 141, 1, 56, 1, 10, 16, 8, 245, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 104, 56, 1, 10, 17, 8, 246, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 186, 1, 56, 1, 10, 16, 8, 247, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 113, 56, 1, 10, 17, 8, 248, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 255, 1, 56, 1, 10, 17, 8, 249, 160, 128, 218, 154, 234, 136, 2, 42, 16, 16, 1, 24, 134, 2, 56, 1, 10, 17, 8, 250, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 163, 1, 56, 1, 10, 17, 8, 251, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 149, 1, 56, 1, 10, 17, 8, 252, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 131, 1, 56, 1, 10, 16, 8, 253, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 13, 56, 1, 10, 17, 8, 254, 160, 128, 218, 15, 4, 234, 136, 242, 16, 16, 1, 24, 220, 1, 56, 1, 10, 17, 8, 255, 160, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 156, 1, 56, 1, 10, 16, 8, 128, 161, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 111, 56, 1, 10, 17, 8, 129, 161, 128, 218, 154, 234, 136, 242, 16, 16, 1, 24, 209, 1, 56, 1, 16, 153, 192, 225, 186, 191, 213, 144, 238, 8, 24, 1, 26, 183, 4, 10, 16, 8, 130, 161, 128, 218, 1, 54, 234, 136, 242, 16, 16, 2, 24, 52, 56, 2, 10, 16, 8, 131, 161, 128, 218, 154, 234, 136, 242, 16, 16, 2, 24, 127, 56, 2, 10, 17, 8, 132, 161, 128, 218, 154, 234, 136, 242, 16, 16, 2, 24, 197, 2, 56, 2, 10, 16, 8, 133, 161, 128, 218, 154, 234, 136, 242, 16, 16, 2, 24, 125, 56, 2, 10, 16, 8, 134, 161, 128, 218, 154, 234, 136, 242, 16, 16, 2, 24, 47, 56, 2, 10, 17, 8, 135, 161, 128, 2, 18, 154, 234, 136, 242, 16, 16, 2, 24, 152, 1, 56, 2, 10, 16, 8, 136, 161, 128, 218, 154, 234, 136, 242, 16, 16, 2, 24, 4, 56, 2, 10, 17, 8, 137, 161, 128, 218, 154, 234, 136, 242, 16, 16, 2, 2, 4, 169, 1, 56, 2, 10, 17, 8, 138, 161, 128, 218, 154, 234, 136, 242, 16, 16, 2, 24, 212, 1, 56, 2, 10, 17, 8, 139, 161, 128, 218, 154, 234, 136, 242, 16, 16, 2, 24, 154, 1, 56, 2, 10, 17, 8, 140, 161, 128, 218, 154, 234, 136, 242, 16, 16, 2, 24, 239, 1, 56, 2, 10, 17, 8, 141, 161, 128, 218, 154, 234, 136, 242, 16, 16, 2, 24, 194, 2, 56, 2, 10, 16, 8, 142, 161, 128, 218, 154, 234, 136, 242, 16, 16, 2, 24, 40, 56, 2, 10, 17, 8, 143, 161, 128, 218, 154, 234, 136, 242, 16, 16, 2, 24, 245, 1, 56, 2, 10, 16, 8, 144, 161, 128, 218, 154, 234, 136, 242, 16, 16, 2, 24, 35, 56, 2, 10, 17, 8, 145, 161, 128, 218, 154, 234, 136, 242, 16, 16, 2, 24, 223, 1, 56, 2, 10, 16, 8, 146, 161, 128, 218, 154, 234, 136, 242, 16, 16, 2, 24, 90, 56, 2, 10, 16, 8, 147, 161, 128, 218, 154, 23, 4, 136, 242, 16, 16, 2, 24, 59, 56, 2, 10, 16, 8, 148, 161, 128, 218, 154, 234, 136, 242, 16, 16, 2, 24, 75, 56, 2, 10, 17, 8, 149, 161, 128, 218, 154, 234, 136, 242, 16, 16, 2, 24, 233, 1, 56, 2, 10, 16, 8, 150, 161, 128, 218, 154, 234, 136, 242, 16, 16, 2, 24, 99, 56, 2, 10, 16, 8, 151, 161, 128, 218, 154, 234, 136, 242, 16, 16, 2, 24, 84, 56, 2, 10, 17, 8, 152, 161, 128, 218, 154, 234, 136, 242, 16, 16, 2, 24, 151, 1, 56, 2, 10, 17, 8, 153, 161, 128, 218, 154, 234, 136, 242, 16, 16, 2, 24, 229, 1, 56, 2, 10, 16, 8, 154, 161, 128, 218, 154, 234, 136, 242, 16, 16, 2, 24, 7, 8, 56, 2, 10, 17, 8, 155, 161, 128, 218, 154, 234, 136, 242, 16, 16, 2, 24, 178, 1, 56, 2, 10, 16, 8, 156, 161, 128, 218, 154, 234, 136, 242, 16, 16, 2, 24, 124, 56, 2, 10, 17, 8, 157, 161, 128, 218, 154, 234, 136, 242, 16, 16, 2, 24, 230, 1, 56, 2, 10, 17, 8, 158, 161, 128, 218, 154, 234, 136, 242, 16, 16, 2, 24, 236, 1, 56, 2, 10, 16, 8, 159, 161, 128, 218, 154, 234, 136, 242, 16, 16, 2, 24, 53, 56, 2, 16, 152, 178, 202, 234, 192, 213, 144, 238, 8, 24, 2, 26, 49, 10, 16, 8, 160, 161, 128, 218, 154, 234, 136, 242, 16, 16, 3, 24, 118, 56, 3, 10, 17, 8, 161, 161, 128, 218, 154, 234, 136, 242, 16, 16, 3, 24, 205, 1, 56, 3, 16, 147, 141, 194, 154, 194, 213, 144, 238, 8, 24, 3, 26, 143, 5, 10, 16, 8, 162, 161, 128, 218, 154, 234, 136, 242, 16, 16, 4, 24, 82, 56, 4, 10, 17, 8, 163, 161, 128, 218, 154, 234, 136, 242, 16, 16, 4, 24, 249, 1, 56, 4, 10, 17, 8, 164, 161, 128, 218, 154, 234, 136, 242, 16, 16, 4, 24, 136, 1, 56, 4, 10, 17, 8, 165, 161, 128, 218, 154, 234, 136, 242, 16, 16, 4, 24, 134, 1, 56, 4, 10, 17, 8, 166, 161, 128, 218, 154, 234, 136, 242, 16, 16, 4, 24, 129, 2, 56, 4, 10, 16, 8, 167, 161, 128, 218, 154, 234, 136, 242, 16, 16, 4, 24, 50, 56, 4, 10, 16, 8, 168, 161, 128, 218, 154, 234, 136, 242, 16, 16, 4, 24, 67, 56, 4, 10, 16, 8, 169, 161, 128, 218, 154, 234, 136, 242, 16, 16, 4, 24, 12, 56, 4, 10, 17, 8, 170, 161, 128, 218, 154, 234, 136, 242, 16, 16, 4, 24, 150, 1, 56, 4, 10, 16, 8, 171, 161, 128, 218, 154, 234, 136, 242, 16, 16, 4, 24, 89, 56, 4, 10, 16, 8, 172, 161, 128, 218, 154, 234, 136, 242, 16, 16, 4, 24, 5, 56, 4, 10, 16, 8, 173, 161, 128, 218, 154, 234, 136, 242, 16, 16, 4, 24, 51, 56, 4, 10, 16, 8, 174, 161, 128, 218, 154, 234, 136, 242, 16, 16, 4, 24, 119, 56, 4, 10, 16, 8, 175, 161, 1, 28, 218, 154, 234, 136, 242, 16, 16, 4, 24, 85, 56, 4, 10, 17, 8, 176, 161, 128, 218, 154, 234, 136, 242, 16, 16, 4, 24, 217, 1, 56, 4, 10, 17, 8, 177, 161, 128, 218, 154, 234, 136, 242, 16, 1, 6, 4, 24, 166, 1, 56, 4, 10, 16, 8, 178, 161, 128, 218, 154, 234, 136, 242, 16, 16, 4, 24, 2, 56, 4, 10, 16, 8, 179, 161, 128, 218, 154, 234, 136, 242, 16, 16, 4, 24, 73, 56, 4, 10, 16, 8, 180, 1, 10, 13, 16, 165, 1, 24, 128, 144, 211, 252, 166, 227, 144, 238, 8}

	got, _ := Encode(input)
	ratio := float64(len(got)) / float64(len(input))
	t.Logf("ratio=%v/%v=%.2f", len(got), len(input), ratio)
}
