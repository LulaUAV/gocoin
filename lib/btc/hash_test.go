package btc

import (
	"bytes"
	"math/rand"
	"testing"
)

func TestHMAC_SHA256(t *testing.T) {
	var keys = []string{
		"\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b",
		"\x4a\x65\x66\x65",
		"\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa",
		"\x01\x02\x03\x04\x05\x06\x07\x08\x09\x0a\x0b\x0c\x0d\x0e\x0f\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19",
		"\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa",
		"\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa\xaa"}
	var inputs = []string{
		"\x48\x69\x20\x54\x68\x65\x72\x65",
		"\x77\x68\x61\x74\x20\x64\x6f\x20\x79\x61\x20\x77\x61\x6e\x74\x20\x66\x6f\x72\x20\x6e\x6f\x74\x68\x69\x6e\x67\x3f",
		"\xdd\xdd\xdd\xdd\xdd\xdd\xdd\xdd\xdd\xdd\xdd\xdd\xdd\xdd\xdd\xdd\xdd\xdd\xdd\xdd\xdd\xdd\xdd\xdd\xdd\xdd\xdd\xdd\xdd\xdd\xdd\xdd\xdd\xdd\xdd\xdd\xdd\xdd\xdd\xdd\xdd\xdd\xdd\xdd\xdd\xdd\xdd\xdd\xdd\xdd",
		"\xcd\xcd\xcd\xcd\xcd\xcd\xcd\xcd\xcd\xcd\xcd\xcd\xcd\xcd\xcd\xcd\xcd\xcd\xcd\xcd\xcd\xcd\xcd\xcd\xcd\xcd\xcd\xcd\xcd\xcd\xcd\xcd\xcd\xcd\xcd\xcd\xcd\xcd\xcd\xcd\xcd\xcd\xcd\xcd\xcd\xcd\xcd\xcd\xcd\xcd",
		"\x54\x65\x73\x74\x20\x55\x73\x69\x6e\x67\x20\x4c\x61\x72\x67\x65\x72\x20\x54\x68\x61\x6e\x20\x42\x6c\x6f\x63\x6b\x2d\x53\x69\x7a\x65\x20\x4b\x65\x79\x20\x2d\x20\x48\x61\x73\x68\x20\x4b\x65\x79\x20\x46\x69\x72\x73\x74",
		"\x54\x68\x69\x73\x20\x69\x73\x20\x61\x20\x74\x65\x73\x74\x20\x75\x73\x69\x6e\x67\x20\x61\x20\x6c\x61\x72\x67\x65\x72\x20\x74\x68\x61\x6e\x20\x62\x6c\x6f\x63\x6b\x2d\x73\x69\x7a\x65\x20\x6b\x65\x79\x20\x61\x6e\x64\x20\x61\x20\x6c\x61\x72\x67\x65\x72\x20\x74\x68\x61\x6e\x20\x62\x6c\x6f\x63\x6b\x2d\x73\x69\x7a\x65\x20\x64\x61\x74\x61\x2e\x20\x54\x68\x65\x20\x6b\x65\x79\x20\x6e\x65\x65\x64\x73\x20\x74\x6f\x20\x62\x65\x20\x68\x61\x73\x68\x65\x64\x20\x62\x65\x66\x6f\x72\x65\x20\x62\x65\x69\x6e\x67\x20\x75\x73\x65\x64\x20\x62\x79\x20\x74\x68\x65\x20\x48\x4d\x41\x43\x20\x61\x6c\x67\x6f\x72\x69\x74\x68\x6d\x2e"}
	var outputs = [][]byte{
		{0xb0, 0x34, 0x4c, 0x61, 0xd8, 0xdb, 0x38, 0x53, 0x5c, 0xa8, 0xaf, 0xce, 0xaf, 0x0b, 0xf1, 0x2b, 0x88, 0x1d, 0xc2, 0x00, 0xc9, 0x83, 0x3d, 0xa7, 0x26, 0xe9, 0x37, 0x6c, 0x2e, 0x32, 0xcf, 0xf7},
		{0x5b, 0xdc, 0xc1, 0x46, 0xbf, 0x60, 0x75, 0x4e, 0x6a, 0x04, 0x24, 0x26, 0x08, 0x95, 0x75, 0xc7, 0x5a, 0x00, 0x3f, 0x08, 0x9d, 0x27, 0x39, 0x83, 0x9d, 0xec, 0x58, 0xb9, 0x64, 0xec, 0x38, 0x43},
		{0x77, 0x3e, 0xa9, 0x1e, 0x36, 0x80, 0x0e, 0x46, 0x85, 0x4d, 0xb8, 0xeb, 0xd0, 0x91, 0x81, 0xa7, 0x29, 0x59, 0x09, 0x8b, 0x3e, 0xf8, 0xc1, 0x22, 0xd9, 0x63, 0x55, 0x14, 0xce, 0xd5, 0x65, 0xfe},
		{0x82, 0x55, 0x8a, 0x38, 0x9a, 0x44, 0x3c, 0x0e, 0xa4, 0xcc, 0x81, 0x98, 0x99, 0xf2, 0x08, 0x3a, 0x85, 0xf0, 0xfa, 0xa3, 0xe5, 0x78, 0xf8, 0x07, 0x7a, 0x2e, 0x3f, 0xf4, 0x67, 0x29, 0x66, 0x5b},
		{0x60, 0xe4, 0x31, 0x59, 0x1e, 0xe0, 0xb6, 0x7f, 0x0d, 0x8a, 0x26, 0xaa, 0xcb, 0xf5, 0xb7, 0x7f, 0x8e, 0x0b, 0xc6, 0x21, 0x37, 0x28, 0xc5, 0x14, 0x05, 0x46, 0x04, 0x0f, 0x0e, 0xe3, 0x7f, 0x54},
		{0x9b, 0x09, 0xff, 0xa7, 0x1b, 0x94, 0x2f, 0xcb, 0x27, 0x63, 0x5f, 0xbc, 0xd5, 0xb0, 0xe9, 0x44, 0xbf, 0xdc, 0x63, 0x64, 0x4f, 0x07, 0x13, 0x93, 0x8a, 0x7f, 0x51, 0x53, 0x5c, 0x3a, 0x35, 0xe2}}
	for i := range keys {
		var out [32]byte
		hasher := HMAC_Init([]byte(keys[i]))
		hasher.Write([]byte(inputs[i]))
		hasher.Finalize(out[:])
		if !bytes.Equal(out[:], outputs[i]) {
			t.Error("Wrong result at index", i)
		}
		if len(inputs[i]) > 0 {
			split := rand.Int31n(int32(len(inputs[i])))
			hasher = HMAC_Init([]byte(keys[i]))
			hasher.Write([]byte(inputs[i][:split]))
			hasher.Write([]byte(inputs[i][split:]))
			hasher.Finalize(out[:])
			if !bytes.Equal(out[:], outputs[i]) {
				t.Error("Wrong split result at index", i)
			}
		}
	}
}

func TestRFC6979_HMAC(t *testing.T) {
	var key1 = []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f, 0x00, 0x4b, 0xf5, 0x12, 0x2f, 0x34, 0x45, 0x54, 0xc5, 0x3b, 0xde, 0x2e, 0xbb, 0x8c, 0xd2, 0xb7, 0xe3, 0xd1, 0x60, 0x0a, 0xd6, 0x31, 0xc3, 0x85, 0xa5, 0xd7, 0xcc, 0xe2, 0x3c, 0x77, 0x85, 0x45, 0x9a, 0}
	var out1 = [][]byte{
		{0x4f, 0xe2, 0x95, 0x25, 0xb2, 0x08, 0x68, 0x09, 0x15, 0x9a, 0xcd, 0xf0, 0x50, 0x6e, 0xfb, 0x86, 0xb0, 0xec, 0x93, 0x2c, 0x7b, 0xa4, 0x42, 0x56, 0xab, 0x32, 0x1e, 0x42, 0x1e, 0x67, 0xe9, 0xfb},
		{0x2b, 0xf0, 0xff, 0xf1, 0xd3, 0xc3, 0x78, 0xa2, 0x2d, 0xc5, 0xde, 0x1d, 0x85, 0x65, 0x22, 0x32, 0x5c, 0x65, 0xb5, 0x04, 0x49, 0x1a, 0x0c, 0xbd, 0x01, 0xcb, 0x8f, 0x3a, 0xa6, 0x7f, 0xfd, 0x4a},
		{0xf5, 0x28, 0xb4, 0x10, 0xcb, 0x54, 0x1f, 0x77, 0x00, 0x0d, 0x7a, 0xfb, 0x6c, 0x5b, 0x53, 0xc5, 0xc4, 0x71, 0xea, 0xb4, 0x3e, 0x46, 0x6d, 0x9a, 0xc5, 0x19, 0x0c, 0x39, 0xc8, 0x2f, 0xd8, 0x2e},
	}
	var key2 = []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xe3, 0xb0, 0xc4, 0x42, 0x98, 0xfc, 0x1c, 0x14, 0x9a, 0xfb, 0xf4, 0xc8, 0x99, 0x6f, 0xb9, 0x24, 0x27, 0xae, 0x41, 0xe4, 0x64, 0x9b, 0x93, 0x4c, 0xa4, 0x95, 0x99, 0x1b, 0x78, 0x52, 0xb8, 0x55}
	var out2 = [][]byte{
		{0x9c, 0x23, 0x6c, 0x16, 0x5b, 0x82, 0xae, 0x0c, 0xd5, 0x90, 0x65, 0x9e, 0x10, 0x0b, 0x6b, 0xab, 0x30, 0x36, 0xe7, 0xba, 0x8b, 0x06, 0x74, 0x9b, 0xaf, 0x69, 0x81, 0xe1, 0x6f, 0x1a, 0x2b, 0x95},
		{0xdf, 0x47, 0x10, 0x61, 0x62, 0x5b, 0xc0, 0xea, 0x14, 0xb6, 0x82, 0xfe, 0xee, 0x2c, 0x9c, 0x02, 0xf2, 0x35, 0xda, 0x04, 0x20, 0x4c, 0x1d, 0x62, 0xa1, 0x53, 0x6c, 0x6e, 0x17, 0xae, 0xd7, 0xa9},
		{0x75, 0x97, 0x88, 0x7c, 0xbd, 0x76, 0x32, 0x1f, 0x32, 0xe3, 0x04, 0x40, 0x67, 0x9a, 0x22, 0xcf, 0x7f, 0x8d, 0x9d, 0x2e, 0xac, 0x39, 0x0e, 0x58, 0x1f, 0xea, 0x09, 0x1c, 0xe2, 0x02, 0xba, 0x94},
	}
	var out [32]byte

	rng := RFC6979_HMAC_Init(key1[:64])
	for i := range out1 {
		rng.Generate(out[:])
		if !bytes.Equal(out[:], out1[i]) {
			t.Error("Wrong key1-64 result at index", i)
		}
	}
	rng.Finalize()

	rng = RFC6979_HMAC_Init(key1[:65])
	for i := range out1 {
		rng.Generate(out[:])
		if bytes.Equal(out[:], out1[i]) {
			t.Error("Wrong key1-65 result at index", i)
		}
	}
	rng.Finalize()

	rng = RFC6979_HMAC_Init(key2)
	for i := range out1 {
		rng.Generate(out[:])
		if !bytes.Equal(out[:], out2[i]) {
			t.Error("Wrong key2 result at index", i)
		}
	}
	rng.Finalize()
}