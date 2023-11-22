/* eslint-disable */
import * as _m0 from "protobufjs/minimal";
import { Observable } from "rxjs";
import { map } from "rxjs/operators";
import Long = require("long");

export const protobufPackage = "event";

export interface IEvent {
  id: number;
  name: string;
  description?: string | undefined;
}

export interface CreateEventDto {
  name: string;
  description?: string | undefined;
}

export interface FindEventDto {
}

export interface GetEventDto {
}

function createBaseIEvent(): IEvent {
  return { id: 0, name: "", description: undefined };
}

export const IEvent = {
  encode(message: IEvent, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).int64(message.id);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.description !== undefined) {
      writer.uint32(26).string(message.description);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): IEvent {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseIEvent();
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
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): IEvent {
    return {
      id: isSet(object.id) ? globalThis.Number(object.id) : 0,
      name: isSet(object.name) ? globalThis.String(object.name) : "",
      description: isSet(object.description) ? globalThis.String(object.description) : undefined,
    };
  },

  toJSON(message: IEvent): unknown {
    const obj: any = {};
    if (message.id !== 0) {
      obj.id = Math.round(message.id);
    }
    if (message.name !== "") {
      obj.name = message.name;
    }
    if (message.description !== undefined) {
      obj.description = message.description;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<IEvent>, I>>(base?: I): IEvent {
    return IEvent.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<IEvent>, I>>(object: I): IEvent {
    const message = createBaseIEvent();
    message.id = object.id ?? 0;
    message.name = object.name ?? "";
    message.description = object.description ?? undefined;
    return message;
  },
};

function createBaseCreateEventDto(): CreateEventDto {
  return { name: "", description: undefined };
}

export const CreateEventDto = {
  encode(message: CreateEventDto, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.description !== undefined) {
      writer.uint32(26).string(message.description);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateEventDto {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateEventDto();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
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
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CreateEventDto {
    return {
      name: isSet(object.name) ? globalThis.String(object.name) : "",
      description: isSet(object.description) ? globalThis.String(object.description) : undefined,
    };
  },

  toJSON(message: CreateEventDto): unknown {
    const obj: any = {};
    if (message.name !== "") {
      obj.name = message.name;
    }
    if (message.description !== undefined) {
      obj.description = message.description;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<CreateEventDto>, I>>(base?: I): CreateEventDto {
    return CreateEventDto.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<CreateEventDto>, I>>(object: I): CreateEventDto {
    const message = createBaseCreateEventDto();
    message.name = object.name ?? "";
    message.description = object.description ?? undefined;
    return message;
  },
};

function createBaseFindEventDto(): FindEventDto {
  return {};
}

export const FindEventDto = {
  encode(_: FindEventDto, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): FindEventDto {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseFindEventDto();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(_: any): FindEventDto {
    return {};
  },

  toJSON(_: FindEventDto): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<FindEventDto>, I>>(base?: I): FindEventDto {
    return FindEventDto.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<FindEventDto>, I>>(_: I): FindEventDto {
    const message = createBaseFindEventDto();
    return message;
  },
};

function createBaseGetEventDto(): GetEventDto {
  return {};
}

export const GetEventDto = {
  encode(_: GetEventDto, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetEventDto {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetEventDto();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(_: any): GetEventDto {
    return {};
  },

  toJSON(_: GetEventDto): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<GetEventDto>, I>>(base?: I): GetEventDto {
    return GetEventDto.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GetEventDto>, I>>(_: I): GetEventDto {
    const message = createBaseGetEventDto();
    return message;
  },
};

export interface IEventService {
  Create(request: CreateEventDto): Promise<IEvent>;
  Find(request: FindEventDto): Observable<IEvent>;
  Get(request: GetEventDto): Promise<IEvent>;
}

export const IEventServiceServiceName = "event.IEventService";
export class IEventServiceClientImpl implements IEventService {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || IEventServiceServiceName;
    this.rpc = rpc;
    this.Create = this.Create.bind(this);
    this.Find = this.Find.bind(this);
    this.Get = this.Get.bind(this);
  }
  Create(request: CreateEventDto): Promise<IEvent> {
    const data = CreateEventDto.encode(request).finish();
    const promise = this.rpc.request(this.service, "Create", data);
    return promise.then((data) => IEvent.decode(_m0.Reader.create(data)));
  }

  Find(request: FindEventDto): Observable<IEvent> {
    const data = FindEventDto.encode(request).finish();
    const result = this.rpc.serverStreamingRequest(this.service, "Find", data);
    return result.pipe(map((data) => IEvent.decode(_m0.Reader.create(data))));
  }

  Get(request: GetEventDto): Promise<IEvent> {
    const data = GetEventDto.encode(request).finish();
    const promise = this.rpc.request(this.service, "Get", data);
    return promise.then((data) => IEvent.decode(_m0.Reader.create(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
  clientStreamingRequest(service: string, method: string, data: Observable<Uint8Array>): Promise<Uint8Array>;
  serverStreamingRequest(service: string, method: string, data: Uint8Array): Observable<Uint8Array>;
  bidirectionalStreamingRequest(service: string, method: string, data: Observable<Uint8Array>): Observable<Uint8Array>;
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
