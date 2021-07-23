/* eslint-disable */
import * as Long from 'long'
import { util, configure, Writer, Reader } from 'protobufjs/minimal'
import { Any } from '../google/protobuf/any'

export const protobufPackage = 'tendermint.spn.launch'

export interface Chain {
  chainID: string
  coordinatorID: number
  createdAt: number
  sourceURL: string
  sourceHash: string
  initialGenesis: Any | undefined
  launchTriggered: boolean
  launchTimestamp: number
}

/** DefaultInitialGenesis specifies using the default CLI-generated genesis as an initial genesis */
export interface DefaultInitialGenesis {}

/** GenesisURL specifies using a custom genesis from a URL as the initial genesis */
export interface GenesisURL {
  url: string
  hash: string
}

const baseChain: object = { chainID: '', coordinatorID: 0, createdAt: 0, sourceURL: '', sourceHash: '', launchTriggered: false, launchTimestamp: 0 }

export const Chain = {
  encode(message: Chain, writer: Writer = Writer.create()): Writer {
    if (message.chainID !== '') {
      writer.uint32(10).string(message.chainID)
    }
    if (message.coordinatorID !== 0) {
      writer.uint32(16).uint64(message.coordinatorID)
    }
    if (message.createdAt !== 0) {
      writer.uint32(24).int64(message.createdAt)
    }
    if (message.sourceURL !== '') {
      writer.uint32(34).string(message.sourceURL)
    }
    if (message.sourceHash !== '') {
      writer.uint32(42).string(message.sourceHash)
    }
    if (message.initialGenesis !== undefined) {
      Any.encode(message.initialGenesis, writer.uint32(50).fork()).ldelim()
    }
    if (message.launchTriggered === true) {
      writer.uint32(56).bool(message.launchTriggered)
    }
    if (message.launchTimestamp !== 0) {
      writer.uint32(64).int64(message.launchTimestamp)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): Chain {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseChain } as Chain
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.chainID = reader.string()
          break
        case 2:
          message.coordinatorID = longToNumber(reader.uint64() as Long)
          break
        case 3:
          message.createdAt = longToNumber(reader.int64() as Long)
          break
        case 4:
          message.sourceURL = reader.string()
          break
        case 5:
          message.sourceHash = reader.string()
          break
        case 6:
          message.initialGenesis = Any.decode(reader, reader.uint32())
          break
        case 7:
          message.launchTriggered = reader.bool()
          break
        case 8:
          message.launchTimestamp = longToNumber(reader.int64() as Long)
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): Chain {
    const message = { ...baseChain } as Chain
    if (object.chainID !== undefined && object.chainID !== null) {
      message.chainID = String(object.chainID)
    } else {
      message.chainID = ''
    }
    if (object.coordinatorID !== undefined && object.coordinatorID !== null) {
      message.coordinatorID = Number(object.coordinatorID)
    } else {
      message.coordinatorID = 0
    }
    if (object.createdAt !== undefined && object.createdAt !== null) {
      message.createdAt = Number(object.createdAt)
    } else {
      message.createdAt = 0
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
    if (object.initialGenesis !== undefined && object.initialGenesis !== null) {
      message.initialGenesis = Any.fromJSON(object.initialGenesis)
    } else {
      message.initialGenesis = undefined
    }
    if (object.launchTriggered !== undefined && object.launchTriggered !== null) {
      message.launchTriggered = Boolean(object.launchTriggered)
    } else {
      message.launchTriggered = false
    }
    if (object.launchTimestamp !== undefined && object.launchTimestamp !== null) {
      message.launchTimestamp = Number(object.launchTimestamp)
    } else {
      message.launchTimestamp = 0
    }
    return message
  },

  toJSON(message: Chain): unknown {
    const obj: any = {}
    message.chainID !== undefined && (obj.chainID = message.chainID)
    message.coordinatorID !== undefined && (obj.coordinatorID = message.coordinatorID)
    message.createdAt !== undefined && (obj.createdAt = message.createdAt)
    message.sourceURL !== undefined && (obj.sourceURL = message.sourceURL)
    message.sourceHash !== undefined && (obj.sourceHash = message.sourceHash)
    message.initialGenesis !== undefined && (obj.initialGenesis = message.initialGenesis ? Any.toJSON(message.initialGenesis) : undefined)
    message.launchTriggered !== undefined && (obj.launchTriggered = message.launchTriggered)
    message.launchTimestamp !== undefined && (obj.launchTimestamp = message.launchTimestamp)
    return obj
  },

  fromPartial(object: DeepPartial<Chain>): Chain {
    const message = { ...baseChain } as Chain
    if (object.chainID !== undefined && object.chainID !== null) {
      message.chainID = object.chainID
    } else {
      message.chainID = ''
    }
    if (object.coordinatorID !== undefined && object.coordinatorID !== null) {
      message.coordinatorID = object.coordinatorID
    } else {
      message.coordinatorID = 0
    }
    if (object.createdAt !== undefined && object.createdAt !== null) {
      message.createdAt = object.createdAt
    } else {
      message.createdAt = 0
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
    if (object.initialGenesis !== undefined && object.initialGenesis !== null) {
      message.initialGenesis = Any.fromPartial(object.initialGenesis)
    } else {
      message.initialGenesis = undefined
    }
    if (object.launchTriggered !== undefined && object.launchTriggered !== null) {
      message.launchTriggered = object.launchTriggered
    } else {
      message.launchTriggered = false
    }
    if (object.launchTimestamp !== undefined && object.launchTimestamp !== null) {
      message.launchTimestamp = object.launchTimestamp
    } else {
      message.launchTimestamp = 0
    }
    return message
  }
}

const baseDefaultInitialGenesis: object = {}

export const DefaultInitialGenesis = {
  encode(_: DefaultInitialGenesis, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): DefaultInitialGenesis {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseDefaultInitialGenesis } as DefaultInitialGenesis
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

  fromJSON(_: any): DefaultInitialGenesis {
    const message = { ...baseDefaultInitialGenesis } as DefaultInitialGenesis
    return message
  },

  toJSON(_: DefaultInitialGenesis): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<DefaultInitialGenesis>): DefaultInitialGenesis {
    const message = { ...baseDefaultInitialGenesis } as DefaultInitialGenesis
    return message
  }
}

const baseGenesisURL: object = { url: '', hash: '' }

export const GenesisURL = {
  encode(message: GenesisURL, writer: Writer = Writer.create()): Writer {
    if (message.url !== '') {
      writer.uint32(10).string(message.url)
    }
    if (message.hash !== '') {
      writer.uint32(18).string(message.hash)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisURL {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseGenesisURL } as GenesisURL
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.url = reader.string()
          break
        case 2:
          message.hash = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): GenesisURL {
    const message = { ...baseGenesisURL } as GenesisURL
    if (object.url !== undefined && object.url !== null) {
      message.url = String(object.url)
    } else {
      message.url = ''
    }
    if (object.hash !== undefined && object.hash !== null) {
      message.hash = String(object.hash)
    } else {
      message.hash = ''
    }
    return message
  },

  toJSON(message: GenesisURL): unknown {
    const obj: any = {}
    message.url !== undefined && (obj.url = message.url)
    message.hash !== undefined && (obj.hash = message.hash)
    return obj
  },

  fromPartial(object: DeepPartial<GenesisURL>): GenesisURL {
    const message = { ...baseGenesisURL } as GenesisURL
    if (object.url !== undefined && object.url !== null) {
      message.url = object.url
    } else {
      message.url = ''
    }
    if (object.hash !== undefined && object.hash !== null) {
      message.hash = object.hash
    } else {
      message.hash = ''
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
