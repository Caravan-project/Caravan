package core

import (
	"blockEmulator/account"
	"bytes"
	"crypto/sha256"
	"fmt"
	"log"
	"math/big"

	"encoding/gob"
	"encoding/hex"
	"encoding/json"
)

type Transaction struct {
	Sender             []byte `json:"sender"`
	Recipient          []byte `json:"recipient"`
	TxHash             []byte
	Id                 int
	Success            bool
	IsRelay            bool
	SenLock            bool
	RecLock            bool
	Value              *big.Int `json:"value"`
	RequestTime        int64
	Second_RequestTime int64
	CommitTime         int64
	LockTime           int64
	UnlockTime         int64
	LockTime2          int64
	UnlockTime2        int64
	HalfLock           bool
	Rec_Suppose_on_chain   int
	Sen_Suppose_on_chain   int
	Relay_Lock         bool
}

////////////////////////////////////////
// MMM 聚合使用Transaction2
type Transaction2 struct {
	Sender             []byte `json:"sender"`
	Recipient          [][]byte `json:"recipient"`        // MMM 多个接收者
	TxHash             []byte
	Id                 int
	Success            bool
	IsRelay            bool
	SenLock            bool
	RecLock            bool
	Value              []*big.Int `json:"value"`          // MMM 多个数值
	RequestTime        int64
	Second_RequestTime int64
	CommitTime         int64
	LockTime           int64
	UnlockTime         int64
	LockTime2          int64
	UnlockTime2        int64
	HalfLock           bool
	Rec_Suppose_on_chain   int
	Sen_Suppose_on_chain   int
	Relay_Lock         bool
}
func (tx *Transaction2) PrintTx() {
	// vals := []interface{}{
	// 	hex.EncodeToString(tx.Sender),
	// 	account.Addr2Shard(hex.EncodeToString(tx.Sender)),
	// 	hex.EncodeToString(tx.Recipient),
	// 	account.Addr2Shard(hex.EncodeToString(tx.Recipient)),
	// 	tx.Value,
	// 	// hex.EncodeToString(tx.TxHash),
	// }
	c, _ := json.Marshal(tx)
	fmt.Printf("%v\n", string(c))
}

func (tx *Transaction2) Hash() []byte {
	hash := sha256.Sum256(tx.Encode())
	return hash[:]
}

// Encode transaction for storing
func (tx *Transaction2) Encode() []byte {
	var buff bytes.Buffer

	enc := gob.NewEncoder(&buff)
	err := enc.Encode(tx)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

func DecodeTx2(to_decode []byte) *Transaction2 {
	var tx Transaction2

	decoder := gob.NewDecoder(bytes.NewReader(to_decode))
	err := decoder.Decode(&tx)
	if err != nil {
		log.Panic(err)
	}

	return &tx
}

// 关键函数，将聚合交易拆分成单独交易
func (tx *Transaction2) Mtx2totx1() []*Transaction2{
	length := len(tx.Recipient)
	txs := make([]*Transaction2, 0, length)
	for i := 0; i < length; i ++{
		tmp_reci := make([][]byte, 1)
		tmp_value := make([]*big.Int, 1)
		tmp_reci[0] = tx.Recipient[i]
		tmp_value[0] = tx.Value[i]
		tmp := &Transaction2{
			Sender:    tx.Sender,
			Recipient: tmp_reci,
			Value:     tmp_value,
		}
		txs = append(txs, tmp)
	}
	return txs
}

//////////MMM


func (tx *Transaction) PrintTx() {
	vals := []interface{}{
		hex.EncodeToString(tx.Sender),
		account.Addr2Shard(hex.EncodeToString(tx.Sender)),
		hex.EncodeToString(tx.Recipient),
		account.Addr2Shard(hex.EncodeToString(tx.Recipient)),
		tx.Value,
		// hex.EncodeToString(tx.TxHash),
	}
	fmt.Printf("%v\n", vals)
}

func (tx *Transaction) Hash() []byte {
	hash := sha256.Sum256(tx.Encode())
	return hash[:]
}

// Encode transaction for storing
func (tx *Transaction) Encode() []byte {
	var buff bytes.Buffer

	enc := gob.NewEncoder(&buff)
	err := enc.Encode(tx)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

func DecodeTx(to_decode []byte) *Transaction {
	var tx Transaction

	decoder := gob.NewDecoder(bytes.NewReader(to_decode))
	err := decoder.Decode(&tx)
	if err != nil {
		log.Panic(err)
	}

	return &tx
}
