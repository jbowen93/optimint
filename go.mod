module github.com/celestiaorg/optimint

go 1.15

require (
	github.com/DataDog/zstd v1.4.5 // indirect
	github.com/dgraph-io/badger/v3 v3.2103.2
	github.com/dgryski/go-farm v0.0.0-20200201041132-a6ae2369ad13 // indirect
	github.com/go-kit/kit v0.12.0
	github.com/gogo/protobuf v1.3.2
	github.com/golang/protobuf v1.5.2
	github.com/gopherjs/gopherjs v0.0.0-20190812055157-5d271430af9f // indirect
	github.com/gorilla/rpc v1.2.0
	github.com/ipfs/go-log v1.0.5
	github.com/libp2p/go-libp2p v0.15.1
	github.com/libp2p/go-libp2p-core v0.9.0
	github.com/libp2p/go-libp2p-discovery v0.5.1
	github.com/libp2p/go-libp2p-kad-dht v0.15.0
	github.com/libp2p/go-libp2p-pubsub v0.5.6
	github.com/libp2p/go-libp2p-quic-transport v0.12.0 // indirect
	github.com/minio/sha256-simd v1.0.0
	github.com/multiformats/go-multiaddr v0.4.1
	github.com/prometheus/client_golang v1.11.0
	github.com/rs/cors v1.8.0
	github.com/smartystreets/assertions v1.0.1 // indirect
	github.com/spf13/cobra v1.2.1
	github.com/spf13/viper v1.9.0
	github.com/stretchr/objx v0.2.0 // indirect
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/tendermint v0.35.0
	go.uber.org/multierr v1.7.0
	golang.org/x/net v0.0.0-20211005001312-d4b1ae081e3b
	google.golang.org/grpc v1.42.0
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
