// Copyright (c) 2024 The Monetarium developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package udb

import (
	"testing"

	"github.com/decred/dcrd/chaincfg/chainhash"
	"github.com/decred/dcrd/cointype"
	"github.com/decred/dcrd/txscript/v4"
	"github.com/decred/dcrd/wire"
)

// createMockSSFeeTx creates a mock SSFee transaction for testing.
func createMockSSFeeTx(ct cointype.CoinType, numOutputs int, amount int64, markerType string) *wire.MsgTx {
	tx := wire.NewMsgTx()
	tx.Version = 3

	// Add null input (like coinbase)
	tx.AddTxIn(&wire.TxIn{
		PreviousOutPoint: wire.OutPoint{
			Hash:  chainhash.Hash{},
			Index: wire.MaxPrevOutIndex,
		},
		SignatureScript: []byte{}, // Empty signature for SSFee
	})

	// Add reward outputs
	for i := 0; i < numOutputs; i++ {
		tx.AddTxOut(&wire.TxOut{
			Value:    amount,
			CoinType: ct,
			PkScript: make([]byte, 25), // Dummy P2PKH script
		})
	}

	// Add OP_RETURN with marker
	var marker []byte
	if markerType == "MF" {
		marker = []byte{0x4D, 0x46} // "MF"
	} else {
		marker = []byte{0x53, 0x46} // "SF"
	}

	opReturnScript := []byte{
		txscript.OP_RETURN, // 0x6a
		0x06,               // OP_DATA_6
	}
	opReturnScript = append(opReturnScript, marker...)
	opReturnScript = append(opReturnScript, []byte{0x00, 0x00, 0x00, 0x00}...) // height

	tx.AddTxOut(&wire.TxOut{
		Value:    0,
		PkScript: opReturnScript,
		CoinType: ct,
	})

	return tx
}

func TestIsSSFeeTx(t *testing.T) {
	tests := []struct {
		name     string
		tx       *wire.MsgTx
		expected bool
	}{
		{
			name:     "SSFee MF transaction",
			tx:       createMockSSFeeTx(cointype.CoinType(1), 3, 1000, "MF"),
			expected: true,
		},
		{
			name:     "SSFee SF transaction",
			tx:       createMockSSFeeTx(cointype.CoinType(1), 3, 1000, "SF"),
			expected: true,
		},
		{
			name: "Regular transaction",
			tx: &wire.MsgTx{
				Version: 1,
				TxIn: []*wire.TxIn{{
					PreviousOutPoint: wire.OutPoint{
						Hash:  chainhash.Hash{1, 2, 3},
						Index: 0,
					},
				}},
				TxOut: []*wire.TxOut{{
					Value:    1000,
					CoinType: cointype.CoinTypeVAR,
				}},
			},
			expected: false,
		},
		{
			name: "Coinbase (not SSFee)",
			tx: &wire.MsgTx{
				Version: 1,
				TxIn: []*wire.TxIn{{
					PreviousOutPoint: wire.OutPoint{
						Hash:  chainhash.Hash{},
						Index: wire.MaxPrevOutIndex,
					},
					SignatureScript: []byte{0x00, 0x01}, // Coinbase signature
				}},
				TxOut: []*wire.TxOut{{
					Value:    5000000000,
					CoinType: cointype.CoinTypeVAR,
				}},
			},
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := isSSFeeTx(test.tx)
			if result != test.expected {
				t.Errorf("%s: isSSFeeTx() = %v, want %v", test.name, result, test.expected)
			}
		})
	}
}

func TestGetSSFeeType(t *testing.T) {
	tests := []struct {
		name     string
		tx       *wire.MsgTx
		expected string
	}{
		{
			name:     "Miner Fee (MF)",
			tx:       createMockSSFeeTx(cointype.CoinType(1), 3, 1000, "MF"),
			expected: "MF",
		},
		{
			name:     "Staker Fee (SF)",
			tx:       createMockSSFeeTx(cointype.CoinType(1), 3, 1000, "SF"),
			expected: "SF",
		},
		{
			name: "Not SSFee",
			tx: &wire.MsgTx{
				Version: 1,
				TxIn: []*wire.TxIn{{
					PreviousOutPoint: wire.OutPoint{
						Hash:  chainhash.Hash{1},
						Index: 0,
					},
				}},
				TxOut: []*wire.TxOut{{Value: 1000}},
			},
			expected: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := getSSFeeType(test.tx)
			if result != test.expected {
				t.Errorf("%s: getSSFeeType() = %q, want %q", test.name, result, test.expected)
			}
		})
	}
}

func TestIsSSFeeMinerTx(t *testing.T) {
	tests := []struct {
		name     string
		tx       *wire.MsgTx
		expected bool
	}{
		{
			name:     "Miner Fee transaction",
			tx:       createMockSSFeeTx(cointype.CoinType(1), 3, 1000, "MF"),
			expected: true,
		},
		{
			name:     "Staker Fee transaction",
			tx:       createMockSSFeeTx(cointype.CoinType(1), 3, 1000, "SF"),
			expected: false,
		},
		{
			name: "Regular transaction",
			tx: &wire.MsgTx{
				Version: 1,
				TxIn:    []*wire.TxIn{{PreviousOutPoint: wire.OutPoint{Hash: chainhash.Hash{1}, Index: 0}}},
				TxOut:   []*wire.TxOut{{Value: 1000}},
			},
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := isSSFeeMinerTx(test.tx)
			if result != test.expected {
				t.Errorf("%s: isSSFeeMinerTx() = %v, want %v", test.name, result, test.expected)
			}
		})
	}
}
