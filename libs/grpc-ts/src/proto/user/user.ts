// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.7.0
//   protoc               v6.30.2
// source: proto/user/user.proto

/* eslint-disable */
import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";
import {
  type CallOptions,
  ChannelCredentials,
  Client,
  type ClientOptions,
  type ClientUnaryCall,
  type handleUnaryCall,
  makeGenericClientConstructor,
  Metadata,
  type ServiceError,
  type UntypedServiceImplementation,
} from "@grpc/grpc-js";

export interface CreateUserRequest {
  oidcId: string;
  studentId: string;
}

export interface CreateUserResponse {
}

export interface GetUserByOidcIdRequest {
  oidcId: string;
}

export interface GetUserByOidcIdResponse {
  id: string;
  oidcId: string;
  studentId: string;
}

function createBaseCreateUserRequest(): CreateUserRequest {
  return { oidcId: "", studentId: "" };
}

export const CreateUserRequest: MessageFns<CreateUserRequest> = {
  encode(message: CreateUserRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.oidcId !== "") {
      writer.uint32(10).string(message.oidcId);
    }
    if (message.studentId !== "") {
      writer.uint32(18).string(message.studentId);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): CreateUserRequest {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateUserRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.oidcId = reader.string();
          continue;
        }
        case 2: {
          if (tag !== 18) {
            break;
          }

          message.studentId = reader.string();
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CreateUserRequest {
    return {
      oidcId: isSet(object.oidcId) ? globalThis.String(object.oidcId) : "",
      studentId: isSet(object.studentId) ? globalThis.String(object.studentId) : "",
    };
  },

  toJSON(message: CreateUserRequest): unknown {
    const obj: any = {};
    if (message.oidcId !== "") {
      obj.oidcId = message.oidcId;
    }
    if (message.studentId !== "") {
      obj.studentId = message.studentId;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<CreateUserRequest>, I>>(base?: I): CreateUserRequest {
    return CreateUserRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<CreateUserRequest>, I>>(object: I): CreateUserRequest {
    const message = createBaseCreateUserRequest();
    message.oidcId = object.oidcId ?? "";
    message.studentId = object.studentId ?? "";
    return message;
  },
};

function createBaseCreateUserResponse(): CreateUserResponse {
  return {};
}

export const CreateUserResponse: MessageFns<CreateUserResponse> = {
  encode(_: CreateUserResponse, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): CreateUserResponse {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateUserResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(_: any): CreateUserResponse {
    return {};
  },

  toJSON(_: CreateUserResponse): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<CreateUserResponse>, I>>(base?: I): CreateUserResponse {
    return CreateUserResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<CreateUserResponse>, I>>(_: I): CreateUserResponse {
    const message = createBaseCreateUserResponse();
    return message;
  },
};

function createBaseGetUserByOidcIdRequest(): GetUserByOidcIdRequest {
  return { oidcId: "" };
}

export const GetUserByOidcIdRequest: MessageFns<GetUserByOidcIdRequest> = {
  encode(message: GetUserByOidcIdRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.oidcId !== "") {
      writer.uint32(10).string(message.oidcId);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): GetUserByOidcIdRequest {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetUserByOidcIdRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.oidcId = reader.string();
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetUserByOidcIdRequest {
    return { oidcId: isSet(object.oidcId) ? globalThis.String(object.oidcId) : "" };
  },

  toJSON(message: GetUserByOidcIdRequest): unknown {
    const obj: any = {};
    if (message.oidcId !== "") {
      obj.oidcId = message.oidcId;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetUserByOidcIdRequest>, I>>(base?: I): GetUserByOidcIdRequest {
    return GetUserByOidcIdRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GetUserByOidcIdRequest>, I>>(object: I): GetUserByOidcIdRequest {
    const message = createBaseGetUserByOidcIdRequest();
    message.oidcId = object.oidcId ?? "";
    return message;
  },
};

function createBaseGetUserByOidcIdResponse(): GetUserByOidcIdResponse {
  return { id: "", oidcId: "", studentId: "" };
}

export const GetUserByOidcIdResponse: MessageFns<GetUserByOidcIdResponse> = {
  encode(message: GetUserByOidcIdResponse, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.oidcId !== "") {
      writer.uint32(18).string(message.oidcId);
    }
    if (message.studentId !== "") {
      writer.uint32(26).string(message.studentId);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): GetUserByOidcIdResponse {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetUserByOidcIdResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.id = reader.string();
          continue;
        }
        case 2: {
          if (tag !== 18) {
            break;
          }

          message.oidcId = reader.string();
          continue;
        }
        case 3: {
          if (tag !== 26) {
            break;
          }

          message.studentId = reader.string();
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetUserByOidcIdResponse {
    return {
      id: isSet(object.id) ? globalThis.String(object.id) : "",
      oidcId: isSet(object.oidcId) ? globalThis.String(object.oidcId) : "",
      studentId: isSet(object.studentId) ? globalThis.String(object.studentId) : "",
    };
  },

  toJSON(message: GetUserByOidcIdResponse): unknown {
    const obj: any = {};
    if (message.id !== "") {
      obj.id = message.id;
    }
    if (message.oidcId !== "") {
      obj.oidcId = message.oidcId;
    }
    if (message.studentId !== "") {
      obj.studentId = message.studentId;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetUserByOidcIdResponse>, I>>(base?: I): GetUserByOidcIdResponse {
    return GetUserByOidcIdResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GetUserByOidcIdResponse>, I>>(object: I): GetUserByOidcIdResponse {
    const message = createBaseGetUserByOidcIdResponse();
    message.id = object.id ?? "";
    message.oidcId = object.oidcId ?? "";
    message.studentId = object.studentId ?? "";
    return message;
  },
};

export type UserServiceService = typeof UserServiceService;
export const UserServiceService = {
  createUser: {
    path: "/user.UserService/CreateUser",
    requestStream: false,
    responseStream: false,
    requestSerialize: (value: CreateUserRequest) => Buffer.from(CreateUserRequest.encode(value).finish()),
    requestDeserialize: (value: Buffer) => CreateUserRequest.decode(value),
    responseSerialize: (value: CreateUserResponse) => Buffer.from(CreateUserResponse.encode(value).finish()),
    responseDeserialize: (value: Buffer) => CreateUserResponse.decode(value),
  },
  getUserByOidcId: {
    path: "/user.UserService/GetUserByOidcId",
    requestStream: false,
    responseStream: false,
    requestSerialize: (value: GetUserByOidcIdRequest) => Buffer.from(GetUserByOidcIdRequest.encode(value).finish()),
    requestDeserialize: (value: Buffer) => GetUserByOidcIdRequest.decode(value),
    responseSerialize: (value: GetUserByOidcIdResponse) => Buffer.from(GetUserByOidcIdResponse.encode(value).finish()),
    responseDeserialize: (value: Buffer) => GetUserByOidcIdResponse.decode(value),
  },
} as const;

export interface UserServiceServer extends UntypedServiceImplementation {
  createUser: handleUnaryCall<CreateUserRequest, CreateUserResponse>;
  getUserByOidcId: handleUnaryCall<GetUserByOidcIdRequest, GetUserByOidcIdResponse>;
}

export interface UserServiceClient extends Client {
  createUser(
    request: CreateUserRequest,
    callback: (error: ServiceError | null, response: CreateUserResponse) => void,
  ): ClientUnaryCall;
  createUser(
    request: CreateUserRequest,
    metadata: Metadata,
    callback: (error: ServiceError | null, response: CreateUserResponse) => void,
  ): ClientUnaryCall;
  createUser(
    request: CreateUserRequest,
    metadata: Metadata,
    options: Partial<CallOptions>,
    callback: (error: ServiceError | null, response: CreateUserResponse) => void,
  ): ClientUnaryCall;
  getUserByOidcId(
    request: GetUserByOidcIdRequest,
    callback: (error: ServiceError | null, response: GetUserByOidcIdResponse) => void,
  ): ClientUnaryCall;
  getUserByOidcId(
    request: GetUserByOidcIdRequest,
    metadata: Metadata,
    callback: (error: ServiceError | null, response: GetUserByOidcIdResponse) => void,
  ): ClientUnaryCall;
  getUserByOidcId(
    request: GetUserByOidcIdRequest,
    metadata: Metadata,
    options: Partial<CallOptions>,
    callback: (error: ServiceError | null, response: GetUserByOidcIdResponse) => void,
  ): ClientUnaryCall;
}

export const UserServiceClient = makeGenericClientConstructor(UserServiceService, "user.UserService") as unknown as {
  new (address: string, credentials: ChannelCredentials, options?: Partial<ClientOptions>): UserServiceClient;
  service: typeof UserServiceService;
  serviceName: string;
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

type DeepPartial<T> = T extends Builtin ? T
  : T extends globalThis.Array<infer U> ? globalThis.Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}

interface MessageFns<T> {
  encode(message: T, writer?: BinaryWriter): BinaryWriter;
  decode(input: BinaryReader | Uint8Array, length?: number): T;
  fromJSON(object: any): T;
  toJSON(message: T): unknown;
  create<I extends Exact<DeepPartial<T>, I>>(base?: I): T;
  fromPartial<I extends Exact<DeepPartial<T>, I>>(object: I): T;
}
