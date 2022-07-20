package defilama

import (
	"fmt"
	"testing"
)

func TestProtocol(t *testing.T) {
	t.Run("Nornal case", func(t *testing.T) {
		tvl := DefaultTvl()
		resp, _ := tvl.Protocol()
		fmt.Printf("protocol now is %+v", resp)
	})
}
