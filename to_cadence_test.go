package godence

import (
	"context"
	"fmt"
	"math/big"
	"testing"

	"github.com/onflow/cadence"
	"github.com/stretchr/testify/assert"
)

func ExampleToCadence_baseType() {
	cadenceValue, err := ToCadence(int(15))
	fmt.Printf("type id: %s, err: %v", cadenceValue.Type().ID(), err)
	//Output: type id: Int, err: <nil>
}

func ExampleToCadence_bigInt() {
	cadenceValue, err := ToCadence(big.NewInt(10))
	fmt.Printf("type id: %s, err: %v", cadenceValue.Type().ID(), err)
	//Output: type id: Int128, err: <nil>
}

func ExampleToCadence_uFix64() {
	cadenceValue, err := ToCadence(UFix64(15))
	fmt.Printf("type id: %s, err: %v", cadenceValue.Type().ID(), err)
	//Output: type id: UFix64, err: <nil>
}

func TestToCadence(t *testing.T) {
	t.Run("unsupport type", func(t *testing.T) {
		assert := assert.New(t)
		cadenceValue, err := ToCadence(unsupportType(""))
		assert.Nil(cadenceValue)
		assert.EqualError(err, "unsupport type: godence.unsupportType")
	})
	t.Run("to int", func(t *testing.T) {
		assert := assert.New(t)
		script := []byte(`pub fun main(arg: Int) { log("the arg is ".concat(arg.toString())) }`)

		cadenceValue, err := ToCadence(int(15))
		assert.NoError(err)
		assert.Equal(cadenceValue.Type().ID(), "Int")

		args := []cadence.Value{cadenceValue}
		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, args)

		assert.NoError(err)
		assert.Equal(ret.Type().ID(), "Void")
	})
	t.Run("to int8", func(t *testing.T) {
		assert := assert.New(t)
		script := []byte(`pub fun main(arg: Int8) { log("the arg is ".concat(arg.toString())) }`)

		cadenceValue, err := ToCadence(int8(15))
		assert.NoError(err)
		assert.Equal(cadenceValue.Type().ID(), "Int8")

		args := []cadence.Value{cadenceValue}
		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, args)

		assert.NoError(err)
		assert.Equal(ret.Type().ID(), "Void")
	})
	t.Run("to int16", func(t *testing.T) {
		assert := assert.New(t)
		script := []byte(`pub fun main(arg: Int16) { log("the arg is ".concat(arg.toString())) }`)

		cadenceValue, err := ToCadence(int16(15))
		assert.NoError(err)
		assert.Equal(cadenceValue.Type().ID(), "Int16")

		args := []cadence.Value{cadenceValue}
		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, args)

		assert.NoError(err)
		assert.Equal(ret.Type().ID(), "Void")
	})
	t.Run("to int32", func(t *testing.T) {
		assert := assert.New(t)
		script := []byte(`pub fun main(arg: Int32) { log("the arg is ".concat(arg.toString())) }`)

		cadenceValue, err := ToCadence(int32(15))
		assert.NoError(err)
		assert.Equal(cadenceValue.Type().ID(), "Int32")

		args := []cadence.Value{cadenceValue}
		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, args)

		assert.NoError(err)
		assert.Equal(ret.Type().ID(), "Void")
	})
	t.Run("to int64", func(t *testing.T) {
		assert := assert.New(t)
		script := []byte(`pub fun main(arg: Int64) { log("the arg is ".concat(arg.toString())) }`)

		cadenceValue, err := ToCadence(int64(15))
		assert.NoError(err)
		assert.Equal(cadenceValue.Type().ID(), "Int64")

		args := []cadence.Value{cadenceValue}
		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, args)

		assert.NoError(err)
		assert.Equal(ret.Type().ID(), "Void")
	})
	t.Run("to int128", func(t *testing.T) {
		assert := assert.New(t)
		script := []byte(`pub fun main(arg: Int128) { log("the arg is ".concat(arg.toString())) }`)
		// make a new big int
		bigInt := &big.Int{}
		bigInt, ok := bigInt.SetString("170141183460469231731687303715884105727", 10)
		assert.True(ok)

		cadenceValue, err := ToCadence(bigInt)
		assert.NoError(err)
		assert.Equal(cadenceValue.Type().ID(), "Int128")

		args := []cadence.Value{cadenceValue}
		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, args)
		assert.NoError(err)
		assert.Equal(ret.Type().ID(), "Void")
	})
	t.Run("to int256", func(t *testing.T) {
		assert := assert.New(t)
		script := []byte(`pub fun main(arg: Int256) { log("the arg is ".concat(arg.toString())) }`)
		// make a new big int
		bigInt := &big.Int{}
		bigInt, ok := bigInt.SetString("57896044618658097711785492504343953926634992332820282019728792003956564819967", 10)
		assert.True(ok)

		cadenceValue, err := ToCadence(bigInt)
		assert.NoError(err)
		assert.Equal(cadenceValue.Type().ID(), "Int256")

		args := []cadence.Value{cadenceValue}
		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, args)
		assert.NoError(err)
		assert.Equal(ret.Type().ID(), "Void")
	})
	t.Run("to uint", func(t *testing.T) {
		assert := assert.New(t)
		script := []byte(`pub fun main(arg: UInt) { log("the arg is ".concat(arg.toString())) }`)

		cadenceValue, err := ToCadence(uint(15))
		assert.NoError(err)
		assert.Equal(cadenceValue.Type().ID(), "UInt")

		args := []cadence.Value{cadenceValue}
		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, args)

		assert.NoError(err)
		assert.Equal(ret.Type().ID(), "Void")
	})
	t.Run("to uint8", func(t *testing.T) {
		assert := assert.New(t)
		script := []byte(`pub fun main(arg: UInt8) { log("the arg is ".concat(arg.toString())) }`)

		cadenceValue, err := ToCadence(uint8(15))
		assert.NoError(err)
		assert.Equal(cadenceValue.Type().ID(), "UInt8")

		args := []cadence.Value{cadenceValue}
		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, args)

		assert.NoError(err)
		assert.Equal(ret.Type().ID(), "Void")
	})
	t.Run("to uint16", func(t *testing.T) {
		assert := assert.New(t)
		script := []byte(`pub fun main(arg: UInt16) { log("the arg is ".concat(arg.toString())) }`)

		cadenceValue, err := ToCadence(uint16(15))
		assert.NoError(err)
		assert.Equal(cadenceValue.Type().ID(), "UInt16")

		args := []cadence.Value{cadenceValue}
		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, args)

		assert.NoError(err)
		assert.Equal(ret.Type().ID(), "Void")
	})
	t.Run("to uint32", func(t *testing.T) {
		assert := assert.New(t)
		script := []byte(`pub fun main(arg: UInt32) { log("the arg is ".concat(arg.toString())) }`)

		cadenceValue, err := ToCadence(uint32(15))
		assert.NoError(err)
		assert.Equal(cadenceValue.Type().ID(), "UInt32")

		args := []cadence.Value{cadenceValue}
		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, args)

		assert.NoError(err)
		assert.Equal(ret.Type().ID(), "Void")
	})
	t.Run("to uint64", func(t *testing.T) {
		assert := assert.New(t)
		script := []byte(`pub fun main(arg: UInt64) { log("the arg is ".concat(arg.toString())) }`)

		cadenceValue, err := ToCadence(uint64(15))
		assert.NoError(err)
		assert.Equal(cadenceValue.Type().ID(), "UInt64")

		args := []cadence.Value{cadenceValue}
		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, args)

		assert.NoError(err)
		assert.Equal(ret.Type().ID(), "Void")
	})
	t.Run("to uint128", func(t *testing.T) {
		assert := assert.New(t)
		script := []byte(`pub fun main(arg: UInt128) { log("the arg is ".concat(arg.toString())) }`)
		// make a new big int
		bigInt := &big.Int{}
		bigInt, ok := bigInt.SetString("340282366920938463463374607431768211455", 10)
		assert.True(ok)

		cadenceValue, err := ToCadence(bigInt)
		assert.NoError(err)
		assert.Equal(cadenceValue.Type().ID(), "UInt128")

		args := []cadence.Value{cadenceValue}
		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, args)
		assert.NoError(err)
		assert.Equal(ret.Type().ID(), "Void")
	})
	t.Run("to uint256", func(t *testing.T) {
		assert := assert.New(t)
		script := []byte(`pub fun main(arg: UInt256) { log("the arg is ".concat(arg.toString())) }`)
		// make a new big int
		bigInt := &big.Int{}
		bigInt, ok := bigInt.SetString("115792089237316195423570985008687907853269984665640564039457584007913129639935", 10)
		assert.True(ok)

		cadenceValue, err := ToCadence(bigInt)
		assert.NoError(err)
		assert.Equal(cadenceValue.Type().ID(), "UInt256")

		args := []cadence.Value{cadenceValue}
		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, args)
		assert.NoError(err)
		assert.Equal(ret.Type().ID(), "Void")
	})
	t.Run("larger than u256 max value, should error", func(t *testing.T) {
		assert := assert.New(t)
		// make a new big int
		bigInt := &big.Int{}
		bigInt, ok := bigInt.SetString("115792089237316195423570985008687907853269984665640564039457584007913129639936", 10)
		assert.True(ok)

		_, err := ToCadence(bigInt)
		assert.Error(err)
	})
	t.Run("to fix64", func(t *testing.T) {
		assert := assert.New(t)
		script := []byte(`pub fun main(arg: Fix64) { log("the arg is ".concat(arg.toString())) }`)

		cadenceValue, err := ToCadence(Fix64(15))
		assert.NoError(err)
		assert.Equal(cadenceValue.Type().ID(), "Fix64")

		args := []cadence.Value{cadenceValue}
		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, args)
		assert.NoError(err)
		assert.Equal(ret.Type().ID(), "Void")
	})
	t.Run("to ufix64", func(t *testing.T) {
		assert := assert.New(t)
		script := []byte(`pub fun main(arg: UFix64) { log("the arg is ".concat(arg.toString())) }`)

		cadenceValue, err := ToCadence(UFix64(15))
		assert.NoError(err)
		assert.Equal(cadenceValue.Type().ID(), "UFix64")

		args := []cadence.Value{cadenceValue}
		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, args)
		assert.NoError(err)
		assert.Equal(ret.Type().ID(), "Void")
	})
	t.Run("to String", func(t *testing.T) {
		assert := assert.New(t)
		script := []byte(`pub fun main(arg: String) { log("the arg is ".concat(arg)) }`)

		cadenceValue, err := ToCadence("hello")
		assert.NoError(err)
		assert.Equal(cadenceValue.Type().ID(), "String")

		args := []cadence.Value{cadenceValue}
		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, args)
		assert.NoError(err)
		assert.Equal(ret.Type().ID(), "Void")
	})
	t.Run("to Address", func(t *testing.T) {
		assert := assert.New(t)
		script := []byte(`pub fun main(arg: Address) { log("the arg is ".concat(arg.toString())) }`)

		cadenceValue, err := ToCadence(Address("0x01234567"))
		assert.NoError(err)
		assert.Equal(cadenceValue.Type().ID(), "Address")

		args := []cadence.Value{cadenceValue}
		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, args)
		assert.NoError(err)
		assert.Equal(ret.Type().ID(), "Void")
	})
	t.Run("to Path", func(t *testing.T) {
		assert := assert.New(t)
		script := []byte(`pub fun main(arg: Path) { log("the arg is ".concat(arg.toString())) }`)

		cadenceValue, err := ToCadence(Path("/public/myCollection"))
		assert.NoError(err)
		assert.Equal(cadenceValue.Type().ID(), "Path")

		args := []cadence.Value{cadenceValue}
		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, args)
		assert.NoError(err)
		assert.Equal(ret.Type().ID(), "Void")
	})
	t.Run("to Character", func(t *testing.T) {
		assert := assert.New(t)
		script := []byte(`pub fun main(arg: Character) { log("the arg is ".concat(arg.toString())) }`)

		cadenceValue, err := ToCadence(Character("a"))
		assert.NoError(err)
		assert.Equal(cadenceValue.Type().ID(), "Character")

		args := []cadence.Value{cadenceValue}
		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, args)
		assert.NoError(err)
		assert.Equal(ret.Type().ID(), "Void")
	})

	t.Run("slice to Array", func(t *testing.T) {
		assert := assert.New(t)
		script := []byte(`pub fun main(arg: [String]): Bool { return arg[0] == "My" && arg[1] == "Name" && arg[2] == "Is" && arg[3] == "LemonNeko" }`)

		cadenceValue, err := ToCadence([]string{"My", "Name", "Is", "LemonNeko"})
		assert.NoError(err)

		args := []cadence.Value{cadenceValue}
		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, args)
		assert.NoError(err)
		assert.Equal(ret.Type().ID(), "Bool")
		assert.True(ret.ToGoValue().(bool))
	})

	t.Run("array to Array", func(t *testing.T) {
		assert := assert.New(t)
		script := []byte(`pub fun main(arg: [String]): Bool { return arg[0] == "My" && arg[1] == "Name" && arg[2] == "Is" && arg[3] == "LemonNeko" }`)

		cadenceValue, err := ToCadence([...]string{"My", "Name", "Is", "LemonNeko"})
		assert.NoError(err)

		args := []cadence.Value{cadenceValue}
		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, args)
		assert.NoError(err)
		assert.Equal(ret.Type().ID(), "Bool")
		assert.True(ret.ToGoValue().(bool))
	})

	t.Run("unsupport to Array", func(t *testing.T) {
		assert := assert.New(t)

		cadenceValue, err := ToCadence([...]unsupportType{"My", "Name", "Is", "LemonNeko"})
		assert.EqualError(err, "unsupport type: godence.unsupportType")
		assert.Nil(cadenceValue)
	})

	t.Run("map to Dictionary", func(t *testing.T) {
		assert := assert.New(t)
		script := []byte(`pub fun main(arg: {String: String}): Bool { return arg["MyName"] == "LemonNeko" }`)

		cadenceValue, err := ToCadence(map[string]string{"MyName": "LemonNeko"})
		assert.NoError(err)

		args := []cadence.Value{cadenceValue}
		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, args)
		assert.NoError(err)
		assert.Equal(ret.Type().ID(), "Bool")
		assert.True(ret.ToGoValue().(bool))
	})

	t.Run("map to Dictionary, unsupport key type", func(t *testing.T) {
		assert := assert.New(t)

		cadenceValue, err := ToCadence(map[unsupportType]string{"MyName": "LemonNeko"})
		assert.EqualError(err, "unsupport type: godence.unsupportType")
		assert.Nil(cadenceValue)
	})

	t.Run("map to Dictionary, unsupport value type", func(t *testing.T) {
		assert := assert.New(t)

		cadenceValue, err := ToCadence(map[string]unsupportType{"MyName": "LemonNeko"})
		assert.EqualError(err, "unsupport type: godence.unsupportType")
		assert.Nil(cadenceValue)
	})
}

func TestToCadenceOptional(t *testing.T) {
	t.Run("to int?", func(t *testing.T) {
		assert := assert.New(t)
		script := []byte(`pub fun main(arg: Int?) { log("the arg is ".concat(arg!.toString())) }`)

		cadenceValue, err := ToCadenceOptional(int(15))
		assert.NoError(err)
		assert.Equal(cadenceValue.Type().ID(), "Int?")

		args := []cadence.Value{cadenceValue}
		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, args)

		assert.NoError(err)
		assert.Equal(ret.Type().ID(), "Void")
	})
	t.Run("to int8?", func(t *testing.T) {
		assert := assert.New(t)
		script := []byte(`pub fun main(arg: Int8?) { log("the arg is ".concat(arg!.toString())) }`)

		cadenceValue, err := ToCadenceOptional(int8(15))
		assert.NoError(err)
		assert.Equal(cadenceValue.Type().ID(), "Int8?")

		args := []cadence.Value{cadenceValue}
		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, args)

		assert.NoError(err)
		assert.Equal(ret.Type().ID(), "Void")
	})
	t.Run("to int16?", func(t *testing.T) {
		assert := assert.New(t)
		script := []byte(`pub fun main(arg: Int16?) { log("the arg is ".concat(arg!.toString())) }`)

		cadenceValue, err := ToCadenceOptional(int16(15))
		assert.NoError(err)
		assert.Equal(cadenceValue.Type().ID(), "Int16?")

		args := []cadence.Value{cadenceValue}
		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, args)

		assert.NoError(err)
		assert.Equal(ret.Type().ID(), "Void")
	})
	t.Run("to int32?", func(t *testing.T) {
		assert := assert.New(t)
		script := []byte(`pub fun main(arg: Int32?) { log("the arg is ".concat(arg!.toString())) }`)

		cadenceValue, err := ToCadenceOptional(int32(15))
		assert.NoError(err)
		assert.Equal(cadenceValue.Type().ID(), "Int32?")

		args := []cadence.Value{cadenceValue}
		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, args)

		assert.NoError(err)
		assert.Equal(ret.Type().ID(), "Void")
	})
	t.Run("to int64?", func(t *testing.T) {
		assert := assert.New(t)
		script := []byte(`pub fun main(arg: Int64?) { log("the arg is ".concat(arg!.toString())) }`)

		cadenceValue, err := ToCadenceOptional(int64(15))
		assert.NoError(err)
		assert.Equal(cadenceValue.Type().ID(), "Int64?")

		args := []cadence.Value{cadenceValue}
		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, args)

		assert.NoError(err)
		assert.Equal(ret.Type().ID(), "Void")
	})
	t.Run("to int128?", func(t *testing.T) {
		assert := assert.New(t)
		script := []byte(`pub fun main(arg: Int128?) { log("the arg is ".concat(arg!.toString())) }`)
		// make a new big int
		bigInt := &big.Int{}
		bigInt, ok := bigInt.SetString("170141183460469231731687303715884105727", 10)
		assert.True(ok)

		cadenceValue, err := ToCadenceOptional(bigInt)
		assert.NoError(err)
		assert.Equal(cadenceValue.Type().ID(), "Int128?")

		args := []cadence.Value{cadenceValue}
		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, args)
		assert.NoError(err)
		assert.Equal(ret.Type().ID(), "Void")
	})
	t.Run("to int256?", func(t *testing.T) {
		assert := assert.New(t)
		script := []byte(`pub fun main(arg: Int256?) { log("the arg is ".concat(arg!.toString())) }`)
		// make a new big int
		bigInt := &big.Int{}
		bigInt, ok := bigInt.SetString("57896044618658097711785492504343953926634992332820282019728792003956564819967", 10)
		assert.True(ok)

		cadenceValue, err := ToCadenceOptional(bigInt)
		assert.NoError(err)
		assert.Equal(cadenceValue.Type().ID(), "Int256?")

		args := []cadence.Value{cadenceValue}
		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, args)
		assert.NoError(err)
		assert.Equal(ret.Type().ID(), "Void")
	})
	t.Run("to uint?", func(t *testing.T) {
		assert := assert.New(t)
		script := []byte(`pub fun main(arg: UInt?) { log("the arg is ".concat(arg!.toString())) }`)

		cadenceValue, err := ToCadenceOptional(uint(15))
		assert.NoError(err)
		assert.Equal(cadenceValue.Type().ID(), "UInt?")

		args := []cadence.Value{cadenceValue}
		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, args)

		assert.NoError(err)
		assert.Equal(ret.Type().ID(), "Void")
	})
	t.Run("to uint8?", func(t *testing.T) {
		assert := assert.New(t)
		script := []byte(`pub fun main(arg: UInt8?) { log("the arg is ".concat(arg!.toString())) }`)

		cadenceValue, err := ToCadenceOptional(uint8(15))
		assert.NoError(err)
		assert.Equal(cadenceValue.Type().ID(), "UInt8?")

		args := []cadence.Value{cadenceValue}
		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, args)

		assert.NoError(err)
		assert.Equal(ret.Type().ID(), "Void")
	})
	t.Run("to uint16?", func(t *testing.T) {
		assert := assert.New(t)
		script := []byte(`pub fun main(arg: UInt16?) { log("the arg is ".concat(arg!.toString())) }`)

		cadenceValue, err := ToCadenceOptional(uint16(15))
		assert.NoError(err)
		assert.Equal(cadenceValue.Type().ID(), "UInt16?")

		args := []cadence.Value{cadenceValue}
		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, args)

		assert.NoError(err)
		assert.Equal(ret.Type().ID(), "Void")
	})
	t.Run("to uint32?", func(t *testing.T) {
		assert := assert.New(t)
		script := []byte(`pub fun main(arg: UInt32?) { log("the arg is ".concat(arg!.toString())) }`)

		cadenceValue, err := ToCadenceOptional(uint32(15))
		assert.NoError(err)
		assert.Equal(cadenceValue.Type().ID(), "UInt32?")

		args := []cadence.Value{cadenceValue}
		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, args)

		assert.NoError(err)
		assert.Equal(ret.Type().ID(), "Void")
	})
	t.Run("to uint64?", func(t *testing.T) {
		assert := assert.New(t)
		script := []byte(`pub fun main(arg: UInt64?) { log("the arg is ".concat(arg!.toString())) }`)

		cadenceValue, err := ToCadenceOptional(uint64(15))
		assert.NoError(err)
		assert.Equal(cadenceValue.Type().ID(), "UInt64?")

		args := []cadence.Value{cadenceValue}
		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, args)

		assert.NoError(err)
		assert.Equal(ret.Type().ID(), "Void")
	})
	t.Run("to uint128?", func(t *testing.T) {
		assert := assert.New(t)
		script := []byte(`pub fun main(arg: UInt128?) { log("the arg is ".concat(arg!.toString())) }`)
		// make a new big int
		bigInt := &big.Int{}
		bigInt, ok := bigInt.SetString("340282366920938463463374607431768211455", 10)
		assert.True(ok)

		cadenceValue, err := ToCadenceOptional(bigInt)
		assert.NoError(err)
		assert.Equal(cadenceValue.Type().ID(), "UInt128?")

		args := []cadence.Value{cadenceValue}
		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, args)
		assert.NoError(err)
		assert.Equal(ret.Type().ID(), "Void")
	})
	t.Run("to uint256?", func(t *testing.T) {
		assert := assert.New(t)
		script := []byte(`pub fun main(arg: UInt256?) { log("the arg is ".concat(arg!.toString())) }`)
		// make a new big int
		bigInt := &big.Int{}
		bigInt, ok := bigInt.SetString("115792089237316195423570985008687907853269984665640564039457584007913129639935", 10)
		assert.True(ok)

		cadenceValue, err := ToCadenceOptional(bigInt)
		assert.NoError(err)
		assert.Equal(cadenceValue.Type().ID(), "UInt256?")

		args := []cadence.Value{cadenceValue}
		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, args)
		assert.NoError(err)
		assert.Equal(ret.Type().ID(), "Void")
	})
	t.Run("to fix64?", func(t *testing.T) {
		assert := assert.New(t)
		script := []byte(`pub fun main(arg: Fix64?) { log("the arg is ".concat(arg!.toString())) }`)

		cadenceValue, err := ToCadenceOptional(Fix64(15))
		assert.NoError(err)
		assert.Equal(cadenceValue.Type().ID(), "Fix64?")

		args := []cadence.Value{cadenceValue}
		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, args)
		assert.NoError(err)
		assert.Equal(ret.Type().ID(), "Void")
	})
	t.Run("to ufix64?", func(t *testing.T) {
		assert := assert.New(t)
		script := []byte(`pub fun main(arg: UFix64?) { log("the arg is ".concat(arg!.toString())) }`)

		cadenceValue, err := ToCadenceOptional(UFix64(15))
		assert.NoError(err)
		assert.Equal(cadenceValue.Type().ID(), "UFix64?")

		args := []cadence.Value{cadenceValue}
		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, args)
		assert.NoError(err)
		assert.Equal(ret.Type().ID(), "Void")
	})
	t.Run("to String?", func(t *testing.T) {
		assert := assert.New(t)
		script := []byte(`pub fun main(arg: String?) { log("the arg is ".concat(arg!)) }`)

		cadenceValue, err := ToCadenceOptional("hello")
		assert.NoError(err)
		assert.Equal(cadenceValue.Type().ID(), "String?")

		args := []cadence.Value{cadenceValue}
		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, args)
		assert.NoError(err)
		assert.Equal(ret.Type().ID(), "Void")
	})
	t.Run("to Bool", func(t *testing.T) {
		assert := assert.New(t)
		script := []byte(`pub fun main(arg: Bool) { log(arg) }`)

		cadenceValue, err := ToCadence(true)
		assert.NoError(err)
		assert.Equal(cadenceValue.Type().ID(), "Bool")

		args := []cadence.Value{cadenceValue}
		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, args)
		assert.NoError(err)
		assert.Equal(ret.Type().ID(), "Void")
	})
	t.Run("to Bool?", func(t *testing.T) {
		assert := assert.New(t)
		script := []byte(`pub fun main(arg: Bool?) { log(arg) }`)

		cadenceValue, err := ToCadenceOptional(true)
		assert.NoError(err)
		assert.Equal(cadenceValue.Type().ID(), "Bool?")

		args := []cadence.Value{cadenceValue}
		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, args)
		assert.NoError(err)
		assert.Equal(ret.Type().ID(), "Void")
	})
	t.Run("error", func(t *testing.T) {
		assert := assert.New(t)

		_, err := ToCadenceOptional(unsupportType(""))
		assert.EqualError(err, "unsupport type: godence.unsupportType")
	})
}

// This test is not for any function
func TestTryToGetSomething(t *testing.T) {
	t.Run("try to get a cadence Struct name", func(t *testing.T) {
		assert := assert.New(t)
		// return a value from cadence, try to get it name
		script := []byte(`
pub struct SimpleStruct {
	pub var A: Int

	init() {
		self.A = 0
	}
}

pub fun main(): SimpleStruct {
	return SimpleStruct()
}`)

		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, nil)
		assert.NoError(err)
		// yeah! i got it.
		assert.Equal(ret.Type().ID(), "s.b37650809c8bd4cb0827a55bc447099f8f8ac555ef5308bfc287f757557430ef.SimpleStruct")
	})
}
