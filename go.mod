module github.com/comrade-coop/scynet-blockchain

go 1.12

// see also https://github.com/cosmos/cosmos-sdk/issues/3129
replace golang.org/x/crypto => github.com/tendermint/crypto v0.0.0-20180820045704-3764759f34a5

require (
	github.com/bartekn/go-bip39 v0.0.0-20171116152956-a05967ea095d // indirect
	github.com/bgentry/speakeasy v0.1.0 // indirect
	github.com/btcsuite/btcd v0.0.0-20190315201642-aa6e0f35703c // indirect
	github.com/btcsuite/btcutil v0.0.0-20190316010144-3ac1210f4b38 // indirect
	github.com/cosmos/cosmos-sdk v0.32.0
	github.com/cosmos/go-bip39 v0.0.0-20180819234021-555e2067c45d // indirect
	github.com/go-kit/kit v0.8.0 // indirect
	github.com/go-logfmt/logfmt v0.4.0 // indirect
	github.com/gogo/protobuf v1.2.1 // indirect
	github.com/golang/protobuf v1.3.1
	github.com/google/uuid v1.1.1
	github.com/gorilla/mux v1.7.0 // indirect
	github.com/gorilla/websocket v1.4.0 // indirect
	github.com/mattn/go-isatty v0.0.7 // indirect
	github.com/pkg/errors v0.8.1 // indirect
	github.com/prometheus/client_golang v0.9.2 // indirect
	github.com/rakyll/statik v0.1.5 // indirect
	github.com/rcrowley/go-metrics v0.0.0-20181016184325-3113b8401b8a // indirect
	github.com/rs/cors v1.6.0 // indirect
	github.com/spf13/cobra v0.0.3
	github.com/spf13/viper v1.3.2 // indirect
	github.com/syndtr/goleveldb v1.0.0 // indirect
	github.com/tendermint/btcd v0.1.1 // indirect
	github.com/tendermint/go-amino v0.14.1
	github.com/tendermint/iavl v0.12.2 // indirect
	github.com/tendermint/tendermint v0.30.2
	golang.org/x/net v0.0.0-20190313220215-9f648a60d977 // indirect
	google.golang.org/grpc v1.19.0
)
