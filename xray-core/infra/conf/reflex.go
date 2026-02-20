package conf

import (
	"github.com/xtls/xray-core/common/errors"
	"github.com/xtls/xray-core/proxy/reflex"
	"google.golang.org/protobuf/proto"
)

type ReflexUserConfig struct {
	ID     string `json:"id"`
	Policy string `json:"policy"`
}

type ReflexFallbackConfig struct {
	Dest uint32 `json:"dest"`
}

type ReflexInboundConfig struct {
	Clients  []*ReflexUserConfig   `json:"clients"`
	Fallback *ReflexFallbackConfig `json:"fallback"`
}

func (c *ReflexInboundConfig) Build() (proto.Message, error) {
	config := &reflex.InboundConfig{}

	for _, rawUser := range c.Clients {
		if rawUser.ID == "" {
			return nil, errors.New("Reflex client: missing id")
		}
		config.Clients = append(config.Clients, &reflex.User{
			Id:     rawUser.ID,
			Policy: rawUser.Policy,
		})
	}

	if c.Fallback != nil {
		config.Fallback = &reflex.Fallback{
			Dest: c.Fallback.Dest,
		}
	}

	return config, nil
}

type ReflexOutboundConfig struct {
	Address string `json:"address"`
	Port    uint32 `json:"port"`
	ID      string `json:"id"`
	Policy  string `json:"policy"`
}

func (c *ReflexOutboundConfig) Build() (proto.Message, error) {
	if c.Address == "" {
		return nil, errors.New("Reflex outbound: missing server address")
	}
	if c.Port == 0 {
		return nil, errors.New("Reflex outbound: missing server port")
	}
	if c.ID == "" {
		return nil, errors.New("Reflex outbound: missing client id")
	}

	return &reflex.OutboundConfig{
		Address: c.Address,
		Port:    c.Port,
		Id:      c.ID,
		Policy:  c.Policy,
	}, nil
}
