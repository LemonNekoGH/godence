package godence

import (
	"fmt"
	"math/big"
	"reflect"

	"github.com/onflow/cadence"
)

// getFieldByName. find cadence field by go field name.
func getFieldByName(name string, value cadence.Value) (cadence.Value, error) {
	switch v := value.(type) {
	case cadence.Struct:
		for index, field := range v.StructType.Fields {
			if name == field.Identifier {
				return v.Fields[index], nil
			}
		}
	case cadence.Event:
		for index, field := range v.EventType.Fields {
			if name == field.Identifier {
				return v.Fields[index], nil
			}
		}
	}
	// not found, return void and error
	return cadence.NewVoid(), fmt.Errorf("cannot find field named %s in cadence struct/event", name)
}

// structEventToGoStruct
func structEventToGoStruct(value cadence.Value, dist any) (err error) {
	distT := reflect.TypeOf(dist)
	distV := reflect.ValueOf(dist)

	defer func() {
		if msg := recover(); msg != nil {
			err = fmt.Errorf("structEventToGoStruct, panic recoverd: %v", msg)
		}
	}()

	// traverse all dist fields.
	for fieldIndex := 0; fieldIndex < distT.Elem().NumField(); fieldIndex++ {
		fieldT := distT.Elem().Field(fieldIndex)
		fieldV := distV.Elem().Field(fieldIndex)
		// cannot set, skip
		if !fieldV.CanSet() {
			continue
		}
		// find cadence field by go field
		fieldName := fieldT.Name
		if tagValue, ok := fieldT.Tag.Lookup("godence"); ok {
			// get cadence field name specified by tag
			fieldName = tagValue
		}
		// if error
		if v, err := getFieldByName(fieldName, value); err == nil {
			if fieldV.Kind() == reflect.Pointer {
				// if big int
				if fieldT.Type.Elem().PkgPath() == "math/big" && fieldT.Type.Elem().Name() == "Int" {
					fieldV.Set(reflect.ValueOf(v.ToGoValue()))
				} else {
					// embedded struct
					// should use type to get kind
					if fieldT.Type.Elem().Kind() == reflect.Struct {
						// check nil
						if fieldV.IsNil() {
							// new a value
							fieldV.Set(reflect.New(fieldT.Type.Elem()))
						}
						toGoStruct(v, fieldV.Interface())
					}
				}
			} else {
				fieldV.Set(reflect.ValueOf(v.ToGoValue()))
			}
		} else if err != nil {
			return err
		}
	}
	return
}

// toGoMap. call this function if type of dist is map kind.
func toGoMap(value cadence.Value, dist any) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("toGoMap, panic recoverd: %v", e)
		}
	}()
	dic := value.(cadence.Dictionary)
	distV := reflect.ValueOf(dist)
	for _, retEntry := range dic.Pairs {
		distV.SetMapIndex(reflect.ValueOf(retEntry.Key.ToGoValue()), reflect.ValueOf(retEntry.Value.ToGoValue()))
	}
	return
}

// toGoSlice. call this function if type of dist is array kind.
func toGoSlice(value cadence.Value, dist any) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("toGoSlice, panic recoverd: %v", e)
		}
	}()
	dic := value.(cadence.Array)
	distV := reflect.ValueOf(dist).Elem()
	for _, retElment := range dic.Values {
		distV.Set(reflect.Append(distV, reflect.ValueOf(retElment.ToGoValue())))
	}
	return
}

// toGoStruct. call this function if type of dist is struct kind.
func toGoStruct(value cadence.Value, dist any) error {
	switch v := value.(type) {
	case cadence.Struct:
		return structEventToGoStruct(v, dist)
	case cadence.Event:
		return structEventToGoStruct(v, dist)
	}
	return fmt.Errorf("to go struct: unsupport cadence type: %s", reflect.TypeOf(value))
}

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
	switch reflect.TypeOf(dist).Kind() {
	// try to convert to struct type
	case reflect.Pointer:
		switch reflect.TypeOf(dist).Elem().Kind() {
		case reflect.Struct:
			return toGoStruct(value, dist)
		case reflect.Slice:
			return toGoSlice(value, dist)
		}
	case reflect.Map:
		return toGoMap(value, dist)
	}
	return fmt.Errorf("unsupport type: %s", reflect.TypeOf(dist))
}
