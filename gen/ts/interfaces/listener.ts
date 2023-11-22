/* eslint-disable */
import * as _m0 from "protobufjs/minimal";
import Long = require("long");

export const protobufPackage = "listener";

export interface Listener {
  id: number;
  name: string;
  description: string;
  eventId: number;
}

export interface Listener_AddListenerDto {
  name: string;
  description: string;
  eventName: string;
}

export interface Listener_GetListenersByEvent {
  eventName: string;
}

export interface Listener_GetListenersByEventResult {
  result: Listener[];
}

function createBaseListener(): Listener {
  return { id: 0, name: "", description: "", eventId: 0 };
}

export const Listener = {
  encode(message: Listener, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).int64(message.id);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.description !== "") {
      writer.uint32(26).string(message.description);
    }
    if (message.eventId !== 0) {
      writer.uint32(32).int64(message.eventId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Listener {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListener();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.id = longToNumber(reader.int64() as Long);
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.name = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.description = reader.string();
          continue;
        case 4:
          if (tag !== 32) {
            break;
          }

          message.eventId = longToNumber(reader.int64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Listener {
    return {
      id: isSet(object.id) ? globalThis.Number(object.id) : 0,
      name: isSet(object.name) ? globalThis.String(object.name) : "",
      description: isSet(object.description) ? globalThis.String(object.description) : "",
      eventId: isSet(object.eventId) ? globalThis.Number(object.eventId) : 0,
    };
  },

  toJSON(message: Listener): unknown {
    const obj: any = {};
    if (message.id !== 0) {
      obj.id = Math.round(message.id);
    }
    if (message.name !== "") {
      obj.name = message.name;
    }
    if (message.description !== "") {
      obj.description = message.description;
    }
    if (message.eventId !== 0) {
      obj.eventId = Math.round(message.eventId);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<Listener>, I>>(base?: I): Listener {
    return Listener.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<Listener>, I>>(object: I): Listener {
    const message = createBaseListener();
    message.id = object.id ?? 0;
    message.name = object.name ?? "";
    message.description = object.description ?? "";
    message.eventId = object.eventId ?? 0;
    return message;
  },
};

function createBaseListener_AddListenerDto(): Listener_AddListenerDto {
  return { name: "", description: "", eventName: "" };
}

export const Listener_AddListenerDto = {
  encode(message: Listener_AddListenerDto, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    if (message.description !== "") {
      writer.uint32(18).string(message.description);
    }
    if (message.eventName !== "") {
      writer.uint32(26).string(message.eventName);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Listener_AddListenerDto {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListener_AddListenerDto();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.name = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.description = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.eventName = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Listener_AddListenerDto {
    return {
      name: isSet(object.name) ? globalThis.String(object.name) : "",
      description: isSet(object.description) ? globalThis.String(object.description) : "",
      eventName: isSet(object.eventName) ? globalThis.String(object.eventName) : "",
    };
  },

  toJSON(message: Listener_AddListenerDto): unknown {
    const obj: any = {};
    if (message.name !== "") {
      obj.name = message.name;
    }
    if (message.description !== "") {
      obj.description = message.description;
    }
    if (message.eventName !== "") {
      obj.eventName = message.eventName;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<Listener_AddListenerDto>, I>>(base?: I): Listener_AddListenerDto {
    return Listener_AddListenerDto.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<Listener_AddListenerDto>, I>>(object: I): Listener_AddListenerDto {
    const message = createBaseListener_AddListenerDto();
    message.name = object.name ?? "";
    message.description = object.description ?? "";
    message.eventName = object.eventName ?? "";
    return message;
  },
};

function createBaseListener_GetListenersByEvent(): Listener_GetListenersByEvent {
  return { eventName: "" };
}

export const Listener_GetListenersByEvent = {
  encode(message: Listener_GetListenersByEvent, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.eventName !== "") {
      writer.uint32(10).string(message.eventName);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Listener_GetListenersByEvent {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListener_GetListenersByEvent();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.eventName = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Listener_GetListenersByEvent {
    return { eventName: isSet(object.eventName) ? globalThis.String(object.eventName) : "" };
  },

  toJSON(message: Listener_GetListenersByEvent): unknown {
    const obj: any = {};
    if (message.eventName !== "") {
      obj.eventName = message.eventName;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<Listener_GetListenersByEvent>, I>>(base?: I): Listener_GetListenersByEvent {
    return Listener_GetListenersByEvent.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<Listener_GetListenersByEvent>, I>>(object: I): Listener_GetListenersByEvent {
    const message = createBaseListener_GetListenersByEvent();
    message.eventName = object.eventName ?? "";
    return message;
  },
};

function createBaseListener_GetListenersByEventResult(): Listener_GetListenersByEventResult {
  return { result: [] };
}

export const Listener_GetListenersByEventResult = {
  encode(message: Listener_GetListenersByEventResult, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.result) {
      Listener.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Listener_GetListenersByEventResult {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListener_GetListenersByEventResult();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.result.push(Listener.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Listener_GetListenersByEventResult {
    return {
      result: globalThis.Array.isArray(object?.result) ? object.result.map((e: any) => Listener.fromJSON(e)) : [],
    };
  },

  toJSON(message: Listener_GetListenersByEventResult): unknown {
    const obj: any = {};
    if (message.result?.length) {
      obj.result = message.result.map((e) => Listener.toJSON(e));
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<Listener_GetListenersByEventResult>, I>>(
    base?: I,
  ): Listener_GetListenersByEventResult {
    return Listener_GetListenersByEventResult.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<Listener_GetListenersByEventResult>, I>>(
    object: I,
  ): Listener_GetListenersByEventResult {
    const message = createBaseListener_GetListenersByEventResult();
    message.result = object.result?.map((e) => Listener.fromPartial(e)) || [];
    return message;
  },
};

export interface IListenerController {
  AddListener(request: Listener_AddListenerDto): Promise<Listener>;
  GetListenersByEvent(request: Listener_GetListenersByEvent): Promise<Listener_GetListenersByEventResult>;
}

export const IListenerControllerServiceName = "listener.IListenerController";
export class IListenerControllerClientImpl implements IListenerController {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || IListenerControllerServiceName;
    this.rpc = rpc;
    this.AddListener = this.AddListener.bind(this);
    this.GetListenersByEvent = this.GetListenersByEvent.bind(this);
  }
  AddListener(request: Listener_AddListenerDto): Promise<Listener> {
    const data = Listener_AddListenerDto.encode(request).finish();
    const promise = this.rpc.request(this.service, "AddListener", data);
    return promise.then((data) => Listener.decode(_m0.Reader.create(data)));
  }

  GetListenersByEvent(request: Listener_GetListenersByEvent): Promise<Listener_GetListenersByEventResult> {
    const data = Listener_GetListenersByEvent.encode(request).finish();
    const promise = this.rpc.request(this.service, "GetListenersByEvent", data);
    return promise.then((data) => Listener_GetListenersByEventResult.decode(_m0.Reader.create(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends globalThis.Array<infer U> ? globalThis.Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function longToNumber(long: Long): number {
  if (long.gt(globalThis.Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
