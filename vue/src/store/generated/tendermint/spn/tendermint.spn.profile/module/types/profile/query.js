/* eslint-disable */
import { Reader, util, configure, Writer } from 'protobufjs/minimal';
import * as Long from 'long';
import { CoordinatorByAddress, Coordinator } from '../profile/coordinator';
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination';
export const protobufPackage = 'tendermint.spn.profile';
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
