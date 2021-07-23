/* eslint-disable */
import { Reader, util, configure, Writer } from 'protobufjs/minimal'
import * as Long from 'long'
import { CoordinatorDescription } from '../profile/coordinator'

export const protobufPackage = 'tendermint.spn.profile'

/** this line is used by starport scaffolding # proto/tx/message */
export interface MsgUpdateValidatorDescription {
  creator: string
  address: string
  description: string
}

export interface MsgUpdateValidatorDescriptionResponse {
  coordinatorId: number
}

export interface MsgCreateCoordinator {
  address: string
  description: CoordinatorDescription | undefined
}

export interface MsgCreateCoordinatorResponse {
  coordinatorId: number
}

const baseMsgUpdateValidatorDescription: object = { creator: '', address: '', description: '' }

export const MsgUpdateValidatorDescription = {
  encode(message: MsgUpdateValidatorDescription, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.address !== '') {
      writer.uint32(18).string(message.address)
    }
    if (message.description !== '') {
      writer.uint32(26).string(message.description)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateValidatorDescription {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgUpdateValidatorDescription } as MsgUpdateValidatorDescription
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
          message.description = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgUpdateValidatorDescription {
    const message = { ...baseMsgUpdateValidatorDescription } as MsgUpdateValidatorDescription
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
    if (object.description !== undefined && object.description !== null) {
      message.description = String(object.description)
    } else {
      message.description = ''
    }
    return message
  },

  toJSON(message: MsgUpdateValidatorDescription): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.address !== undefined && (obj.address = message.address)
    message.description !== undefined && (obj.description = message.description)
    return obj
  },

  fromPartial(object: DeepPartial<MsgUpdateValidatorDescription>): MsgUpdateValidatorDescription {
    const message = { ...baseMsgUpdateValidatorDescription } as MsgUpdateValidatorDescription
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
    if (object.description !== undefined && object.description !== null) {
      message.description = object.description
    } else {
      message.description = ''
    }
    return message
  }
}

const baseMsgUpdateValidatorDescriptionResponse: object = { coordinatorId: 0 }

export const MsgUpdateValidatorDescriptionResponse = {
  encode(message: MsgUpdateValidatorDescriptionResponse, writer: Writer = Writer.create()): Writer {
    if (message.coordinatorId !== 0) {
      writer.uint32(8).uint64(message.coordinatorId)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateValidatorDescriptionResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgUpdateValidatorDescriptionResponse } as MsgUpdateValidatorDescriptionResponse
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

  fromJSON(object: any): MsgUpdateValidatorDescriptionResponse {
    const message = { ...baseMsgUpdateValidatorDescriptionResponse } as MsgUpdateValidatorDescriptionResponse
    if (object.coordinatorId !== undefined && object.coordinatorId !== null) {
      message.coordinatorId = Number(object.coordinatorId)
    } else {
      message.coordinatorId = 0
    }
    return message
  },

  toJSON(message: MsgUpdateValidatorDescriptionResponse): unknown {
    const obj: any = {}
    message.coordinatorId !== undefined && (obj.coordinatorId = message.coordinatorId)
    return obj
  },

  fromPartial(object: DeepPartial<MsgUpdateValidatorDescriptionResponse>): MsgUpdateValidatorDescriptionResponse {
    const message = { ...baseMsgUpdateValidatorDescriptionResponse } as MsgUpdateValidatorDescriptionResponse
    if (object.coordinatorId !== undefined && object.coordinatorId !== null) {
      message.coordinatorId = object.coordinatorId
    } else {
      message.coordinatorId = 0
    }
    return message
  }
}

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

/** Msg defines the Msg service. */
export interface Msg {
  /** this line is used by starport scaffolding # proto/tx/rpc */
  UpdateValidatorDescription(request: MsgUpdateValidatorDescription): Promise<MsgUpdateValidatorDescriptionResponse>
  CreateCoordinator(request: MsgCreateCoordinator): Promise<MsgCreateCoordinatorResponse>
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc
  constructor(rpc: Rpc) {
    this.rpc = rpc
  }
  UpdateValidatorDescription(request: MsgUpdateValidatorDescription): Promise<MsgUpdateValidatorDescriptionResponse> {
    const data = MsgUpdateValidatorDescription.encode(request).finish()
    const promise = this.rpc.request('tendermint.spn.profile.Msg', 'UpdateValidatorDescription', data)
    return promise.then((data) => MsgUpdateValidatorDescriptionResponse.decode(new Reader(data)))
  }

  CreateCoordinator(request: MsgCreateCoordinator): Promise<MsgCreateCoordinatorResponse> {
    const data = MsgCreateCoordinator.encode(request).finish()
    const promise = this.rpc.request('tendermint.spn.profile.Msg', 'CreateCoordinator', data)
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
