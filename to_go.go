package godence

import (
	"fmt"
	"math/big"
	"reflect"

	"github.com/onflow/cadence"
)

// ToGo. Convert cadence types to go.
// Param 1: cadence value to convert.
// Param 2: go pointer.
// Address convert to string, will have 0x prefix.
func ToGo(value cadence.Value, dist any) (err error) {
	// type cast may be failed, should recover panic
	defer func() {
		if rec := recover(); rec != nil {
			err = fmt.Errorf("panic recovered: %v", rec)
		}
		// defer function has no return expression.
		// should use named return value.
	}()
	switch v := dist.(type) {
	// integers
	case **big.Int: // Cadence Int, Int128, Int256, UInt, UInt128, UInt256
		*v = value.ToGoValue().(*big.Int)
		return nil
	case *int8:
		*v = value.ToGoValue().(int8)
		return nil
	case *int16:
		*v = value.ToGoValue().(int16)
		return nil
	case *int32:
		*v = value.ToGoValue().(int32)
		return nil
	case *int64: // Cadence Int64, Fix64
		*v = value.ToGoValue().(int64)
		return nil
	// unsigned integers
	case *uint8:
		*v = value.ToGoValue().(uint8)
		return nil
	case *uint16:
		*v = value.ToGoValue().(uint16)
		return nil
	case *uint32:
		*v = value.ToGoValue().(uint32)
		return nil
	case *uint64: // Cadence UInt64, UFix64
		*v = value.ToGoValue().(uint64)
		return nil
	// other
	case *string: // Cadence String, Address
		switch cv := value.(type) {
		case cadence.Address:
			*v = cv.String()
			return
		}
		*v = value.ToGoValue().(string)
		return nil
	case *[8]uint8: // Address
		*v = value.ToGoValue().([8]uint8)
		return nil
	case *cadence.Address: // Address
		*v = value.(cadence.Address)
		return nil
	case *bool:
		*v = value.ToGoValue().(bool)
		return nil
	}
	// check if panic recovered
	if err != nil {
		return err
	}
	return fmt.Errorf("unsupport type: %s", reflect.TypeOf(dist))
}
