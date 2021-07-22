/* eslint-disable */
import { Reader, util, configure, Writer } from 'protobufjs/minimal';
import * as Long from 'long';
export const protobufPackage = 'tendermint.spn.profile';
const baseMsgUpdateCoordinatorAddress = { address: '', newAddress: '' };
export const MsgUpdateCoordinatorAddress = {
    encode(message, writer = Writer.create()) {
        if (message.address !== '') {
            writer.uint32(10).string(message.address);
        }
        if (message.newAddress !== '') {
            writer.uint32(18).string(message.newAddress);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgUpdateCoordinatorAddress };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.address = reader.string();
                    break;
                case 2:
                    message.newAddress = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgUpdateCoordinatorAddress };
        if (object.address !== undefined && object.address !== null) {
            message.address = String(object.address);
        }
        else {
            message.address = '';
        }
        if (object.newAddress !== undefined && object.newAddress !== null) {
            message.newAddress = String(object.newAddress);
        }
        else {
            message.newAddress = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.address !== undefined && (obj.address = message.address);
        message.newAddress !== undefined && (obj.newAddress = message.newAddress);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgUpdateCoordinatorAddress };
        if (object.address !== undefined && object.address !== null) {
            message.address = object.address;
        }
        else {
            message.address = '';
        }
        if (object.newAddress !== undefined && object.newAddress !== null) {
            message.newAddress = object.newAddress;
        }
        else {
            message.newAddress = '';
        }
        return message;
    }
};
const baseMsgUpdateCoordinatorAddressResponse = { coordinatorId: 0 };
export const MsgUpdateCoordinatorAddressResponse = {
    encode(message, writer = Writer.create()) {
        if (message.coordinatorId !== 0) {
            writer.uint32(8).uint64(message.coordinatorId);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgUpdateCoordinatorAddressResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
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
        const message = { ...baseMsgUpdateCoordinatorAddressResponse };
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
        message.coordinatorId !== undefined && (obj.coordinatorId = message.coordinatorId);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgUpdateCoordinatorAddressResponse };
        if (object.coordinatorId !== undefined && object.coordinatorId !== null) {
            message.coordinatorId = object.coordinatorId;
        }
        else {
            message.coordinatorId = 0;
        }
        return message;
    }
};
export class MsgClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
    }
    UpdateCoordinatorAddress(request) {
        const data = MsgUpdateCoordinatorAddress.encode(request).finish();
        const promise = this.rpc.request('tendermint.spn.profile.Msg', 'UpdateCoordinatorAddress', data);
        return promise.then((data) => MsgUpdateCoordinatorAddressResponse.decode(new Reader(data)));
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
