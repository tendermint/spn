/* eslint-disable */
import * as Long from 'long';
import { util, configure, Writer, Reader } from 'protobufjs/minimal';
import { CoordinatorByAddress, Coordinator } from '../account/coordinator';
export const protobufPackage = 'tendermint.spn.account';
const baseGenesisState = { coordinatorCount: 0 };
export const GenesisState = {
    encode(message, writer = Writer.create()) {
        for (const v of message.coordinatorByAddressList) {
            CoordinatorByAddress.encode(v, writer.uint32(26).fork()).ldelim();
        }
        for (const v of message.coordinatorList) {
            Coordinator.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.coordinatorCount !== 0) {
            writer.uint32(16).uint64(message.coordinatorCount);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseGenesisState };
        message.coordinatorByAddressList = [];
        message.coordinatorList = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 3:
                    message.coordinatorByAddressList.push(CoordinatorByAddress.decode(reader, reader.uint32()));
                    break;
                case 1:
                    message.coordinatorList.push(Coordinator.decode(reader, reader.uint32()));
                    break;
                case 2:
                    message.coordinatorCount = longToNumber(reader.uint64());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseGenesisState };
        message.coordinatorByAddressList = [];
        message.coordinatorList = [];
        if (object.coordinatorByAddressList !== undefined && object.coordinatorByAddressList !== null) {
            for (const e of object.coordinatorByAddressList) {
                message.coordinatorByAddressList.push(CoordinatorByAddress.fromJSON(e));
            }
        }
        if (object.coordinatorList !== undefined && object.coordinatorList !== null) {
            for (const e of object.coordinatorList) {
                message.coordinatorList.push(Coordinator.fromJSON(e));
            }
        }
        if (object.coordinatorCount !== undefined && object.coordinatorCount !== null) {
            message.coordinatorCount = Number(object.coordinatorCount);
        }
        else {
            message.coordinatorCount = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.coordinatorByAddressList) {
            obj.coordinatorByAddressList = message.coordinatorByAddressList.map((e) => (e ? CoordinatorByAddress.toJSON(e) : undefined));
        }
        else {
            obj.coordinatorByAddressList = [];
        }
        if (message.coordinatorList) {
            obj.coordinatorList = message.coordinatorList.map((e) => (e ? Coordinator.toJSON(e) : undefined));
        }
        else {
            obj.coordinatorList = [];
        }
        message.coordinatorCount !== undefined && (obj.coordinatorCount = message.coordinatorCount);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseGenesisState };
        message.coordinatorByAddressList = [];
        message.coordinatorList = [];
        if (object.coordinatorByAddressList !== undefined && object.coordinatorByAddressList !== null) {
            for (const e of object.coordinatorByAddressList) {
                message.coordinatorByAddressList.push(CoordinatorByAddress.fromPartial(e));
            }
        }
        if (object.coordinatorList !== undefined && object.coordinatorList !== null) {
            for (const e of object.coordinatorList) {
                message.coordinatorList.push(Coordinator.fromPartial(e));
            }
        }
        if (object.coordinatorCount !== undefined && object.coordinatorCount !== null) {
            message.coordinatorCount = object.coordinatorCount;
        }
        else {
            message.coordinatorCount = 0;
        }
        return message;
    }
};
var globalThis = (() => {
    if (typeof globalThis !== 'undefined')
        return globalThis;
    if (typeof self !== 'undefined')
        return self;
    if (typeof window !== 'undefined')
        return window;
    if (typeof global !== 'undefined')
        return global;
    throw 'Unable to locate global object';
})();
function longToNumber(long) {
    if (long.gt(Number.MAX_SAFE_INTEGER)) {
        throw new globalThis.Error('Value is larger than Number.MAX_SAFE_INTEGER');
    }
    return long.toNumber();
}
if (util.Long !== Long) {
    util.Long = Long;
    configure();
}
