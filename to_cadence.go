package godence

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"reflect"
	"strings"

	"github.com/onflow/cadence"
)

// helper for ufix64
type UFix64 uint64

// helper for fix64
type Fix64 int64

// helper for Address
type Address string

// helper for Address
type Path string

// helper for Character
type Character string

// bigIntToCadence
func bigIntToCadence(i *big.Int) (cadence.Value, error) {
	// should from small to big
	i128, err := cadence.NewInt128FromBig(i)
	if err == nil {
		return i128, nil
	}
	u128, err := cadence.NewUInt128FromBig(i)
	if err == nil {
		return u128, nil
	}
	i256, err := cadence.NewInt256FromBig(i)
	if err == nil {
		return i256, nil
	}
	u256, err := cadence.NewUInt256FromBig(i)
	if err == nil {
		return u256, nil
	}
	return nil, fmt.Errorf("unsupport big.Int value: %s", i.Text(10))
}

// arrayOrSliceToCadence
func arrayOrSliceToCadence(value any) (cadence.Value, error) {
	ret := []cadence.Value{}
	v := reflect.ValueOf(value)
	for i := 0; i < v.Len(); i++ {
		// convert all elements of slice/array
		cv, err := ToCadence(v.Index(i).Interface())
		if err != nil {
			return nil, err
		}
		ret = append(ret, cv)
	}
	return cadence.NewArray(ret), nil
}

// mapToCadence
func mapToCadence(value any) (cadence.Value, error) {
	ret := []cadence.KeyValuePair{}
	v := reflect.ValueOf(value)
	// convert all entry to KeyValuePair
	for _, key := range v.MapKeys() {
		ck, err := ToCadence(key.Interface())
		if err != nil {
			return nil, err
		}
		cv, err := ToCadence(v.MapIndex(key).Interface())
		if err != nil {
			return nil, err
		}
		ret = append(ret, cadence.KeyValuePair{
			Key:   ck,
			Value: cv,
		})
	}
	return cadence.NewDictionary(ret), nil
}

// ToCadence Convert any go value to cadence value.
// Type uint64 will convert to UInt64, if you want to convert to UFix64,
// you should wrap the field in a struct and tagged it with `godence:"type:UFix64"`, or use our UFix64 type.
func ToCadence(value any) (cadence.Value, error) {
	switch v := value.(type) {
	// integer
	case int:
		return cadence.NewInt(v), nil
	case int8:
		return cadence.NewInt8(v), nil
	case int16:
		return cadence.NewInt16(v), nil
	case int32:
		return cadence.NewInt32(v), nil
	case int64:
		return cadence.NewInt64(v), nil
	case uint:
		return cadence.NewUInt(v), nil
	case uint8:
		return cadence.NewUInt8(v), nil
	case uint16:
		return cadence.NewUInt16(v), nil
	case uint32:
		return cadence.NewUInt32(v), nil
	case uint64:
		return cadence.NewUInt64(v), nil
	case Fix64:
		return cadence.NewFix64(fmt.Sprintf("%d.0", int64(v)))
	case UFix64:
		return cadence.NewUFix64(fmt.Sprintf("%d.0", uint64(v)))
	case *big.Int:
		return bigIntToCadence(v)
	// TODO: float, should float convert to Fix64?
	// case float32:
	// case float64:
	case string:
		return cadence.NewString(v)
	case Address:
		decoded, err := hex.DecodeString(strings.TrimPrefix(string(v), "0x"))
		return cadence.BytesToAddress(decoded), err
	case Path:
		part := strings.Split(string(v), "/")
		return cadence.NewPath(part[1], part[2]), nil
	case Character:
		return cadence.NewCharacter(string(v))
	case bool:
		return cadence.NewBool(v), nil
	}
	switch reflect.TypeOf(value).Kind() {
	// array or slice
	case reflect.Slice, reflect.Array:
		return arrayOrSliceToCadence(value)
	// map
	case reflect.Map:
		return mapToCadence(value)
	}
	return nil, fmt.Errorf("unsupport type: %s", reflect.TypeOf(value))
}

// ToCadence Convert any go value to cadence optional value.
func ToCadenceOptional(value any) (cadence.Value, error) {
	v, err := ToCadence(value)
	if err != nil {
		return nil, err
	}
	return cadence.NewOptional(v), nil
}
