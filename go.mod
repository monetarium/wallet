module github.com/monetarium/monetarium-wallet

go 1.23

require (
	decred.org/cspp/v2 v2.4.0
	github.com/decred/dcrd/txscript/v4 v4.1.1
	github.com/decred/go-socks v1.1.0
	github.com/decred/slog v1.2.0
	github.com/decred/vspd/client/v4 v4.0.1
	github.com/decred/vspd/types/v3 v3.0.0
	github.com/gorilla/websocket v1.5.1
	github.com/jessevdk/go-flags v1.5.0
	github.com/jrick/bitset v1.0.0
	github.com/jrick/logrotate v1.0.0
	github.com/jrick/wsrpc/v2 v2.3.8
	github.com/monetarium/monetarium-node/addrmgr v1.0.6
	github.com/monetarium/monetarium-node/blockchain v1.0.6
	github.com/monetarium/monetarium-node/blockchain/stake v1.0.6
	github.com/monetarium/monetarium-node/blockchain/standalone v1.0.6
	github.com/monetarium/monetarium-node/certgen v1.0.6
	github.com/monetarium/monetarium-node/chaincfg v1.0.6
	github.com/monetarium/monetarium-node/chaincfg/chainhash v1.0.6
	github.com/monetarium/monetarium-node/cointype v1.0.6
	github.com/monetarium/monetarium-node/connmgr v1.0.6
	github.com/monetarium/monetarium-node/crypto/blake256 v1.0.6
	github.com/monetarium/monetarium-node/crypto/rand v1.0.6
	github.com/monetarium/monetarium-node/crypto/ripemd160 v1.0.6
	github.com/monetarium/monetarium-node/dcrec v1.0.6
	github.com/monetarium/monetarium-node/dcrec/secp256k1 v1.0.6
	github.com/monetarium/monetarium-node/dcrjson v1.0.6
	github.com/monetarium/monetarium-node/dcrutil v1.0.6
	github.com/monetarium/monetarium-node/gcs v1.0.6
	github.com/monetarium/monetarium-node/hdkeychain v1.0.6
	github.com/monetarium/monetarium-node/mixing v1.0.6
	github.com/monetarium/monetarium-node/rpc/jsonrpc/types v1.0.6
	github.com/monetarium/monetarium-node/rpcclient v1.0.6
	github.com/monetarium/monetarium-node/txscript v1.0.6
	github.com/monetarium/monetarium-node/wire v1.0.6
	go.etcd.io/bbolt v1.3.11
	golang.org/x/crypto v0.33.0
	golang.org/x/sync v0.11.0
	golang.org/x/term v0.29.0
	google.golang.org/grpc v1.71.0
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.5.1
	google.golang.org/protobuf v1.36.5
)

require (
	github.com/agl/ed25519 v0.0.0-20170116200512-5312a6153412 // indirect
	github.com/companyzero/sntrup4591761 v0.0.0-20220309191932-9e0f3af2f07a // indirect
	github.com/dchest/siphash v1.2.3 // indirect
	github.com/decred/base58 v1.0.6 // indirect
	github.com/decred/dcrd/chaincfg/chainhash v1.0.4 // indirect
	github.com/decred/dcrd/crypto/blake256 v1.1.0 // indirect
	github.com/decred/dcrd/crypto/ripemd160 v1.0.2 // indirect
	github.com/decred/dcrd/dcrec v1.0.1 // indirect
	github.com/decred/dcrd/dcrec/edwards/v2 v2.0.3 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.3.0 // indirect
	github.com/decred/dcrd/wire v1.7.0 // indirect
	github.com/klauspost/cpuid/v2 v2.2.8 // indirect
	github.com/monetarium/monetarium-node/container/lru v1.0.6 // indirect
	github.com/monetarium/monetarium-node/database v1.0.6 // indirect
	github.com/monetarium/monetarium-node/dcrec/edwards v1.0.6 // indirect
	golang.org/x/net v0.34.0 // indirect
	golang.org/x/sys v0.30.0 // indirect
	golang.org/x/text v0.22.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250115164207-1a7da9e5054f // indirect
	lukechampine.com/blake3 v1.3.0 // indirect
)

// Local development: replace with local monetarium-node packages
// TODO: Remove these replace directives when packages are published to GitHub with new tags
replace (
	github.com/monetarium/monetarium-node/addrmgr => ../monetarium-node/addrmgr
	github.com/monetarium/monetarium-node/blockchain => ../monetarium-node/blockchain
	github.com/monetarium/monetarium-node/blockchain/stake => ../monetarium-node/blockchain/stake
	github.com/monetarium/monetarium-node/blockchain/standalone => ../monetarium-node/blockchain/standalone
	github.com/monetarium/monetarium-node/certgen => ../monetarium-node/certgen
	github.com/monetarium/monetarium-node/chaincfg => ../monetarium-node/chaincfg
	github.com/monetarium/monetarium-node/chaincfg/chainhash => ../monetarium-node/chaincfg/chainhash
	github.com/monetarium/monetarium-node/cointype => ../monetarium-node/cointype
	github.com/monetarium/monetarium-node/connmgr => ../monetarium-node/connmgr
	github.com/monetarium/monetarium-node/container/lru => ../monetarium-node/container/lru
	github.com/monetarium/monetarium-node/crypto/blake256 => ../monetarium-node/crypto/blake256
	github.com/monetarium/monetarium-node/crypto/rand => ../monetarium-node/crypto/rand
	github.com/monetarium/monetarium-node/crypto/ripemd160 => ../monetarium-node/crypto/ripemd160
	github.com/monetarium/monetarium-node/database => ../monetarium-node/database
	github.com/monetarium/monetarium-node/dcrec => ../monetarium-node/dcrec
	github.com/monetarium/monetarium-node/dcrec/edwards => ../monetarium-node/dcrec/edwards
	github.com/monetarium/monetarium-node/dcrec/secp256k1 => ../monetarium-node/dcrec/secp256k1
	github.com/monetarium/monetarium-node/dcrjson => ../monetarium-node/dcrjson
	github.com/monetarium/monetarium-node/dcrutil => ../monetarium-node/dcrutil
	github.com/monetarium/monetarium-node/gcs => ../monetarium-node/gcs
	github.com/monetarium/monetarium-node/hdkeychain => ../monetarium-node/hdkeychain
	github.com/monetarium/monetarium-node/mixing => ../monetarium-node/mixing
	github.com/monetarium/monetarium-node/rpc/jsonrpc/types => ../monetarium-node/rpc/jsonrpc/types
	github.com/monetarium/monetarium-node/rpcclient => ../monetarium-node/rpcclient
	github.com/monetarium/monetarium-node/txscript => ../monetarium-node/txscript
	github.com/monetarium/monetarium-node/wire => ../monetarium-node/wire
)
