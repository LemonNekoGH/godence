package godence

import (
	"context"
	"math/big"
	"testing"

	"github.com/onflow/cadence"
	"github.com/stretchr/testify/assert"
)

func TestToGo(t *testing.T) {
	t.Run("unsupport type", func(t *testing.T) {
		assert := assert.New(t)
		script := []byte(`pub fun main(): Int { return 15 }`)

		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, nil)
		assert.NoError(err)
		assert.Equal(ret.Type().ID(), "Int")

		dist := UnsupportType("")
		err = ToGo(ret, dist)
		assert.EqualError(err, "unsupport type: godence.UnsupportType")
	})
	t.Run("type cast failed", func(t *testing.T) {
		assert := assert.New(t)
		script := []byte(`pub fun main(): Int16 { return 32767 }`)

		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, nil)
		assert.NoError(err)
		assert.Equal(ret.Type().ID(), "Int16")

		var dist int8
		err = ToGo(ret, &dist)
		assert.EqualError(err, "panic recovered: interface conversion: interface {} is int16, not int8")
	})
	// =============
	// integers
	// =============
	t.Run("cadence Int is big.Int", func(t *testing.T) {
		assert := assert.New(t)
		script := []byte(`pub fun main(): Int { return 127 }`)

		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, nil)
		assert.NoError(err)
		assert.Equal(ret.Type().ID(), "Int")

		var dist *big.Int
		err = ToGo(ret, &dist)
		expect := big.NewInt(127)
		assert.NoError(err)
		assert.Equal(expect, dist)
	})
	t.Run("convert to int8", func(t *testing.T) {
		assert := assert.New(t)
		script := []byte(`pub fun main(): Int8 { return 127 }`)

		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, nil)
		assert.NoError(err)
		assert.Equal(ret.Type().ID(), "Int8")

		var dist int8
		err = ToGo(ret, &dist)
		assert.NoError(err)
		assert.Equal(dist, int8(127))
	})
	t.Run("convert to int16", func(t *testing.T) {
		assert := assert.New(t)
		script := []byte(`pub fun main(): Int16 { return 127 }`)

		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, nil)
		assert.NoError(err)
		assert.Equal(ret.Type().ID(), "Int16")

		var dist int16
		err = ToGo(ret, &dist)
		assert.NoError(err)
		assert.Equal(dist, int16(127))
	})
	t.Run("convert to int32", func(t *testing.T) {
		assert := assert.New(t)
		script := []byte(`pub fun main(): Int32 { return 127 }`)

		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, nil)
		assert.NoError(err)
		assert.Equal(ret.Type().ID(), "Int32")

		var dist int32
		err = ToGo(ret, &dist)
		assert.NoError(err)
		assert.Equal(dist, int32(127))
	})
	t.Run("convert to int64", func(t *testing.T) {
		assert := assert.New(t)
		script := []byte(`pub fun main(): Int64 { return 127 }`)

		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, nil)
		assert.NoError(err)
		assert.Equal(ret.Type().ID(), "Int64")

		var dist int64
		err = ToGo(ret, &dist)
		assert.NoError(err)
		assert.Equal(dist, int64(127))
	})
	t.Run("cadence Int128 is big.Int", func(t *testing.T) {
		assert := assert.New(t)
		script := []byte(`pub fun main(): Int128 { return 127 }`)

		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, nil)
		assert.NoError(err)
		assert.Equal(ret.Type().ID(), "Int128")

		var dist *big.Int
		err = ToGo(ret, &dist)
		expect := big.NewInt(127)
		assert.NoError(err)
		assert.Equal(expect, dist)
	})
	t.Run("cadence Int256 is big.Int", func(t *testing.T) {
		assert := assert.New(t)
		script := []byte(`pub fun main(): Int256 { return 127 }`)

		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, nil)
		assert.NoError(err)
		assert.Equal(ret.Type().ID(), "Int256")

		var dist *big.Int
		err = ToGo(ret, &dist)
		expect := big.NewInt(127)
		assert.NoError(err)
		assert.Equal(expect, dist)
	})
	t.Run("Fix64 convert to int64", func(t *testing.T) {
		assert := assert.New(t)
		script := []byte(`pub fun main(): Fix64 { return 127.0 }`)

		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, nil)
		assert.NoError(err)
		assert.Equal(ret.Type().ID(), "Fix64")

		var dist int64
		err = ToGo(ret, &dist)
		assert.NoError(err)
		assert.Equal(dist, int64(12700000000))
	})
	// =============
	// unsigned integers
	// =============
	t.Run("cadence UInt is big.Int", func(t *testing.T) {
		assert := assert.New(t)
		script := []byte(`pub fun main(): UInt { return 127 }`)

		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, nil)
		assert.NoError(err)
		assert.Equal(ret.Type().ID(), "UInt")

		var dist *big.Int
		err = ToGo(ret, &dist)
		expect := big.NewInt(127)
		assert.NoError(err)
		assert.Equal(expect, dist)
	})
	t.Run("convert to uint8", func(t *testing.T) {
		assert := assert.New(t)
		script := []byte(`pub fun main(): UInt8 { return 127 }`)

		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, nil)
		assert.NoError(err)
		assert.Equal(ret.Type().ID(), "UInt8")

		var dist uint8
		err = ToGo(ret, &dist)
		assert.NoError(err)
		assert.Equal(dist, uint8(127))
	})
	t.Run("convert to uint16", func(t *testing.T) {
		assert := assert.New(t)
		script := []byte(`pub fun main(): UInt16 { return 127 }`)

		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, nil)
		assert.NoError(err)
		assert.Equal(ret.Type().ID(), "UInt16")

		var dist uint16
		err = ToGo(ret, &dist)
		assert.NoError(err)
		assert.Equal(dist, uint16(127))
	})
	t.Run("convert to uint32", func(t *testing.T) {
		assert := assert.New(t)
		script := []byte(`pub fun main(): UInt32 { return 127 }`)

		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, nil)
		assert.NoError(err)
		assert.Equal(ret.Type().ID(), "UInt32")

		var dist uint32
		err = ToGo(ret, &dist)
		assert.NoError(err)
		assert.Equal(dist, uint32(127))
	})
	t.Run("convert to uint64", func(t *testing.T) {
		assert := assert.New(t)
		script := []byte(`pub fun main(): UInt64 { return 127 }`)

		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, nil)
		assert.NoError(err)
		assert.Equal(ret.Type().ID(), "UInt64")

		var dist uint64
		err = ToGo(ret, &dist)
		assert.NoError(err)
		assert.Equal(dist, uint64(127))
	})
	t.Run("cadence UInt128 is big.Int", func(t *testing.T) {
		assert := assert.New(t)
		script := []byte(`pub fun main(): UInt128 { return 127 }`)

		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, nil)
		assert.NoError(err)
		assert.Equal(ret.Type().ID(), "UInt128")

		var dist *big.Int
		err = ToGo(ret, &dist)
		expect := big.NewInt(127)
		assert.NoError(err)
		assert.Equal(expect, dist)
	})
	t.Run("cadence UInt256 is big.Int", func(t *testing.T) {
		assert := assert.New(t)
		script := []byte(`pub fun main(): UInt256 { return 127 }`)

		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, nil)
		assert.NoError(err)
		assert.Equal(ret.Type().ID(), "UInt256")

		var dist *big.Int
		err = ToGo(ret, &dist)
		expect := big.NewInt(127)
		assert.NoError(err)
		assert.Equal(expect, dist)
	})
	t.Run("UFix64 convert to uint64", func(t *testing.T) {
		assert := assert.New(t)
		script := []byte(`pub fun main(): UFix64 { return 127.0 }`)

		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, nil)
		assert.NoError(err)
		assert.Equal(ret.Type().ID(), "UFix64")

		var dist uint64
		err = ToGo(ret, &dist)
		assert.NoError(err)
		assert.Equal(dist, uint64(12700000000))
	})
	// =============
	// strings
	// =============
	t.Run("convert to string", func(t *testing.T) {
		assert := assert.New(t)
		script := []byte(`pub fun main(): String { return "Hello" }`)

		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, nil)
		assert.NoError(err)
		assert.Equal(ret.Type().ID(), "String")

		var dist string
		err = ToGo(ret, &dist)
		assert.NoError(err)
		assert.Equal("Hello", dist)
	})
	t.Run("Address convert to string", func(t *testing.T) {
		assert := assert.New(t)
		script := []byte(`pub fun main(): Address { return 0x0 }`)

		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, nil)
		assert.NoError(err)
		assert.Equal(ret.Type().ID(), "Address")

		var dist string
		err = ToGo(ret, &dist)
		assert.NoError(err)
		assert.Equal("0x0000000000000000", dist)
	})
	t.Run("Address convert to [8]uint8", func(t *testing.T) {
		assert := assert.New(t)
		script := []byte(`pub fun main(): Address { return 0x0 }`)

		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, nil)
		assert.NoError(err)
		assert.Equal(ret.Type().ID(), "Address")

		var dist [8]uint8
		err = ToGo(ret, &dist)
		assert.NoError(err)
		assert.Equal([8]uint8{0, 0, 0, 0, 0, 0, 0, 0}, dist)
	})
	t.Run("Address convert to cadence.Address", func(t *testing.T) {
		assert := assert.New(t)
		script := []byte(`pub fun main(): Address { return 0x0 }`)

		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, nil)
		assert.NoError(err)
		assert.Equal(ret.Type().ID(), "Address")

		var dist cadence.Address
		err = ToGo(ret, &dist)
		assert.NoError(err)
		assert.Equal(cadence.Address{0, 0, 0, 0, 0, 0, 0, 0}, dist)
	})
	t.Run("Bool convert to bool", func(t *testing.T) {
		assert := assert.New(t)
		script := []byte(`pub fun main(): Bool { return true }`)

		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, nil)
		assert.NoError(err)
		assert.Equal(ret.Type().ID(), "Bool")

		var dist bool
		err = ToGo(ret, &dist)
		assert.NoError(err)
		assert.True(dist)
	})
}
