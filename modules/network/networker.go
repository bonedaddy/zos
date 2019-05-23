package network

import (
	"net"

	"github.com/threefoldtech/zosv2/modules"
)

type networker struct {
	netResAlloc NetResourceAllocator
}

func NewNetworker(allocator NetResourceAllocator) modules.Networker {
	return &networker{allocator}
}

var _ modules.Networker = (*networker)(nil)

func (n *networker) GetNetResource(id string) (modules.NetResource, error) {
	// TODO check signature
	return n.netResAlloc.Get(id)
}

func (n *networker) ApplyNetResource(resource modules.NetResource) error {
	return nil
}

type NetResourceAllocator interface {
	Get(txID string) (modules.NetResource, error)
}

// type httpNetResourceAllocator struct {
// 	baseURL string
// }

// func NewHTTPNetResourceAllocator(url string) *httpNetResourceAllocator {
// 	return &httpNetResourceAllocator{url}
// }

// func (a *httpNetResourceAllocator) Get(txID string) (modules.NetResource, error) {
// 	netRes := modules.NetResource{}

// 	url, err := joinURL(a.baseURL, txID)

// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return netRes, err
// 	}
// 	defer resp.Body.Close()

// 	if err := json.NewDecoder(resp.Body).Decode(&netRes); err != nil {
// 		return netRes, err
// 	}

// 	return netRes, nil
// }

// func joinURL(base, path string) (string, error) {
// 	u, err := url.Parse(base)
// 	if err != nil {
// 		return "nil", err
// 	}
// 	u.Path = filepath.Join(u.Path, path)
// 	return u.String(), nil
// }

type TestNetResourceAllocator struct {
	Resource modules.NetResource
}

func NewTestNetResourceAllocator() NetResourceAllocator {
	return &TestNetResourceAllocator{
		Resource: modules.NetResource{
			NodeID: modules.NodeID("supernode"),
			Prefix: MustParseCIDR("2a02:1802:5e:f002::/64"),
			Connected: []modules.Connected{
				{
					Type:   modules.ConnTypeWireguard,
					Prefix: MustParseCIDR("2a02:1802:5e:cc02::/64"),
					Connection: modules.Wireguard{
						NICName: "cc02",
						Peer:    net.ParseIP("2001::1"),
						Key:     []byte{},
						// LinkLocal: net.
					},
				},
			},
		},
	}
}

func (a *TestNetResourceAllocator) Get(txID string) (modules.NetResource, error) {
	return a.Resource, nil
}

func MustParseCIDR(cidr string) net.IPNet {
	_, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		panic(err)
	}
	return *ipnet
}