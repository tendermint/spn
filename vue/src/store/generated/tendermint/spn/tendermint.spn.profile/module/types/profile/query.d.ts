import { Reader, Writer } from 'protobufjs/minimal';
import { ConsensusKeyNonce, ValidatorByConsAddress, ValidatorByAddress } from '../profile/validator';
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination';
import { CoordinatorByAddress, Coordinator } from '../profile/coordinator';
export declare const protobufPackage = "tendermint.spn.profile";
/** this line is used by starport scaffolding # 3 */
export interface QueryGetConsensusKeyNonceRequest {
    consAddress: string;
}
export interface QueryGetConsensusKeyNonceResponse {
    consensusKeyNonce: ConsensusKeyNonce | undefined;
}
export interface QueryGetValidatorByConsAddressRequest {
    consAddress: string;
}
export interface QueryGetValidatorByConsAddressResponse {
    validatorByConsAddress: ValidatorByConsAddress | undefined;
}
export interface QueryGetValidatorByAddressRequest {
    address: string;
}
export interface QueryGetValidatorByAddressResponse {
    validatorByAddress: ValidatorByAddress | undefined;
}
export interface QueryAllValidatorRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllValidatorResponse {
    validatorByAddress: ValidatorByAddress[];
    pagination: PageResponse | undefined;
}
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
export declare const QueryGetConsensusKeyNonceRequest: {
    encode(message: QueryGetConsensusKeyNonceRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetConsensusKeyNonceRequest;
    fromJSON(object: any): QueryGetConsensusKeyNonceRequest;
    toJSON(message: QueryGetConsensusKeyNonceRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetConsensusKeyNonceRequest>): QueryGetConsensusKeyNonceRequest;
};
export declare const QueryGetConsensusKeyNonceResponse: {
    encode(message: QueryGetConsensusKeyNonceResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetConsensusKeyNonceResponse;
    fromJSON(object: any): QueryGetConsensusKeyNonceResponse;
    toJSON(message: QueryGetConsensusKeyNonceResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetConsensusKeyNonceResponse>): QueryGetConsensusKeyNonceResponse;
};
export declare const QueryGetValidatorByConsAddressRequest: {
    encode(message: QueryGetValidatorByConsAddressRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetValidatorByConsAddressRequest;
    fromJSON(object: any): QueryGetValidatorByConsAddressRequest;
    toJSON(message: QueryGetValidatorByConsAddressRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetValidatorByConsAddressRequest>): QueryGetValidatorByConsAddressRequest;
};
export declare const QueryGetValidatorByConsAddressResponse: {
    encode(message: QueryGetValidatorByConsAddressResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetValidatorByConsAddressResponse;
    fromJSON(object: any): QueryGetValidatorByConsAddressResponse;
    toJSON(message: QueryGetValidatorByConsAddressResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetValidatorByConsAddressResponse>): QueryGetValidatorByConsAddressResponse;
};
export declare const QueryGetValidatorByAddressRequest: {
    encode(message: QueryGetValidatorByAddressRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetValidatorByAddressRequest;
    fromJSON(object: any): QueryGetValidatorByAddressRequest;
    toJSON(message: QueryGetValidatorByAddressRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetValidatorByAddressRequest>): QueryGetValidatorByAddressRequest;
};
export declare const QueryGetValidatorByAddressResponse: {
    encode(message: QueryGetValidatorByAddressResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetValidatorByAddressResponse;
    fromJSON(object: any): QueryGetValidatorByAddressResponse;
    toJSON(message: QueryGetValidatorByAddressResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetValidatorByAddressResponse>): QueryGetValidatorByAddressResponse;
};
export declare const QueryAllValidatorRequest: {
    encode(message: QueryAllValidatorRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllValidatorRequest;
    fromJSON(object: any): QueryAllValidatorRequest;
    toJSON(message: QueryAllValidatorRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllValidatorRequest>): QueryAllValidatorRequest;
};
export declare const QueryAllValidatorResponse: {
    encode(message: QueryAllValidatorResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllValidatorResponse;
    fromJSON(object: any): QueryAllValidatorResponse;
    toJSON(message: QueryAllValidatorResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllValidatorResponse>): QueryAllValidatorResponse;
};
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
    /** Queries a consensusKeyNonce by index. */
    ConsensusKeyNonce(request: QueryGetConsensusKeyNonceRequest): Promise<QueryGetConsensusKeyNonceResponse>;
    /** Queries a validatorByConsAddress by index. */
    ValidatorByConsAddress(request: QueryGetValidatorByConsAddressRequest): Promise<QueryGetValidatorByConsAddressResponse>;
    /** Queries a validatorByAddress by index. */
    ValidatorByAddress(request: QueryGetValidatorByAddressRequest): Promise<QueryGetValidatorByAddressResponse>;
    /** Queries a list of validatorByAddress items. */
    ValidatorByAddressAll(request: QueryAllValidatorRequest): Promise<QueryAllValidatorResponse>;
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
    ConsensusKeyNonce(request: QueryGetConsensusKeyNonceRequest): Promise<QueryGetConsensusKeyNonceResponse>;
    ValidatorByConsAddress(request: QueryGetValidatorByConsAddressRequest): Promise<QueryGetValidatorByConsAddressResponse>;
    ValidatorByAddress(request: QueryGetValidatorByAddressRequest): Promise<QueryGetValidatorByAddressResponse>;
    ValidatorByAddressAll(request: QueryAllValidatorRequest): Promise<QueryAllValidatorResponse>;
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
