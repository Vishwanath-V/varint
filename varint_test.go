package varint

import (
	"fmt"
	"testing"
)

func TestVarIntEncoding(t *testing.T) {
	fmt.Println(EncodeUInt64(2334))
	fmt.Println(DecodeToUInt64(EncodeUInt64(2334)))
}
