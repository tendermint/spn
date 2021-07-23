/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal';
export const protobufPackage = 'tendermint.spn.launch';
const baseMsgCreateChain = { coordinator: '', chainName: '', sourceURL: '', sourceHash: '', genesisURL: '', genesisHash: '' };
export const MsgCreateChain = {
    encode(message, writer = Writer.create()) {
        if (message.coordinator !== '') {
            writer.uint32(10).string(message.coordinator);
        }
        if (message.chainName !== '') {
            writer.uint32(18).string(message.chainName);
        }
        if (message.sourceURL !== '') {
            writer.uint32(26).string(message.sourceURL);
        }
        if (message.sourceHash !== '') {
            writer.uint32(34).string(message.sourceHash);
        }
        if (message.genesisURL !== '') {
            writer.uint32(42).string(message.genesisURL);
        }
        if (message.genesisHash !== '') {
            writer.uint32(50).string(message.genesisHash);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgCreateChain };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.coordinator = reader.string();
                    break;
                case 2:
                    message.chainName = reader.string();
                    break;
                case 3:
                    message.sourceURL = reader.string();
                    break;
                case 4:
                    message.sourceHash = reader.string();
                    break;
                case 5:
                    message.genesisURL = reader.string();
                    break;
                case 6:
                    message.genesisHash = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgCreateChain };
        if (object.coordinator !== undefined && object.coordinator !== null) {
            message.coordinator = String(object.coordinator);
        }
        else {
            message.coordinator = '';
        }
        if (object.chainName !== undefined && object.chainName !== null) {
            message.chainName = String(object.chainName);
        }
        else {
            message.chainName = '';
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
        if (object.genesisURL !== undefined && object.genesisURL !== null) {
            message.genesisURL = String(object.genesisURL);
        }
        else {
            message.genesisURL = '';
        }
        if (object.genesisHash !== undefined && object.genesisHash !== null) {
            message.genesisHash = String(object.genesisHash);
        }
        else {
            message.genesisHash = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.coordinator !== undefined && (obj.coordinator = message.coordinator);
        message.chainName !== undefined && (obj.chainName = message.chainName);
        message.sourceURL !== undefined && (obj.sourceURL = message.sourceURL);
        message.sourceHash !== undefined && (obj.sourceHash = message.sourceHash);
        message.genesisURL !== undefined && (obj.genesisURL = message.genesisURL);
        message.genesisHash !== undefined && (obj.genesisHash = message.genesisHash);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgCreateChain };
        if (object.coordinator !== undefined && object.coordinator !== null) {
            message.coordinator = object.coordinator;
        }
        else {
            message.coordinator = '';
        }
        if (object.chainName !== undefined && object.chainName !== null) {
            message.chainName = object.chainName;
        }
        else {
            message.chainName = '';
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
        if (object.genesisURL !== undefined && object.genesisURL !== null) {
            message.genesisURL = object.genesisURL;
        }
        else {
            message.genesisURL = '';
        }
        if (object.genesisHash !== undefined && object.genesisHash !== null) {
            message.genesisHash = object.genesisHash;
        }
        else {
            message.genesisHash = '';
        }
        return message;
    }
};
const baseMsgCreateChainResponse = { chainID: '' };
export const MsgCreateChainResponse = {
    encode(message, writer = Writer.create()) {
        if (message.chainID !== '') {
            writer.uint32(10).string(message.chainID);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgCreateChainResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.chainID = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgCreateChainResponse };
        if (object.chainID !== undefined && object.chainID !== null) {
            message.chainID = String(object.chainID);
        }
        else {
            message.chainID = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.chainID !== undefined && (obj.chainID = message.chainID);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgCreateChainResponse };
        if (object.chainID !== undefined && object.chainID !== null) {
            message.chainID = object.chainID;
        }
        else {
            message.chainID = '';
        }
        return message;
    }
};
export class MsgClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
    }
    CreateChain(request) {
        const data = MsgCreateChain.encode(request).finish();
        const promise = this.rpc.request('tendermint.spn.launch.Msg', 'CreateChain', data);
        return promise.then((data) => MsgCreateChainResponse.decode(new Reader(data)));
    }
}
