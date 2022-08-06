pub contract ForTest {
    pub struct ForEmbedded {
        pub var myName: String

        init() {
            self.myName = "LemonNeko"
        }
    }

    pub resource ForEmbeddedR {
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

    // a simple resource
    pub resource SimpleR {
        pub var MyName: String

        init() {
            self.MyName = "LemonNeko"
        }
    }

    // a simple resource with name tag.
    // a simple resource with wrong name tag.
    pub resource SimpleR2 {
        pub var myName: String

        init() {
            self.myName = "LemonNeko"
        }
    }

    // a simple resource with wrong type.
    pub resource SimpleR3 {
        pub var myName: UInt8

        init() {
            self.myName = 127
        }
    }

    // a resource contains many type
    pub resource ResourceContainsManyType {
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

    pub resource AStructEmbeddedInAResource {
        pub var forEmbedded: ForEmbedded

        init() {
            self.forEmbedded = ForEmbedded()
        }
    }

    pub resource AResourceEmbeddedInAResource {
        pub var forEmbedded: @ForEmbeddedR

        init() {
            self.forEmbedded <- create ForEmbeddedR()
        }

        destroy() {
            destroy self.forEmbedded
        }
    }

    init() {
        self.account.save(<- create SimpleR(), to: /storage/simpleR)
        self.account.link<&SimpleR>(/public/simpleR, target: /storage/simpleR)

        self.account.save(<- create SimpleR2(), to: /storage/simpleR2)
        self.account.link<&SimpleR2>(/public/simpleR2, target: /storage/simpleR2)

        self.account.save(<- create SimpleR3(), to: /storage/simpleR3)
        self.account.link<&SimpleR3>(/public/simpleR3, target: /storage/simpleR3)

        self.account.save(<- create ResourceContainsManyType(), to: /storage/rMany)
        self.account.link<&ResourceContainsManyType>(/public/rMany, target: /storage/rMany)
        
        self.account.save(<- create AStructEmbeddedInAResource(), to: /storage/embeddedStruct)
        self.account.link<&AStructEmbeddedInAResource>(/public/embeddedStruct, target: /storage/embeddedStruct)

        self.account.save(<- create AResourceEmbeddedInAResource(), to: /storage/embeddedResource)
        self.account.link<&AResourceEmbeddedInAResource>(/public/embeddedResource, target: /storage/embeddedResource)
    }
}