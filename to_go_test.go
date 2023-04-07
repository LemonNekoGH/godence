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

		dist := unsupportType("")
		err = ToGo(ret, dist)
		assert.EqualError(err, "unsupport type: godence.unsupportType")
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
	// Optional
	// =============
	t.Run("Optional", func(t *testing.T) {
		a := assert.New(t)
		script := []byte(`pub fun main(): String? { return "Test" }`)

		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, nil)
		a.NoError(err)
		a.Equal(ret.Type().ID(), "String?")

		var dist string
		err = ToGo(ret, &dist)
		a.Nil(err)
		a.Equal("Test", dist)
	})
	t.Run("Optional, nil", func(t *testing.T) {
		a := assert.New(t)
		script := []byte(`pub fun main(): String? { return nil }`)

		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, nil)
		a.NoError(err)
		a.Equal(ret.Type().ID(), "Never?")

		var dist string
		err = ToGo(ret, &dist)
		a.Nil(err)
		a.Equal("", dist)
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
	t.Run("Path covert to string", func(t *testing.T) {
		assert := assert.New(t)
		script := []byte(`pub fun main(): Path { return /public/collectionStoragePath }`)

		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, nil)
		assert.NoError(err)
		assert.Equal(ret.Type().ID(), "Path")

		var dist string
		err = ToGo(ret, &dist)
		assert.NoError(err)
		assert.Equal("/public/collectionStoragePath", dist)
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
	t.Run("Character convert to string", func(t *testing.T) {
		assert := assert.New(t)
		script := []byte(`pub fun main(): Character { return "a" }`)

		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, nil)
		assert.NoError(err)
		assert.Equal(ret.Type().ID(), "Character")

		var dist string
		err = ToGo(ret, &dist)
		assert.NoError(err)
		assert.Equal("a", dist)
	})

	t.Run("a simple struct", func(t *testing.T) {
		type simpleStruct struct {
			MyName string
		}
		assert := assert.New(t)
		script := []byte(`
pub struct SimpleStruct {
	pub var MyName: String

	init() {
		self.MyName = "LemonNeko"
	}
}
pub fun main(): SimpleStruct {
	return SimpleStruct()
}`)
		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, nil)
		assert.NoError(err)

		dist := simpleStruct{}
		err = ToGo(ret, &dist)
		assert.NoError(err)
		assert.Equal("LemonNeko", dist.MyName)
	})

	t.Run("a simple optional struct, nil", func(t *testing.T) {
		type simpleStruct struct {
			MyName string
		}
		assert := assert.New(t)
		script := []byte(`
pub struct SimpleStruct {
	pub var MyName: String

	init() {
		self.MyName = "LemonNeko"
	}
}
pub fun main(): SimpleStruct? {
	return nil
}`)
		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, nil)
		assert.NoError(err)

		dist := simpleStruct{}
		err = ToGo(ret, &dist)
		assert.NoError(err)
		assert.Zero(dist)
	})

	t.Run("a simple optional struct, not nil", func(t *testing.T) {
		type simpleStruct struct {
			MyName string
		}
		assert := assert.New(t)
		script := []byte(`
pub struct SimpleStruct {
	pub var MyName: String

	init() {
		self.MyName = "LemonNeko"
	}
}
pub fun main(): SimpleStruct? {
	return SimpleStruct()
}`)
		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, nil)
		assert.NoError(err)

		dist := simpleStruct{}
		err = ToGo(ret, &dist)
		assert.NoError(err)
		assert.Equal("LemonNeko", dist.MyName)
	})

	t.Run("a string string dictionary", func(t *testing.T) {
		assert := assert.New(t)
		defer func() {
			err := recover()
			assert.Nil(err)
		}()

		script := []byte(`
pub fun main(): {String: String} {
	return {
		"MyName": "LemonNeko"
	}
}`)
		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, nil)
		assert.NoError(err)

		dist := map[string]string{}
		err = ToGo(ret, dist)
		assert.NoError(err)
		assert.Equal("LemonNeko", dist["MyName"])
	})

	t.Run("a string slice", func(t *testing.T) {
		assert := assert.New(t)
		script := []byte(`
pub fun main(): [String] {
	return ["MyName", "LemonNeko"]
}`)
		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, nil)
		assert.NoError(err)

		dist := []string{}
		err = ToGo(ret, &dist)
		assert.NoError(err)
		assert.Equal([]string{"MyName", "LemonNeko"}, dist)
	})
}

// test for structToGoStruct
func TestToGoStruct(t *testing.T) {
	type structContainsManyType struct {
		IntValue     *big.Int `godence:"intValue"`
		Int8Value    int8     `godence:"int8Value"`
		Int16Value   int16    `godence:"int16Value"`
		Int32Value   int32    `godence:"int32Value"`
		Int64Value   int64    `godence:"int64Value"`
		Int128Value  *big.Int `godence:"int128Value"`
		Int256Value  *big.Int `godence:"int256Value"`
		UIntValue    *big.Int `godence:"uintValue"`
		UInt8Value   uint8    `godence:"uint8Value"`
		UInt16Value  uint16   `godence:"uint16Value"`
		UInt32Value  uint32   `godence:"uint32Value"`
		UInt64Value  uint64   `godence:"uint64Value"`
		UInt128Value *big.Int `godence:"uint128Value"`
		UInt256Value *big.Int `godence:"uint256Value"`
		StringValue  string   `godence:"stringValue"`
		AddressValue string   `godence:"addressValue"`
		BoolValue    bool     `godence:"boolValue"`
	}

	t.Run("not a struct", func(t *testing.T) {
		type simpleStruct struct {
			MyName string
		}
		assert := assert.New(t)
		script := []byte(`
pub fun main(): String {
	return "LemonNeko"
}`)
		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, nil)
		assert.NoError(err)

		dist := simpleStruct{}
		err = toGoStruct(ret, &dist)
		assert.EqualError(err, "to go struct: unsupport cadence type: cadence.String")
	})

	t.Run("a simple struct", func(t *testing.T) {
		type simpleStruct struct {
			MyName string
		}
		assert := assert.New(t)
		script := []byte(`
pub struct SimpleStruct {
	pub var MyName: String

	init() {
		self.MyName = "LemonNeko"
	}
}
pub fun main(): SimpleStruct {
	return SimpleStruct()
}`)
		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, nil)
		assert.NoError(err)

		dist := simpleStruct{}
		err = toGoStruct(ret, &dist)
		assert.NoError(err)
		assert.Equal("LemonNeko", dist.MyName)
	})

	t.Run("a simple struct, but no any field exported", func(t *testing.T) {
		type simpleStruct struct {
			myName string
		}
		assert := assert.New(t)
		script := []byte(`
pub struct SimpleStruct {
	pub var MyName: String

	init() {
		self.MyName = "LemonNeko"
	}
}
pub fun main(): SimpleStruct {
	return SimpleStruct()
}`)
		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, nil)
		assert.NoError(err)

		dist := simpleStruct{}
		err = toGoStruct(ret, &dist)
		assert.NoError(err)
		assert.Equal("", dist.myName)
	})

	t.Run("a simple struct, with field name tag", func(t *testing.T) {
		type simpleStruct struct {
			MyName string `godence:"myName"`
		}
		assert := assert.New(t)
		script := []byte(`
pub struct SimpleStruct {
	pub var myName: String

	init() {
		self.myName = "LemonNeko"
	}
}
pub fun main(): SimpleStruct {
	return SimpleStruct()
}`)
		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, nil)
		assert.NoError(err)

		dist := simpleStruct{}
		err = toGoStruct(ret, &dist)
		assert.NoError(err)
		assert.Equal("LemonNeko", dist.MyName)
	})

	t.Run("a simple struct, with wrong field name tag", func(t *testing.T) {
		type simpleStruct struct {
			MyName string `godence:"none"`
		}
		assert := assert.New(t)
		script := []byte(`
pub struct SimpleStruct {
	pub var myName: String

	init() {
		self.myName = "LemonNeko"
	}
}
pub fun main(): SimpleStruct {
	return SimpleStruct()
}`)
		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, nil)
		assert.NoError(err)

		dist := simpleStruct{}
		err = toGoStruct(ret, &dist)
		assert.EqualError(err, "cannot find field named none in cadence struct/event/resource")
	})

	t.Run("a simple struct, with wrong type", func(t *testing.T) {
		type simpleStruct struct {
			MyName string `godence:"myName"`
		}
		assert := assert.New(t)
		script := []byte(`
pub struct SimpleStruct {
	pub var myName: UInt8

	init() {
		self.myName = 127
	}
}
pub fun main(): SimpleStruct {
	return SimpleStruct()
}`)
		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, nil)
		assert.NoError(err)

		dist := simpleStruct{}
		err = toGoStruct(ret, &dist)
		assert.EqualError(err, "structEventResourceToGoStruct, panic recoverd: reflect.Set: value of type uint8 is not assignable to type string")
	})

	t.Run("a struct contains many type", func(t *testing.T) {
		assert := assert.New(t)
		script := []byte(`
pub struct StructContainsManyType {
	pub var intValue: Int
	pub var int8Value: Int8
	pub var int16Value: Int16
	pub var int32Value: Int32
	pub var int64Value: Int64
	pub var int128Value: Int128
	pub var int256Value: Int256
	pub var uintValue: UInt
	pub var uint8Value: UInt8
	pub var uint16Value: UInt16
	pub var uint32Value: UInt32
	pub var uint64Value: UInt64
	pub var uint128Value: UInt128
	pub var uint256Value: UInt256
	pub var stringValue: String
	pub var addressValue: Address
	pub var boolValue: Bool

	init() {
		self.intValue = 127
		self.int8Value = 127
		self.int16Value = 127
		self.int32Value = 127
		self.int64Value = 127
		self.int128Value = 127
		self.int256Value = 127
		self.uintValue = 127
		self.uint8Value = 127
		self.uint16Value = 127
		self.uint32Value = 127
		self.uint64Value = 127
		self.uint128Value = 127
		self.uint256Value = 127
		self.stringValue = "LemonNeko"
		self.addressValue = 0x0
		self.boolValue = true
	}
}
pub fun main(): StructContainsManyType {
	return StructContainsManyType()
}`)
		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, nil)
		assert.NoError(err)

		dist := structContainsManyType{}
		err = toGoStruct(ret, &dist)
		assert.NoError(err)
		assert.Equal(big.NewInt(127), dist.IntValue)
		assert.Equal(int8(127), dist.Int8Value)
		assert.Equal(int16(127), dist.Int16Value)
		assert.Equal(int32(127), dist.Int32Value)
		assert.Equal(int64(127), dist.Int64Value)
		assert.Equal(big.NewInt(127), dist.Int128Value)
		assert.Equal(big.NewInt(127), dist.Int256Value)
		assert.Equal(big.NewInt(127), dist.UIntValue)
		assert.Equal(uint8(127), dist.UInt8Value)
		assert.Equal(uint16(127), dist.UInt16Value)
		assert.Equal(uint32(127), dist.UInt32Value)
		assert.Equal(uint64(127), dist.UInt64Value)
		assert.Equal(big.NewInt(127), dist.UInt128Value)
		assert.Equal(big.NewInt(127), dist.UInt256Value)
		assert.Equal("LemonNeko", dist.StringValue)
		assert.Equal("0x0000000000000000", dist.AddressValue)
		assert.True(dist.BoolValue)
	})

	t.Run("embedded struct", func(t *testing.T) {
		type embeddedStruct struct {
			Inter *struct {
				MyName string `godence:"myName"`
			} `godence:"inter"`
		}
		assert := assert.New(t)
		script := []byte(`
	pub struct Inter {
		pub var myName: String

		init() {
			self.myName = "LemonNeko"
		}
	}

	pub struct EmbeddedStruct {
		pub var inter: Inter

		init() {
			self.inter = Inter()
		}
	}
	pub fun main(): EmbeddedStruct {
		return EmbeddedStruct()
	}`)
		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, nil)
		assert.NoError(err)

		dist := embeddedStruct{}
		err = toGoStruct(ret, &dist)
		assert.NoError(err)
		assert.Equal("LemonNeko", dist.Inter.MyName)
	})

	t.Run("a simple event", func(t *testing.T) {
		type simpleStruct struct {
			MyName string
		}
		assert := assert.New(t)
		script := []byte(`
import ForTest from 0xf8d6e0586b0a20c7

transaction {
	prepare() {
		ForTest.emitSimple()
	}
}
		`)
		tx := buildSimpleTx(script, assert)
		err := flowCli.SendTransaction(context.Background(), *tx)
		assert.NoError(err)

		result := waitForTransactionSealed(tx, assert)
		assert.NoError(result.Error)
		assert.Equal(1, len(result.Events))
		assert.Equal("A.f8d6e0586b0a20c7.ForTest.Simple", result.Events[0].Value.Type().ID())

		dist := simpleStruct{}
		err = toGoStruct(result.Events[0].Value, &dist)
		assert.NoError(err)
		assert.Equal("LemonNeko", dist.MyName)
	})

	t.Run("a simple event, with field name tag", func(t *testing.T) {
		type simpleStruct struct {
			MyName string `godence:"myName"`
		}
		assert := assert.New(t)
		script := []byte(`
import ForTest from 0xf8d6e0586b0a20c7

transaction {
	prepare() {
		ForTest.emitSimple2()
	}
}
		`)
		tx := buildSimpleTx(script, assert)
		err := flowCli.SendTransaction(context.Background(), *tx)
		assert.NoError(err)

		result := waitForTransactionSealed(tx, assert)
		assert.NoError(result.Error)
		assert.Equal(1, len(result.Events))
		assert.Equal("A.f8d6e0586b0a20c7.ForTest.Simple2", result.Events[0].Value.Type().ID())

		dist := simpleStruct{}
		err = toGoStruct(result.Events[0].Value, &dist)
		assert.NoError(err)
		assert.Equal("LemonNeko", dist.MyName)
	})

	t.Run("a simple event, with wrong field name tag", func(t *testing.T) {
		type simpleStruct struct {
			MyName string `godence:"none"`
		}
		assert := assert.New(t)
		script := []byte(`
import ForTest from 0xf8d6e0586b0a20c7

transaction {
	prepare() {
		ForTest.emitSimple2()
	}
}
		`)
		tx := buildSimpleTx(script, assert)
		err := flowCli.SendTransaction(context.Background(), *tx)
		assert.NoError(err)

		result := waitForTransactionSealed(tx, assert)
		assert.NoError(result.Error)
		assert.Equal(1, len(result.Events))
		assert.Equal("A.f8d6e0586b0a20c7.ForTest.Simple2", result.Events[0].Value.Type().ID())

		dist := simpleStruct{}
		err = toGoStruct(result.Events[0].Value, &dist)
		assert.EqualError(err, "cannot find field named none in cadence struct/event/resource")
	})

	t.Run("a simple event, with wrong type", func(t *testing.T) {
		type simpleStruct struct {
			MyName string `godence:"myName"`
		}
		assert := assert.New(t)
		script := []byte(`
import ForTest from 0xf8d6e0586b0a20c7

transaction {
	prepare() {
		ForTest.emitSimple3()
	}
}
		`)
		tx := buildSimpleTx(script, assert)
		err := flowCli.SendTransaction(context.Background(), *tx)
		assert.NoError(err)

		result := waitForTransactionSealed(tx, assert)
		assert.NoError(result.Error)
		assert.Equal(1, len(result.Events))
		assert.Equal("A.f8d6e0586b0a20c7.ForTest.Simple3", result.Events[0].Value.Type().ID())

		dist := simpleStruct{}
		err = toGoStruct(result.Events[0].Value, &dist)
		assert.EqualError(err, "structEventResourceToGoStruct, panic recoverd: reflect.Set: value of type uint8 is not assignable to type string")
	})

	t.Run("a event contains many type", func(t *testing.T) {
		assert := assert.New(t)
		script := []byte(`
import ForTest from 0xf8d6e0586b0a20c7

transaction {
	prepare() {
		ForTest.emitManyType()
	}
}
		`)
		tx := buildSimpleTx(script, assert)
		err := flowCli.SendTransaction(context.Background(), *tx)
		assert.NoError(err)

		result := waitForTransactionSealed(tx, assert)
		assert.NoError(result.Error)
		assert.Equal(1, len(result.Events))
		assert.Equal("A.f8d6e0586b0a20c7.ForTest.ManyType", result.Events[0].Value.Type().ID())

		dist := structContainsManyType{}
		err = toGoStruct(result.Events[0].Value, &dist)
		assert.NoError(err)
		assert.Equal(big.NewInt(127), dist.IntValue)
		assert.Equal(int8(127), dist.Int8Value)
		assert.Equal(int16(127), dist.Int16Value)
		assert.Equal(int32(127), dist.Int32Value)
		assert.Equal(int64(127), dist.Int64Value)
		assert.Equal(big.NewInt(127), dist.Int128Value)
		assert.Equal(big.NewInt(127), dist.Int256Value)
		assert.Equal(big.NewInt(127), dist.UIntValue)
		assert.Equal(uint8(127), dist.UInt8Value)
		assert.Equal(uint16(127), dist.UInt16Value)
		assert.Equal(uint32(127), dist.UInt32Value)
		assert.Equal(uint64(127), dist.UInt64Value)
		assert.Equal(big.NewInt(127), dist.UInt128Value)
		assert.Equal(big.NewInt(127), dist.UInt256Value)
		assert.Equal("LemonNeko", dist.StringValue)
		assert.Equal("0x0000000000000000", dist.AddressValue)
		assert.True(dist.BoolValue)
	})

	t.Run("embedded struct", func(t *testing.T) {
		type structInParam struct {
			P *struct {
				MyName string `godence:"myName"`
			} `godence:"p"`
		}
		assert := assert.New(t)
		script := []byte(`
		import ForTest from 0xf8d6e0586b0a20c7

		transaction {
			prepare() {
				ForTest.emitStructInParam()
			}
		}`)

		tx := buildSimpleTx(script, assert)
		err := flowCli.SendTransaction(context.Background(), *tx)
		assert.NoError(err)

		result := waitForTransactionSealed(tx, assert)
		assert.NoError(result.Error)
		assert.Equal(1, len(result.Events))
		assert.Equal("A.f8d6e0586b0a20c7.ForTest.StructInParam", result.Events[0].Value.Type().ID())

		dist := structInParam{}
		err = toGoStruct(result.Events[0].Value, &dist)
		assert.NoError(err)
		assert.Equal("LemonNeko", dist.P.MyName)
	})

	t.Run("a simple resource", func(t *testing.T) {
		type simpleStruct struct {
			MyName string
		}
		assert := assert.New(t)
		script := []byte(`
import ForTest from 0xf8d6e0586b0a20c7

pub fun main(): &ForTest.SimpleR {
	return getAccount(0xf8d6e0586b0a20c7).getCapability<&ForTest.SimpleR>(/public/simpleR).borrow()!
}`)
		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, nil)
		assert.NoError(err)

		dist := simpleStruct{}
		err = toGoStruct(ret, &dist)
		assert.NoError(err)
		assert.Equal("LemonNeko", dist.MyName)
	})

	t.Run("a simple resource with name tag", func(t *testing.T) {
		type simpleStruct struct {
			MyName string `godence:"myName"`
		}
		assert := assert.New(t)
		script := []byte(`
import ForTest from 0xf8d6e0586b0a20c7

pub fun main(): &ForTest.SimpleR2 {
	return getAccount(0xf8d6e0586b0a20c7).getCapability<&ForTest.SimpleR2>(/public/simpleR2).borrow()!
}`)
		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, nil)
		assert.NoError(err)

		dist := simpleStruct{}
		err = toGoStruct(ret, &dist)
		assert.NoError(err)
		assert.Equal("LemonNeko", dist.MyName)
	})

	t.Run("a simple resource with wrong name tag", func(t *testing.T) {
		type simpleStruct struct {
			MyName string `godence:"none"`
		}
		assert := assert.New(t)
		script := []byte(`
import ForTest from 0xf8d6e0586b0a20c7

pub fun main(): &ForTest.SimpleR2 {
	return getAccount(0xf8d6e0586b0a20c7).getCapability<&ForTest.SimpleR2>(/public/simpleR2).borrow()!
}`)
		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, nil)
		assert.NoError(err)

		dist := simpleStruct{}
		err = toGoStruct(ret, &dist)
		assert.EqualError(err, "cannot find field named none in cadence struct/event/resource")
	})

	t.Run("a resource contains many type", func(t *testing.T) {
		assert := assert.New(t)
		script := []byte(`
import ForTest from 0xf8d6e0586b0a20c7

pub fun main(): &ForTest.ResourceContainsManyType {
	return getAccount(0xf8d6e0586b0a20c7).getCapability<&ForTest.ResourceContainsManyType>(/public/rMany).borrow()!
}`)
		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, nil)
		assert.NoError(err)

		dist := structContainsManyType{}
		err = toGoStruct(ret, &dist)
		assert.NoError(err)
		assert.Equal(big.NewInt(127), dist.IntValue)
		assert.Equal(int8(127), dist.Int8Value)
		assert.Equal(int16(127), dist.Int16Value)
		assert.Equal(int32(127), dist.Int32Value)
		assert.Equal(int64(127), dist.Int64Value)
		assert.Equal(big.NewInt(127), dist.Int128Value)
		assert.Equal(big.NewInt(127), dist.Int256Value)
		assert.Equal(big.NewInt(127), dist.UIntValue)
		assert.Equal(uint8(127), dist.UInt8Value)
		assert.Equal(uint16(127), dist.UInt16Value)
		assert.Equal(uint32(127), dist.UInt32Value)
		assert.Equal(uint64(127), dist.UInt64Value)
		assert.Equal(big.NewInt(127), dist.UInt128Value)
		assert.Equal(big.NewInt(127), dist.UInt256Value)
		assert.Equal("LemonNeko", dist.StringValue)
		assert.Equal("0x0000000000000000", dist.AddressValue)
		assert.True(dist.BoolValue)
	})

	t.Run("a struct embedded resource", func(t *testing.T) {
		type structInParam struct {
			P *struct {
				MyName string `godence:"myName"`
			} `godence:"forEmbedded"`
		}
		assert := assert.New(t)
		script := []byte(`
import ForTest from 0xf8d6e0586b0a20c7

pub fun main(): &ForTest.AStructEmbeddedInAResource {
	return getAccount(0xf8d6e0586b0a20c7).getCapability<&ForTest.AStructEmbeddedInAResource>(/public/embeddedStruct).borrow()!
}`)
		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, nil)
		assert.NoError(err)

		dist := structInParam{}
		err = toGoStruct(ret, &dist)
		assert.NoError(err)
		assert.Equal("LemonNeko", dist.P.MyName)
	})

	t.Run("a resource embedded resource", func(t *testing.T) {
		type structInParam struct {
			P *struct {
				MyName string `godence:"myName"`
			} `godence:"forEmbedded"`
		}
		assert := assert.New(t)
		script := []byte(`
import ForTest from 0xf8d6e0586b0a20c7

pub fun main(): &ForTest.AResourceEmbeddedInAResource {
	return getAccount(0xf8d6e0586b0a20c7).getCapability<&ForTest.AResourceEmbeddedInAResource>(/public/embeddedResource).borrow()!
}`)
		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, nil)
		assert.NoError(err)

		dist := structInParam{}
		err = toGoStruct(ret, &dist)
		assert.NoError(err)
		assert.Equal("LemonNeko", dist.P.MyName)
	})
}

func TestToGoMap(t *testing.T) {
	t.Run("a string string dictionary", func(t *testing.T) {
		assert := assert.New(t)
		defer func() {
			err := recover()
			assert.Nil(err)
		}()

		script := []byte(`
pub fun main(): {String: String} {
	return {
		"MyName": "LemonNeko"
	}
}`)
		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, nil)
		assert.NoError(err)

		dist := map[string]string{}
		err = toGoMap(ret, dist)
		assert.NoError(err)
		assert.Equal("LemonNeko", dist["MyName"])
	})

	t.Run("a int64 int64 dictionary", func(t *testing.T) {
		assert := assert.New(t)
		defer func() {
			err := recover()
			assert.Nil(err)
		}()

		script := []byte(`
pub fun main(): {Int64: Int64} {
	return {
		88: 64
	}
}`)
		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, nil)
		assert.NoError(err)

		dist := map[int64]int64{}
		err = toGoMap(ret, dist)
		assert.NoError(err)
		assert.Equal(int64(64), dist[88])
	})

	t.Run("dictionary type mismatched", func(t *testing.T) {
		assert := assert.New(t)
		defer func() {
			err := recover()
			assert.Nil(err)
		}()

		script := []byte(`
pub fun main(): {String: String} {
	return {"MyName": "LemonNeko"}
}`)
		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, nil)
		assert.NoError(err)

		dist := map[uint64]uint64{}
		err = toGoMap(ret, dist)
		assert.EqualError(err, "toGoMap, panic recoverd: reflect.Value.SetMapIndex: value of type string is not assignable to type uint64")
	})

	t.Run("not a dictionary", func(t *testing.T) {
		assert := assert.New(t)
		defer func() {
			err := recover()
			assert.Nil(err)
		}()

		script := []byte(`
pub fun main(): String {
	return "LemonNeko"
}`)
		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, nil)
		assert.NoError(err)

		dist := map[string]string{}
		err = toGoMap(ret, dist)
		assert.EqualError(err, "toGoMap, panic recoverd: interface conversion: cadence.Value is cadence.String, not cadence.Dictionary")
	})
}

func TestToGoSlice(t *testing.T) {
	t.Run("a string slice", func(t *testing.T) {
		assert := assert.New(t)
		script := []byte(`
pub fun main(): [String] {
	return ["MyName", "LemonNeko"]
}`)
		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, nil)
		assert.NoError(err)

		dist := []string{}
		err = toGoSlice(ret, &dist)
		assert.NoError(err)
		assert.Equal([]string{"MyName", "LemonNeko"}, dist)
	})

	t.Run("a int64 slice", func(t *testing.T) {
		assert := assert.New(t)
		script := []byte(`
pub fun main(): [Int64] {
	return [88, 64]
}`)
		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, nil)
		assert.NoError(err)

		dist := []int64{}
		err = toGoSlice(ret, &dist)
		assert.NoError(err)
		assert.Equal([]int64{88, 64}, dist)
	})

	t.Run("array type mismatched", func(t *testing.T) {
		assert := assert.New(t)
		script := []byte(`
pub fun main(): [String] {
	return ["MyName", "LemonNeko"]
}`)
		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, nil)
		assert.NoError(err)

		dist := []uint64{}
		err = toGoSlice(ret, &dist)
		assert.EqualError(err, "toGoSlice, panic recoverd: reflect.Set: value of type string is not assignable to type uint64")
	})

	t.Run("not a slice", func(t *testing.T) {
		assert := assert.New(t)

		script := []byte(`
pub fun main(): String {
	return "LemonNeko"
}`)
		ret, err := flowCli.ExecuteScriptAtLatestBlock(context.Background(), script, nil)
		assert.NoError(err)

		dist := []string{}
		err = toGoSlice(ret, &dist)
		assert.EqualError(err, "toGoSlice, panic recoverd: interface conversion: cadence.Value is cadence.String, not cadence.Array")
	})
}

func Test_isValueAddressOrPath(t *testing.T) {
	tests := []struct {
		name string
		arg  cadence.Value
		want bool
	}{
		{name: "any", arg: cadence.NewVoid(), want: false},
		{name: "bool", arg: cadence.NewBool(false), want: false},
		{name: "string", arg: cadence.NewAddress([8]byte{}), want: true},
		{name: "string", arg: cadence.NewPath("", ""), want: true},
		{name: "*big.Int", arg: cadence.NewInt(0), want: false},
		{name: "int8", arg: cadence.NewInt8(0), want: false},
		{name: "int16", arg: cadence.NewInt16(0), want: false},
		{name: "int32", arg: cadence.NewInt32(0), want: false},
		{name: "int64", arg: cadence.NewInt64(0), want: false},
		{name: "*big.Int", arg: cadence.NewUInt(0), want: false},
		{name: "uint8", arg: cadence.NewUInt8(0), want: false},
		{name: "uint16", arg: cadence.NewUInt16(0), want: false},
		{name: "uint32", arg: cadence.NewUInt32(0), want: false},
		{name: "uint64", arg: cadence.NewUInt64(0), want: false},
		{name: "uint8", arg: cadence.NewWord8(0), want: false},
		{name: "uint32", arg: cadence.NewWord32(0), want: false},
		{name: "uint64", arg: cadence.NewWord64(0), want: false},
		{name: "*big.Int", arg: cadence.NewInt128(0), want: false},
		{name: "*big.Int", arg: cadence.NewUInt128(0), want: false},
		{name: "*big.Int", arg: cadence.NewInt256(0), want: false},
		{name: "*big.Int", arg: cadence.NewUInt256(0), want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValueAddressOrPath(tt.arg); got != tt.want {
				t.Errorf("isValueAddressOrPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
