package defilama

import (
	"encoding/json"
	"fmt"
	"github.com/sang89vh/go-defilama-sdk/pkg/restclient"
	"github.com/sang89vh/go-defilama-sdk/types"
	"sync"
)

const ApiBasePath = "https://api.llama.fi/"

type Tvl struct {
	path string
}

var one sync.Once
var instance *Tvl

func DefaultTvl() *Tvl {
	one.Do(func() {
		i := Tvl{path: "protocol/aave"}
		instance = &i

	})

	return instance
}
func (tvl *Tvl) Protocol() (*types.Protocol, error) {
	url := fmt.Sprintf("%s%s", ApiBasePath, tvl.path)
	resp, err := restclient.MakeReq(url)
	if err != nil {
		return nil, err
	}
	var data *types.Protocol
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
