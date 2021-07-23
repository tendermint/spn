import { Coin } from '../cosmos/base/v1beta1/coin';
import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "tendermint.spn.launch";
export interface GenesisAccount {
    chainID: string;
    address: string;
    Coins: Coin[];
}
export declare const GenesisAccount: {
    encode(message: GenesisAccount, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): GenesisAccount;
    fromJSON(object: any): GenesisAccount;
    toJSON(message: GenesisAccount): unknown;
    fromPartial(object: DeepPartial<GenesisAccount>): GenesisAccount;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
