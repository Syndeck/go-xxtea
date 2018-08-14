package xxtea

func bytesToInt32(data []byte) []int32 {
	if l := len(data); l%4 != 0 {
		data = append(data, 0, 0, 0, 0)
	}

	// Unsigned => Signed
	signed := make([]int32, len(data))
	for i, v := range data {
		signed[i] = int32(v)
	}

	rv := make([]int32, len(signed)/4)

	for i := 0; i < len(signed)/4; i++ {
		rv[i] = int32(data[i*4 + 3]) | int32(data[i*4 + 2]) << 8 | int32(data[i*4 + 1]) << 16 | int32(data[i*4 + 0]) << 24
	}
	return rv
}

func int32ToBytes(data []int32) []byte {
	l := len(data)
	rv := make([]byte, 4*l)
	for i := 0; i < len(data); i++ {
		rv[4*i + 3] = byte(data[i])
		rv[4*i + 2] = byte(data[i] >> 8)
		rv[4*i + 1] = byte(data[i] >> 16)
		rv[4*i + 0] = byte(data[i] >> 24)
	}
	return rv
}