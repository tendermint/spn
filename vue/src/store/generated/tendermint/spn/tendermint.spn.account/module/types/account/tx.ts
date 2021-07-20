/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal'

export const protobufPackage = 'tendermint.spn.account'

/** this line is used by starport scaffolding # proto/tx/message */
export interface MsgDeleteCoordinator {
  address: string
}

export interface MsgDeleteCoordinatorResponse {}

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
  DeleteCoordinator(request: MsgDeleteCoordinator): Promise<MsgDeleteCoordinatorResponse>
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc
  constructor(rpc: Rpc) {
    this.rpc = rpc
  }
  DeleteCoordinator(request: MsgDeleteCoordinator): Promise<MsgDeleteCoordinatorResponse> {
    const data = MsgDeleteCoordinator.encode(request).finish()
    const promise = this.rpc.request('tendermint.spn.account.Msg', 'DeleteCoordinator', data)
    return promise.then((data) => MsgDeleteCoordinatorResponse.decode(new Reader(data)))
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
