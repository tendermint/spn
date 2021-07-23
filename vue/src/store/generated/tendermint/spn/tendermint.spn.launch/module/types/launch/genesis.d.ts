import { GenesisAccount } from '../launch/genesis_account';
import { Chain } from '../launch/chain';
import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "tendermint.spn.launch";
/** GenesisState defines the launch module's genesis state. */
export interface GenesisState {
    /** this line is used by starport scaffolding # genesis/proto/state */
    genesisAccountList: GenesisAccount[];
    /** this line is used by starport scaffolding # genesis/proto/stateField */
    chainList: Chain[];
}
export declare const GenesisState: {
    encode(message: GenesisState, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): GenesisState;
    fromJSON(object: any): GenesisState;
    toJSON(message: GenesisState): unknown;
    fromPartial(object: DeepPartial<GenesisState>): GenesisState;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
