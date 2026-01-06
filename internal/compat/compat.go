package compat

import (
	"github.com/monetarium/node/blockchain/standalone"
	"github.com/monetarium/node/hdkeychain"
	"github.com/monetarium/node/txscript/stdaddr"
	"github.com/monetarium/node/wire"
)

func HD2Address(k *hdkeychain.ExtendedKey, params stdaddr.AddressParams) (*stdaddr.AddressPubKeyHashEcdsaSecp256k1V0, error) {
	pk := k.SerializedPubKey()
	hash := stdaddr.Hash160(pk)
	return stdaddr.NewAddressPubKeyHashEcdsaSecp256k1V0(hash, params)
}

// IsEitherCoinBaseTx verifies if a transaction is either a coinbase prior to
// the treasury agenda activation or a coinbse after treasury agenda
// activation.
func IsEitherCoinBaseTx(tx *wire.MsgTx) bool {
	if standalone.IsCoinBaseTx(tx, false) {
		return true
	}
	if standalone.IsCoinBaseTx(tx, true) {
		return true
	}
	return false
}
