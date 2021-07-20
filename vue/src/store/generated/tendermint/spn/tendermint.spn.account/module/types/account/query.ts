/* eslint-disable */
import { Reader, util, configure, Writer } from 'protobufjs/minimal'
import * as Long from 'long'
import { CoordinatorByAddress, Coordinator } from '../account/coordinator'
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination'

export const protobufPackage = 'tendermint.spn.account'

/** this line is used by starport scaffolding # 3 */
export interface QueryGetCoordinatorByAddressRequest {
  address: string
}

export interface QueryGetCoordinatorByAddressResponse {
  coordinatorByAddress: CoordinatorByAddress | undefined
}

export interface QueryAllCoordinatorByAddressRequest {
  pagination: PageRequest | undefined
}

export interface QueryAllCoordinatorByAddressResponse {
  coordinatorByAddress: CoordinatorByAddress[]
  pagination: PageResponse | undefined
}

export interface QueryGetCoordinatorRequest {
  id: number
}

export interface QueryGetCoordinatorResponse {
  Coordinator: Coordinator | undefined
}

export interface QueryAllCoordinatorRequest {
  pagination: PageRequest | undefined
}

export interface QueryAllCoordinatorResponse {
  Coordinator: Coordinator[]
  pagination: PageResponse | undefined
}

const baseQueryGetCoordinatorByAddressRequest: object = { address: '' }

export const QueryGetCoordinatorByAddressRequest = {
  encode(message: QueryGetCoordinatorByAddressRequest, writer: Writer = Writer.create()): Writer {
    if (message.address !== '') {
      writer.uint32(10).string(message.address)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetCoordinatorByAddressRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetCoordinatorByAddressRequest } as QueryGetCoordinatorByAddressRequest
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

  fromJSON(object: any): QueryGetCoordinatorByAddressRequest {
    const message = { ...baseQueryGetCoordinatorByAddressRequest } as QueryGetCoordinatorByAddressRequest
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address)
    } else {
      message.address = ''
    }
    return message
  },

  toJSON(message: QueryGetCoordinatorByAddressRequest): unknown {
    const obj: any = {}
    message.address !== undefined && (obj.address = message.address)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetCoordinatorByAddressRequest>): QueryGetCoordinatorByAddressRequest {
    const message = { ...baseQueryGetCoordinatorByAddressRequest } as QueryGetCoordinatorByAddressRequest
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address
    } else {
      message.address = ''
    }
    return message
  }
}

const baseQueryGetCoordinatorByAddressResponse: object = {}

export const QueryGetCoordinatorByAddressResponse = {
  encode(message: QueryGetCoordinatorByAddressResponse, writer: Writer = Writer.create()): Writer {
    if (message.coordinatorByAddress !== undefined) {
      CoordinatorByAddress.encode(message.coordinatorByAddress, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetCoordinatorByAddressResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetCoordinatorByAddressResponse } as QueryGetCoordinatorByAddressResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.coordinatorByAddress = CoordinatorByAddress.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetCoordinatorByAddressResponse {
    const message = { ...baseQueryGetCoordinatorByAddressResponse } as QueryGetCoordinatorByAddressResponse
    if (object.coordinatorByAddress !== undefined && object.coordinatorByAddress !== null) {
      message.coordinatorByAddress = CoordinatorByAddress.fromJSON(object.coordinatorByAddress)
    } else {
      message.coordinatorByAddress = undefined
    }
    return message
  },

  toJSON(message: QueryGetCoordinatorByAddressResponse): unknown {
    const obj: any = {}
    message.coordinatorByAddress !== undefined &&
      (obj.coordinatorByAddress = message.coordinatorByAddress ? CoordinatorByAddress.toJSON(message.coordinatorByAddress) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetCoordinatorByAddressResponse>): QueryGetCoordinatorByAddressResponse {
    const message = { ...baseQueryGetCoordinatorByAddressResponse } as QueryGetCoordinatorByAddressResponse
    if (object.coordinatorByAddress !== undefined && object.coordinatorByAddress !== null) {
      message.coordinatorByAddress = CoordinatorByAddress.fromPartial(object.coordinatorByAddress)
    } else {
      message.coordinatorByAddress = undefined
    }
    return message
  }
}

const baseQueryAllCoordinatorByAddressRequest: object = {}

export const QueryAllCoordinatorByAddressRequest = {
  encode(message: QueryAllCoordinatorByAddressRequest, writer: Writer = Writer.create()): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllCoordinatorByAddressRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllCoordinatorByAddressRequest } as QueryAllCoordinatorByAddressRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryAllCoordinatorByAddressRequest {
    const message = { ...baseQueryAllCoordinatorByAddressRequest } as QueryAllCoordinatorByAddressRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllCoordinatorByAddressRequest): unknown {
    const obj: any = {}
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllCoordinatorByAddressRequest>): QueryAllCoordinatorByAddressRequest {
    const message = { ...baseQueryAllCoordinatorByAddressRequest } as QueryAllCoordinatorByAddressRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseQueryAllCoordinatorByAddressResponse: object = {}

export const QueryAllCoordinatorByAddressResponse = {
  encode(message: QueryAllCoordinatorByAddressResponse, writer: Writer = Writer.create()): Writer {
    for (const v of message.coordinatorByAddress) {
      CoordinatorByAddress.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllCoordinatorByAddressResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllCoordinatorByAddressResponse } as QueryAllCoordinatorByAddressResponse
    message.coordinatorByAddress = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.coordinatorByAddress.push(CoordinatorByAddress.decode(reader, reader.uint32()))
          break
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryAllCoordinatorByAddressResponse {
    const message = { ...baseQueryAllCoordinatorByAddressResponse } as QueryAllCoordinatorByAddressResponse
    message.coordinatorByAddress = []
    if (object.coordinatorByAddress !== undefined && object.coordinatorByAddress !== null) {
      for (const e of object.coordinatorByAddress) {
        message.coordinatorByAddress.push(CoordinatorByAddress.fromJSON(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllCoordinatorByAddressResponse): unknown {
    const obj: any = {}
    if (message.coordinatorByAddress) {
      obj.coordinatorByAddress = message.coordinatorByAddress.map((e) => (e ? CoordinatorByAddress.toJSON(e) : undefined))
    } else {
      obj.coordinatorByAddress = []
    }
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllCoordinatorByAddressResponse>): QueryAllCoordinatorByAddressResponse {
    const message = { ...baseQueryAllCoordinatorByAddressResponse } as QueryAllCoordinatorByAddressResponse
    message.coordinatorByAddress = []
    if (object.coordinatorByAddress !== undefined && object.coordinatorByAddress !== null) {
      for (const e of object.coordinatorByAddress) {
        message.coordinatorByAddress.push(CoordinatorByAddress.fromPartial(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseQueryGetCoordinatorRequest: object = { id: 0 }

export const QueryGetCoordinatorRequest = {
  encode(message: QueryGetCoordinatorRequest, writer: Writer = Writer.create()): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetCoordinatorRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetCoordinatorRequest } as QueryGetCoordinatorRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long)
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetCoordinatorRequest {
    const message = { ...baseQueryGetCoordinatorRequest } as QueryGetCoordinatorRequest
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id)
    } else {
      message.id = 0
    }
    return message
  },

  toJSON(message: QueryGetCoordinatorRequest): unknown {
    const obj: any = {}
    message.id !== undefined && (obj.id = message.id)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetCoordinatorRequest>): QueryGetCoordinatorRequest {
    const message = { ...baseQueryGetCoordinatorRequest } as QueryGetCoordinatorRequest
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id
    } else {
      message.id = 0
    }
    return message
  }
}

const baseQueryGetCoordinatorResponse: object = {}

export const QueryGetCoordinatorResponse = {
  encode(message: QueryGetCoordinatorResponse, writer: Writer = Writer.create()): Writer {
    if (message.Coordinator !== undefined) {
      Coordinator.encode(message.Coordinator, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetCoordinatorResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetCoordinatorResponse } as QueryGetCoordinatorResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.Coordinator = Coordinator.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetCoordinatorResponse {
    const message = { ...baseQueryGetCoordinatorResponse } as QueryGetCoordinatorResponse
    if (object.Coordinator !== undefined && object.Coordinator !== null) {
      message.Coordinator = Coordinator.fromJSON(object.Coordinator)
    } else {
      message.Coordinator = undefined
    }
    return message
  },

  toJSON(message: QueryGetCoordinatorResponse): unknown {
    const obj: any = {}
    message.Coordinator !== undefined && (obj.Coordinator = message.Coordinator ? Coordinator.toJSON(message.Coordinator) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetCoordinatorResponse>): QueryGetCoordinatorResponse {
    const message = { ...baseQueryGetCoordinatorResponse } as QueryGetCoordinatorResponse
    if (object.Coordinator !== undefined && object.Coordinator !== null) {
      message.Coordinator = Coordinator.fromPartial(object.Coordinator)
    } else {
      message.Coordinator = undefined
    }
    return message
  }
}

const baseQueryAllCoordinatorRequest: object = {}

export const QueryAllCoordinatorRequest = {
  encode(message: QueryAllCoordinatorRequest, writer: Writer = Writer.create()): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllCoordinatorRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllCoordinatorRequest } as QueryAllCoordinatorRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryAllCoordinatorRequest {
    const message = { ...baseQueryAllCoordinatorRequest } as QueryAllCoordinatorRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllCoordinatorRequest): unknown {
    const obj: any = {}
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllCoordinatorRequest>): QueryAllCoordinatorRequest {
    const message = { ...baseQueryAllCoordinatorRequest } as QueryAllCoordinatorRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseQueryAllCoordinatorResponse: object = {}

export const QueryAllCoordinatorResponse = {
  encode(message: QueryAllCoordinatorResponse, writer: Writer = Writer.create()): Writer {
    for (const v of message.Coordinator) {
      Coordinator.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllCoordinatorResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllCoordinatorResponse } as QueryAllCoordinatorResponse
    message.Coordinator = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.Coordinator.push(Coordinator.decode(reader, reader.uint32()))
          break
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryAllCoordinatorResponse {
    const message = { ...baseQueryAllCoordinatorResponse } as QueryAllCoordinatorResponse
    message.Coordinator = []
    if (object.Coordinator !== undefined && object.Coordinator !== null) {
      for (const e of object.Coordinator) {
        message.Coordinator.push(Coordinator.fromJSON(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllCoordinatorResponse): unknown {
    const obj: any = {}
    if (message.Coordinator) {
      obj.Coordinator = message.Coordinator.map((e) => (e ? Coordinator.toJSON(e) : undefined))
    } else {
      obj.Coordinator = []
    }
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllCoordinatorResponse>): QueryAllCoordinatorResponse {
    const message = { ...baseQueryAllCoordinatorResponse } as QueryAllCoordinatorResponse
    message.Coordinator = []
    if (object.Coordinator !== undefined && object.Coordinator !== null) {
      for (const e of object.Coordinator) {
        message.Coordinator.push(Coordinator.fromPartial(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

/** Query defines the gRPC querier service. */
export interface Query {
  /** Queries a coordinatorByAddress by index. */
  CoordinatorByAddress(request: QueryGetCoordinatorByAddressRequest): Promise<QueryGetCoordinatorByAddressResponse>
  /** Queries a coordinator by id. */
  Coordinator(request: QueryGetCoordinatorRequest): Promise<QueryGetCoordinatorResponse>
  /** Queries a list of coordinator items. */
  CoordinatorAll(request: QueryAllCoordinatorRequest): Promise<QueryAllCoordinatorResponse>
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc
  constructor(rpc: Rpc) {
    this.rpc = rpc
  }
  CoordinatorByAddress(request: QueryGetCoordinatorByAddressRequest): Promise<QueryGetCoordinatorByAddressResponse> {
    const data = QueryGetCoordinatorByAddressRequest.encode(request).finish()
    const promise = this.rpc.request('tendermint.spn.account.Query', 'CoordinatorByAddress', data)
    return promise.then((data) => QueryGetCoordinatorByAddressResponse.decode(new Reader(data)))
  }

  Coordinator(request: QueryGetCoordinatorRequest): Promise<QueryGetCoordinatorResponse> {
    const data = QueryGetCoordinatorRequest.encode(request).finish()
    const promise = this.rpc.request('tendermint.spn.account.Query', 'Coordinator', data)
    return promise.then((data) => QueryGetCoordinatorResponse.decode(new Reader(data)))
  }

  CoordinatorAll(request: QueryAllCoordinatorRequest): Promise<QueryAllCoordinatorResponse> {
    const data = QueryAllCoordinatorRequest.encode(request).finish()
    const promise = this.rpc.request('tendermint.spn.account.Query', 'CoordinatorAll', data)
    return promise.then((data) => QueryAllCoordinatorResponse.decode(new Reader(data)))
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
