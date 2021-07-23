/* eslint-disable */
import { Coin } from '../cosmos/base/v1beta1/coin'
import { Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'tendermint.spn.launch'

export interface GenesisAccount {
  chainID: string
  address: string
  Coins: Coin[]
}

const baseGenesisAccount: object = { chainID: '', address: '' }

export const GenesisAccount = {
  encode(message: GenesisAccount, writer: Writer = Writer.create()): Writer {
    if (message.chainID !== '') {
      writer.uint32(10).string(message.chainID)
    }
    if (message.address !== '') {
      writer.uint32(18).string(message.address)
    }
    for (const v of message.Coins) {
      Coin.encode(v!, writer.uint32(26).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisAccount {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseGenesisAccount } as GenesisAccount
    message.Coins = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.chainID = reader.string()
          break
        case 2:
          message.address = reader.string()
          break
        case 3:
          message.Coins.push(Coin.decode(reader, reader.uint32()))
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): GenesisAccount {
    const message = { ...baseGenesisAccount } as GenesisAccount
    message.Coins = []
    if (object.chainID !== undefined && object.chainID !== null) {
      message.chainID = String(object.chainID)
    } else {
      message.chainID = ''
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address)
    } else {
      message.address = ''
    }
    if (object.Coins !== undefined && object.Coins !== null) {
      for (const e of object.Coins) {
        message.Coins.push(Coin.fromJSON(e))
      }
    }
    return message
  },

  toJSON(message: GenesisAccount): unknown {
    const obj: any = {}
    message.chainID !== undefined && (obj.chainID = message.chainID)
    message.address !== undefined && (obj.address = message.address)
    if (message.Coins) {
      obj.Coins = message.Coins.map((e) => (e ? Coin.toJSON(e) : undefined))
    } else {
      obj.Coins = []
    }
    return obj
  },

  fromPartial(object: DeepPartial<GenesisAccount>): GenesisAccount {
    const message = { ...baseGenesisAccount } as GenesisAccount
    message.Coins = []
    if (object.chainID !== undefined && object.chainID !== null) {
      message.chainID = object.chainID
    } else {
      message.chainID = ''
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address
    } else {
      message.address = ''
    }
    if (object.Coins !== undefined && object.Coins !== null) {
      for (const e of object.Coins) {
        message.Coins.push(Coin.fromPartial(e))
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
