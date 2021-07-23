/* eslint-disable */
import * as Long from 'long';
import { util, configure, Writer, Reader } from 'protobufjs/minimal';
export const protobufPackage = 'tendermint.spn.launch';
const baseChainNameCount = { creator: '', chainName: '', count: 0 };
export const ChainNameCount = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== '') {
            writer.uint32(10).string(message.creator);
        }
        if (message.chainName !== '') {
            writer.uint32(18).string(message.chainName);
        }
        if (message.count !== 0) {
            writer.uint32(24).uint64(message.count);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseChainNameCount };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.chainName = reader.string();
                    break;
                case 3:
                    message.count = longToNumber(reader.uint64());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseChainNameCount };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = '';
        }
        if (object.chainName !== undefined && object.chainName !== null) {
            message.chainName = String(object.chainName);
        }
        else {
            message.chainName = '';
        }
        if (object.count !== undefined && object.count !== null) {
            message.count = Number(object.count);
        }
        else {
            message.count = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.chainName !== undefined && (obj.chainName = message.chainName);
        message.count !== undefined && (obj.count = message.count);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseChainNameCount };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = '';
        }
        if (object.chainName !== undefined && object.chainName !== null) {
            message.chainName = object.chainName;
        }
        else {
            message.chainName = '';
        }
        if (object.count !== undefined && object.count !== null) {
            message.count = object.count;
        }
        else {
            message.count = 0;
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
