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
	case **big.Int:
		*v = value.ToGoValue().(*big.Int)
		return nil
	case *int8:
		*v = value.ToGoValue().(int8)
		return nil
	case *string:
		*v = value.ToGoValue().(string)
		return nil
	}
	// check if panic recovered
	if err != nil {
		return err
	}
	return fmt.Errorf("unsupport type: %s", reflect.TypeOf(dist))
}
