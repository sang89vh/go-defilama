package defilama

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sang89vh/go-defilama-sdk/pkg/restclient"
	"github.com/sang89vh/go-defilama-sdk/types"
	"log"
	"os"
	"sync"
)

const (
	pathProtocolList   = "protocols"
	pathProtocolDetail = "protocol/%s"
	pathChart          = "charts"
	pathChartDetail    = "charts/%s"
	pathTvlProtocol    = "tvl/%s"
	pathChains         = "chains"
)

type Tvl struct {
	ApiBasePath string
}

var one sync.Once
var instance *Tvl

func DefaultTvl() *Tvl {
	one.Do(func() {
		i := Tvl{}
		// load .env file
		err := godotenv.Load("./../.env")

		if err != nil {
			log.Fatalf("Error loading .env file")
		}
		i.ApiBasePath = os.Getenv("DEFILAMA_API_URL")
		instance = &i

	})

	return instance
}
func (tvl *Tvl) GetAllProtocol() ([]types.Protocol, error) {
	url := fmt.Sprintf("%s%s", tvl.ApiBasePath, pathProtocolList)
	resp, err := restclient.MakeReq(url)
	if err != nil {
		return nil, err
	}
	var data []types.Protocol
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (tvl *Tvl) Protocol(protolcol string) (*types.ProtocolDetail, error) {
	url := fmt.Sprintf(pathProtocolDetail, protolcol)
	url = fmt.Sprintf("%s%s", tvl.ApiBasePath, url)
	resp, err := restclient.MakeReq(url)
	if err != nil {
		return nil, err
	}
	var data *types.ProtocolDetail
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (tvl *Tvl) Chart() ([]types.Tvl, error) {
	url := fmt.Sprintf("%s%s", tvl.ApiBasePath, pathChart)
	resp, err := restclient.MakeReq(url)
	if err != nil {
		return nil, err
	}
	var data []types.Tvl
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (tvl *Tvl) ChartChain(chain string) ([]types.Tvl, error) {
	url := fmt.Sprintf(pathChartDetail, chain)
	url = fmt.Sprintf("%s%s", tvl.ApiBasePath, url)
	resp, err := restclient.MakeReq(url)
	if err != nil {
		return nil, err
	}
	var data []types.Tvl
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (tvl *Tvl) Chains() ([]types.ChainTvlNow, error) {
	url := fmt.Sprintf("%s%s", tvl.ApiBasePath, pathChains)
	resp, err := restclient.MakeReq(url)
	if err != nil {
		return nil, err
	}
	var data []types.ChainTvlNow
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
