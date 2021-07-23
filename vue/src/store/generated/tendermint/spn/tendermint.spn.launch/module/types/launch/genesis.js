/* eslint-disable */
import { GenesisAccount } from '../launch/genesis_account';
import { Chain } from '../launch/chain';
import { Writer, Reader } from 'protobufjs/minimal';
export const protobufPackage = 'tendermint.spn.launch';
const baseGenesisState = {};
export const GenesisState = {
    encode(message, writer = Writer.create()) {
        for (const v of message.genesisAccountList) {
            GenesisAccount.encode(v, writer.uint32(18).fork()).ldelim();
        }
        for (const v of message.chainList) {
            Chain.encode(v, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseGenesisState };
        message.genesisAccountList = [];
        message.chainList = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 2:
                    message.genesisAccountList.push(GenesisAccount.decode(reader, reader.uint32()));
                    break;
                case 1:
                    message.chainList.push(Chain.decode(reader, reader.uint32()));
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseGenesisState };
        message.genesisAccountList = [];
        message.chainList = [];
        if (object.genesisAccountList !== undefined && object.genesisAccountList !== null) {
            for (const e of object.genesisAccountList) {
                message.genesisAccountList.push(GenesisAccount.fromJSON(e));
            }
        }
        if (object.chainList !== undefined && object.chainList !== null) {
            for (const e of object.chainList) {
                message.chainList.push(Chain.fromJSON(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.genesisAccountList) {
            obj.genesisAccountList = message.genesisAccountList.map((e) => (e ? GenesisAccount.toJSON(e) : undefined));
        }
        else {
            obj.genesisAccountList = [];
        }
        if (message.chainList) {
            obj.chainList = message.chainList.map((e) => (e ? Chain.toJSON(e) : undefined));
        }
        else {
            obj.chainList = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseGenesisState };
        message.genesisAccountList = [];
        message.chainList = [];
        if (object.genesisAccountList !== undefined && object.genesisAccountList !== null) {
            for (const e of object.genesisAccountList) {
                message.genesisAccountList.push(GenesisAccount.fromPartial(e));
            }
        }
        if (object.chainList !== undefined && object.chainList !== null) {
            for (const e of object.chainList) {
                message.chainList.push(Chain.fromPartial(e));
            }
        }
        return message;
    }
};
