import { Reader, Writer } from 'protobufjs/minimal';
import { CoordinatorDescription } from '../profile/coordinator';
export declare const protobufPackage = "tendermint.spn.profile";
/** this line is used by starport scaffolding # proto/tx/message */
export interface MsgCreateCoordinator {
    address: string;
    description: CoordinatorDescription | undefined;
}
export interface MsgCreateCoordinatorResponse {
    coordinatorId: number;
}
export interface MsgDeleteCoordinator {
    address: string;
}
export interface MsgDeleteCoordinatorResponse {
    coordinatorId: number;
}
export declare const MsgCreateCoordinator: {
    encode(message: MsgCreateCoordinator, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateCoordinator;
    fromJSON(object: any): MsgCreateCoordinator;
    toJSON(message: MsgCreateCoordinator): unknown;
    fromPartial(object: DeepPartial<MsgCreateCoordinator>): MsgCreateCoordinator;
};
export declare const MsgCreateCoordinatorResponse: {
    encode(message: MsgCreateCoordinatorResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateCoordinatorResponse;
    fromJSON(object: any): MsgCreateCoordinatorResponse;
    toJSON(message: MsgCreateCoordinatorResponse): unknown;
    fromPartial(object: DeepPartial<MsgCreateCoordinatorResponse>): MsgCreateCoordinatorResponse;
};
export declare const MsgDeleteCoordinator: {
    encode(message: MsgDeleteCoordinator, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteCoordinator;
    fromJSON(object: any): MsgDeleteCoordinator;
    toJSON(message: MsgDeleteCoordinator): unknown;
    fromPartial(object: DeepPartial<MsgDeleteCoordinator>): MsgDeleteCoordinator;
};
export declare const MsgDeleteCoordinatorResponse: {
    encode(message: MsgDeleteCoordinatorResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteCoordinatorResponse;
    fromJSON(object: any): MsgDeleteCoordinatorResponse;
    toJSON(message: MsgDeleteCoordinatorResponse): unknown;
    fromPartial(object: DeepPartial<MsgDeleteCoordinatorResponse>): MsgDeleteCoordinatorResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    /** this line is used by starport scaffolding # proto/tx/rpc */
    CreateCoordinator(request: MsgCreateCoordinator): Promise<MsgCreateCoordinatorResponse>;
    DeleteCoordinator(request: MsgDeleteCoordinator): Promise<MsgDeleteCoordinatorResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    CreateCoordinator(request: MsgCreateCoordinator): Promise<MsgCreateCoordinatorResponse>;
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
