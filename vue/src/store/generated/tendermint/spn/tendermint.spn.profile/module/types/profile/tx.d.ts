import { Reader, Writer } from 'protobufjs/minimal';
export declare const protobufPackage = "tendermint.spn.profile";
/** this line is used by starport scaffolding # proto/tx/message */
export interface MsgUpdateCoordinatorAddress {
    address: string;
    newAddress: string;
}
export interface MsgUpdateCoordinatorAddressResponse {
    coordinatorId: number;
}
export declare const MsgUpdateCoordinatorAddress: {
    encode(message: MsgUpdateCoordinatorAddress, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateCoordinatorAddress;
    fromJSON(object: any): MsgUpdateCoordinatorAddress;
    toJSON(message: MsgUpdateCoordinatorAddress): unknown;
    fromPartial(object: DeepPartial<MsgUpdateCoordinatorAddress>): MsgUpdateCoordinatorAddress;
};
export declare const MsgUpdateCoordinatorAddressResponse: {
    encode(message: MsgUpdateCoordinatorAddressResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateCoordinatorAddressResponse;
    fromJSON(object: any): MsgUpdateCoordinatorAddressResponse;
    toJSON(message: MsgUpdateCoordinatorAddressResponse): unknown;
    fromPartial(object: DeepPartial<MsgUpdateCoordinatorAddressResponse>): MsgUpdateCoordinatorAddressResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    /** this line is used by starport scaffolding # proto/tx/rpc */
    UpdateCoordinatorAddress(request: MsgUpdateCoordinatorAddress): Promise<MsgUpdateCoordinatorAddressResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    UpdateCoordinatorAddress(request: MsgUpdateCoordinatorAddress): Promise<MsgUpdateCoordinatorAddressResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
