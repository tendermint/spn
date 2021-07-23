/* eslint-disable */
import { ChainNameCount, Chain } from '../launch/chain'
import { GenesisAccount } from '../launch/genesis_account'
import { Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'tendermint.spn.launch'

/** GenesisState defines the launch module's genesis state. */
export interface GenesisState {
  /** this line is used by starport scaffolding # genesis/proto/state */
  chainNameCountList: ChainNameCount[]
  /** this line is used by starport scaffolding # genesis/proto/stateField */
  genesisAccountList: GenesisAccount[]
  /** this line is used by starport scaffolding # genesis/proto/stateField */
  chainList: Chain[]
}

const baseGenesisState: object = {}

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    for (const v of message.chainNameCountList) {
      ChainNameCount.encode(v!, writer.uint32(26).fork()).ldelim()
    }
    for (const v of message.genesisAccountList) {
      GenesisAccount.encode(v!, writer.uint32(18).fork()).ldelim()
    }
    for (const v of message.chainList) {
      Chain.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseGenesisState } as GenesisState
    message.chainNameCountList = []
    message.genesisAccountList = []
    message.chainList = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 3:
          message.chainNameCountList.push(ChainNameCount.decode(reader, reader.uint32()))
          break
        case 2:
          message.genesisAccountList.push(GenesisAccount.decode(reader, reader.uint32()))
          break
        case 1:
          message.chainList.push(Chain.decode(reader, reader.uint32()))
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): GenesisState {
    const message = { ...baseGenesisState } as GenesisState
    message.chainNameCountList = []
    message.genesisAccountList = []
    message.chainList = []
    if (object.chainNameCountList !== undefined && object.chainNameCountList !== null) {
      for (const e of object.chainNameCountList) {
        message.chainNameCountList.push(ChainNameCount.fromJSON(e))
      }
    }
    if (object.genesisAccountList !== undefined && object.genesisAccountList !== null) {
      for (const e of object.genesisAccountList) {
        message.genesisAccountList.push(GenesisAccount.fromJSON(e))
      }
    }
    if (object.chainList !== undefined && object.chainList !== null) {
      for (const e of object.chainList) {
        message.chainList.push(Chain.fromJSON(e))
      }
    }
    return message
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {}
    if (message.chainNameCountList) {
      obj.chainNameCountList = message.chainNameCountList.map((e) => (e ? ChainNameCount.toJSON(e) : undefined))
    } else {
      obj.chainNameCountList = []
    }
    if (message.genesisAccountList) {
      obj.genesisAccountList = message.genesisAccountList.map((e) => (e ? GenesisAccount.toJSON(e) : undefined))
    } else {
      obj.genesisAccountList = []
    }
    if (message.chainList) {
      obj.chainList = message.chainList.map((e) => (e ? Chain.toJSON(e) : undefined))
    } else {
      obj.chainList = []
    }
    return obj
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState
    message.chainNameCountList = []
    message.genesisAccountList = []
    message.chainList = []
    if (object.chainNameCountList !== undefined && object.chainNameCountList !== null) {
      for (const e of object.chainNameCountList) {
        message.chainNameCountList.push(ChainNameCount.fromPartial(e))
      }
    }
    if (object.genesisAccountList !== undefined && object.genesisAccountList !== null) {
      for (const e of object.genesisAccountList) {
        message.genesisAccountList.push(GenesisAccount.fromPartial(e))
      }
    }
    if (object.chainList !== undefined && object.chainList !== null) {
      for (const e of object.chainList) {
        message.chainList.push(Chain.fromPartial(e))
      }
    }
    return message
  }
}

type Builtin = Date | Function | Uint8Array | string | number | undefined
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>
