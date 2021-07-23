/* eslint-disable */
import { Reader, util, configure, Writer } from 'protobufjs/minimal';
import * as Long from 'long';
import { ConsensusKeyNonce, ValidatorByConsAddress, ValidatorByAddress } from '../profile/validator';
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination';
import { CoordinatorByAddress, Coordinator } from '../profile/coordinator';
export const protobufPackage = 'tendermint.spn.profile';
const baseQueryGetConsensusKeyNonceRequest = { consAddress: '' };
export const QueryGetConsensusKeyNonceRequest = {
    encode(message, writer = Writer.create()) {
        if (message.consAddress !== '') {
            writer.uint32(10).string(message.consAddress);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetConsensusKeyNonceRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.consAddress = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetConsensusKeyNonceRequest };
        if (object.consAddress !== undefined && object.consAddress !== null) {
            message.consAddress = String(object.consAddress);
        }
        else {
            message.consAddress = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.consAddress !== undefined && (obj.consAddress = message.consAddress);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetConsensusKeyNonceRequest };
        if (object.consAddress !== undefined && object.consAddress !== null) {
            message.consAddress = object.consAddress;
        }
        else {
            message.consAddress = '';
        }
        return message;
    }
};
const baseQueryGetConsensusKeyNonceResponse = {};
export const QueryGetConsensusKeyNonceResponse = {
    encode(message, writer = Writer.create()) {
        if (message.consensusKeyNonce !== undefined) {
            ConsensusKeyNonce.encode(message.consensusKeyNonce, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetConsensusKeyNonceResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.consensusKeyNonce = ConsensusKeyNonce.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetConsensusKeyNonceResponse };
        if (object.consensusKeyNonce !== undefined && object.consensusKeyNonce !== null) {
            message.consensusKeyNonce = ConsensusKeyNonce.fromJSON(object.consensusKeyNonce);
        }
        else {
            message.consensusKeyNonce = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.consensusKeyNonce !== undefined &&
            (obj.consensusKeyNonce = message.consensusKeyNonce ? ConsensusKeyNonce.toJSON(message.consensusKeyNonce) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetConsensusKeyNonceResponse };
        if (object.consensusKeyNonce !== undefined && object.consensusKeyNonce !== null) {
            message.consensusKeyNonce = ConsensusKeyNonce.fromPartial(object.consensusKeyNonce);
        }
        else {
            message.consensusKeyNonce = undefined;
        }
        return message;
    }
};
const baseQueryGetValidatorByConsAddressRequest = { consAddress: '' };
export const QueryGetValidatorByConsAddressRequest = {
    encode(message, writer = Writer.create()) {
        if (message.consAddress !== '') {
            writer.uint32(10).string(message.consAddress);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetValidatorByConsAddressRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.consAddress = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetValidatorByConsAddressRequest };
        if (object.consAddress !== undefined && object.consAddress !== null) {
            message.consAddress = String(object.consAddress);
        }
        else {
            message.consAddress = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.consAddress !== undefined && (obj.consAddress = message.consAddress);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetValidatorByConsAddressRequest };
        if (object.consAddress !== undefined && object.consAddress !== null) {
            message.consAddress = object.consAddress;
        }
        else {
            message.consAddress = '';
        }
        return message;
    }
};
const baseQueryGetValidatorByConsAddressResponse = {};
export const QueryGetValidatorByConsAddressResponse = {
    encode(message, writer = Writer.create()) {
        if (message.validatorByConsAddress !== undefined) {
            ValidatorByConsAddress.encode(message.validatorByConsAddress, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetValidatorByConsAddressResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.validatorByConsAddress = ValidatorByConsAddress.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetValidatorByConsAddressResponse };
        if (object.validatorByConsAddress !== undefined && object.validatorByConsAddress !== null) {
            message.validatorByConsAddress = ValidatorByConsAddress.fromJSON(object.validatorByConsAddress);
        }
        else {
            message.validatorByConsAddress = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.validatorByConsAddress !== undefined &&
            (obj.validatorByConsAddress = message.validatorByConsAddress ? ValidatorByConsAddress.toJSON(message.validatorByConsAddress) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetValidatorByConsAddressResponse };
        if (object.validatorByConsAddress !== undefined && object.validatorByConsAddress !== null) {
            message.validatorByConsAddress = ValidatorByConsAddress.fromPartial(object.validatorByConsAddress);
        }
        else {
            message.validatorByConsAddress = undefined;
        }
        return message;
    }
};
const baseQueryGetValidatorByAddressRequest = { address: '' };
export const QueryGetValidatorByAddressRequest = {
    encode(message, writer = Writer.create()) {
        if (message.address !== '') {
            writer.uint32(10).string(message.address);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetValidatorByAddressRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
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
        const message = { ...baseQueryGetValidatorByAddressRequest };
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
        message.address !== undefined && (obj.address = message.address);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetValidatorByAddressRequest };
        if (object.address !== undefined && object.address !== null) {
            message.address = object.address;
        }
        else {
            message.address = '';
        }
        return message;
    }
};
const baseQueryGetValidatorByAddressResponse = {};
export const QueryGetValidatorByAddressResponse = {
    encode(message, writer = Writer.create()) {
        if (message.validatorByAddress !== undefined) {
            ValidatorByAddress.encode(message.validatorByAddress, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetValidatorByAddressResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.validatorByAddress = ValidatorByAddress.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetValidatorByAddressResponse };
        if (object.validatorByAddress !== undefined && object.validatorByAddress !== null) {
            message.validatorByAddress = ValidatorByAddress.fromJSON(object.validatorByAddress);
        }
        else {
            message.validatorByAddress = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.validatorByAddress !== undefined &&
            (obj.validatorByAddress = message.validatorByAddress ? ValidatorByAddress.toJSON(message.validatorByAddress) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetValidatorByAddressResponse };
        if (object.validatorByAddress !== undefined && object.validatorByAddress !== null) {
            message.validatorByAddress = ValidatorByAddress.fromPartial(object.validatorByAddress);
        }
        else {
            message.validatorByAddress = undefined;
        }
        return message;
    }
};
const baseQueryAllValidatorRequest = {};
export const QueryAllValidatorRequest = {
    encode(message, writer = Writer.create()) {
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllValidatorRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.pagination = PageRequest.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryAllValidatorRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllValidatorRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
const baseQueryAllValidatorResponse = {};
export const QueryAllValidatorResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.validatorByAddress) {
            ValidatorByAddress.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllValidatorResponse };
        message.validatorByAddress = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.validatorByAddress.push(ValidatorByAddress.decode(reader, reader.uint32()));
                    break;
                case 2:
                    message.pagination = PageResponse.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryAllValidatorResponse };
        message.validatorByAddress = [];
        if (object.validatorByAddress !== undefined && object.validatorByAddress !== null) {
            for (const e of object.validatorByAddress) {
                message.validatorByAddress.push(ValidatorByAddress.fromJSON(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.validatorByAddress) {
            obj.validatorByAddress = message.validatorByAddress.map((e) => (e ? ValidatorByAddress.toJSON(e) : undefined));
        }
        else {
            obj.validatorByAddress = [];
        }
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllValidatorResponse };
        message.validatorByAddress = [];
        if (object.validatorByAddress !== undefined && object.validatorByAddress !== null) {
            for (const e of object.validatorByAddress) {
                message.validatorByAddress.push(ValidatorByAddress.fromPartial(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
const baseQueryGetCoordinatorByAddressRequest = { address: '' };
export const QueryGetCoordinatorByAddressRequest = {
    encode(message, writer = Writer.create()) {
        if (message.address !== '') {
            writer.uint32(10).string(message.address);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetCoordinatorByAddressRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
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
        const message = { ...baseQueryGetCoordinatorByAddressRequest };
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
        message.address !== undefined && (obj.address = message.address);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetCoordinatorByAddressRequest };
        if (object.address !== undefined && object.address !== null) {
            message.address = object.address;
        }
        else {
            message.address = '';
        }
        return message;
    }
};
const baseQueryGetCoordinatorByAddressResponse = {};
export const QueryGetCoordinatorByAddressResponse = {
    encode(message, writer = Writer.create()) {
        if (message.coordinatorByAddress !== undefined) {
            CoordinatorByAddress.encode(message.coordinatorByAddress, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetCoordinatorByAddressResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.coordinatorByAddress = CoordinatorByAddress.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetCoordinatorByAddressResponse };
        if (object.coordinatorByAddress !== undefined && object.coordinatorByAddress !== null) {
            message.coordinatorByAddress = CoordinatorByAddress.fromJSON(object.coordinatorByAddress);
        }
        else {
            message.coordinatorByAddress = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.coordinatorByAddress !== undefined &&
            (obj.coordinatorByAddress = message.coordinatorByAddress ? CoordinatorByAddress.toJSON(message.coordinatorByAddress) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetCoordinatorByAddressResponse };
        if (object.coordinatorByAddress !== undefined && object.coordinatorByAddress !== null) {
            message.coordinatorByAddress = CoordinatorByAddress.fromPartial(object.coordinatorByAddress);
        }
        else {
            message.coordinatorByAddress = undefined;
        }
        return message;
    }
};
const baseQueryAllCoordinatorByAddressRequest = {};
export const QueryAllCoordinatorByAddressRequest = {
    encode(message, writer = Writer.create()) {
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllCoordinatorByAddressRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.pagination = PageRequest.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryAllCoordinatorByAddressRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllCoordinatorByAddressRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
const baseQueryAllCoordinatorByAddressResponse = {};
export const QueryAllCoordinatorByAddressResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.coordinatorByAddress) {
            CoordinatorByAddress.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllCoordinatorByAddressResponse };
        message.coordinatorByAddress = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.coordinatorByAddress.push(CoordinatorByAddress.decode(reader, reader.uint32()));
                    break;
                case 2:
                    message.pagination = PageResponse.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryAllCoordinatorByAddressResponse };
        message.coordinatorByAddress = [];
        if (object.coordinatorByAddress !== undefined && object.coordinatorByAddress !== null) {
            for (const e of object.coordinatorByAddress) {
                message.coordinatorByAddress.push(CoordinatorByAddress.fromJSON(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.coordinatorByAddress) {
            obj.coordinatorByAddress = message.coordinatorByAddress.map((e) => (e ? CoordinatorByAddress.toJSON(e) : undefined));
        }
        else {
            obj.coordinatorByAddress = [];
        }
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllCoordinatorByAddressResponse };
        message.coordinatorByAddress = [];
        if (object.coordinatorByAddress !== undefined && object.coordinatorByAddress !== null) {
            for (const e of object.coordinatorByAddress) {
                message.coordinatorByAddress.push(CoordinatorByAddress.fromPartial(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
const baseQueryGetCoordinatorRequest = { id: 0 };
export const QueryGetCoordinatorRequest = {
    encode(message, writer = Writer.create()) {
        if (message.id !== 0) {
            writer.uint32(8).uint64(message.id);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetCoordinatorRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.id = longToNumber(reader.uint64());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetCoordinatorRequest };
        if (object.id !== undefined && object.id !== null) {
            message.id = Number(object.id);
        }
        else {
            message.id = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.id !== undefined && (obj.id = message.id);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetCoordinatorRequest };
        if (object.id !== undefined && object.id !== null) {
            message.id = object.id;
        }
        else {
            message.id = 0;
        }
        return message;
    }
};
const baseQueryGetCoordinatorResponse = {};
export const QueryGetCoordinatorResponse = {
    encode(message, writer = Writer.create()) {
        if (message.Coordinator !== undefined) {
            Coordinator.encode(message.Coordinator, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetCoordinatorResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.Coordinator = Coordinator.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetCoordinatorResponse };
        if (object.Coordinator !== undefined && object.Coordinator !== null) {
            message.Coordinator = Coordinator.fromJSON(object.Coordinator);
        }
        else {
            message.Coordinator = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.Coordinator !== undefined && (obj.Coordinator = message.Coordinator ? Coordinator.toJSON(message.Coordinator) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetCoordinatorResponse };
        if (object.Coordinator !== undefined && object.Coordinator !== null) {
            message.Coordinator = Coordinator.fromPartial(object.Coordinator);
        }
        else {
            message.Coordinator = undefined;
        }
        return message;
    }
};
const baseQueryAllCoordinatorRequest = {};
export const QueryAllCoordinatorRequest = {
    encode(message, writer = Writer.create()) {
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllCoordinatorRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.pagination = PageRequest.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryAllCoordinatorRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllCoordinatorRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
const baseQueryAllCoordinatorResponse = {};
export const QueryAllCoordinatorResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.Coordinator) {
            Coordinator.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllCoordinatorResponse };
        message.Coordinator = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.Coordinator.push(Coordinator.decode(reader, reader.uint32()));
                    break;
                case 2:
                    message.pagination = PageResponse.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryAllCoordinatorResponse };
        message.Coordinator = [];
        if (object.Coordinator !== undefined && object.Coordinator !== null) {
            for (const e of object.Coordinator) {
                message.Coordinator.push(Coordinator.fromJSON(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.Coordinator) {
            obj.Coordinator = message.Coordinator.map((e) => (e ? Coordinator.toJSON(e) : undefined));
        }
        else {
            obj.Coordinator = [];
        }
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllCoordinatorResponse };
        message.Coordinator = [];
        if (object.Coordinator !== undefined && object.Coordinator !== null) {
            for (const e of object.Coordinator) {
                message.Coordinator.push(Coordinator.fromPartial(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
export class QueryClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
    }
    ConsensusKeyNonce(request) {
        const data = QueryGetConsensusKeyNonceRequest.encode(request).finish();
        const promise = this.rpc.request('tendermint.spn.profile.Query', 'ConsensusKeyNonce', data);
        return promise.then((data) => QueryGetConsensusKeyNonceResponse.decode(new Reader(data)));
    }
    ValidatorByConsAddress(request) {
        const data = QueryGetValidatorByConsAddressRequest.encode(request).finish();
        const promise = this.rpc.request('tendermint.spn.profile.Query', 'ValidatorByConsAddress', data);
        return promise.then((data) => QueryGetValidatorByConsAddressResponse.decode(new Reader(data)));
    }
    ValidatorByAddress(request) {
        const data = QueryGetValidatorByAddressRequest.encode(request).finish();
        const promise = this.rpc.request('tendermint.spn.profile.Query', 'ValidatorByAddress', data);
        return promise.then((data) => QueryGetValidatorByAddressResponse.decode(new Reader(data)));
    }
    ValidatorByAddressAll(request) {
        const data = QueryAllValidatorRequest.encode(request).finish();
        const promise = this.rpc.request('tendermint.spn.profile.Query', 'ValidatorByAddressAll', data);
        return promise.then((data) => QueryAllValidatorResponse.decode(new Reader(data)));
    }
    CoordinatorByAddress(request) {
        const data = QueryGetCoordinatorByAddressRequest.encode(request).finish();
        const promise = this.rpc.request('tendermint.spn.profile.Query', 'CoordinatorByAddress', data);
        return promise.then((data) => QueryGetCoordinatorByAddressResponse.decode(new Reader(data)));
    }
    Coordinator(request) {
        const data = QueryGetCoordinatorRequest.encode(request).finish();
        const promise = this.rpc.request('tendermint.spn.profile.Query', 'Coordinator', data);
        return promise.then((data) => QueryGetCoordinatorResponse.decode(new Reader(data)));
    }
    CoordinatorAll(request) {
        const data = QueryAllCoordinatorRequest.encode(request).finish();
        const promise = this.rpc.request('tendermint.spn.profile.Query', 'CoordinatorAll', data);
        return promise.then((data) => QueryAllCoordinatorResponse.decode(new Reader(data)));
    }
}
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
