import { Writer, Reader } from 'protobufjs/minimal';
import { ConsensusKeyNonce, ValidatorByConsAddress, ValidatorByAddress } from '../profile/validator';
import { CoordinatorByAddress, Coordinator } from '../profile/coordinator';
export declare const protobufPackage = "tendermint.spn.profile";
/** GenesisState defines the profile module's genesis state. */
export interface GenesisState {
    /** this line is used by starport scaffolding # genesis/proto/state */
    consensusKeyNonceList: ConsensusKeyNonce[];
    /** this line is used by starport scaffolding # genesis/proto/stateField */
    validatorByConsAddressList: ValidatorByConsAddress[];
    /** this line is used by starport scaffolding # genesis/proto/stateField */
    validatorByAddressList: ValidatorByAddress[];
    /** this line is used by starport scaffolding # genesis/proto/stateField */
    coordinatorByAddressList: CoordinatorByAddress[];
    /** this line is used by starport scaffolding # genesis/proto/stateField */
    coordinatorList: Coordinator[];
    /** this line is used by starport scaffolding # genesis/proto/stateField */
    coordinatorCount: number;
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
