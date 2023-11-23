/* eslint-disable */
import * as _m0 from "protobufjs/minimal";
import Long = require("long");

export const protobufPackage = "mail";

export interface IMail {
  id: number;
  /** Email to send */
  to: string;
  /** Subject of mail */
  subject: string;
  /** Html of mail */
  html: string;
}

export interface SendMailDto {
  to: string;
  subject: string;
  template?: string | undefined;
  html?: string | undefined;
}

function createBaseIMail(): IMail {
  return { id: 0, to: "", subject: "", html: "" };
}

export const IMail = {
  encode(message: IMail, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).int64(message.id);
    }
    if (message.to !== "") {
      writer.uint32(18).string(message.to);
    }
    if (message.subject !== "") {
      writer.uint32(26).string(message.subject);
    }
    if (message.html !== "") {
      writer.uint32(34).string(message.html);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): IMail {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseIMail();
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

          message.to = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.subject = reader.string();
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.html = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): IMail {
    return {
      id: isSet(object.id) ? globalThis.Number(object.id) : 0,
      to: isSet(object.to) ? globalThis.String(object.to) : "",
      subject: isSet(object.subject) ? globalThis.String(object.subject) : "",
      html: isSet(object.html) ? globalThis.String(object.html) : "",
    };
  },

  toJSON(message: IMail): unknown {
    const obj: any = {};
    if (message.id !== 0) {
      obj.id = Math.round(message.id);
    }
    if (message.to !== "") {
      obj.to = message.to;
    }
    if (message.subject !== "") {
      obj.subject = message.subject;
    }
    if (message.html !== "") {
      obj.html = message.html;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<IMail>, I>>(base?: I): IMail {
    return IMail.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<IMail>, I>>(object: I): IMail {
    const message = createBaseIMail();
    message.id = object.id ?? 0;
    message.to = object.to ?? "";
    message.subject = object.subject ?? "";
    message.html = object.html ?? "";
    return message;
  },
};

function createBaseSendMailDto(): SendMailDto {
  return { to: "", subject: "", template: undefined, html: undefined };
}

export const SendMailDto = {
  encode(message: SendMailDto, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.to !== "") {
      writer.uint32(18).string(message.to);
    }
    if (message.subject !== "") {
      writer.uint32(26).string(message.subject);
    }
    if (message.template !== undefined) {
      writer.uint32(34).string(message.template);
    }
    if (message.html !== undefined) {
      writer.uint32(42).string(message.html);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SendMailDto {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSendMailDto();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 2:
          if (tag !== 18) {
            break;
          }

          message.to = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.subject = reader.string();
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.template = reader.string();
          continue;
        case 5:
          if (tag !== 42) {
            break;
          }

          message.html = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): SendMailDto {
    return {
      to: isSet(object.to) ? globalThis.String(object.to) : "",
      subject: isSet(object.subject) ? globalThis.String(object.subject) : "",
      template: isSet(object.template) ? globalThis.String(object.template) : undefined,
      html: isSet(object.html) ? globalThis.String(object.html) : undefined,
    };
  },

  toJSON(message: SendMailDto): unknown {
    const obj: any = {};
    if (message.to !== "") {
      obj.to = message.to;
    }
    if (message.subject !== "") {
      obj.subject = message.subject;
    }
    if (message.template !== undefined) {
      obj.template = message.template;
    }
    if (message.html !== undefined) {
      obj.html = message.html;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<SendMailDto>, I>>(base?: I): SendMailDto {
    return SendMailDto.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<SendMailDto>, I>>(object: I): SendMailDto {
    const message = createBaseSendMailDto();
    message.to = object.to ?? "";
    message.subject = object.subject ?? "";
    message.template = object.template ?? undefined;
    message.html = object.html ?? undefined;
    return message;
  },
};

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
