import { Reader, Writer } from 'protobufjs/minimal';
import { ValidatorDescription } from '../profile/validator';
import { CoordinatorDescription } from '../profile/coordinator';
export declare const protobufPackage = "tendermint.spn.profile";
/** this line is used by starport scaffolding # proto/tx/message */
export interface MsgUpdateValidatorDescription {
    address: string;
    description: ValidatorDescription | undefined;
}
export interface MsgUpdateValidatorDescriptionResponse {
}
export interface MsgCreateCoordinator {
    address: string;
    description: CoordinatorDescription | undefined;
}
export interface MsgCreateCoordinatorResponse {
    coordinatorId: number;
}
export declare const MsgUpdateValidatorDescription: {
    encode(message: MsgUpdateValidatorDescription, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateValidatorDescription;
    fromJSON(object: any): MsgUpdateValidatorDescription;
    toJSON(message: MsgUpdateValidatorDescription): unknown;
    fromPartial(object: DeepPartial<MsgUpdateValidatorDescription>): MsgUpdateValidatorDescription;
};
export declare const MsgUpdateValidatorDescriptionResponse: {
    encode(_: MsgUpdateValidatorDescriptionResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateValidatorDescriptionResponse;
    fromJSON(_: any): MsgUpdateValidatorDescriptionResponse;
    toJSON(_: MsgUpdateValidatorDescriptionResponse): unknown;
    fromPartial(_: DeepPartial<MsgUpdateValidatorDescriptionResponse>): MsgUpdateValidatorDescriptionResponse;
};
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
/** Msg defines the Msg service. */
export interface Msg {
    /** this line is used by starport scaffolding # proto/tx/rpc */
    UpdateValidatorDescription(request: MsgUpdateValidatorDescription): Promise<MsgUpdateValidatorDescriptionResponse>;
    CreateCoordinator(request: MsgCreateCoordinator): Promise<MsgCreateCoordinatorResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    UpdateValidatorDescription(request: MsgUpdateValidatorDescription): Promise<MsgUpdateValidatorDescriptionResponse>;
    CreateCoordinator(request: MsgCreateCoordinator): Promise<MsgCreateCoordinatorResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
