import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "tendermint.spn.launch";
export interface ChainNameCount {
    creator: string;
    chainName: string;
    count: number;
}
export declare const ChainNameCount: {
    encode(message: ChainNameCount, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): ChainNameCount;
    fromJSON(object: any): ChainNameCount;
    toJSON(message: ChainNameCount): unknown;
    fromPartial(object: DeepPartial<ChainNameCount>): ChainNameCount;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
