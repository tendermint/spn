import { Writer, Reader } from 'protobufjs/minimal';
import { Any } from '../google/protobuf/any';
export declare const protobufPackage = "tendermint.spn.launch";
export interface Chain {
    chainID: string;
    coordinatorID: number;
    createdAt: number;
    sourceURL: string;
    sourceHash: string;
    initialGenesis: Any | undefined;
    launchTriggered: boolean;
    launchTimestamp: number;
}
/** DefaultInitialGenesis specifies using the default CLI-generated genesis as an initial genesis */
export interface DefaultInitialGenesis {
}
/** GenesisURL specifies using a custom genesis from a URL as the initial genesis */
export interface GenesisURL {
    url: string;
    hash: string;
}
export declare const Chain: {
    encode(message: Chain, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Chain;
    fromJSON(object: any): Chain;
    toJSON(message: Chain): unknown;
    fromPartial(object: DeepPartial<Chain>): Chain;
};
export declare const DefaultInitialGenesis: {
    encode(_: DefaultInitialGenesis, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): DefaultInitialGenesis;
    fromJSON(_: any): DefaultInitialGenesis;
    toJSON(_: DefaultInitialGenesis): unknown;
    fromPartial(_: DeepPartial<DefaultInitialGenesis>): DefaultInitialGenesis;
};
export declare const GenesisURL: {
    encode(message: GenesisURL, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): GenesisURL;
    fromJSON(object: any): GenesisURL;
    toJSON(message: GenesisURL): unknown;
    fromPartial(object: DeepPartial<GenesisURL>): GenesisURL;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
