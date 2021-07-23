import { Reader, Writer } from 'protobufjs/minimal';
export declare const protobufPackage = "tendermint.spn.launch";
/** this line is used by starport scaffolding # proto/tx/message */
export interface MsgCreateChain {
    coordinator: string;
    chainName: string;
    sourceURL: string;
    sourceHash: string;
    genesisURL: string;
    genesisHash: string;
}
export interface MsgCreateChainResponse {
    chainID: string;
}
export declare const MsgCreateChain: {
    encode(message: MsgCreateChain, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateChain;
    fromJSON(object: any): MsgCreateChain;
    toJSON(message: MsgCreateChain): unknown;
    fromPartial(object: DeepPartial<MsgCreateChain>): MsgCreateChain;
};
export declare const MsgCreateChainResponse: {
    encode(message: MsgCreateChainResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateChainResponse;
    fromJSON(object: any): MsgCreateChainResponse;
    toJSON(message: MsgCreateChainResponse): unknown;
    fromPartial(object: DeepPartial<MsgCreateChainResponse>): MsgCreateChainResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    /** this line is used by starport scaffolding # proto/tx/rpc */
    CreateChain(request: MsgCreateChain): Promise<MsgCreateChainResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    CreateChain(request: MsgCreateChain): Promise<MsgCreateChainResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
