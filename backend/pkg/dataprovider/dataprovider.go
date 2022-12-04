package dataprovider

import whensmybin "github.com/aaronclaydon/whensmybin/pkg"

type DataProvider interface {
	ProviderName() string
	GetBinData(whensmybin.Address) BinData
}
