/* eslint-disable */
import { Reader, util, configure, Writer } from 'protobufjs/minimal'
import * as Long from 'long'

export const protobufPackage = 'tendermint.spn.profile'

/** this line is used by starport scaffolding # proto/tx/message */
export interface MsgUpdateCoordinatorAddress {
  address: string
  newAddress: string
}

export interface MsgUpdateCoordinatorAddressResponse {
  coordinatorId: number
}

const baseMsgUpdateCoordinatorAddress: object = { address: '', newAddress: '' }

export const MsgUpdateCoordinatorAddress = {
  encode(message: MsgUpdateCoordinatorAddress, writer: Writer = Writer.create()): Writer {
    if (message.address !== '') {
      writer.uint32(10).string(message.address)
    }
    if (message.newAddress !== '') {
      writer.uint32(18).string(message.newAddress)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateCoordinatorAddress {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgUpdateCoordinatorAddress } as MsgUpdateCoordinatorAddress
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.address = reader.string()
          break
        case 2:
          message.newAddress = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgUpdateCoordinatorAddress {
    const message = { ...baseMsgUpdateCoordinatorAddress } as MsgUpdateCoordinatorAddress
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address)
    } else {
      message.address = ''
    }
    if (object.newAddress !== undefined && object.newAddress !== null) {
      message.newAddress = String(object.newAddress)
    } else {
      message.newAddress = ''
    }
    return message
  },

  toJSON(message: MsgUpdateCoordinatorAddress): unknown {
    const obj: any = {}
    message.address !== undefined && (obj.address = message.address)
    message.newAddress !== undefined && (obj.newAddress = message.newAddress)
    return obj
  },

  fromPartial(object: DeepPartial<MsgUpdateCoordinatorAddress>): MsgUpdateCoordinatorAddress {
    const message = { ...baseMsgUpdateCoordinatorAddress } as MsgUpdateCoordinatorAddress
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address
    } else {
      message.address = ''
    }
    if (object.newAddress !== undefined && object.newAddress !== null) {
      message.newAddress = object.newAddress
    } else {
      message.newAddress = ''
    }
    return message
  }
}

const baseMsgUpdateCoordinatorAddressResponse: object = { coordinatorId: 0 }

export const MsgUpdateCoordinatorAddressResponse = {
  encode(message: MsgUpdateCoordinatorAddressResponse, writer: Writer = Writer.create()): Writer {
    if (message.coordinatorId !== 0) {
      writer.uint32(8).uint64(message.coordinatorId)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateCoordinatorAddressResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgUpdateCoordinatorAddressResponse } as MsgUpdateCoordinatorAddressResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.coordinatorId = longToNumber(reader.uint64() as Long)
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgUpdateCoordinatorAddressResponse {
    const message = { ...baseMsgUpdateCoordinatorAddressResponse } as MsgUpdateCoordinatorAddressResponse
    if (object.coordinatorId !== undefined && object.coordinatorId !== null) {
      message.coordinatorId = Number(object.coordinatorId)
    } else {
      message.coordinatorId = 0
    }
    return message
  },

  toJSON(message: MsgUpdateCoordinatorAddressResponse): unknown {
    const obj: any = {}
    message.coordinatorId !== undefined && (obj.coordinatorId = message.coordinatorId)
    return obj
  },

  fromPartial(object: DeepPartial<MsgUpdateCoordinatorAddressResponse>): MsgUpdateCoordinatorAddressResponse {
    const message = { ...baseMsgUpdateCoordinatorAddressResponse } as MsgUpdateCoordinatorAddressResponse
    if (object.coordinatorId !== undefined && object.coordinatorId !== null) {
      message.coordinatorId = object.coordinatorId
    } else {
      message.coordinatorId = 0
    }
    return message
  }
}

/** Msg defines the Msg service. */
export interface Msg {
  /** this line is used by starport scaffolding # proto/tx/rpc */
  UpdateCoordinatorAddress(request: MsgUpdateCoordinatorAddress): Promise<MsgUpdateCoordinatorAddressResponse>
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc
  constructor(rpc: Rpc) {
    this.rpc = rpc
  }
  UpdateCoordinatorAddress(request: MsgUpdateCoordinatorAddress): Promise<MsgUpdateCoordinatorAddressResponse> {
    const data = MsgUpdateCoordinatorAddress.encode(request).finish()
    const promise = this.rpc.request('tendermint.spn.profile.Msg', 'UpdateCoordinatorAddress', data)
    return promise.then((data) => MsgUpdateCoordinatorAddressResponse.decode(new Reader(data)))
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>
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
