/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal';
import { ChainNameCount } from '../launch/chain_name_count';
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination';
import { GenesisAccount } from '../launch/genesis_account';
import { Chain } from '../launch/chain';
export const protobufPackage = 'tendermint.spn.launch';
const baseQueryGetChainNameCountRequest = { chainName: '' };
export const QueryGetChainNameCountRequest = {
    encode(message, writer = Writer.create()) {
        if (message.chainName !== '') {
            writer.uint32(10).string(message.chainName);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetChainNameCountRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.chainName = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetChainNameCountRequest };
        if (object.chainName !== undefined && object.chainName !== null) {
            message.chainName = String(object.chainName);
        }
        else {
            message.chainName = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.chainName !== undefined && (obj.chainName = message.chainName);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetChainNameCountRequest };
        if (object.chainName !== undefined && object.chainName !== null) {
            message.chainName = object.chainName;
        }
        else {
            message.chainName = '';
        }
        return message;
    }
};
const baseQueryGetChainNameCountResponse = {};
export const QueryGetChainNameCountResponse = {
    encode(message, writer = Writer.create()) {
        if (message.chainNameCount !== undefined) {
            ChainNameCount.encode(message.chainNameCount, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetChainNameCountResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.chainNameCount = ChainNameCount.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetChainNameCountResponse };
        if (object.chainNameCount !== undefined && object.chainNameCount !== null) {
            message.chainNameCount = ChainNameCount.fromJSON(object.chainNameCount);
        }
        else {
            message.chainNameCount = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.chainNameCount !== undefined && (obj.chainNameCount = message.chainNameCount ? ChainNameCount.toJSON(message.chainNameCount) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetChainNameCountResponse };
        if (object.chainNameCount !== undefined && object.chainNameCount !== null) {
            message.chainNameCount = ChainNameCount.fromPartial(object.chainNameCount);
        }
        else {
            message.chainNameCount = undefined;
        }
        return message;
    }
};
const baseQueryAllChainNameCountRequest = {};
export const QueryAllChainNameCountRequest = {
    encode(message, writer = Writer.create()) {
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllChainNameCountRequest };
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
        const message = { ...baseQueryAllChainNameCountRequest };
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
        const message = { ...baseQueryAllChainNameCountRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
const baseQueryAllChainNameCountResponse = {};
export const QueryAllChainNameCountResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.chainNameCount) {
            ChainNameCount.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllChainNameCountResponse };
        message.chainNameCount = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.chainNameCount.push(ChainNameCount.decode(reader, reader.uint32()));
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
        const message = { ...baseQueryAllChainNameCountResponse };
        message.chainNameCount = [];
        if (object.chainNameCount !== undefined && object.chainNameCount !== null) {
            for (const e of object.chainNameCount) {
                message.chainNameCount.push(ChainNameCount.fromJSON(e));
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
        if (message.chainNameCount) {
            obj.chainNameCount = message.chainNameCount.map((e) => (e ? ChainNameCount.toJSON(e) : undefined));
        }
        else {
            obj.chainNameCount = [];
        }
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllChainNameCountResponse };
        message.chainNameCount = [];
        if (object.chainNameCount !== undefined && object.chainNameCount !== null) {
            for (const e of object.chainNameCount) {
                message.chainNameCount.push(ChainNameCount.fromPartial(e));
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
const baseQueryGetGenesisAccountRequest = { chainID: '', address: '' };
export const QueryGetGenesisAccountRequest = {
    encode(message, writer = Writer.create()) {
        if (message.chainID !== '') {
            writer.uint32(10).string(message.chainID);
        }
        if (message.address !== '') {
            writer.uint32(18).string(message.address);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetGenesisAccountRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.chainID = reader.string();
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
        const message = { ...baseQueryGetGenesisAccountRequest };
        if (object.chainID !== undefined && object.chainID !== null) {
            message.chainID = String(object.chainID);
        }
        else {
            message.chainID = '';
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
        message.chainID !== undefined && (obj.chainID = message.chainID);
        message.address !== undefined && (obj.address = message.address);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetGenesisAccountRequest };
        if (object.chainID !== undefined && object.chainID !== null) {
            message.chainID = object.chainID;
        }
        else {
            message.chainID = '';
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
const baseQueryGetGenesisAccountResponse = {};
export const QueryGetGenesisAccountResponse = {
    encode(message, writer = Writer.create()) {
        if (message.genesisAccount !== undefined) {
            GenesisAccount.encode(message.genesisAccount, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetGenesisAccountResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.genesisAccount = GenesisAccount.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetGenesisAccountResponse };
        if (object.genesisAccount !== undefined && object.genesisAccount !== null) {
            message.genesisAccount = GenesisAccount.fromJSON(object.genesisAccount);
        }
        else {
            message.genesisAccount = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.genesisAccount !== undefined && (obj.genesisAccount = message.genesisAccount ? GenesisAccount.toJSON(message.genesisAccount) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetGenesisAccountResponse };
        if (object.genesisAccount !== undefined && object.genesisAccount !== null) {
            message.genesisAccount = GenesisAccount.fromPartial(object.genesisAccount);
        }
        else {
            message.genesisAccount = undefined;
        }
        return message;
    }
};
const baseQueryAllGenesisAccountRequest = {};
export const QueryAllGenesisAccountRequest = {
    encode(message, writer = Writer.create()) {
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllGenesisAccountRequest };
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
        const message = { ...baseQueryAllGenesisAccountRequest };
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
        const message = { ...baseQueryAllGenesisAccountRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
const baseQueryAllGenesisAccountResponse = {};
export const QueryAllGenesisAccountResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.genesisAccount) {
            GenesisAccount.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllGenesisAccountResponse };
        message.genesisAccount = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.genesisAccount.push(GenesisAccount.decode(reader, reader.uint32()));
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
        const message = { ...baseQueryAllGenesisAccountResponse };
        message.genesisAccount = [];
        if (object.genesisAccount !== undefined && object.genesisAccount !== null) {
            for (const e of object.genesisAccount) {
                message.genesisAccount.push(GenesisAccount.fromJSON(e));
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
        if (message.genesisAccount) {
            obj.genesisAccount = message.genesisAccount.map((e) => (e ? GenesisAccount.toJSON(e) : undefined));
        }
        else {
            obj.genesisAccount = [];
        }
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllGenesisAccountResponse };
        message.genesisAccount = [];
        if (object.genesisAccount !== undefined && object.genesisAccount !== null) {
            for (const e of object.genesisAccount) {
                message.genesisAccount.push(GenesisAccount.fromPartial(e));
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
const baseQueryGetChainRequest = { chainID: '' };
export const QueryGetChainRequest = {
    encode(message, writer = Writer.create()) {
        if (message.chainID !== '') {
            writer.uint32(10).string(message.chainID);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetChainRequest };
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
        const message = { ...baseQueryGetChainRequest };
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
        const message = { ...baseQueryGetChainRequest };
        if (object.chainID !== undefined && object.chainID !== null) {
            message.chainID = object.chainID;
        }
        else {
            message.chainID = '';
        }
        return message;
    }
};
const baseQueryGetChainResponse = {};
export const QueryGetChainResponse = {
    encode(message, writer = Writer.create()) {
        if (message.chain !== undefined) {
            Chain.encode(message.chain, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetChainResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.chain = Chain.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetChainResponse };
        if (object.chain !== undefined && object.chain !== null) {
            message.chain = Chain.fromJSON(object.chain);
        }
        else {
            message.chain = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.chain !== undefined && (obj.chain = message.chain ? Chain.toJSON(message.chain) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetChainResponse };
        if (object.chain !== undefined && object.chain !== null) {
            message.chain = Chain.fromPartial(object.chain);
        }
        else {
            message.chain = undefined;
        }
        return message;
    }
};
const baseQueryAllChainRequest = {};
export const QueryAllChainRequest = {
    encode(message, writer = Writer.create()) {
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllChainRequest };
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
        const message = { ...baseQueryAllChainRequest };
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
        const message = { ...baseQueryAllChainRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
const baseQueryAllChainResponse = {};
export const QueryAllChainResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.chain) {
            Chain.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllChainResponse };
        message.chain = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.chain.push(Chain.decode(reader, reader.uint32()));
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
        const message = { ...baseQueryAllChainResponse };
        message.chain = [];
        if (object.chain !== undefined && object.chain !== null) {
            for (const e of object.chain) {
                message.chain.push(Chain.fromJSON(e));
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
        if (message.chain) {
            obj.chain = message.chain.map((e) => (e ? Chain.toJSON(e) : undefined));
        }
        else {
            obj.chain = [];
        }
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllChainResponse };
        message.chain = [];
        if (object.chain !== undefined && object.chain !== null) {
            for (const e of object.chain) {
                message.chain.push(Chain.fromPartial(e));
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
    ChainNameCount(request) {
        const data = QueryGetChainNameCountRequest.encode(request).finish();
        const promise = this.rpc.request('tendermint.spn.launch.Query', 'ChainNameCount', data);
        return promise.then((data) => QueryGetChainNameCountResponse.decode(new Reader(data)));
    }
    ChainNameCountAll(request) {
        const data = QueryAllChainNameCountRequest.encode(request).finish();
        const promise = this.rpc.request('tendermint.spn.launch.Query', 'ChainNameCountAll', data);
        return promise.then((data) => QueryAllChainNameCountResponse.decode(new Reader(data)));
    }
    GenesisAccount(request) {
        const data = QueryGetGenesisAccountRequest.encode(request).finish();
        const promise = this.rpc.request('tendermint.spn.launch.Query', 'GenesisAccount', data);
        return promise.then((data) => QueryGetGenesisAccountResponse.decode(new Reader(data)));
    }
    GenesisAccountAll(request) {
        const data = QueryAllGenesisAccountRequest.encode(request).finish();
        const promise = this.rpc.request('tendermint.spn.launch.Query', 'GenesisAccountAll', data);
        return promise.then((data) => QueryAllGenesisAccountResponse.decode(new Reader(data)));
    }
    Chain(request) {
        const data = QueryGetChainRequest.encode(request).finish();
        const promise = this.rpc.request('tendermint.spn.launch.Query', 'Chain', data);
        return promise.then((data) => QueryGetChainResponse.decode(new Reader(data)));
    }
    ChainAll(request) {
        const data = QueryAllChainRequest.encode(request).finish();
        const promise = this.rpc.request('tendermint.spn.launch.Query', 'ChainAll', data);
        return promise.then((data) => QueryAllChainResponse.decode(new Reader(data)));
    }
}
