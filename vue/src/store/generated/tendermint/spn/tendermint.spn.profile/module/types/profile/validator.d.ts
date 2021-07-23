import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "tendermint.spn.profile";
export interface ValidatorByAddress {
    address: string;
    consensusAddress: string;
    description: ValidatorDescription | undefined;
}
export interface ValidatorDescription {
    identity: string;
    moniker: string;
    website: string;
    securityContact: string;
    details: string;
}
export interface ValidatorByConsAddress {
    consAddress: string;
    address: string;
}
export interface ConsensusKeyNonce {
    consAddress: string;
    nonce: number;
}
export declare const ValidatorByAddress: {
    encode(message: ValidatorByAddress, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): ValidatorByAddress;
    fromJSON(object: any): ValidatorByAddress;
    toJSON(message: ValidatorByAddress): unknown;
    fromPartial(object: DeepPartial<ValidatorByAddress>): ValidatorByAddress;
};
export declare const ValidatorDescription: {
    encode(message: ValidatorDescription, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): ValidatorDescription;
    fromJSON(object: any): ValidatorDescription;
    toJSON(message: ValidatorDescription): unknown;
    fromPartial(object: DeepPartial<ValidatorDescription>): ValidatorDescription;
};
export declare const ValidatorByConsAddress: {
    encode(message: ValidatorByConsAddress, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): ValidatorByConsAddress;
    fromJSON(object: any): ValidatorByConsAddress;
    toJSON(message: ValidatorByConsAddress): unknown;
    fromPartial(object: DeepPartial<ValidatorByConsAddress>): ValidatorByConsAddress;
};
export declare const ConsensusKeyNonce: {
    encode(message: ConsensusKeyNonce, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): ConsensusKeyNonce;
    fromJSON(object: any): ConsensusKeyNonce;
    toJSON(message: ConsensusKeyNonce): unknown;
    fromPartial(object: DeepPartial<ConsensusKeyNonce>): ConsensusKeyNonce;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
