# Godence
[![codecov](https://codecov.io/gh/LemonNekoGH/godence/branch/main/graph/badge.svg?token=KT1RNHTIQZ)](https://codecov.io/gh/LemonNekoGH/godence)
![Unit Test](https://github.com/LemonNekoGH/godence/actions/workflows/main_unittest.yaml/badge.svg)

## Get Started
install by go get.
```
go get github.com/LemonNekoGH/godence
```
### Convert Cadence value to Go value
If you have the following Cadence struct.
```cadence
pub struct Person {
    pub var age: UInt8
    pub var Name: String
}
```
```go
import "github.com/LemonNekoGH/godence"

type Person struct {
    Age uint8 `godence:"age"` // you can specify field name in cadence by tag.
    Name string
}

func main() {
    dist := &Person{}
    // Omit: You got a Cadence value from a script or transaction. And return value named 'ret'.
    godence.ToGo(ret, dist)
}
```
### Convert Go value to Cadence value
Only support for base types now.
## Testing
### Requirements
- [Flow CLI](https://docs.onflow.org/flow-cli/): Use to emulate flow network.

### Steps
1. Start the flow emulator with command: `flow emulator`
2. Open another terminal and run `go clean -testcache && go test .`

## TODO-List: Go to Cadence
### Integers
- [x] Go `int` to Cadence `Int`
- [x] Go `int8` to Cadence `Int8`
- [x] Go `int16` to Cadence `Int16`
- [x] Go `int32` to Cadence `Int32`
- [x] Go `int64` to Cadence `Int64`
- [x] Go `*big.Int` to Cadence `Int128`
- [x] Go `*big.Int` to Cadence `Int256`
- [x] Go `uint` to Cadence `UInt`
- [x] Go `uint8` to Cadence `UInt8`
- [x] Go `uint16` to Cadence `UInt16`
- [x] Go `uint32` to Cadence `UInt32`
- [x] Go `uint64` to Cadence `UInt64`
- [x] Go `*big.Int` to Cadence `UInt128`
- [x] Go `*big.Int` to Cadence `UInt256`
### Fixed-Point Numbers
- [x] Go `int64` to Cadence `Fix64`
- [x] Go `uint64` to Cadence `UFix64`
### Other
- [x] Go `string` to Cadence `String`
- [ ] Go `?` to Cadence `Path`
- [ ] Go `?` to Cadence `Address`
- [x] Go `bool` to Cadence `Bool`
- [ ] Go `?` to Cadence `Array`
- [ ] Go `?` to Cadence `Struct`
- [ ] Go `?` to Cadence `Character`
- [ ] Go `?` to Cadence `Resource`
- [ ] Go `?` to Cadence `Dictionary`
- [ ] Go `?` to Cadence `Event`

## TODO-List: Cadence to go
### Integers
- [x] Cadence `Int` to Go `*big.Int`
- [x] Cadence `Int8` to Go `int8`
- [x] Cadence `Int16` to Go `int16`
- [x] Cadence `Int32` to Go `int32`
- [x] Cadence `Int64` to Go `int64`
- [x] Cadence `Int128` to Go `*big.Int`
- [x] Cadence `Int256` to Go `*big.Int`
- [x] Cadence `UInt` to Go `*big.Int`
- [x] Cadence `UInt8` to Go `uint8`
- [x] Cadence `UInt16` to Go `uint16`
- [x] Cadence `UInt32` to Go `uint32`
- [x] Cadence `UInt64` to Go `uint64`
- [x] Cadence `UInt128` to Go `*big.Int`
- [x] Cadence `UInt256` to Go `*big.Int`
### Fixed-Point Numbers
- [x] Cadence `Fix64` to Go `int64`
- [x] Cadence `UFix64` to Go `uint64`
### Other
- [x] Cadence `String` to Go `string`
- [x] Cadence `Path` to Go `string`
- [x] Cadence `Address` to Go `string` or `cadence.Address` or `[8]uint8`
- [x] Cadence `Bool` to Go `bool`
- [x] Cadence `Array` to Go `slice`
- [x] Cadence `Struct` to Go `struct`
- [x] Cadence `Character` to Go `string`
- [ ] Cadence `Resource` to Go `?`
- [x] Cadence `Dictionary` to Go `map`
- [x] Cadence `Event` to Go `struct`