package common

import (
	"encoding/json"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"strconv"
	"strings"
)

type UID struct {
	sequence   uint32 // 32 bits
	objectType int    // 10 bits
	shardID    uint32 // 18 bits
}

func generateUID(sequence uint32, objectType int, shardID uint32) uint64 {
	sequenceBitsShifted := uint64(sequence) << 28
	objectTypeBitsShifted := uint64(objectType) << 18
	shardIDBitsShifted := uint64(shardID) << 0

	return sequenceBitsShifted | objectTypeBitsShifted | shardIDBitsShifted
}

func (u UID) String() string {
	val := generateUID(u.sequence, u.objectType, u.shardID)
	return base58.Encode([]byte(fmt.Sprintf("%v", val)))
}

func (u UID) GetShardID() uint32 {
	return u.shardID
}

func (u UID) GetSequence() uint32 {
	return u.sequence
}

func (u UID) GetObjectType() int {
	return u.objectType
}

func UIDFromString(s string) (UID, error) {
	val, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return UID{}, err
	}
	// 0x3FF or 0x3FFFF means fully 1111111...
	sequence := uint32(val >> 28)
	objectType := int((val >> 18) & 0x3FF)
	shardID := uint32((val >> 0) & 0x3FFFF)

	return UID{
		sequence:   sequence,
		objectType: objectType,
		shardID:    shardID,
	}, nil
}

func UIDFromBase58(uidString string) (UID, error) {
	decoded := base58.Decode(uidString)
	//if len(decoded) != 10 {
	//	return UID{}, fmt.Errorf("invalid UID string: %s", uidString)
	//}
	return UIDFromString(string(decoded))
}

func NewUID(sequence uint32, objectType int, shardID uint32) UID {
	return UID{
		sequence:   sequence,
		objectType: objectType,
		shardID:    shardID,
	}
}

func (u UID) MarshalJSON() ([]byte, error) {
	return json.Marshal(fmt.Sprintf("%s", u.String()))
}

func (u *UID) UnmarshalJSON(data []byte) error {
	str := strings.ReplaceAll(string(data), "\"", "")
	decoded, err := UIDFromBase58(str)
	if err != nil {
		return err
	}
	u.sequence = decoded.sequence
	u.objectType = decoded.objectType
	u.shardID = decoded.shardID
	return nil
}
