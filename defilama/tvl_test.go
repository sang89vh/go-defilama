package defilama

import (
	"fmt"
	"testing"
)

func TestProtocol(t *testing.T) {
	t.Run("Nornal case", func(t *testing.T) {
		tvl := DefaultTvl()
		resp, _ := tvl.Protocol("aave")
		fmt.Printf("protocol now is %+v", resp)
	})
}

func TestGetAllProtocol(t *testing.T) {
	t.Run("Nornal case", func(t *testing.T) {
		tvl := DefaultTvl()
		resp, err := tvl.GetAllProtocol()
		if err != nil {
			fmt.Printf("protocol has an error %+v", err.Error())
		} else {
			for i, p := range resp {
				fmt.Printf("%d protocol now is %+v \n\r", i, p)
			}
		}
	})
}

func TestChart(t *testing.T) {
	t.Run("Nornal case", func(t *testing.T) {
		tvl := DefaultTvl()
		resp, err := tvl.Chart()
		if err != nil {
			fmt.Printf("protocol has an error %+v", err.Error())
		} else {
			for i, p := range resp {
				fmt.Printf("%d chart %+v \n\r", i, p)
			}
		}
	})
}
func TestChartChain(t *testing.T) {
	t.Run("Nornal case", func(t *testing.T) {
		tvl := DefaultTvl()
		resp, err := tvl.ChartChain("Ethereum")
		if err != nil {
			fmt.Printf("protocol has an error %+v", err.Error())
		} else {
			for i, p := range resp {
				fmt.Printf("%d chart %+v \n\r", i, p)
			}
		}
	})
}

func TestChains(t *testing.T) {
	t.Run("Nornal case", func(t *testing.T) {
		tvl := DefaultTvl()
		resp, err := tvl.Chains()
		if err != nil {
			fmt.Printf("protocol has an error %+v", err.Error())
		} else {
			for i, p := range resp {
				fmt.Printf("%d chart %+v \n\r", i, p)
			}
		}
	})
}
