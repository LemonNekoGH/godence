pub contract ForTest {
    pub struct ForEmbedded {
        pub var myName: String

        init() {
            self.myName = "LemonNeko"
        }
    }
    // a simple event
    pub event Simple(_ MyName: String)
    // a simple event, with field name tag
    // a simple event, with wrong field name tag
    pub event Simple2(_ myName: String)
    // a simple event, with wrong type
    pub event Simple3(_ myName: UInt8)
    // a event contains many type
    pub event ManyType(
        intValue: Int,
	    int8Value: Int8,
	    int16Value: Int16,
	    int32Value: Int32,
	    int64Value: Int64,
	    int128Value: Int128,
	    int256Value: Int256,
	    uintValue: UInt,
	    uint8Value: UInt8,
	    uint16Value: UInt16,
	    uint32Value: UInt32,
	    uint64Value: UInt64,
	    uint128Value: UInt128,
	    uint256Value: UInt256,
	    stringValue: String,
	    addressValue: Address,
	    boolValue: Bool,
    )
    // a struct in a event param
    pub event StructInParam(p: ForEmbedded)

    // transaction cannot emit event in import type
    pub fun emitSimple() {
        emit Simple("LemonNeko")
    }

    pub fun emitSimple2() {
        emit Simple2("LemonNeko")
    }

    pub fun emitSimple3() {
        emit Simple3(64)
    }

    pub fun emitManyType() {
        emit ManyType(
            intValue: 127,
            int8Value: 127,
            int16Value: 127,
            int32Value: 127,
            int64Value: 127,
            int128Value: 127,
            int256Value: 127,
            uintValue: 127,
            uint8Value: 127,
            uint16Value: 127,
            uint32Value: 127,
            uint64Value: 127,
            uint128Value: 127,
            uint256Value: 127,
            stringValue: "LemonNeko",
	        addressValue: 0x0,
	        boolValue: true,
        )
    }

    pub fun emitStructInParam() {
        emit StructInParam(p: ForEmbedded())
    }
}