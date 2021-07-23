/* eslint-disable */
import * as Long from 'long';
import { util, configure, Writer, Reader } from 'protobufjs/minimal';
import { Any } from '../google/protobuf/any';
export const protobufPackage = 'tendermint.spn.launch';
const baseChain = { chainID: '', coordinatorID: 0, createdAt: 0, sourceURL: '', sourceHash: '', launchTriggered: false, launchTimestamp: 0 };
export const Chain = {
    encode(message, writer = Writer.create()) {
        if (message.chainID !== '') {
            writer.uint32(10).string(message.chainID);
        }
        if (message.coordinatorID !== 0) {
            writer.uint32(16).uint64(message.coordinatorID);
        }
        if (message.createdAt !== 0) {
            writer.uint32(24).int64(message.createdAt);
        }
        if (message.sourceURL !== '') {
            writer.uint32(34).string(message.sourceURL);
        }
        if (message.sourceHash !== '') {
            writer.uint32(42).string(message.sourceHash);
        }
        if (message.initialGenesis !== undefined) {
            Any.encode(message.initialGenesis, writer.uint32(50).fork()).ldelim();
        }
        if (message.launchTriggered === true) {
            writer.uint32(56).bool(message.launchTriggered);
        }
        if (message.launchTimestamp !== 0) {
            writer.uint32(64).int64(message.launchTimestamp);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseChain };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.chainID = reader.string();
                    break;
                case 2:
                    message.coordinatorID = longToNumber(reader.uint64());
                    break;
                case 3:
                    message.createdAt = longToNumber(reader.int64());
                    break;
                case 4:
                    message.sourceURL = reader.string();
                    break;
                case 5:
                    message.sourceHash = reader.string();
                    break;
                case 6:
                    message.initialGenesis = Any.decode(reader, reader.uint32());
                    break;
                case 7:
                    message.launchTriggered = reader.bool();
                    break;
                case 8:
                    message.launchTimestamp = longToNumber(reader.int64());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseChain };
        if (object.chainID !== undefined && object.chainID !== null) {
            message.chainID = String(object.chainID);
        }
        else {
            message.chainID = '';
        }
        if (object.coordinatorID !== undefined && object.coordinatorID !== null) {
            message.coordinatorID = Number(object.coordinatorID);
        }
        else {
            message.coordinatorID = 0;
        }
        if (object.createdAt !== undefined && object.createdAt !== null) {
            message.createdAt = Number(object.createdAt);
        }
        else {
            message.createdAt = 0;
        }
        if (object.sourceURL !== undefined && object.sourceURL !== null) {
            message.sourceURL = String(object.sourceURL);
        }
        else {
            message.sourceURL = '';
        }
        if (object.sourceHash !== undefined && object.sourceHash !== null) {
            message.sourceHash = String(object.sourceHash);
        }
        else {
            message.sourceHash = '';
        }
        if (object.initialGenesis !== undefined && object.initialGenesis !== null) {
            message.initialGenesis = Any.fromJSON(object.initialGenesis);
        }
        else {
            message.initialGenesis = undefined;
        }
        if (object.launchTriggered !== undefined && object.launchTriggered !== null) {
            message.launchTriggered = Boolean(object.launchTriggered);
        }
        else {
            message.launchTriggered = false;
        }
        if (object.launchTimestamp !== undefined && object.launchTimestamp !== null) {
            message.launchTimestamp = Number(object.launchTimestamp);
        }
        else {
            message.launchTimestamp = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.chainID !== undefined && (obj.chainID = message.chainID);
        message.coordinatorID !== undefined && (obj.coordinatorID = message.coordinatorID);
        message.createdAt !== undefined && (obj.createdAt = message.createdAt);
        message.sourceURL !== undefined && (obj.sourceURL = message.sourceURL);
        message.sourceHash !== undefined && (obj.sourceHash = message.sourceHash);
        message.initialGenesis !== undefined && (obj.initialGenesis = message.initialGenesis ? Any.toJSON(message.initialGenesis) : undefined);
        message.launchTriggered !== undefined && (obj.launchTriggered = message.launchTriggered);
        message.launchTimestamp !== undefined && (obj.launchTimestamp = message.launchTimestamp);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseChain };
        if (object.chainID !== undefined && object.chainID !== null) {
            message.chainID = object.chainID;
        }
        else {
            message.chainID = '';
        }
        if (object.coordinatorID !== undefined && object.coordinatorID !== null) {
            message.coordinatorID = object.coordinatorID;
        }
        else {
            message.coordinatorID = 0;
        }
        if (object.createdAt !== undefined && object.createdAt !== null) {
            message.createdAt = object.createdAt;
        }
        else {
            message.createdAt = 0;
        }
        if (object.sourceURL !== undefined && object.sourceURL !== null) {
            message.sourceURL = object.sourceURL;
        }
        else {
            message.sourceURL = '';
        }
        if (object.sourceHash !== undefined && object.sourceHash !== null) {
            message.sourceHash = object.sourceHash;
        }
        else {
            message.sourceHash = '';
        }
        if (object.initialGenesis !== undefined && object.initialGenesis !== null) {
            message.initialGenesis = Any.fromPartial(object.initialGenesis);
        }
        else {
            message.initialGenesis = undefined;
        }
        if (object.launchTriggered !== undefined && object.launchTriggered !== null) {
            message.launchTriggered = object.launchTriggered;
        }
        else {
            message.launchTriggered = false;
        }
        if (object.launchTimestamp !== undefined && object.launchTimestamp !== null) {
            message.launchTimestamp = object.launchTimestamp;
        }
        else {
            message.launchTimestamp = 0;
        }
        return message;
    }
};
const baseDefaultInitialGenesis = {};
export const DefaultInitialGenesis = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseDefaultInitialGenesis };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(_) {
        const message = { ...baseDefaultInitialGenesis };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseDefaultInitialGenesis };
        return message;
    }
};
const baseGenesisURL = { url: '', hash: '' };
export const GenesisURL = {
    encode(message, writer = Writer.create()) {
        if (message.url !== '') {
            writer.uint32(10).string(message.url);
        }
        if (message.hash !== '') {
            writer.uint32(18).string(message.hash);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseGenesisURL };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.url = reader.string();
                    break;
                case 2:
                    message.hash = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseGenesisURL };
        if (object.url !== undefined && object.url !== null) {
            message.url = String(object.url);
        }
        else {
            message.url = '';
        }
        if (object.hash !== undefined && object.hash !== null) {
            message.hash = String(object.hash);
        }
        else {
            message.hash = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.url !== undefined && (obj.url = message.url);
        message.hash !== undefined && (obj.hash = message.hash);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseGenesisURL };
        if (object.url !== undefined && object.url !== null) {
            message.url = object.url;
        }
        else {
            message.url = '';
        }
        if (object.hash !== undefined && object.hash !== null) {
            message.hash = object.hash;
        }
        else {
            message.hash = '';
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
