/* eslint-disable */
import * as _m0 from "protobufjs/minimal";
import Long = require("long");

export const protobufPackage = "product";

export interface IdDto {
  id: number;
}

export interface dtoUpdateAmount {
  alias: string;
  amount: number;
}

export interface IdsDto {
  ids: number[];
}

export interface AliasDto {
  alias: string;
}

export interface ManyAliasDto {
  aliases: string[];
}

export interface ResponseFindOne {
  id: number;
  name: string;
  alias: string;
  price: number;
  totalSold: number;
  discount: number;
  description: string;
}

export interface ResponseFindMany {
  products: ResponseFindOne[];
}

function createBaseIdDto(): IdDto {
  return { id: 0 };
}

export const IdDto = {
  encode(message: IdDto, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).int64(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): IdDto {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseIdDto();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.id = longToNumber(reader.int64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): IdDto {
    return { id: isSet(object.id) ? globalThis.Number(object.id) : 0 };
  },

  toJSON(message: IdDto): unknown {
    const obj: any = {};
    if (message.id !== 0) {
      obj.id = Math.round(message.id);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<IdDto>, I>>(base?: I): IdDto {
    return IdDto.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<IdDto>, I>>(object: I): IdDto {
    const message = createBaseIdDto();
    message.id = object.id ?? 0;
    return message;
  },
};

function createBasedtoUpdateAmount(): dtoUpdateAmount {
  return { alias: "", amount: 0 };
}

export const dtoUpdateAmount = {
  encode(message: dtoUpdateAmount, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.alias !== "") {
      writer.uint32(10).string(message.alias);
    }
    if (message.amount !== 0) {
      writer.uint32(16).int64(message.amount);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): dtoUpdateAmount {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasedtoUpdateAmount();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.alias = reader.string();
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.amount = longToNumber(reader.int64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): dtoUpdateAmount {
    return {
      alias: isSet(object.alias) ? globalThis.String(object.alias) : "",
      amount: isSet(object.amount) ? globalThis.Number(object.amount) : 0,
    };
  },

  toJSON(message: dtoUpdateAmount): unknown {
    const obj: any = {};
    if (message.alias !== "") {
      obj.alias = message.alias;
    }
    if (message.amount !== 0) {
      obj.amount = Math.round(message.amount);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<dtoUpdateAmount>, I>>(base?: I): dtoUpdateAmount {
    return dtoUpdateAmount.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<dtoUpdateAmount>, I>>(object: I): dtoUpdateAmount {
    const message = createBasedtoUpdateAmount();
    message.alias = object.alias ?? "";
    message.amount = object.amount ?? 0;
    return message;
  },
};

function createBaseIdsDto(): IdsDto {
  return { ids: [] };
}

export const IdsDto = {
  encode(message: IdsDto, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    writer.uint32(10).fork();
    for (const v of message.ids) {
      writer.int64(v);
    }
    writer.ldelim();
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): IdsDto {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseIdsDto();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag === 8) {
            message.ids.push(longToNumber(reader.int64() as Long));

            continue;
          }

          if (tag === 10) {
            const end2 = reader.uint32() + reader.pos;
            while (reader.pos < end2) {
              message.ids.push(longToNumber(reader.int64() as Long));
            }

            continue;
          }

          break;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): IdsDto {
    return { ids: globalThis.Array.isArray(object?.ids) ? object.ids.map((e: any) => globalThis.Number(e)) : [] };
  },

  toJSON(message: IdsDto): unknown {
    const obj: any = {};
    if (message.ids?.length) {
      obj.ids = message.ids.map((e) => Math.round(e));
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<IdsDto>, I>>(base?: I): IdsDto {
    return IdsDto.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<IdsDto>, I>>(object: I): IdsDto {
    const message = createBaseIdsDto();
    message.ids = object.ids?.map((e) => e) || [];
    return message;
  },
};

function createBaseAliasDto(): AliasDto {
  return { alias: "" };
}

export const AliasDto = {
  encode(message: AliasDto, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.alias !== "") {
      writer.uint32(10).string(message.alias);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): AliasDto {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAliasDto();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.alias = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): AliasDto {
    return { alias: isSet(object.alias) ? globalThis.String(object.alias) : "" };
  },

  toJSON(message: AliasDto): unknown {
    const obj: any = {};
    if (message.alias !== "") {
      obj.alias = message.alias;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<AliasDto>, I>>(base?: I): AliasDto {
    return AliasDto.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<AliasDto>, I>>(object: I): AliasDto {
    const message = createBaseAliasDto();
    message.alias = object.alias ?? "";
    return message;
  },
};

function createBaseManyAliasDto(): ManyAliasDto {
  return { aliases: [] };
}

export const ManyAliasDto = {
  encode(message: ManyAliasDto, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.aliases) {
      writer.uint32(10).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ManyAliasDto {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseManyAliasDto();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.aliases.push(reader.string());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ManyAliasDto {
    return {
      aliases: globalThis.Array.isArray(object?.aliases) ? object.aliases.map((e: any) => globalThis.String(e)) : [],
    };
  },

  toJSON(message: ManyAliasDto): unknown {
    const obj: any = {};
    if (message.aliases?.length) {
      obj.aliases = message.aliases;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<ManyAliasDto>, I>>(base?: I): ManyAliasDto {
    return ManyAliasDto.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<ManyAliasDto>, I>>(object: I): ManyAliasDto {
    const message = createBaseManyAliasDto();
    message.aliases = object.aliases?.map((e) => e) || [];
    return message;
  },
};

function createBaseResponseFindOne(): ResponseFindOne {
  return { id: 0, name: "", alias: "", price: 0, totalSold: 0, discount: 0, description: "" };
}

export const ResponseFindOne = {
  encode(message: ResponseFindOne, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).int64(message.id);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.alias !== "") {
      writer.uint32(26).string(message.alias);
    }
    if (message.price !== 0) {
      writer.uint32(32).int64(message.price);
    }
    if (message.totalSold !== 0) {
      writer.uint32(40).int64(message.totalSold);
    }
    if (message.discount !== 0) {
      writer.uint32(48).int64(message.discount);
    }
    if (message.description !== "") {
      writer.uint32(58).string(message.description);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ResponseFindOne {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseResponseFindOne();
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

          message.alias = reader.string();
          continue;
        case 4:
          if (tag !== 32) {
            break;
          }

          message.price = longToNumber(reader.int64() as Long);
          continue;
        case 5:
          if (tag !== 40) {
            break;
          }

          message.totalSold = longToNumber(reader.int64() as Long);
          continue;
        case 6:
          if (tag !== 48) {
            break;
          }

          message.discount = longToNumber(reader.int64() as Long);
          continue;
        case 7:
          if (tag !== 58) {
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

  fromJSON(object: any): ResponseFindOne {
    return {
      id: isSet(object.id) ? globalThis.Number(object.id) : 0,
      name: isSet(object.name) ? globalThis.String(object.name) : "",
      alias: isSet(object.alias) ? globalThis.String(object.alias) : "",
      price: isSet(object.price) ? globalThis.Number(object.price) : 0,
      totalSold: isSet(object.totalSold) ? globalThis.Number(object.totalSold) : 0,
      discount: isSet(object.discount) ? globalThis.Number(object.discount) : 0,
      description: isSet(object.description) ? globalThis.String(object.description) : "",
    };
  },

  toJSON(message: ResponseFindOne): unknown {
    const obj: any = {};
    if (message.id !== 0) {
      obj.id = Math.round(message.id);
    }
    if (message.name !== "") {
      obj.name = message.name;
    }
    if (message.alias !== "") {
      obj.alias = message.alias;
    }
    if (message.price !== 0) {
      obj.price = Math.round(message.price);
    }
    if (message.totalSold !== 0) {
      obj.totalSold = Math.round(message.totalSold);
    }
    if (message.discount !== 0) {
      obj.discount = Math.round(message.discount);
    }
    if (message.description !== "") {
      obj.description = message.description;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<ResponseFindOne>, I>>(base?: I): ResponseFindOne {
    return ResponseFindOne.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<ResponseFindOne>, I>>(object: I): ResponseFindOne {
    const message = createBaseResponseFindOne();
    message.id = object.id ?? 0;
    message.name = object.name ?? "";
    message.alias = object.alias ?? "";
    message.price = object.price ?? 0;
    message.totalSold = object.totalSold ?? 0;
    message.discount = object.discount ?? 0;
    message.description = object.description ?? "";
    return message;
  },
};

function createBaseResponseFindMany(): ResponseFindMany {
  return { products: [] };
}

export const ResponseFindMany = {
  encode(message: ResponseFindMany, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.products) {
      ResponseFindOne.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ResponseFindMany {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseResponseFindMany();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.products.push(ResponseFindOne.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ResponseFindMany {
    return {
      products: globalThis.Array.isArray(object?.products)
        ? object.products.map((e: any) => ResponseFindOne.fromJSON(e))
        : [],
    };
  },

  toJSON(message: ResponseFindMany): unknown {
    const obj: any = {};
    if (message.products?.length) {
      obj.products = message.products.map((e) => ResponseFindOne.toJSON(e));
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<ResponseFindMany>, I>>(base?: I): ResponseFindMany {
    return ResponseFindMany.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<ResponseFindMany>, I>>(object: I): ResponseFindMany {
    const message = createBaseResponseFindMany();
    message.products = object.products?.map((e) => ResponseFindOne.fromPartial(e)) || [];
    return message;
  },
};

export interface IProductController {
  findOneById(request: IdDto): Promise<ResponseFindOne>;
  findManyByIds(request: IdsDto): Promise<ResponseFindMany>;
  findOneByAlias(request: AliasDto): Promise<ResponseFindOne>;
  findManyByAlias(request: ManyAliasDto): Promise<ResponseFindMany>;
  incrementProduct(request: dtoUpdateAmount): Promise<ResponseFindOne>;
  descrementProduct(request: dtoUpdateAmount): Promise<ResponseFindOne>;
}

export const IProductControllerServiceName = "product.IProductController";
export class IProductControllerClientImpl implements IProductController {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || IProductControllerServiceName;
    this.rpc = rpc;
    this.findOneById = this.findOneById.bind(this);
    this.findManyByIds = this.findManyByIds.bind(this);
    this.findOneByAlias = this.findOneByAlias.bind(this);
    this.findManyByAlias = this.findManyByAlias.bind(this);
    this.incrementProduct = this.incrementProduct.bind(this);
    this.descrementProduct = this.descrementProduct.bind(this);
  }
  findOneById(request: IdDto): Promise<ResponseFindOne> {
    const data = IdDto.encode(request).finish();
    const promise = this.rpc.request(this.service, "findOneById", data);
    return promise.then((data) => ResponseFindOne.decode(_m0.Reader.create(data)));
  }

  findManyByIds(request: IdsDto): Promise<ResponseFindMany> {
    const data = IdsDto.encode(request).finish();
    const promise = this.rpc.request(this.service, "findManyByIds", data);
    return promise.then((data) => ResponseFindMany.decode(_m0.Reader.create(data)));
  }

  findOneByAlias(request: AliasDto): Promise<ResponseFindOne> {
    const data = AliasDto.encode(request).finish();
    const promise = this.rpc.request(this.service, "findOneByAlias", data);
    return promise.then((data) => ResponseFindOne.decode(_m0.Reader.create(data)));
  }

  findManyByAlias(request: ManyAliasDto): Promise<ResponseFindMany> {
    const data = ManyAliasDto.encode(request).finish();
    const promise = this.rpc.request(this.service, "findManyByAlias", data);
    return promise.then((data) => ResponseFindMany.decode(_m0.Reader.create(data)));
  }

  incrementProduct(request: dtoUpdateAmount): Promise<ResponseFindOne> {
    const data = dtoUpdateAmount.encode(request).finish();
    const promise = this.rpc.request(this.service, "incrementProduct", data);
    return promise.then((data) => ResponseFindOne.decode(_m0.Reader.create(data)));
  }

  descrementProduct(request: dtoUpdateAmount): Promise<ResponseFindOne> {
    const data = dtoUpdateAmount.encode(request).finish();
    const promise = this.rpc.request(this.service, "descrementProduct", data);
    return promise.then((data) => ResponseFindOne.decode(_m0.Reader.create(data)));
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
