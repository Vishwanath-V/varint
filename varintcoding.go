package varint

const mask = uint64(1<<7) - 1

func EncodeUInt64(input uint64) []byte {
	res := []byte{} //bytes to hold the encoded bits
	for {
		res = append(res, byte(input&mask))
		input >>= 7 //rightshift processed bits
		if input > 0 {
			res[len(res)-1] |= 1 << 7 //lasft byte in the result needs to be or'ed with 256 (10000000) to set the 8th bit flag if there are more bytes to be processed
		} else {
			break //no more bytes to process
		}
	}
	return res //return little endian format of the output
}

func DecodeToUInt64(input []byte) uint64 {
	var o uint64 = 0
	for i := 0; i < len(input); i++ {
		o |= ((uint64(input[i]) & mask) << (7 * i))
	}
	return o
}
