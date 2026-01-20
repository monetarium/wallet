// Copyright (c) 2016 The btcsuite developers
// Copyright (c) 2016-2019 The Decred developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package txsizes

import "github.com/monetarium/monetarium-node/wire"

// Worst case script and input/output size estimates.
const (
	// RedeemP2PKSigScriptSize is the worst case (largest) serialize size
	// of a transaction input script that redeems a compressed P2PK output.
	// It is calculated as:
	//
	//   - OP_DATA_73
	//   - 72 bytes DER signature + 1 byte sighash
	RedeemP2PKSigScriptSize = 1 + 73

	// RedeemP2PKHSigScriptSize is the worst case (largest) serialize size
	// of a transaction input script that redeems a compressed P2PKH output.
	// It is calculated as:
	//
	//   - OP_DATA_73
	//   - 72 bytes DER signature + 1 byte sighash
	//   - OP_DATA_33
	//   - 33 bytes serialized compressed pubkey
	RedeemP2PKHSigScriptSize = 1 + 73 + 1 + 33

	// RedeemP2SHSigScriptSize is the worst case (largest) serialize size
	// of a transaction input script that redeems a P2SH output.
	// It is calculated as:
	//
	//  - OP_DATA_73
	//  - 73-byte signature
	//  - OP_DATA_35
	//  - OP_DATA_33
	//  - 33 bytes serialized compressed pubkey
	//  - OP_CHECKSIG
	RedeemP2SHSigScriptSize = 1 + 73 + 1 + 1 + 33 + 1

	// RedeemP2PKHInputSize is the worst case (largest) serialize size of a
	// transaction input redeeming a compressed P2PKH output.  It is
	// calculated as:
	//
	//   - 32 bytes previous tx
	//   - 4 bytes output index
	//   - 1 byte tree
	//   - 8 bytes amount
	//   - 4 bytes block height
	//   - 4 bytes block index
	//   - 1 byte compact int encoding value 107
	//   - 107 bytes signature script
	//   - 4 bytes sequence
	RedeemP2PKHInputSize = 32 + 4 + 1 + 8 + 4 + 4 + 1 + RedeemP2PKHSigScriptSize + 4

	// P2PKHPkScriptSize is the size of a transaction output script that
	// pays to a compressed pubkey hash.  It is calculated as:
	//
	//   - OP_DUP
	//   - OP_HASH160
	//   - OP_DATA_20
	//   - 20 bytes pubkey hash
	//   - OP_EQUALVERIFY
	//   - OP_CHECKSIG
	P2PKHPkScriptSize = 1 + 1 + 1 + 20 + 1 + 1

	// P2PKHPkTreasruryScriptSize is the size of a transaction output
	// script that pays stake change to a compressed pubkey hash.  This is
	// used when a user sends coins to the treasury via OP_TADD.  It is
	// calculated as:
	//
	//   - OP_SSTXCHANGE
	//   - OP_DUP
	//   - OP_HASH160
	//   - OP_DATA_20
	//   - 20 bytes pubkey hash
	//   - OP_EQUALVERIFY
	//   - OP_CHECKSIG
	P2PKHPkTreasruryScriptSize = 1 + 1 + 1 + 1 + 20 + 1 + 1

	// P2SHPkScriptSize is the size of a transaction output script that
	// pays to a script hash.  It is calculated as:
	//
	//   - OP_HASH160
	//   - OP_DATA_20
	//   - 20 bytes script hash
	//   - OP_EQUAL
	P2SHPkScriptSize = 1 + 1 + 20 + 1

	// TicketCommitmentScriptSize is the size of a ticket purchase commitment
	// script. It is calculated as:
	//
	//   - OP_RETURN
	//   - OP_DATA_30
	//   - 20 bytes P2SH/P2PKH
	//   - 8 byte amount
	//   - 2 byte fee range limits
	TicketCommitmentScriptSize = 1 + 1 + 20 + 8 + 2

	// P2PKHOutputSize is the serialize size of a transaction output with a
	// P2PKH output script.  It is calculated as:
	//
	//   - 8 bytes output value
	//   - 1 byte coin type (dual-coin support)
	//   - 2 bytes version
	//   - 1 byte compact int encoding value 25
	//   - 25 bytes P2PKH output script
	P2PKHOutputSize = 8 + 1 + 2 + 1 + 25

	// TSPENDInputSize
	//
	//   - OP_DATA_73
	//   - 73 bytes signature
	//   - OP_DATA_33
	//   - 33 bytes serialized compressed pubkey
	//   - 1 byte OP_TSPEND
	TSPENDInputSize = 1 + 73 + 1 + 33 + 1
)

func sumOutputSerializeSizes(outputs []*wire.TxOut) (serializeSize int) {
	for _, txOut := range outputs {
		serializeSize += txOut.SerializeSize()
	}
	return serializeSize
}

// EstimateSerializeSize returns a worst case serialize size estimate for a
// signed transaction that spends a number of outputs and contains each
// transaction output from txOuts. The estimated size is incremented for an
// additional change output if changeScriptSize is greater than 0. Passing 0
// does not add a change output.
func EstimateSerializeSize(scriptSizes []int, txOuts []*wire.TxOut, changeScriptSize int) int {
	return estimateSerializeSizeInternal(scriptSizes, txOuts, changeScriptSize, false)
}

// EstimateSerializeSizeSKA returns a worst case serialize size estimate for a
// signed SKA transaction. SKA outputs have a slightly different wire format
// (1-byte length prefix for the value), so change output estimation differs.
func EstimateSerializeSizeSKA(scriptSizes []int, txOuts []*wire.TxOut, changeScriptSize int) int {
	return estimateSerializeSizeInternal(scriptSizes, txOuts, changeScriptSize, true)
}

func estimateSerializeSizeInternal(scriptSizes []int, txOuts []*wire.TxOut, changeScriptSize int, isSKA bool) int {
	inputCount := len(scriptSizes)
	outputCount := len(txOuts)
	changeSize := 0
	if changeScriptSize != 0 {
		if isSKA {
			changeSize = EstimateOutputSizeSKA(changeScriptSize)
		} else {
			changeSize = EstimateOutputSize(changeScriptSize)
		}
		outputCount++
	}

	// Calculate size for TxSerializeFull format (prefix + witness)
	// This matches the format used in wire.MsgTx.SerializeSize() for TxSerializeFull

	// Base: Version 4 bytes + LockTime 4 bytes + Expiry 4 bytes = 12 bytes
	// Plus varint sizes for input count (x2) and output count
	baseSize := 12 + wire.VarIntSerializeSize(uint64(inputCount)) +
		wire.VarIntSerializeSize(uint64(inputCount)) +
		wire.VarIntSerializeSize(uint64(outputCount))

	// Calculate prefix input sizes (without witness data)
	prefixInputsSize := 0
	for range scriptSizes {
		prefixInputsSize += EstimateInputPrefixSize()
	}

	// Calculate witness input sizes (signature scripts)
	// V13 format: [ValueIn:8][SKAValueInLen:1][SKAValueIn:N][BlockHeight:4][BlockIndex:4][SigScript:var]
	// For SKA inputs, SKAValueIn can be up to 16 bytes (worst case)
	witnessInputsSize := 0
	for _, scriptSize := range scriptSizes {
		if isSKA {
			witnessInputsSize += EstimateInputWitnessSizeSKA(scriptSize)
		} else {
			witnessInputsSize += EstimateInputWitnessSize(scriptSize)
		}
	}

	// Calculate output sizes (includes CoinType field for dual-coin)
	outputsSize := sumOutputSerializeSizes(txOuts) + changeSize

	return baseSize + prefixInputsSize + witnessInputsSize + outputsSize
}

// EstimateSerializeSizeFromScriptSizes returns a worst case serialize size
// estimate for a signed transaction that spends len(inputSizes) previous
// outputs and pays to len(outputSizes) outputs with scripts of the provided
// worst-case sizes. The estimated size is incremented for an additional
// change output if changeScriptSize is greater than 0. Passing 0 does not
// add a change output.
func EstimateSerializeSizeFromScriptSizes(inputSizes []int, outputSizes []int, changeScriptSize int) int {
	// Generate and sum up the estimated sizes of the inputs.
	txInsSize := 0
	for _, inputSize := range inputSizes {
		txInsSize += EstimateInputSize(inputSize)
	}

	// Generate and sum up the estimated sizes of the outputs.
	txOutsSize := 0
	for _, outputSize := range outputSizes {
		txOutsSize += EstimateOutputSize(outputSize)
	}

	inputCount := len(inputSizes)
	outputCount := len(outputSizes)
	changeSize := 0
	if changeScriptSize > 0 {
		changeSize = EstimateOutputSize(changeScriptSize)
		outputCount++
	}

	// 12 additional bytes are for version, locktime and expiry.
	return 12 + (2 * wire.VarIntSerializeSize(uint64(inputCount))) +
		wire.VarIntSerializeSize(uint64(outputCount)) +
		txInsSize + txOutsSize + changeSize
}

// EstimateInputSize returns the worst case serialize size estimate for a tx input
//   - 32 bytes previous tx
//   - 4 bytes output index
//   - 1 byte tree
//   - 8 bytes amount
//   - 4 bytes block height
//   - 4 bytes block index
//   - the compact int representation of the script size
//   - the supplied script size
//   - 4 bytes sequence
func EstimateInputSize(scriptSize int) int {
	return 32 + 4 + 1 + 8 + 4 + 4 + wire.VarIntSerializeSize(uint64(scriptSize)) + scriptSize + 4
}

// EstimateOutputSize returns the worst case serialize size estimate for a tx output
//   - 8 bytes amount
//   - 1 byte coin type (dual-coin support)
//   - 2 bytes version
//   - the compact int representation of the script size
//   - the supplied script size
func EstimateOutputSize(scriptSize int) int {
	return 8 + 1 + 2 + wire.VarIntSerializeSize(uint64(scriptSize)) + scriptSize
}

// EstimateOutputSizeSKA returns the serialize size estimate for an SKA tx output.
// SKA outputs have a different format from VAR:
//   - 1 byte coin type
//   - 1 byte value length prefix
//   - N bytes value (up to 16 bytes for large amounts)
//   - 2 bytes version
//   - the compact int representation of the script size
//   - the supplied script size
//
// We use worst-case 16 bytes for value to ensure fee is never underestimated.
// SKA amounts can be very large (900T * 1e18 = ~14 bytes), so we round up.
func EstimateOutputSizeSKA(scriptSize int) int {
	// SKA format: CoinType(1) + ValLen(1) + Value(16 max) + Version(2) + VarInt + PkScript
	// Always overestimate to avoid fee rejection
	return 1 + 1 + 16 + 2 + wire.VarIntSerializeSize(uint64(scriptSize)) + scriptSize
}

// EstimateInputPrefixSize returns the serialize size estimate for a tx input prefix
//   - 32 bytes previous tx
//   - 4 bytes output index
//   - 1 byte tree
//   - 4 bytes sequence
func EstimateInputPrefixSize() int {
	return 32 + 4 + 1 + 4
}

// EstimateInputWitnessSize returns the serialize size estimate for a tx input witness
// V13 format: [ValueIn:8][SKAValueInLen:1][SKAValueIn:N][BlockHeight:4][BlockIndex:4][SigScript:var]
//   - 8 bytes amount (ValueIn for fraud proofs)
//   - 1 byte SKAValueInLen (V13: always present, 0 for VAR inputs)
//   - 4 bytes block height
//   - 4 bytes block index
//   - the compact int representation of the script size
//   - the supplied script size
func EstimateInputWitnessSize(scriptSize int) int {
	// V13 format includes SKAValueInLen (1 byte), which is 0 for VAR inputs
	return 8 + 1 + 4 + 4 + wire.VarIntSerializeSize(uint64(scriptSize)) + scriptSize
}

// EstimateInputWitnessSizeSKA returns the serialize size estimate for an SKA tx input witness.
// SKA inputs include SKAValueIn which can be up to 16 bytes for large amounts.
// We use worst-case 16 bytes to ensure fee is never underestimated.
func EstimateInputWitnessSizeSKA(scriptSize int) int {
	// V13 SKA format: ValueIn(8) + SKAValueInLen(1) + SKAValueIn(16 max) + BlockHeight(4) + BlockIndex(4) + VarInt + SigScript
	return 8 + 1 + 16 + 4 + 4 + wire.VarIntSerializeSize(uint64(scriptSize)) + scriptSize
}
