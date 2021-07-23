/* eslint-disable */
import * as Long from 'long'
import { util, configure, Writer, Reader } from 'protobufjs/minimal'
import { ConsensusKeyNonce, ValidatorByConsAddress, ValidatorByAddress } from '../profile/validator'
import { CoordinatorByAddress, Coordinator } from '../profile/coordinator'

export const protobufPackage = 'tendermint.spn.profile'

/** GenesisState defines the profile module's genesis state. */
export interface GenesisState {
  /** this line is used by starport scaffolding # genesis/proto/state */
  consensusKeyNonceList: ConsensusKeyNonce[]
  /** this line is used by starport scaffolding # genesis/proto/stateField */
  validatorByConsAddressList: ValidatorByConsAddress[]
  /** this line is used by starport scaffolding # genesis/proto/stateField */
  validatorByAddressList: ValidatorByAddress[]
  /** this line is used by starport scaffolding # genesis/proto/stateField */
  coordinatorByAddressList: CoordinatorByAddress[]
  /** this line is used by starport scaffolding # genesis/proto/stateField */
  coordinatorList: Coordinator[]
  /** this line is used by starport scaffolding # genesis/proto/stateField */
  coordinatorCount: number
}

const baseGenesisState: object = { coordinatorCount: 0 }

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    for (const v of message.consensusKeyNonceList) {
      ConsensusKeyNonce.encode(v!, writer.uint32(50).fork()).ldelim()
    }
    for (const v of message.validatorByConsAddressList) {
      ValidatorByConsAddress.encode(v!, writer.uint32(42).fork()).ldelim()
    }
    for (const v of message.validatorByAddressList) {
      ValidatorByAddress.encode(v!, writer.uint32(34).fork()).ldelim()
    }
    for (const v of message.coordinatorByAddressList) {
      CoordinatorByAddress.encode(v!, writer.uint32(26).fork()).ldelim()
    }
    for (const v of message.coordinatorList) {
      Coordinator.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    if (message.coordinatorCount !== 0) {
      writer.uint32(16).uint64(message.coordinatorCount)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseGenesisState } as GenesisState
    message.consensusKeyNonceList = []
    message.validatorByConsAddressList = []
    message.validatorByAddressList = []
    message.coordinatorByAddressList = []
    message.coordinatorList = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 6:
          message.consensusKeyNonceList.push(ConsensusKeyNonce.decode(reader, reader.uint32()))
          break
        case 5:
          message.validatorByConsAddressList.push(ValidatorByConsAddress.decode(reader, reader.uint32()))
          break
        case 4:
          message.validatorByAddressList.push(ValidatorByAddress.decode(reader, reader.uint32()))
          break
        case 3:
          message.coordinatorByAddressList.push(CoordinatorByAddress.decode(reader, reader.uint32()))
          break
        case 1:
          message.coordinatorList.push(Coordinator.decode(reader, reader.uint32()))
          break
        case 2:
          message.coordinatorCount = longToNumber(reader.uint64() as Long)
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
    message.consensusKeyNonceList = []
    message.validatorByConsAddressList = []
    message.validatorByAddressList = []
    message.coordinatorByAddressList = []
    message.coordinatorList = []
    if (object.consensusKeyNonceList !== undefined && object.consensusKeyNonceList !== null) {
      for (const e of object.consensusKeyNonceList) {
        message.consensusKeyNonceList.push(ConsensusKeyNonce.fromJSON(e))
      }
    }
    if (object.validatorByConsAddressList !== undefined && object.validatorByConsAddressList !== null) {
      for (const e of object.validatorByConsAddressList) {
        message.validatorByConsAddressList.push(ValidatorByConsAddress.fromJSON(e))
      }
    }
    if (object.validatorByAddressList !== undefined && object.validatorByAddressList !== null) {
      for (const e of object.validatorByAddressList) {
        message.validatorByAddressList.push(ValidatorByAddress.fromJSON(e))
      }
    }
    if (object.coordinatorByAddressList !== undefined && object.coordinatorByAddressList !== null) {
      for (const e of object.coordinatorByAddressList) {
        message.coordinatorByAddressList.push(CoordinatorByAddress.fromJSON(e))
      }
    }
    if (object.coordinatorList !== undefined && object.coordinatorList !== null) {
      for (const e of object.coordinatorList) {
        message.coordinatorList.push(Coordinator.fromJSON(e))
      }
    }
    if (object.coordinatorCount !== undefined && object.coordinatorCount !== null) {
      message.coordinatorCount = Number(object.coordinatorCount)
    } else {
      message.coordinatorCount = 0
    }
    return message
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {}
    if (message.consensusKeyNonceList) {
      obj.consensusKeyNonceList = message.consensusKeyNonceList.map((e) => (e ? ConsensusKeyNonce.toJSON(e) : undefined))
    } else {
      obj.consensusKeyNonceList = []
    }
    if (message.validatorByConsAddressList) {
      obj.validatorByConsAddressList = message.validatorByConsAddressList.map((e) => (e ? ValidatorByConsAddress.toJSON(e) : undefined))
    } else {
      obj.validatorByConsAddressList = []
    }
    if (message.validatorByAddressList) {
      obj.validatorByAddressList = message.validatorByAddressList.map((e) => (e ? ValidatorByAddress.toJSON(e) : undefined))
    } else {
      obj.validatorByAddressList = []
    }
    if (message.coordinatorByAddressList) {
      obj.coordinatorByAddressList = message.coordinatorByAddressList.map((e) => (e ? CoordinatorByAddress.toJSON(e) : undefined))
    } else {
      obj.coordinatorByAddressList = []
    }
    if (message.coordinatorList) {
      obj.coordinatorList = message.coordinatorList.map((e) => (e ? Coordinator.toJSON(e) : undefined))
    } else {
      obj.coordinatorList = []
    }
    message.coordinatorCount !== undefined && (obj.coordinatorCount = message.coordinatorCount)
    return obj
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState
    message.consensusKeyNonceList = []
    message.validatorByConsAddressList = []
    message.validatorByAddressList = []
    message.coordinatorByAddressList = []
    message.coordinatorList = []
    if (object.consensusKeyNonceList !== undefined && object.consensusKeyNonceList !== null) {
      for (const e of object.consensusKeyNonceList) {
        message.consensusKeyNonceList.push(ConsensusKeyNonce.fromPartial(e))
      }
    }
    if (object.validatorByConsAddressList !== undefined && object.validatorByConsAddressList !== null) {
      for (const e of object.validatorByConsAddressList) {
        message.validatorByConsAddressList.push(ValidatorByConsAddress.fromPartial(e))
      }
    }
    if (object.validatorByAddressList !== undefined && object.validatorByAddressList !== null) {
      for (const e of object.validatorByAddressList) {
        message.validatorByAddressList.push(ValidatorByAddress.fromPartial(e))
      }
    }
    if (object.coordinatorByAddressList !== undefined && object.coordinatorByAddressList !== null) {
      for (const e of object.coordinatorByAddressList) {
        message.coordinatorByAddressList.push(CoordinatorByAddress.fromPartial(e))
      }
    }
    if (object.coordinatorList !== undefined && object.coordinatorList !== null) {
      for (const e of object.coordinatorList) {
        message.coordinatorList.push(Coordinator.fromPartial(e))
      }
    }
    if (object.coordinatorCount !== undefined && object.coordinatorCount !== null) {
      message.coordinatorCount = object.coordinatorCount
    } else {
      message.coordinatorCount = 0
    }
    return message
  }
}

declare var self: any | undefined
declare var window: any | undefined
var globalThis: any = (() => {
  if (typeof globalThis !== 'undefined') return globalThis
  if (typeof self !== 'undefined') return self
  if (typeof window !== 'undefined') return window
  if (typeof global !== 'undefined') return global
  throw 'Unable to locate global object'
})()

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

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error('Value is larger than Number.MAX_SAFE_INTEGER')
  }
  return long.toNumber()
}

if (util.Long !== Long) {
  util.Long = Long as any
  configure()
}
