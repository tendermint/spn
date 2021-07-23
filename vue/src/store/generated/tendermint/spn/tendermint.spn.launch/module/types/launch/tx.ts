/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal'

export const protobufPackage = 'tendermint.spn.launch'

/** this line is used by starport scaffolding # proto/tx/message */
export interface MsgCreateChain {
  creator: string
  chainName: string
  sourceURL: string
  sourceHash: string
  genesisURL: string
  genesisHash: string
}

export interface MsgCreateChainResponse {
  chainID: string
}

const baseMsgCreateChain: object = { creator: '', chainName: '', sourceURL: '', sourceHash: '', genesisURL: '', genesisHash: '' }

export const MsgCreateChain = {
  encode(message: MsgCreateChain, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.chainName !== '') {
      writer.uint32(18).string(message.chainName)
    }
    if (message.sourceURL !== '') {
      writer.uint32(26).string(message.sourceURL)
    }
    if (message.sourceHash !== '') {
      writer.uint32(34).string(message.sourceHash)
    }
    if (message.genesisURL !== '') {
      writer.uint32(42).string(message.genesisURL)
    }
    if (message.genesisHash !== '') {
      writer.uint32(50).string(message.genesisHash)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateChain {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgCreateChain } as MsgCreateChain
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string()
          break
        case 2:
          message.chainName = reader.string()
          break
        case 3:
          message.sourceURL = reader.string()
          break
        case 4:
          message.sourceHash = reader.string()
          break
        case 5:
          message.genesisURL = reader.string()
          break
        case 6:
          message.genesisHash = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgCreateChain {
    const message = { ...baseMsgCreateChain } as MsgCreateChain
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator)
    } else {
      message.creator = ''
    }
    if (object.chainName !== undefined && object.chainName !== null) {
      message.chainName = String(object.chainName)
    } else {
      message.chainName = ''
    }
    if (object.sourceURL !== undefined && object.sourceURL !== null) {
      message.sourceURL = String(object.sourceURL)
    } else {
      message.sourceURL = ''
    }
    if (object.sourceHash !== undefined && object.sourceHash !== null) {
      message.sourceHash = String(object.sourceHash)
    } else {
      message.sourceHash = ''
    }
    if (object.genesisURL !== undefined && object.genesisURL !== null) {
      message.genesisURL = String(object.genesisURL)
    } else {
      message.genesisURL = ''
    }
    if (object.genesisHash !== undefined && object.genesisHash !== null) {
      message.genesisHash = String(object.genesisHash)
    } else {
      message.genesisHash = ''
    }
    return message
  },

  toJSON(message: MsgCreateChain): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.chainName !== undefined && (obj.chainName = message.chainName)
    message.sourceURL !== undefined && (obj.sourceURL = message.sourceURL)
    message.sourceHash !== undefined && (obj.sourceHash = message.sourceHash)
    message.genesisURL !== undefined && (obj.genesisURL = message.genesisURL)
    message.genesisHash !== undefined && (obj.genesisHash = message.genesisHash)
    return obj
  },

  fromPartial(object: DeepPartial<MsgCreateChain>): MsgCreateChain {
    const message = { ...baseMsgCreateChain } as MsgCreateChain
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator
    } else {
      message.creator = ''
    }
    if (object.chainName !== undefined && object.chainName !== null) {
      message.chainName = object.chainName
    } else {
      message.chainName = ''
    }
    if (object.sourceURL !== undefined && object.sourceURL !== null) {
      message.sourceURL = object.sourceURL
    } else {
      message.sourceURL = ''
    }
    if (object.sourceHash !== undefined && object.sourceHash !== null) {
      message.sourceHash = object.sourceHash
    } else {
      message.sourceHash = ''
    }
    if (object.genesisURL !== undefined && object.genesisURL !== null) {
      message.genesisURL = object.genesisURL
    } else {
      message.genesisURL = ''
    }
    if (object.genesisHash !== undefined && object.genesisHash !== null) {
      message.genesisHash = object.genesisHash
    } else {
      message.genesisHash = ''
    }
    return message
  }
}

const baseMsgCreateChainResponse: object = { chainID: '' }

export const MsgCreateChainResponse = {
  encode(message: MsgCreateChainResponse, writer: Writer = Writer.create()): Writer {
    if (message.chainID !== '') {
      writer.uint32(10).string(message.chainID)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateChainResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgCreateChainResponse } as MsgCreateChainResponse
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

  fromJSON(object: any): MsgCreateChainResponse {
    const message = { ...baseMsgCreateChainResponse } as MsgCreateChainResponse
    if (object.chainID !== undefined && object.chainID !== null) {
      message.chainID = String(object.chainID)
    } else {
      message.chainID = ''
    }
    return message
  },

  toJSON(message: MsgCreateChainResponse): unknown {
    const obj: any = {}
    message.chainID !== undefined && (obj.chainID = message.chainID)
    return obj
  },

  fromPartial(object: DeepPartial<MsgCreateChainResponse>): MsgCreateChainResponse {
    const message = { ...baseMsgCreateChainResponse } as MsgCreateChainResponse
    if (object.chainID !== undefined && object.chainID !== null) {
      message.chainID = object.chainID
    } else {
      message.chainID = ''
    }
    return message
  }
}

/** Msg defines the Msg service. */
export interface Msg {
  /** this line is used by starport scaffolding # proto/tx/rpc */
  CreateChain(request: MsgCreateChain): Promise<MsgCreateChainResponse>
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc
  constructor(rpc: Rpc) {
    this.rpc = rpc
  }
  CreateChain(request: MsgCreateChain): Promise<MsgCreateChainResponse> {
    const data = MsgCreateChain.encode(request).finish()
    const promise = this.rpc.request('tendermint.spn.launch.Msg', 'CreateChain', data)
    return promise.then((data) => MsgCreateChainResponse.decode(new Reader(data)))
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
