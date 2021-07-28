import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "tendermint.spn.profile";
export interface Coordinator {
    coordinatorId: number;
    address: string;
    description: CoordinatorDescription | undefined;
}
export interface CoordinatorDescription {
    identity: string;
    website: string;
    details: string;
}
export interface CoordinatorByAddress {
    address: string;
    coordinatorId: number;
}
export declare const Coordinator: {
    encode(message: Coordinator, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Coordinator;
    fromJSON(object: any): Coordinator;
    toJSON(message: Coordinator): unknown;
    fromPartial(object: DeepPartial<Coordinator>): Coordinator;
};
export declare const CoordinatorDescription: {
    encode(message: CoordinatorDescription, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): CoordinatorDescription;
    fromJSON(object: any): CoordinatorDescription;
    toJSON(message: CoordinatorDescription): unknown;
    fromPartial(object: DeepPartial<CoordinatorDescription>): CoordinatorDescription;
};
export declare const CoordinatorByAddress: {
    encode(message: CoordinatorByAddress, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): CoordinatorByAddress;
    fromJSON(object: any): CoordinatorByAddress;
    toJSON(message: CoordinatorByAddress): unknown;
    fromPartial(object: DeepPartial<CoordinatorByAddress>): CoordinatorByAddress;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
