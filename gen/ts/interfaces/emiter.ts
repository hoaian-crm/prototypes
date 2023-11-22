/* eslint-disable */
import * as _m0 from "protobufjs/minimal";

export const protobufPackage = "event";

export interface UserRegisterDto {
  email: string;
  otpCode: string;
}

function createBaseUserRegisterDto(): UserRegisterDto {
  return { email: "", otpCode: "" };
}

export const UserRegisterDto = {
  encode(message: UserRegisterDto, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.email !== "") {
      writer.uint32(10).string(message.email);
    }
    if (message.otpCode !== "") {
      writer.uint32(18).string(message.otpCode);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UserRegisterDto {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUserRegisterDto();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.email = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.otpCode = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): UserRegisterDto {
    return {
      email: isSet(object.email) ? globalThis.String(object.email) : "",
      otpCode: isSet(object.otpCode) ? globalThis.String(object.otpCode) : "",
    };
  },

  toJSON(message: UserRegisterDto): unknown {
    const obj: any = {};
    if (message.email !== "") {
      obj.email = message.email;
    }
    if (message.otpCode !== "") {
      obj.otpCode = message.otpCode;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<UserRegisterDto>, I>>(base?: I): UserRegisterDto {
    return UserRegisterDto.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<UserRegisterDto>, I>>(object: I): UserRegisterDto {
    const message = createBaseUserRegisterDto();
    message.email = object.email ?? "";
    message.otpCode = object.otpCode ?? "";
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

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
