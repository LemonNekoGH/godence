package godence

import (
	"context"
	"math/big"
	"testing"

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
}
