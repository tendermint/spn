import { Reader, Writer } from 'protobufjs/minimal';
import { CoordinatorDescription } from '../account/coordinator';
export declare const protobufPackage = "tendermint.spn.account";
/** this line is used by starport scaffolding # proto/tx/message */
export interface MsgUpdateCoordinatorDescription {
    address: string;
    description: CoordinatorDescription | undefined;
}
export interface MsgUpdateCoordinatorDescriptionResponse {
    coordinatorId: number;
}
export declare const MsgUpdateCoordinatorDescription: {
    encode(message: MsgUpdateCoordinatorDescription, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateCoordinatorDescription;
    fromJSON(object: any): MsgUpdateCoordinatorDescription;
    toJSON(message: MsgUpdateCoordinatorDescription): unknown;
    fromPartial(object: DeepPartial<MsgUpdateCoordinatorDescription>): MsgUpdateCoordinatorDescription;
};
export declare const MsgUpdateCoordinatorDescriptionResponse: {
    encode(message: MsgUpdateCoordinatorDescriptionResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateCoordinatorDescriptionResponse;
    fromJSON(object: any): MsgUpdateCoordinatorDescriptionResponse;
    toJSON(message: MsgUpdateCoordinatorDescriptionResponse): unknown;
    fromPartial(object: DeepPartial<MsgUpdateCoordinatorDescriptionResponse>): MsgUpdateCoordinatorDescriptionResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    /** this line is used by starport scaffolding # proto/tx/rpc */
    UpdateCoordinatorDescription(request: MsgUpdateCoordinatorDescription): Promise<MsgUpdateCoordinatorDescriptionResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    UpdateCoordinatorDescription(request: MsgUpdateCoordinatorDescription): Promise<MsgUpdateCoordinatorDescriptionResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
