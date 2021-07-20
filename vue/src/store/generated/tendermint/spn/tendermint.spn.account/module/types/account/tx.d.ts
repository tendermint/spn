import { Reader, Writer } from 'protobufjs/minimal';
export declare const protobufPackage = "tendermint.spn.account";
/** this line is used by starport scaffolding # proto/tx/message */
export interface MsgDeleteCoordinator {
    address: string;
}
export interface MsgDeleteCoordinatorResponse {
}
export declare const MsgDeleteCoordinator: {
    encode(message: MsgDeleteCoordinator, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteCoordinator;
    fromJSON(object: any): MsgDeleteCoordinator;
    toJSON(message: MsgDeleteCoordinator): unknown;
    fromPartial(object: DeepPartial<MsgDeleteCoordinator>): MsgDeleteCoordinator;
};
export declare const MsgDeleteCoordinatorResponse: {
    encode(_: MsgDeleteCoordinatorResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteCoordinatorResponse;
    fromJSON(_: any): MsgDeleteCoordinatorResponse;
    toJSON(_: MsgDeleteCoordinatorResponse): unknown;
    fromPartial(_: DeepPartial<MsgDeleteCoordinatorResponse>): MsgDeleteCoordinatorResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    /** this line is used by starport scaffolding # proto/tx/rpc */
    DeleteCoordinator(request: MsgDeleteCoordinator): Promise<MsgDeleteCoordinatorResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    DeleteCoordinator(request: MsgDeleteCoordinator): Promise<MsgDeleteCoordinatorResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
