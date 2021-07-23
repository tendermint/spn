import { Reader, Writer } from 'protobufjs/minimal';
import { GenesisAccount } from '../launch/genesis_account';
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination';
import { Chain } from '../launch/chain';
export declare const protobufPackage = "tendermint.spn.launch";
/** this line is used by starport scaffolding # 3 */
export interface QueryGetGenesisAccountRequest {
    chainID: string;
    address: string;
}
export interface QueryGetGenesisAccountResponse {
    genesisAccount: GenesisAccount | undefined;
}
export interface QueryAllGenesisAccountRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllGenesisAccountResponse {
    genesisAccount: GenesisAccount[];
    pagination: PageResponse | undefined;
}
export interface QueryGetChainRequest {
    chainID: string;
}
export interface QueryGetChainResponse {
    chain: Chain | undefined;
}
export interface QueryAllChainRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllChainResponse {
    chain: Chain[];
    pagination: PageResponse | undefined;
}
export declare const QueryGetGenesisAccountRequest: {
    encode(message: QueryGetGenesisAccountRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetGenesisAccountRequest;
    fromJSON(object: any): QueryGetGenesisAccountRequest;
    toJSON(message: QueryGetGenesisAccountRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetGenesisAccountRequest>): QueryGetGenesisAccountRequest;
};
export declare const QueryGetGenesisAccountResponse: {
    encode(message: QueryGetGenesisAccountResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetGenesisAccountResponse;
    fromJSON(object: any): QueryGetGenesisAccountResponse;
    toJSON(message: QueryGetGenesisAccountResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetGenesisAccountResponse>): QueryGetGenesisAccountResponse;
};
export declare const QueryAllGenesisAccountRequest: {
    encode(message: QueryAllGenesisAccountRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllGenesisAccountRequest;
    fromJSON(object: any): QueryAllGenesisAccountRequest;
    toJSON(message: QueryAllGenesisAccountRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllGenesisAccountRequest>): QueryAllGenesisAccountRequest;
};
export declare const QueryAllGenesisAccountResponse: {
    encode(message: QueryAllGenesisAccountResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllGenesisAccountResponse;
    fromJSON(object: any): QueryAllGenesisAccountResponse;
    toJSON(message: QueryAllGenesisAccountResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllGenesisAccountResponse>): QueryAllGenesisAccountResponse;
};
export declare const QueryGetChainRequest: {
    encode(message: QueryGetChainRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetChainRequest;
    fromJSON(object: any): QueryGetChainRequest;
    toJSON(message: QueryGetChainRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetChainRequest>): QueryGetChainRequest;
};
export declare const QueryGetChainResponse: {
    encode(message: QueryGetChainResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetChainResponse;
    fromJSON(object: any): QueryGetChainResponse;
    toJSON(message: QueryGetChainResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetChainResponse>): QueryGetChainResponse;
};
export declare const QueryAllChainRequest: {
    encode(message: QueryAllChainRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllChainRequest;
    fromJSON(object: any): QueryAllChainRequest;
    toJSON(message: QueryAllChainRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllChainRequest>): QueryAllChainRequest;
};
export declare const QueryAllChainResponse: {
    encode(message: QueryAllChainResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllChainResponse;
    fromJSON(object: any): QueryAllChainResponse;
    toJSON(message: QueryAllChainResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllChainResponse>): QueryAllChainResponse;
};
/** Query defines the gRPC querier service. */
export interface Query {
    /** Queries a genesisAccount by index. */
    GenesisAccount(request: QueryGetGenesisAccountRequest): Promise<QueryGetGenesisAccountResponse>;
    /** Queries a list of genesisAccount items. */
    GenesisAccountAll(request: QueryAllGenesisAccountRequest): Promise<QueryAllGenesisAccountResponse>;
    /** Queries a chain by index. */
    Chain(request: QueryGetChainRequest): Promise<QueryGetChainResponse>;
    /** Queries a list of chain items. */
    ChainAll(request: QueryAllChainRequest): Promise<QueryAllChainResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    GenesisAccount(request: QueryGetGenesisAccountRequest): Promise<QueryGetGenesisAccountResponse>;
    GenesisAccountAll(request: QueryAllGenesisAccountRequest): Promise<QueryAllGenesisAccountResponse>;
    Chain(request: QueryGetChainRequest): Promise<QueryGetChainResponse>;
    ChainAll(request: QueryAllChainRequest): Promise<QueryAllChainResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
