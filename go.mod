module github.com/konstellation/konstellation

go 1.15

require (
	github.com/CosmWasm/wasmd v0.15.0
	github.com/Masterminds/semver v1.5.0 // indirect
	github.com/Masterminds/vcs v1.13.1 // indirect
	github.com/allegro/bigcache v1.2.1 // indirect
	github.com/aristanetworks/goarista v0.0.0-20191001182449-186a6201b8ef // indirect
	github.com/armon/go-radix v1.0.0 // indirect
	github.com/bartekn/go-bip39 v0.0.0-20171116152956-a05967ea095d // indirect
	github.com/boltdb/bolt v1.3.1 // indirect
	github.com/cespare/cp v1.1.1 // indirect
	github.com/cosmos/cosmos-sdk v0.41.0
	github.com/deckarep/golang-set v1.7.1 // indirect
	github.com/elastic/gosigar v0.10.5 // indirect
	github.com/ethereum/go-ethereum v1.9.11 // indirect
	github.com/fsnotify/fsnotify v1.4.9
	github.com/golang/dep v0.5.4 // indirect
	github.com/google/uuid v1.1.2 // indirect
	github.com/gorilla/mux v1.8.0
	github.com/hashicorp/go-version v1.2.0 // indirect
	github.com/jmank88/nuts v0.4.0 // indirect
	github.com/konstellation/kn-sdk v0.1.14
	github.com/mitchellh/gox v1.0.1 // indirect
	github.com/mitchellh/mapstructure v1.1.2
	github.com/nightlyone/lockfile v0.0.0-20200124072040-edb130adc195 // indirect
	github.com/pborman/uuid v1.2.0 // indirect
	github.com/pkg/errors v0.9.1
	github.com/rakyll/statik v0.1.7
	github.com/rjeczalik/notify v0.9.2 // indirect
	github.com/rs/cors v1.7.0
	github.com/rs/zerolog v1.20.0
	github.com/sdboyer/constext v0.0.0-20170321163424-836a14457353 // indirect
	github.com/spf13/cast v1.3.1
	github.com/spf13/cobra v1.1.1
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.7.1
	github.com/stevenmatthewt/semantics v2.0.4+incompatible // indirect
	github.com/stretchr/testify v1.7.0
	github.com/stumble/gorocksdb v0.0.3 // indirect
	github.com/tendermint/crypto v0.0.0-20191022145703-50d29ede1e15 // indirect
	github.com/tendermint/go-amino v0.16.0
	github.com/tendermint/iavl v0.12.4 // indirect
	github.com/tendermint/tendermint v0.34.3
	github.com/tendermint/tm-db v0.6.3
	go.opencensus.io v0.22.6 // indirect
	golang.org/x/sys v0.0.0-20210124154548-22da62e12c0c // indirect
)

replace github.com/konstellation/kn-sdk => ../kn-sdk

replace google.golang.org/grpc => google.golang.org/grpc v1.33.2

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
