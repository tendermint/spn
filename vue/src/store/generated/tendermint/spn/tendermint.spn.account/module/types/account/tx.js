/* eslint-disable */
import { Reader, util, configure, Writer } from 'protobufjs/minimal';
import * as Long from 'long';
import { CoordinatorDescription } from '../account/coordinator';
export const protobufPackage = 'tendermint.spn.account';
const baseMsgCreateCoordinator = {};
export const MsgCreateCoordinator = {
    encode(message, writer = Writer.create()) {
        if (message.description !== undefined) {
            CoordinatorDescription.encode(message.description, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgCreateCoordinator };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
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
        const message = { ...baseMsgCreateCoordinator };
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
        message.description !== undefined && (obj.description = message.description ? CoordinatorDescription.toJSON(message.description) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgCreateCoordinator };
        if (object.description !== undefined && object.description !== null) {
            message.description = CoordinatorDescription.fromPartial(object.description);
        }
        else {
            message.description = undefined;
        }
        return message;
    }
};
const baseMsgCreateCoordinatorResponse = { coordinatorId: 0 };
export const MsgCreateCoordinatorResponse = {
    encode(message, writer = Writer.create()) {
        if (message.coordinatorId !== 0) {
            writer.uint32(8).uint64(message.coordinatorId);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgCreateCoordinatorResponse };
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
        const message = { ...baseMsgCreateCoordinatorResponse };
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
        const message = { ...baseMsgCreateCoordinatorResponse };
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
    CreateCoordinator(request) {
        const data = MsgCreateCoordinator.encode(request).finish();
        const promise = this.rpc.request('tendermint.spn.account.Msg', 'CreateCoordinator', data);
        return promise.then((data) => MsgCreateCoordinatorResponse.decode(new Reader(data)));
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
