package whitespacews

import (
	whensmybin "github.com/aaronclaydon/whensmybin/pkg"
	"github.com/aaronclaydon/whensmybin/pkg/dataprovider"
)

type WhiteSpaceWSProvider struct {
	Endpoint string
}

func (provider WhiteSpaceWSProvider) ProviderName() string {
	return provider.Endpoint
}

func (provider WhiteSpaceWSProvider) GetBinData(whensmybin.Address) dataprovider.BinData {
	return dataprovider.BinData{}
}
