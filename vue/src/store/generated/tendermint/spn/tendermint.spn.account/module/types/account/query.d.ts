import { Reader, Writer } from 'protobufjs/minimal';
import { CoordinatorByAddress, Coordinator } from '../account/coordinator';
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination';
export declare const protobufPackage = "tendermint.spn.account";
/** this line is used by starport scaffolding # 3 */
export interface QueryGetCoordinatorByAddressRequest {
    address: string;
}
export interface QueryGetCoordinatorByAddressResponse {
    coordinatorByAddress: CoordinatorByAddress | undefined;
}
export interface QueryAllCoordinatorByAddressRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllCoordinatorByAddressResponse {
    coordinatorByAddress: CoordinatorByAddress[];
    pagination: PageResponse | undefined;
}
export interface QueryGetCoordinatorRequest {
    id: number;
}
export interface QueryGetCoordinatorResponse {
    Coordinator: Coordinator | undefined;
}
export interface QueryAllCoordinatorRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllCoordinatorResponse {
    Coordinator: Coordinator[];
    pagination: PageResponse | undefined;
}
export declare const QueryGetCoordinatorByAddressRequest: {
    encode(message: QueryGetCoordinatorByAddressRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetCoordinatorByAddressRequest;
    fromJSON(object: any): QueryGetCoordinatorByAddressRequest;
    toJSON(message: QueryGetCoordinatorByAddressRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetCoordinatorByAddressRequest>): QueryGetCoordinatorByAddressRequest;
};
export declare const QueryGetCoordinatorByAddressResponse: {
    encode(message: QueryGetCoordinatorByAddressResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetCoordinatorByAddressResponse;
    fromJSON(object: any): QueryGetCoordinatorByAddressResponse;
    toJSON(message: QueryGetCoordinatorByAddressResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetCoordinatorByAddressResponse>): QueryGetCoordinatorByAddressResponse;
};
export declare const QueryAllCoordinatorByAddressRequest: {
    encode(message: QueryAllCoordinatorByAddressRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllCoordinatorByAddressRequest;
    fromJSON(object: any): QueryAllCoordinatorByAddressRequest;
    toJSON(message: QueryAllCoordinatorByAddressRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllCoordinatorByAddressRequest>): QueryAllCoordinatorByAddressRequest;
};
export declare const QueryAllCoordinatorByAddressResponse: {
    encode(message: QueryAllCoordinatorByAddressResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllCoordinatorByAddressResponse;
    fromJSON(object: any): QueryAllCoordinatorByAddressResponse;
    toJSON(message: QueryAllCoordinatorByAddressResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllCoordinatorByAddressResponse>): QueryAllCoordinatorByAddressResponse;
};
export declare const QueryGetCoordinatorRequest: {
    encode(message: QueryGetCoordinatorRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetCoordinatorRequest;
    fromJSON(object: any): QueryGetCoordinatorRequest;
    toJSON(message: QueryGetCoordinatorRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetCoordinatorRequest>): QueryGetCoordinatorRequest;
};
export declare const QueryGetCoordinatorResponse: {
    encode(message: QueryGetCoordinatorResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetCoordinatorResponse;
    fromJSON(object: any): QueryGetCoordinatorResponse;
    toJSON(message: QueryGetCoordinatorResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetCoordinatorResponse>): QueryGetCoordinatorResponse;
};
export declare const QueryAllCoordinatorRequest: {
    encode(message: QueryAllCoordinatorRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllCoordinatorRequest;
    fromJSON(object: any): QueryAllCoordinatorRequest;
    toJSON(message: QueryAllCoordinatorRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllCoordinatorRequest>): QueryAllCoordinatorRequest;
};
export declare const QueryAllCoordinatorResponse: {
    encode(message: QueryAllCoordinatorResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllCoordinatorResponse;
    fromJSON(object: any): QueryAllCoordinatorResponse;
    toJSON(message: QueryAllCoordinatorResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllCoordinatorResponse>): QueryAllCoordinatorResponse;
};
/** Query defines the gRPC querier service. */
export interface Query {
    /** Queries a coordinatorByAddress by index. */
    CoordinatorByAddress(request: QueryGetCoordinatorByAddressRequest): Promise<QueryGetCoordinatorByAddressResponse>;
    /** Queries a coordinator by id. */
    Coordinator(request: QueryGetCoordinatorRequest): Promise<QueryGetCoordinatorResponse>;
    /** Queries a list of coordinator items. */
    CoordinatorAll(request: QueryAllCoordinatorRequest): Promise<QueryAllCoordinatorResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    CoordinatorByAddress(request: QueryGetCoordinatorByAddressRequest): Promise<QueryGetCoordinatorByAddressResponse>;
    Coordinator(request: QueryGetCoordinatorRequest): Promise<QueryGetCoordinatorResponse>;
    CoordinatorAll(request: QueryAllCoordinatorRequest): Promise<QueryAllCoordinatorResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
