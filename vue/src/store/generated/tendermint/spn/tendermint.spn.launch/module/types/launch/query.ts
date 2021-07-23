/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal'
import { GenesisAccount } from '../launch/genesis_account'
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination'
import { Chain } from '../launch/chain'

export const protobufPackage = 'tendermint.spn.launch'

/** this line is used by starport scaffolding # 3 */
export interface QueryGetGenesisAccountRequest {
  chainID: string
  address: string
}

export interface QueryGetGenesisAccountResponse {
  genesisAccount: GenesisAccount | undefined
}

export interface QueryAllGenesisAccountRequest {
  pagination: PageRequest | undefined
}

export interface QueryAllGenesisAccountResponse {
  genesisAccount: GenesisAccount[]
  pagination: PageResponse | undefined
}

export interface QueryGetChainRequest {
  chainID: string
}

export interface QueryGetChainResponse {
  chain: Chain | undefined
}

export interface QueryAllChainRequest {
  pagination: PageRequest | undefined
}

export interface QueryAllChainResponse {
  chain: Chain[]
  pagination: PageResponse | undefined
}

const baseQueryGetGenesisAccountRequest: object = { chainID: '', address: '' }

export const QueryGetGenesisAccountRequest = {
  encode(message: QueryGetGenesisAccountRequest, writer: Writer = Writer.create()): Writer {
    if (message.chainID !== '') {
      writer.uint32(10).string(message.chainID)
    }
    if (message.address !== '') {
      writer.uint32(18).string(message.address)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetGenesisAccountRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetGenesisAccountRequest } as QueryGetGenesisAccountRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.chainID = reader.string()
          break
        case 2:
          message.address = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetGenesisAccountRequest {
    const message = { ...baseQueryGetGenesisAccountRequest } as QueryGetGenesisAccountRequest
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
    return message
  },

  toJSON(message: QueryGetGenesisAccountRequest): unknown {
    const obj: any = {}
    message.chainID !== undefined && (obj.chainID = message.chainID)
    message.address !== undefined && (obj.address = message.address)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetGenesisAccountRequest>): QueryGetGenesisAccountRequest {
    const message = { ...baseQueryGetGenesisAccountRequest } as QueryGetGenesisAccountRequest
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
    return message
  }
}

const baseQueryGetGenesisAccountResponse: object = {}

export const QueryGetGenesisAccountResponse = {
  encode(message: QueryGetGenesisAccountResponse, writer: Writer = Writer.create()): Writer {
    if (message.genesisAccount !== undefined) {
      GenesisAccount.encode(message.genesisAccount, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetGenesisAccountResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetGenesisAccountResponse } as QueryGetGenesisAccountResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.genesisAccount = GenesisAccount.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetGenesisAccountResponse {
    const message = { ...baseQueryGetGenesisAccountResponse } as QueryGetGenesisAccountResponse
    if (object.genesisAccount !== undefined && object.genesisAccount !== null) {
      message.genesisAccount = GenesisAccount.fromJSON(object.genesisAccount)
    } else {
      message.genesisAccount = undefined
    }
    return message
  },

  toJSON(message: QueryGetGenesisAccountResponse): unknown {
    const obj: any = {}
    message.genesisAccount !== undefined && (obj.genesisAccount = message.genesisAccount ? GenesisAccount.toJSON(message.genesisAccount) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetGenesisAccountResponse>): QueryGetGenesisAccountResponse {
    const message = { ...baseQueryGetGenesisAccountResponse } as QueryGetGenesisAccountResponse
    if (object.genesisAccount !== undefined && object.genesisAccount !== null) {
      message.genesisAccount = GenesisAccount.fromPartial(object.genesisAccount)
    } else {
      message.genesisAccount = undefined
    }
    return message
  }
}

const baseQueryAllGenesisAccountRequest: object = {}

export const QueryAllGenesisAccountRequest = {
  encode(message: QueryAllGenesisAccountRequest, writer: Writer = Writer.create()): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllGenesisAccountRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllGenesisAccountRequest } as QueryAllGenesisAccountRequest
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

  fromJSON(object: any): QueryAllGenesisAccountRequest {
    const message = { ...baseQueryAllGenesisAccountRequest } as QueryAllGenesisAccountRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllGenesisAccountRequest): unknown {
    const obj: any = {}
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllGenesisAccountRequest>): QueryAllGenesisAccountRequest {
    const message = { ...baseQueryAllGenesisAccountRequest } as QueryAllGenesisAccountRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseQueryAllGenesisAccountResponse: object = {}

export const QueryAllGenesisAccountResponse = {
  encode(message: QueryAllGenesisAccountResponse, writer: Writer = Writer.create()): Writer {
    for (const v of message.genesisAccount) {
      GenesisAccount.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllGenesisAccountResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllGenesisAccountResponse } as QueryAllGenesisAccountResponse
    message.genesisAccount = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.genesisAccount.push(GenesisAccount.decode(reader, reader.uint32()))
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

  fromJSON(object: any): QueryAllGenesisAccountResponse {
    const message = { ...baseQueryAllGenesisAccountResponse } as QueryAllGenesisAccountResponse
    message.genesisAccount = []
    if (object.genesisAccount !== undefined && object.genesisAccount !== null) {
      for (const e of object.genesisAccount) {
        message.genesisAccount.push(GenesisAccount.fromJSON(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllGenesisAccountResponse): unknown {
    const obj: any = {}
    if (message.genesisAccount) {
      obj.genesisAccount = message.genesisAccount.map((e) => (e ? GenesisAccount.toJSON(e) : undefined))
    } else {
      obj.genesisAccount = []
    }
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllGenesisAccountResponse>): QueryAllGenesisAccountResponse {
    const message = { ...baseQueryAllGenesisAccountResponse } as QueryAllGenesisAccountResponse
    message.genesisAccount = []
    if (object.genesisAccount !== undefined && object.genesisAccount !== null) {
      for (const e of object.genesisAccount) {
        message.genesisAccount.push(GenesisAccount.fromPartial(e))
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

const baseQueryGetChainRequest: object = { chainID: '' }

export const QueryGetChainRequest = {
  encode(message: QueryGetChainRequest, writer: Writer = Writer.create()): Writer {
    if (message.chainID !== '') {
      writer.uint32(10).string(message.chainID)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetChainRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetChainRequest } as QueryGetChainRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.chainID = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetChainRequest {
    const message = { ...baseQueryGetChainRequest } as QueryGetChainRequest
    if (object.chainID !== undefined && object.chainID !== null) {
      message.chainID = String(object.chainID)
    } else {
      message.chainID = ''
    }
    return message
  },

  toJSON(message: QueryGetChainRequest): unknown {
    const obj: any = {}
    message.chainID !== undefined && (obj.chainID = message.chainID)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetChainRequest>): QueryGetChainRequest {
    const message = { ...baseQueryGetChainRequest } as QueryGetChainRequest
    if (object.chainID !== undefined && object.chainID !== null) {
      message.chainID = object.chainID
    } else {
      message.chainID = ''
    }
    return message
  }
}

const baseQueryGetChainResponse: object = {}

export const QueryGetChainResponse = {
  encode(message: QueryGetChainResponse, writer: Writer = Writer.create()): Writer {
    if (message.chain !== undefined) {
      Chain.encode(message.chain, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetChainResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetChainResponse } as QueryGetChainResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.chain = Chain.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetChainResponse {
    const message = { ...baseQueryGetChainResponse } as QueryGetChainResponse
    if (object.chain !== undefined && object.chain !== null) {
      message.chain = Chain.fromJSON(object.chain)
    } else {
      message.chain = undefined
    }
    return message
  },

  toJSON(message: QueryGetChainResponse): unknown {
    const obj: any = {}
    message.chain !== undefined && (obj.chain = message.chain ? Chain.toJSON(message.chain) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetChainResponse>): QueryGetChainResponse {
    const message = { ...baseQueryGetChainResponse } as QueryGetChainResponse
    if (object.chain !== undefined && object.chain !== null) {
      message.chain = Chain.fromPartial(object.chain)
    } else {
      message.chain = undefined
    }
    return message
  }
}

const baseQueryAllChainRequest: object = {}

export const QueryAllChainRequest = {
  encode(message: QueryAllChainRequest, writer: Writer = Writer.create()): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllChainRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllChainRequest } as QueryAllChainRequest
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

  fromJSON(object: any): QueryAllChainRequest {
    const message = { ...baseQueryAllChainRequest } as QueryAllChainRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllChainRequest): unknown {
    const obj: any = {}
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllChainRequest>): QueryAllChainRequest {
    const message = { ...baseQueryAllChainRequest } as QueryAllChainRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseQueryAllChainResponse: object = {}

export const QueryAllChainResponse = {
  encode(message: QueryAllChainResponse, writer: Writer = Writer.create()): Writer {
    for (const v of message.chain) {
      Chain.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllChainResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllChainResponse } as QueryAllChainResponse
    message.chain = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.chain.push(Chain.decode(reader, reader.uint32()))
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

  fromJSON(object: any): QueryAllChainResponse {
    const message = { ...baseQueryAllChainResponse } as QueryAllChainResponse
    message.chain = []
    if (object.chain !== undefined && object.chain !== null) {
      for (const e of object.chain) {
        message.chain.push(Chain.fromJSON(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllChainResponse): unknown {
    const obj: any = {}
    if (message.chain) {
      obj.chain = message.chain.map((e) => (e ? Chain.toJSON(e) : undefined))
    } else {
      obj.chain = []
    }
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllChainResponse>): QueryAllChainResponse {
    const message = { ...baseQueryAllChainResponse } as QueryAllChainResponse
    message.chain = []
    if (object.chain !== undefined && object.chain !== null) {
      for (const e of object.chain) {
        message.chain.push(Chain.fromPartial(e))
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
  /** Queries a genesisAccount by index. */
  GenesisAccount(request: QueryGetGenesisAccountRequest): Promise<QueryGetGenesisAccountResponse>
  /** Queries a list of genesisAccount items. */
  GenesisAccountAll(request: QueryAllGenesisAccountRequest): Promise<QueryAllGenesisAccountResponse>
  /** Queries a chain by index. */
  Chain(request: QueryGetChainRequest): Promise<QueryGetChainResponse>
  /** Queries a list of chain items. */
  ChainAll(request: QueryAllChainRequest): Promise<QueryAllChainResponse>
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc
  constructor(rpc: Rpc) {
    this.rpc = rpc
  }
  GenesisAccount(request: QueryGetGenesisAccountRequest): Promise<QueryGetGenesisAccountResponse> {
    const data = QueryGetGenesisAccountRequest.encode(request).finish()
    const promise = this.rpc.request('tendermint.spn.launch.Query', 'GenesisAccount', data)
    return promise.then((data) => QueryGetGenesisAccountResponse.decode(new Reader(data)))
  }

  GenesisAccountAll(request: QueryAllGenesisAccountRequest): Promise<QueryAllGenesisAccountResponse> {
    const data = QueryAllGenesisAccountRequest.encode(request).finish()
    const promise = this.rpc.request('tendermint.spn.launch.Query', 'GenesisAccountAll', data)
    return promise.then((data) => QueryAllGenesisAccountResponse.decode(new Reader(data)))
  }

  Chain(request: QueryGetChainRequest): Promise<QueryGetChainResponse> {
    const data = QueryGetChainRequest.encode(request).finish()
    const promise = this.rpc.request('tendermint.spn.launch.Query', 'Chain', data)
    return promise.then((data) => QueryGetChainResponse.decode(new Reader(data)))
  }

  ChainAll(request: QueryAllChainRequest): Promise<QueryAllChainResponse> {
    const data = QueryAllChainRequest.encode(request).finish()
    const promise = this.rpc.request('tendermint.spn.launch.Query', 'ChainAll', data)
    return promise.then((data) => QueryAllChainResponse.decode(new Reader(data)))
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>
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
