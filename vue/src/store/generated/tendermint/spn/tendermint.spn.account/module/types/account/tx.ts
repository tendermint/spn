/* eslint-disable */
import { Reader, util, configure, Writer } from 'protobufjs/minimal'
import * as Long from 'long'

export const protobufPackage = 'tendermint.spn.account'

/** this line is used by starport scaffolding # proto/tx/message */
export interface MsgCreateCoordinator {
  creator: string
  address: string
  identity: string
  website: string
  details: string
}

export interface MsgCreateCoordinatorResponse {
  coordinatorId: number
}

const baseMsgCreateCoordinator: object = { creator: '', address: '', identity: '', website: '', details: '' }

export const MsgCreateCoordinator = {
  encode(message: MsgCreateCoordinator, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.address !== '') {
      writer.uint32(18).string(message.address)
    }
    if (message.identity !== '') {
      writer.uint32(26).string(message.identity)
    }
    if (message.website !== '') {
      writer.uint32(34).string(message.website)
    }
    if (message.details !== '') {
      writer.uint32(42).string(message.details)
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
          message.creator = reader.string()
          break
        case 2:
          message.address = reader.string()
          break
        case 3:
          message.identity = reader.string()
          break
        case 4:
          message.website = reader.string()
          break
        case 5:
          message.details = reader.string()
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
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator)
    } else {
      message.creator = ''
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address)
    } else {
      message.address = ''
    }
    if (object.identity !== undefined && object.identity !== null) {
      message.identity = String(object.identity)
    } else {
      message.identity = ''
    }
    if (object.website !== undefined && object.website !== null) {
      message.website = String(object.website)
    } else {
      message.website = ''
    }
    if (object.details !== undefined && object.details !== null) {
      message.details = String(object.details)
    } else {
      message.details = ''
    }
    return message
  },

  toJSON(message: MsgCreateCoordinator): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.address !== undefined && (obj.address = message.address)
    message.identity !== undefined && (obj.identity = message.identity)
    message.website !== undefined && (obj.website = message.website)
    message.details !== undefined && (obj.details = message.details)
    return obj
  },

  fromPartial(object: DeepPartial<MsgCreateCoordinator>): MsgCreateCoordinator {
    const message = { ...baseMsgCreateCoordinator } as MsgCreateCoordinator
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator
    } else {
      message.creator = ''
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address
    } else {
      message.address = ''
    }
    if (object.identity !== undefined && object.identity !== null) {
      message.identity = object.identity
    } else {
      message.identity = ''
    }
    if (object.website !== undefined && object.website !== null) {
      message.website = object.website
    } else {
      message.website = ''
    }
    if (object.details !== undefined && object.details !== null) {
      message.details = object.details
    } else {
      message.details = ''
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

/** Msg defines the Msg service. */
export interface Msg {
  /** this line is used by starport scaffolding # proto/tx/rpc */
  CreateCoordinator(request: MsgCreateCoordinator): Promise<MsgCreateCoordinatorResponse>
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc
  constructor(rpc: Rpc) {
    this.rpc = rpc
  }
  CreateCoordinator(request: MsgCreateCoordinator): Promise<MsgCreateCoordinatorResponse> {
    const data = MsgCreateCoordinator.encode(request).finish()
    const promise = this.rpc.request('tendermint.spn.account.Msg', 'CreateCoordinator', data)
    return promise.then((data) => MsgCreateCoordinatorResponse.decode(new Reader(data)))
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
