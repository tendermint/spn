/* eslint-disable */
import * as Long from 'long';
import { util, configure, Writer, Reader } from 'protobufjs/minimal';
export const protobufPackage = 'tendermint.spn.profile';
const baseValidatorByAddress = { address: '', consensusAddress: '' };
export const ValidatorByAddress = {
    encode(message, writer = Writer.create()) {
        if (message.address !== '') {
            writer.uint32(10).string(message.address);
        }
        if (message.consensusAddress !== '') {
            writer.uint32(18).string(message.consensusAddress);
        }
        if (message.description !== undefined) {
            ValidatorDescription.encode(message.description, writer.uint32(26).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseValidatorByAddress };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.address = reader.string();
                    break;
                case 2:
                    message.consensusAddress = reader.string();
                    break;
                case 3:
                    message.description = ValidatorDescription.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseValidatorByAddress };
        if (object.address !== undefined && object.address !== null) {
            message.address = String(object.address);
        }
        else {
            message.address = '';
        }
        if (object.consensusAddress !== undefined && object.consensusAddress !== null) {
            message.consensusAddress = String(object.consensusAddress);
        }
        else {
            message.consensusAddress = '';
        }
        if (object.description !== undefined && object.description !== null) {
            message.description = ValidatorDescription.fromJSON(object.description);
        }
        else {
            message.description = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.address !== undefined && (obj.address = message.address);
        message.consensusAddress !== undefined && (obj.consensusAddress = message.consensusAddress);
        message.description !== undefined && (obj.description = message.description ? ValidatorDescription.toJSON(message.description) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseValidatorByAddress };
        if (object.address !== undefined && object.address !== null) {
            message.address = object.address;
        }
        else {
            message.address = '';
        }
        if (object.consensusAddress !== undefined && object.consensusAddress !== null) {
            message.consensusAddress = object.consensusAddress;
        }
        else {
            message.consensusAddress = '';
        }
        if (object.description !== undefined && object.description !== null) {
            message.description = ValidatorDescription.fromPartial(object.description);
        }
        else {
            message.description = undefined;
        }
        return message;
    }
};
const baseValidatorDescription = { identity: '', moniker: '', website: '', securityContact: '', details: '' };
export const ValidatorDescription = {
    encode(message, writer = Writer.create()) {
        if (message.identity !== '') {
            writer.uint32(10).string(message.identity);
        }
        if (message.moniker !== '') {
            writer.uint32(18).string(message.moniker);
        }
        if (message.website !== '') {
            writer.uint32(26).string(message.website);
        }
        if (message.securityContact !== '') {
            writer.uint32(34).string(message.securityContact);
        }
        if (message.details !== '') {
            writer.uint32(42).string(message.details);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseValidatorDescription };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.identity = reader.string();
                    break;
                case 2:
                    message.moniker = reader.string();
                    break;
                case 3:
                    message.website = reader.string();
                    break;
                case 4:
                    message.securityContact = reader.string();
                    break;
                case 5:
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
        const message = { ...baseValidatorDescription };
        if (object.identity !== undefined && object.identity !== null) {
            message.identity = String(object.identity);
        }
        else {
            message.identity = '';
        }
        if (object.moniker !== undefined && object.moniker !== null) {
            message.moniker = String(object.moniker);
        }
        else {
            message.moniker = '';
        }
        if (object.website !== undefined && object.website !== null) {
            message.website = String(object.website);
        }
        else {
            message.website = '';
        }
        if (object.securityContact !== undefined && object.securityContact !== null) {
            message.securityContact = String(object.securityContact);
        }
        else {
            message.securityContact = '';
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
        message.moniker !== undefined && (obj.moniker = message.moniker);
        message.website !== undefined && (obj.website = message.website);
        message.securityContact !== undefined && (obj.securityContact = message.securityContact);
        message.details !== undefined && (obj.details = message.details);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseValidatorDescription };
        if (object.identity !== undefined && object.identity !== null) {
            message.identity = object.identity;
        }
        else {
            message.identity = '';
        }
        if (object.moniker !== undefined && object.moniker !== null) {
            message.moniker = object.moniker;
        }
        else {
            message.moniker = '';
        }
        if (object.website !== undefined && object.website !== null) {
            message.website = object.website;
        }
        else {
            message.website = '';
        }
        if (object.securityContact !== undefined && object.securityContact !== null) {
            message.securityContact = object.securityContact;
        }
        else {
            message.securityContact = '';
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
const baseValidatorByConsAddress = { consAddress: '', address: '' };
export const ValidatorByConsAddress = {
    encode(message, writer = Writer.create()) {
        if (message.consAddress !== '') {
            writer.uint32(10).string(message.consAddress);
        }
        if (message.address !== '') {
            writer.uint32(18).string(message.address);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseValidatorByConsAddress };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.consAddress = reader.string();
                    break;
                case 2:
                    message.address = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseValidatorByConsAddress };
        if (object.consAddress !== undefined && object.consAddress !== null) {
            message.consAddress = String(object.consAddress);
        }
        else {
            message.consAddress = '';
        }
        if (object.address !== undefined && object.address !== null) {
            message.address = String(object.address);
        }
        else {
            message.address = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.consAddress !== undefined && (obj.consAddress = message.consAddress);
        message.address !== undefined && (obj.address = message.address);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseValidatorByConsAddress };
        if (object.consAddress !== undefined && object.consAddress !== null) {
            message.consAddress = object.consAddress;
        }
        else {
            message.consAddress = '';
        }
        if (object.address !== undefined && object.address !== null) {
            message.address = object.address;
        }
        else {
            message.address = '';
        }
        return message;
    }
};
const baseConsensusKeyNonce = { consAddress: '', nonce: 0 };
export const ConsensusKeyNonce = {
    encode(message, writer = Writer.create()) {
        if (message.consAddress !== '') {
            writer.uint32(10).string(message.consAddress);
        }
        if (message.nonce !== 0) {
            writer.uint32(16).uint64(message.nonce);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseConsensusKeyNonce };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.consAddress = reader.string();
                    break;
                case 2:
                    message.nonce = longToNumber(reader.uint64());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseConsensusKeyNonce };
        if (object.consAddress !== undefined && object.consAddress !== null) {
            message.consAddress = String(object.consAddress);
        }
        else {
            message.consAddress = '';
        }
        if (object.nonce !== undefined && object.nonce !== null) {
            message.nonce = Number(object.nonce);
        }
        else {
            message.nonce = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.consAddress !== undefined && (obj.consAddress = message.consAddress);
        message.nonce !== undefined && (obj.nonce = message.nonce);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseConsensusKeyNonce };
        if (object.consAddress !== undefined && object.consAddress !== null) {
            message.consAddress = object.consAddress;
        }
        else {
            message.consAddress = '';
        }
        if (object.nonce !== undefined && object.nonce !== null) {
            message.nonce = object.nonce;
        }
        else {
            message.nonce = 0;
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
