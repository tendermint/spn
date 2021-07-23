/* eslint-disable */
import { Reader, util, configure, Writer } from 'protobufjs/minimal';
import * as Long from 'long';
import { CoordinatorDescription } from '../profile/coordinator';
export const protobufPackage = 'tendermint.spn.profile';
const baseMsgUpdateCoordinatorDescription = { address: '' };
export const MsgUpdateCoordinatorDescription = {
    encode(message, writer = Writer.create()) {
        if (message.address !== '') {
            writer.uint32(10).string(message.address);
        }
        if (message.description !== undefined) {
            CoordinatorDescription.encode(message.description, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgUpdateCoordinatorDescription };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.address = reader.string();
                    break;
                case 2:
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
        const message = { ...baseMsgUpdateCoordinatorDescription };
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
        message.address !== undefined && (obj.address = message.address);
        message.description !== undefined && (obj.description = message.description ? CoordinatorDescription.toJSON(message.description) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgUpdateCoordinatorDescription };
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
const baseMsgUpdateCoordinatorDescriptionResponse = { coordinatorId: 0 };
export const MsgUpdateCoordinatorDescriptionResponse = {
    encode(message, writer = Writer.create()) {
        if (message.coordinatorId !== 0) {
            writer.uint32(8).uint64(message.coordinatorId);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgUpdateCoordinatorDescriptionResponse };
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
        const message = { ...baseMsgUpdateCoordinatorDescriptionResponse };
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
        const message = { ...baseMsgUpdateCoordinatorDescriptionResponse };
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
    UpdateCoordinatorDescription(request) {
        const data = MsgUpdateCoordinatorDescription.encode(request).finish();
        const promise = this.rpc.request('tendermint.spn.profile.Msg', 'UpdateCoordinatorDescription', data);
        return promise.then((data) => MsgUpdateCoordinatorDescriptionResponse.decode(new Reader(data)));
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
