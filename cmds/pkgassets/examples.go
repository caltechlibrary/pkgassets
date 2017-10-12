package main

var (

    // Examples is a map to asset files associated with main package
    Examples = map[string][]byte{
    "help": []byte{0xa,0x23,0x20,0x43,0x72,0x65,0x61,0x74,0x69,0x6e,0x67,0x20,0x68,0x65,0x6c,0x70,0x20,0x64,0x6f,0x63,0x73,0xa,0xa,0x48,0x65,0x72,0x65,0x20,0x69,0x73,0x20,0x74,0x68,0x65,0x20,0x62,0x61,0x73,0x69,0x63,0x20,0x72,0x65,0x63,0x69,0x70,0x65,0x20,0x66,0x6f,0x72,0x20,0x74,0x75,0x72,0x6e,0x69,0x6e,0x67,0x20,0x74,0x65,0x78,0x74,0x20,0x6f,0x72,0x20,0x6d,0x61,0x72,0x6b,0x64,0x6f,0x77,0x6e,0x20,0x64,0x6f,0x63,0x75,0x6d,0x65,0x6e,0x74,0x73,0x20,0x69,0x6e,0x74,0x6f,0x20,0x6d,0x61,0x70,0x5b,0x73,0x74,0x72,0x69,0x6e,0x67,0x5d,0x5b,0x5d,0x62,0x79,0x74,0x65,0x20,0x6f,0x62,0x6a,0x65,0x63,0x74,0xa,0x74,0x68,0x61,0x74,0x20,0x63,0x61,0x6e,0x20,0x64,0x72,0x69,0x76,0x65,0x20,0x79,0x6f,0x75,0x72,0x20,0x68,0x65,0x6c,0x70,0x20,0x74,0x6f,0x70,0x69,0x63,0x73,0x2e,0xa,0xa,0x31,0x2e,0x20,0x6d,0x6b,0x64,0x69,0x72,0x20,0x61,0x20,0x64,0x6f,0x63,0x73,0x20,0x64,0x69,0x72,0x65,0x63,0x74,0x6f,0x72,0x79,0xa,0x32,0x2e,0x20,0x63,0x72,0x65,0x61,0x74,0x65,0x20,0x6d,0x61,0x72,0x6b,0x64,0x6f,0x77,0x6e,0x20,0x28,0x2a,0x2e,0x6d,0x64,0x29,0x20,0x6f,0x72,0x20,0x74,0x65,0x78,0x74,0x20,0x28,0x2a,0x2e,0x74,0x78,0x74,0x29,0x20,0x64,0x6f,0x63,0x75,0x6d,0x65,0x6e,0x74,0x73,0x20,0x69,0x6e,0x20,0x74,0x68,0x61,0x74,0x20,0x64,0x69,0x72,0x65,0x63,0x74,0x6f,0x72,0x79,0x20,0x66,0x6f,0x72,0x20,0x79,0x6f,0x75,0x72,0x20,0x68,0x65,0x6c,0x70,0x20,0x74,0x6f,0x70,0x69,0x63,0x73,0xa,0x20,0x20,0x20,0x20,0x2b,0x20,0x66,0x69,0x6c,0x65,0x6e,0x61,0x6d,0x65,0x20,0x77,0x69,0x74,0x68,0x6f,0x75,0x74,0x20,0x65,0x78,0x74,0x65,0x6e,0x73,0x69,0x6f,0x6e,0x20,0x77,0x69,0x6c,0x6c,0x20,0x62,0x65,0x63,0x6f,0x6d,0x65,0x20,0x6b,0x65,0x79,0x77,0x6f,0x72,0x64,0x20,0x75,0x73,0x65,0x64,0x20,0x77,0x69,0x74,0x68,0x20,0x74,0x68,0x65,0x20,0x68,0x65,0x6c,0x70,0x20,0x63,0x6f,0x6d,0x6d,0x61,0x6e,0x64,0xa,0x33,0x2e,0x20,0x72,0x75,0x6e,0x20,0x5f,0x70,0x6b,0x67,0x61,0x73,0x73,0x65,0x74,0x73,0x5f,0x20,0x6f,0x76,0x65,0x72,0x20,0x74,0x68,0x65,0x20,0x64,0x69,0x72,0x65,0x63,0x74,0x6f,0x72,0x79,0x20,0x67,0x65,0x6e,0x65,0x72,0x61,0x74,0x69,0x6e,0x67,0x20,0x61,0x6e,0x20,0x2a,0x61,0x73,0x73,0x65,0x74,0x2e,0x67,0x6f,0x2a,0x20,0x66,0x69,0x6c,0x65,0x20,0x69,0x6e,0x20,0x74,0x68,0x65,0x20,0x73,0x61,0x6d,0x65,0x20,0x66,0x6f,0x6c,0x64,0x65,0x72,0x20,0x61,0x73,0x20,0x79,0x6f,0x75,0x72,0x20,0x63,0x6c,0x69,0x20,0x70,0x72,0x6f,0x67,0x72,0x61,0x6d,0xa,0x34,0x2e,0x20,0x63,0x6f,0x6d,0x70,0x69,0x6c,0x65,0x20,0x79,0x6f,0x75,0x72,0x20,0x63,0x6c,0x69,0x20,0x70,0x72,0x6f,0x67,0x61,0x6d,0xa,0xa,0x23,0x23,0x20,0x45,0x78,0x61,0x6d,0x70,0x6c,0x65,0xa,0xa,0x49,0x6e,0x20,0x74,0x68,0x69,0x73,0x20,0x65,0x78,0x61,0x6d,0x70,0x6c,0x65,0x20,0x74,0x68,0x65,0x20,0x2a,0x63,0x6d,0x64,0x2f,0x68,0x65,0x6c,0x6c,0x6f,0x77,0x6f,0x72,0x6c,0x64,0x2f,0x68,0x65,0x6c,0x6c,0x6f,0x77,0x6f,0x72,0x6c,0x64,0x2e,0x67,0x6f,0x2a,0x20,0x77,0x6f,0x75,0x6c,0x64,0x20,0x63,0x6f,0x6e,0x74,0x61,0x69,0x6e,0x73,0x20,0x74,0x68,0x65,0x20,0x22,0x6d,0x61,0x69,0x6e,0x22,0x20,0x70,0x61,0x63,0x6b,0x61,0x67,0x65,0x20,0x66,0x6f,0x72,0x20,0x61,0x20,0x63,0x6c,0x69,0x20,0x70,0x72,0x6f,0x67,0x72,0x61,0x6d,0xa,0x79,0x6f,0x75,0x27,0x72,0x65,0x20,0x67,0x6f,0x69,0x6e,0x67,0x20,0x74,0x6f,0x20,0x62,0x75,0x69,0x6c,0x64,0x2e,0x20,0x20,0x54,0x68,0x65,0x20,0x64,0x6f,0x63,0x75,0x6d,0x65,0x6e,0x74,0x61,0x74,0x69,0x6f,0x6e,0x20,0x66,0x6f,0x72,0x20,0x5f,0x68,0x65,0x6c,0x6c,0x6f,0x77,0x6f,0x72,0x6c,0x64,0x5f,0x20,0x69,0x73,0x20,0x69,0x6e,0x20,0x61,0x20,0x66,0x6f,0x6c,0x64,0x65,0x72,0x20,0x63,0x61,0x6c,0x6c,0x65,0x64,0x20,0x64,0x6f,0x63,0x73,0x2e,0xa,0xa,0x60,0x60,0x60,0xa,0x20,0x20,0x20,0x20,0x70,0x6b,0x67,0x61,0x73,0x73,0x65,0x74,0x73,0x20,0x2d,0x6f,0x20,0x63,0x6d,0x64,0x2f,0x68,0x65,0x6c,0x6c,0x6f,0x77,0x6f,0x72,0x6c,0x64,0x2f,0x61,0x73,0x73,0x65,0x74,0x73,0x2e,0x67,0x6f,0x20,0x2d,0x70,0x20,0x6d,0x61,0x69,0x6e,0x20,0x2d,0x73,0x74,0x72,0x69,0x70,0x2d,0x70,0x72,0x65,0x66,0x69,0x78,0x3d,0x22,0x2f,0x22,0x20,0x48,0x65,0x6c,0x70,0x20,0x64,0x6f,0x63,0x73,0xa,0x60,0x60,0x60,0xa,0xa,0x54,0x68,0x69,0x73,0x20,0x77,0x69,0x6c,0x6c,0x20,0x63,0x72,0x65,0x61,0x74,0x65,0x20,0x74,0x68,0x65,0x20,0x5f,0x61,0x73,0x73,0x65,0x74,0x73,0x2e,0x67,0x6f,0x5f,0x20,0x66,0x69,0x6c,0x65,0x20,0x77,0x68,0x69,0x63,0x68,0x20,0x77,0x69,0x6c,0x6c,0x20,0x63,0x6f,0x6e,0x74,0x61,0x69,0x6e,0x20,0x61,0x20,0x6d,0x61,0x70,0x5b,0x73,0x74,0x72,0x69,0x6e,0x67,0x5d,0x5b,0x5d,0x62,0x79,0x74,0x65,0x20,0x6f,0x66,0x20,0x79,0x6f,0x75,0x72,0x20,0x68,0x65,0x6c,0x70,0x20,0x64,0x6f,0x63,0x73,0x2e,0xa,0x59,0x6f,0x75,0x20,0x63,0x61,0x6e,0x20,0x74,0x68,0x65,0x6e,0x20,0x63,0x6f,0x6d,0x70,0x69,0x6c,0x65,0x20,0x5f,0x68,0x65,0x6c,0x6c,0x6f,0x77,0x6f,0x72,0x6c,0x64,0x5f,0x20,0x6e,0x6f,0x72,0x6d,0x61,0x6c,0x79,0x20,0x77,0x69,0x74,0x68,0x20,0x5f,0x67,0x6f,0x5f,0x2e,0x20,0x4e,0x6f,0x74,0x65,0x20,0x74,0x68,0x65,0x20,0x70,0x6b,0x67,0x61,0x73,0x73,0x65,0x74,0x73,0x20,0x73,0x74,0x72,0x69,0x70,0x73,0x20,0x74,0x68,0x65,0x20,0x22,0x64,0x6f,0x63,0x73,0x22,0x20,0x66,0x72,0x6f,0x6d,0xa,0x74,0x68,0x65,0x20,0x76,0x61,0x6c,0x75,0x65,0x20,0x70,0x61,0x73,0x73,0x65,0x64,0x20,0x69,0x6e,0x20,0x61,0x73,0x20,0x74,0x68,0x65,0x20,0x6b,0x65,0x79,0x20,0x62,0x69,0x74,0x20,0x6e,0x6f,0x74,0x20,0x74,0x68,0x65,0x20,0x22,0x2f,0x22,0x2e,0x20,0x54,0x68,0x69,0x73,0x20,0x69,0x73,0x20,0x74,0x6f,0x20,0x73,0x75,0x70,0x70,0x6f,0x72,0x74,0x20,0x75,0x73,0x69,0x6e,0x67,0x20,0x60,0x6d,0x61,0x70,0x5b,0x73,0x74,0x72,0x69,0x6e,0x67,0x5d,0x5b,0x5d,0x62,0x79,0x74,0x65,0x60,0xa,0x61,0x73,0x20,0x68,0x6f,0x6c,0x64,0x65,0x72,0x73,0x20,0x6f,0x66,0x20,0x77,0x65,0x62,0x20,0x63,0x6f,0x6e,0x74,0x65,0x6e,0x74,0x2e,0x20,0x57,0x65,0x20,0x75,0x73,0x65,0x20,0x74,0x68,0x65,0x20,0x61,0x64,0x64,0x69,0x74,0x69,0x6f,0x6e,0x61,0x6c,0x20,0x6f,0x70,0x74,0x69,0x6f,0x6e,0x20,0x22,0x2d,0x73,0x74,0x72,0x69,0x70,0x2d,0x70,0x72,0x65,0x66,0x69,0x78,0x22,0x20,0x74,0x6f,0x20,0x72,0x65,0x6d,0x6f,0x76,0x65,0x20,0x74,0x68,0x65,0x20,0x6c,0x65,0x61,0x64,0x69,0x6e,0x67,0xa,0x73,0x6c,0x61,0x73,0x68,0x20,0x6c,0x65,0x61,0x76,0x69,0x6e,0x67,0x20,0x74,0x68,0x65,0x20,0x72,0x65,0x6e,0x61,0x6d,0x69,0x6e,0x67,0x20,0x66,0x69,0x6c,0x65,0x6e,0x61,0x6d,0x65,0x20,0x61,0x73,0x20,0x74,0x68,0x65,0x20,0x6b,0x65,0x79,0x20,0x69,0x6e,0x20,0x66,0x6f,0x72,0x20,0x74,0x68,0x65,0x20,0x6d,0x61,0x70,0x70,0x65,0x64,0x20,0x68,0x65,0x6c,0x70,0x20,0x70,0x61,0x67,0x65,0x2e,0xa,0xa},

    "htdocs": []byte{0xa,0x23,0x20,0x54,0x75,0x72,0x6e,0x69,0x6e,0x67,0x20,0x61,0x6e,0x20,0x68,0x74,0x64,0x6f,0x63,0x73,0x20,0x64,0x69,0x72,0x65,0x63,0x74,0x6f,0x72,0x79,0x20,0x69,0x6e,0x74,0x6f,0x20,0x61,0x6e,0x20,0x61,0x73,0x73,0x65,0x74,0xa,0xa,0x53,0x61,0x79,0x20,0x79,0x6f,0x75,0x20,0x77,0x61,0x6e,0x74,0x20,0x74,0x6f,0x20,0x63,0x72,0x65,0x61,0x74,0x65,0x20,0x61,0x20,0x73,0x74,0x61,0x6e,0x64,0x61,0x6c,0x6f,0x6e,0x65,0x20,0x62,0x69,0x6e,0x61,0x72,0x79,0x20,0x63,0x6f,0x6e,0x74,0x61,0x69,0x6e,0x69,0x6e,0x67,0x20,0x74,0x68,0x65,0x20,0x63,0x6f,0x6e,0x74,0x65,0x6e,0x74,0x73,0x20,0x6f,0x66,0x20,0x61,0x6e,0x20,0x68,0x74,0x64,0x6f,0x63,0x73,0x20,0x64,0x69,0x72,0x65,0x63,0x74,0x6f,0x72,0x79,0x2e,0x20,0x54,0x68,0x65,0x20,0x70,0x72,0x6f,0x67,0x72,0x61,0x6d,0xa,0x6e,0x61,0x6d,0x65,0x20,0x77,0x69,0x6c,0x6c,0x20,0x62,0x65,0x20,0x63,0x61,0x6c,0x6c,0x65,0x64,0x20,0x5f,0x68,0x65,0x6c,0x6c,0x6f,0x73,0x65,0x72,0x76,0x65,0x72,0x5f,0x20,0x77,0x69,0x74,0x68,0x20,0x69,0x74,0x73,0x20,0x6d,0x61,0x69,0x6e,0x20,0x64,0x65,0x66,0x69,0x6e,0x65,0x64,0x20,0x69,0x6e,0x20,0x2a,0x63,0x6d,0x64,0x2f,0x68,0x65,0x6c,0x6c,0x6f,0x73,0x65,0x72,0x76,0x65,0x72,0x2f,0x68,0x65,0x6c,0x6c,0x6f,0x73,0x65,0x72,0x76,0x65,0x72,0x2e,0x67,0x6f,0x2a,0x2e,0xa,0xa,0x42,0x75,0x69,0x6c,0x64,0x69,0x6e,0x67,0x20,0x74,0x68,0x65,0x20,0x70,0x61,0x63,0x6b,0x61,0x67,0x65,0x20,0x66,0x6f,0x72,0x20,0x74,0x68,0x65,0x20,0x61,0x6c,0x6c,0x20,0x74,0x68,0x65,0x20,0x64,0x69,0x72,0x65,0x63,0x74,0x6f,0x72,0x79,0x20,0x28,0x69,0x6e,0x63,0x6c,0x75,0x64,0x69,0x6e,0x67,0x20,0x73,0x75,0x62,0x20,0x64,0x69,0x72,0x65,0x63,0x74,0x6f,0x72,0x69,0x65,0x73,0x20,0x6f,0x66,0x20,0x68,0x74,0x64,0x6f,0x63,0x73,0x20,0x63,0x61,0x6e,0x20,0x62,0x65,0x20,0x64,0x6f,0x6e,0x65,0xa,0x6c,0x69,0x6b,0x65,0xa,0xa,0x60,0x60,0x60,0xa,0x20,0x20,0x20,0x20,0x70,0x6b,0x67,0x61,0x73,0x73,0x65,0x74,0x73,0x20,0x2d,0x6f,0x20,0x63,0x6d,0x64,0x2f,0x68,0x65,0x6c,0x6c,0x6f,0x73,0x65,0x72,0x76,0x65,0x72,0x2f,0x61,0x73,0x73,0x65,0x74,0x73,0x2e,0x67,0x6f,0x20,0x2d,0x70,0x20,0x6d,0x61,0x69,0x6e,0x20,0x48,0x74,0x64,0x6f,0x63,0x73,0x20,0x68,0x74,0x64,0x6f,0x63,0x73,0xa,0x60,0x60,0x60,0xa,0xa,0x49,0x6e,0x20,0x74,0x68,0x65,0x20,0x66,0x69,0x6c,0x65,0x20,0x2a,0x68,0x65,0x6c,0x6c,0x6f,0x73,0x65,0x72,0x76,0x65,0x72,0x2e,0x67,0x6f,0x2a,0x20,0x79,0x6f,0x75,0x27,0x64,0x20,0x72,0x65,0x66,0x65,0x72,0x65,0x6e,0x63,0x65,0x20,0x74,0x68,0x65,0x20,0x63,0x6f,0x6e,0x74,0x61,0x69,0x6e,0x73,0x20,0x6f,0x66,0x20,0x74,0x68,0x65,0x20,0x70,0x61,0x63,0x6b,0x61,0x67,0x65,0x20,0x76,0x61,0x72,0x69,0x61,0x62,0x6c,0x65,0x20,0x61,0x73,0x20,0x60,0x6d,0x61,0x69,0x6e,0x2e,0x48,0x74,0x64,0x6f,0x63,0x73,0x60,0x20,0x70,0x61,0x73,0x73,0x69,0x6e,0x67,0x20,0x69,0x6e,0x20,0x74,0x68,0x65,0xa,0x70,0x61,0x74,0x68,0x20,0x72,0x65,0x71,0x75,0x65,0x73,0x74,0x65,0x64,0x20,0x62,0x79,0x20,0x79,0x6f,0x75,0x72,0x20,0x73,0x65,0x72,0x76,0x65,0x72,0xa,0xa,0x60,0x60,0x60,0xa,0x20,0x20,0x20,0x20,0x66,0x75,0x6e,0x63,0x20,0x48,0x61,0x6e,0x64,0x6c,0x65,0x48,0x74,0x64,0x6f,0x63,0x28,0x72,0x65,0x73,0x20,0x68,0x74,0x74,0x70,0x2e,0x52,0x65,0x73,0x70,0x6f,0x6e,0x73,0x65,0x57,0x72,0x69,0x74,0x65,0x72,0x2c,0x20,0x72,0x65,0x71,0x20,0x2a,0x68,0x74,0x74,0x70,0x2e,0x52,0x65,0x71,0x75,0x65,0x73,0x74,0x29,0x20,0x7b,0xa,0x20,0x20,0x20,0x20,0x20,0x20,0x20,0x20,0x70,0x20,0x3a,0x3d,0x20,0x72,0x65,0x71,0x2e,0x55,0x52,0x4c,0x2e,0x50,0x61,0x74,0x68,0xa,0x20,0x20,0x20,0x20,0x20,0x20,0x20,0x20,0x69,0x66,0x20,0x62,0x75,0x66,0x2c,0x20,0x6f,0x6b,0x20,0x3a,0x3d,0x20,0x6d,0x61,0x69,0x6e,0x2e,0x48,0x74,0x64,0x6f,0x63,0x73,0x5b,0x70,0x5d,0x3b,0x20,0x6f,0x6b,0x20,0x3d,0x3d,0x20,0x74,0x72,0x75,0x65,0x20,0x7b,0xa,0x20,0x20,0x20,0x20,0x20,0x20,0x20,0x20,0x20,0x20,0x20,0x20,0x2f,0x2f,0x20,0x50,0x72,0x6f,0x62,0x61,0x62,0x6c,0x65,0x20,0x77,0x61,0x6e,0x74,0x20,0x74,0x6f,0x20,0x73,0x65,0x74,0x20,0x43,0x6f,0x6e,0x74,0x65,0x6e,0x74,0x2d,0x54,0x79,0x70,0x65,0x2c,0x20,0x65,0x74,0x63,0x20,0x62,0x65,0x66,0x6f,0x72,0x65,0x20,0x68,0x61,0x6e,0x64,0x69,0x6e,0x67,0x20,0x62,0x61,0x63,0x6b,0x20,0x74,0x68,0x65,0x20,0x64,0x61,0x74,0x61,0xa,0x20,0x20,0x20,0x20,0x20,0x20,0x20,0x20,0x20,0x20,0x20,0x20,0x69,0x6f,0x2e,0x57,0x72,0x69,0x74,0x65,0x28,0x72,0x65,0x73,0x2c,0x20,0x62,0x75,0x66,0x29,0xa,0x20,0x20,0x20,0x20,0x20,0x20,0x20,0x20,0x7d,0x20,0x65,0x6c,0x73,0x65,0x20,0x7b,0xa,0x20,0x20,0x20,0x20,0x20,0x20,0x20,0x20,0x20,0x20,0x20,0x20,0x2f,0x2f,0x20,0x48,0x61,0x6e,0x64,0x6c,0x65,0x20,0x79,0x6f,0x75,0x72,0x20,0x65,0x72,0x72,0x6f,0x72,0x20,0x68,0x65,0x72,0x65,0xa,0x20,0x20,0x20,0x20,0x20,0x20,0x20,0x20,0x7d,0xa,0x20,0x20,0x20,0x20,0x7d,0xa,0x60,0x60,0x60,0xa,0xa,0x4e,0x6f,0x74,0x69,0x63,0x65,0x20,0x74,0x68,0x65,0x20,0x5f,0x70,0x6b,0x67,0x61,0x73,0x73,0x65,0x74,0x73,0x5f,0x20,0x62,0x79,0x20,0x64,0x65,0x66,0x61,0x75,0x6c,0x74,0x20,0x73,0x74,0x72,0x69,0x70,0x73,0x20,0x74,0x68,0x65,0x20,0x69,0x6e,0x69,0x74,0x69,0x61,0x6c,0x20,0x64,0x69,0x72,0x65,0x63,0x74,0x6f,0x72,0x79,0x20,0x6e,0x61,0x6d,0x65,0x20,0x66,0x72,0x6f,0x6d,0x20,0x74,0x68,0x65,0x20,0x70,0x61,0x74,0x68,0x20,0x6f,0x66,0x20,0x74,0x68,0x65,0x20,0x76,0x61,0x6c,0x75,0x65,0x20,0x73,0x74,0x6f,0x72,0x65,0x64,0x2e,0x20,0x54,0x68,0x69,0x73,0x20,0x69,0x73,0xa,0x73,0x6f,0x20,0x74,0x68,0x65,0x20,0x70,0x61,0x74,0x68,0x20,0x6d,0x61,0x74,0x63,0x68,0x65,0x73,0x20,0x65,0x61,0x73,0x69,0x6c,0x79,0x20,0x77,0x68,0x61,0x74,0x20,0x69,0x73,0x20,0x70,0x61,0x73,0x73,0x65,0x64,0x20,0x69,0x6e,0x20,0x76,0x69,0x61,0x20,0x60,0x72,0x65,0x71,0x2e,0x55,0x52,0x4c,0x2e,0x50,0x61,0x74,0x68,0x60,0x2e,0xa},

	}

)

