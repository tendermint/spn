/* eslint-disable */
import { Reader, util, configure, Writer } from 'protobufjs/minimal'
import * as Long from 'long'
import { CoordinatorDescription } from '../profile/coordinator'

export const protobufPackage = 'tendermint.spn.profile'

/** this line is used by starport scaffolding # proto/tx/message */
export interface MsgCreateCoordinator {
  address: string
  description: CoordinatorDescription | undefined
}

export interface MsgCreateCoordinatorResponse {
  coordinatorId: number
}

export interface MsgDeleteCoordinator {
  address: string
}

export interface MsgDeleteCoordinatorResponse {}

const baseMsgCreateCoordinator: object = { address: '' }

export const MsgCreateCoordinator = {
  encode(message: MsgCreateCoordinator, writer: Writer = Writer.create()): Writer {
    if (message.address !== '') {
      writer.uint32(10).string(message.address)
    }
    if (message.description !== undefined) {
      CoordinatorDescription.encode(message.description, writer.uint32(18).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateCoordinator {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgCreateCoordinator } as MsgCreateCoordinator
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.address = reader.string()
          break
        case 2:
          message.description = CoordinatorDescription.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgCreateCoordinator {
    const message = { ...baseMsgCreateCoordinator } as MsgCreateCoordinator
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address)
    } else {
      message.address = ''
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = CoordinatorDescription.fromJSON(object.description)
    } else {
      message.description = undefined
    }
    return message
  },

  toJSON(message: MsgCreateCoordinator): unknown {
    const obj: any = {}
    message.address !== undefined && (obj.address = message.address)
    message.description !== undefined && (obj.description = message.description ? CoordinatorDescription.toJSON(message.description) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<MsgCreateCoordinator>): MsgCreateCoordinator {
    const message = { ...baseMsgCreateCoordinator } as MsgCreateCoordinator
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address
    } else {
      message.address = ''
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = CoordinatorDescription.fromPartial(object.description)
    } else {
      message.description = undefined
    }
    return message
  }
}

const baseMsgCreateCoordinatorResponse: object = { coordinatorId: 0 }

export const MsgCreateCoordinatorResponse = {
  encode(message: MsgCreateCoordinatorResponse, writer: Writer = Writer.create()): Writer {
    if (message.coordinatorId !== 0) {
      writer.uint32(8).uint64(message.coordinatorId)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateCoordinatorResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgCreateCoordinatorResponse } as MsgCreateCoordinatorResponse
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

  fromJSON(object: any): MsgCreateCoordinatorResponse {
    const message = { ...baseMsgCreateCoordinatorResponse } as MsgCreateCoordinatorResponse
    if (object.coordinatorId !== undefined && object.coordinatorId !== null) {
      message.coordinatorId = Number(object.coordinatorId)
    } else {
      message.coordinatorId = 0
    }
    return message
  },

  toJSON(message: MsgCreateCoordinatorResponse): unknown {
    const obj: any = {}
    message.coordinatorId !== undefined && (obj.coordinatorId = message.coordinatorId)
    return obj
  },

  fromPartial(object: DeepPartial<MsgCreateCoordinatorResponse>): MsgCreateCoordinatorResponse {
    const message = { ...baseMsgCreateCoordinatorResponse } as MsgCreateCoordinatorResponse
    if (object.coordinatorId !== undefined && object.coordinatorId !== null) {
      message.coordinatorId = object.coordinatorId
    } else {
      message.coordinatorId = 0
    }
    return message
  }
}

const baseMsgDeleteCoordinator: object = { address: '' }

export const MsgDeleteCoordinator = {
  encode(message: MsgDeleteCoordinator, writer: Writer = Writer.create()): Writer {
    if (message.address !== '') {
      writer.uint32(10).string(message.address)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeleteCoordinator {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgDeleteCoordinator } as MsgDeleteCoordinator
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.address = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgDeleteCoordinator {
    const message = { ...baseMsgDeleteCoordinator } as MsgDeleteCoordinator
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address)
    } else {
      message.address = ''
    }
    return message
  },

  toJSON(message: MsgDeleteCoordinator): unknown {
    const obj: any = {}
    message.address !== undefined && (obj.address = message.address)
    return obj
  },

  fromPartial(object: DeepPartial<MsgDeleteCoordinator>): MsgDeleteCoordinator {
    const message = { ...baseMsgDeleteCoordinator } as MsgDeleteCoordinator
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address
    } else {
      message.address = ''
    }
    return message
  }
}

const baseMsgDeleteCoordinatorResponse: object = {}

export const MsgDeleteCoordinatorResponse = {
  encode(_: MsgDeleteCoordinatorResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeleteCoordinatorResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgDeleteCoordinatorResponse } as MsgDeleteCoordinatorResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(_: any): MsgDeleteCoordinatorResponse {
    const message = { ...baseMsgDeleteCoordinatorResponse } as MsgDeleteCoordinatorResponse
    return message
  },

  toJSON(_: MsgDeleteCoordinatorResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgDeleteCoordinatorResponse>): MsgDeleteCoordinatorResponse {
    const message = { ...baseMsgDeleteCoordinatorResponse } as MsgDeleteCoordinatorResponse
    return message
  }
}

/** Msg defines the Msg service. */
export interface Msg {
  /** this line is used by starport scaffolding # proto/tx/rpc */
  CreateCoordinator(request: MsgCreateCoordinator): Promise<MsgCreateCoordinatorResponse>
  DeleteCoordinator(request: MsgDeleteCoordinator): Promise<MsgDeleteCoordinatorResponse>
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc
  constructor(rpc: Rpc) {
    this.rpc = rpc
  }
  CreateCoordinator(request: MsgCreateCoordinator): Promise<MsgCreateCoordinatorResponse> {
    const data = MsgCreateCoordinator.encode(request).finish()
    const promise = this.rpc.request('tendermint.spn.profile.Msg', 'CreateCoordinator', data)
    return promise.then((data) => MsgCreateCoordinatorResponse.decode(new Reader(data)))
  }

  DeleteCoordinator(request: MsgDeleteCoordinator): Promise<MsgDeleteCoordinatorResponse> {
    const data = MsgDeleteCoordinator.encode(request).finish()
    const promise = this.rpc.request('tendermint.spn.profile.Msg', 'DeleteCoordinator', data)
    return promise.then((data) => MsgDeleteCoordinatorResponse.decode(new Reader(data)))
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
