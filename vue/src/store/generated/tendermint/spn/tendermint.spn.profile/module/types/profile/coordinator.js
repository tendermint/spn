/* eslint-disable */
import * as Long from 'long';
import { util, configure, Writer, Reader } from 'protobufjs/minimal';
export const protobufPackage = 'tendermint.spn.profile';
const baseCoordinator = { coordinatorId: 0, address: '' };
export const Coordinator = {
    encode(message, writer = Writer.create()) {
        if (message.coordinatorId !== 0) {
            writer.uint32(8).uint64(message.coordinatorId);
        }
        if (message.address !== '') {
            writer.uint32(18).string(message.address);
        }
        if (message.description !== undefined) {
            CoordinatorDescription.encode(message.description, writer.uint32(26).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseCoordinator };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.coordinatorId = longToNumber(reader.uint64());
                    break;
                case 2:
                    message.address = reader.string();
                    break;
                case 3:
                    message.description = CoordinatorDescription.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseCoordinator };
        if (object.coordinatorId !== undefined && object.coordinatorId !== null) {
            message.coordinatorId = Number(object.coordinatorId);
        }
        else {
            message.coordinatorId = 0;
        }
        if (object.address !== undefined && object.address !== null) {
            message.address = String(object.address);
        }
        else {
            message.address = '';
        }
        if (object.description !== undefined && object.description !== null) {
            message.description = CoordinatorDescription.fromJSON(object.description);
        }
        else {
            message.description = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.coordinatorId !== undefined && (obj.coordinatorId = message.coordinatorId);
        message.address !== undefined && (obj.address = message.address);
        message.description !== undefined && (obj.description = message.description ? CoordinatorDescription.toJSON(message.description) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseCoordinator };
        if (object.coordinatorId !== undefined && object.coordinatorId !== null) {
            message.coordinatorId = object.coordinatorId;
        }
        else {
            message.coordinatorId = 0;
        }
        if (object.address !== undefined && object.address !== null) {
            message.address = object.address;
        }
        else {
            message.address = '';
        }
        if (object.description !== undefined && object.description !== null) {
            message.description = CoordinatorDescription.fromPartial(object.description);
        }
        else {
            message.description = undefined;
        }
        return message;
    }
};
const baseCoordinatorDescription = { identity: '', website: '', details: '' };
export const CoordinatorDescription = {
    encode(message, writer = Writer.create()) {
        if (message.identity !== '') {
            writer.uint32(10).string(message.identity);
        }
        if (message.website !== '') {
            writer.uint32(18).string(message.website);
        }
        if (message.details !== '') {
            writer.uint32(26).string(message.details);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseCoordinatorDescription };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.identity = reader.string();
                    break;
                case 2:
                    message.website = reader.string();
                    break;
                case 3:
                    message.details = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseCoordinatorDescription };
        if (object.identity !== undefined && object.identity !== null) {
            message.identity = String(object.identity);
        }
        else {
            message.identity = '';
        }
        if (object.website !== undefined && object.website !== null) {
            message.website = String(object.website);
        }
        else {
            message.website = '';
        }
        if (object.details !== undefined && object.details !== null) {
            message.details = String(object.details);
        }
        else {
            message.details = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.identity !== undefined && (obj.identity = message.identity);
        message.website !== undefined && (obj.website = message.website);
        message.details !== undefined && (obj.details = message.details);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseCoordinatorDescription };
        if (object.identity !== undefined && object.identity !== null) {
            message.identity = object.identity;
        }
        else {
            message.identity = '';
        }
        if (object.website !== undefined && object.website !== null) {
            message.website = object.website;
        }
        else {
            message.website = '';
        }
        if (object.details !== undefined && object.details !== null) {
            message.details = object.details;
        }
        else {
            message.details = '';
        }
        return message;
    }
};
const baseCoordinatorByAddress = { address: '', coordinatorId: 0 };
export const CoordinatorByAddress = {
    encode(message, writer = Writer.create()) {
        if (message.address !== '') {
            writer.uint32(10).string(message.address);
        }
        if (message.coordinatorId !== 0) {
            writer.uint32(16).uint64(message.coordinatorId);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseCoordinatorByAddress };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.address = reader.string();
                    break;
                case 2:
                    message.coordinatorId = longToNumber(reader.uint64());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseCoordinatorByAddress };
        if (object.address !== undefined && object.address !== null) {
            message.address = String(object.address);
        }
        else {
            message.address = '';
        }
        if (object.coordinatorId !== undefined && object.coordinatorId !== null) {
            message.coordinatorId = Number(object.coordinatorId);
        }
        else {
            message.coordinatorId = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.address !== undefined && (obj.address = message.address);
        message.coordinatorId !== undefined && (obj.coordinatorId = message.coordinatorId);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseCoordinatorByAddress };
        if (object.address !== undefined && object.address !== null) {
            message.address = object.address;
        }
        else {
            message.address = '';
        }
        if (object.coordinatorId !== undefined && object.coordinatorId !== null) {
            message.coordinatorId = object.coordinatorId;
        }
        else {
            message.coordinatorId = 0;
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
