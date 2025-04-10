// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.7.0
//   protoc               v6.30.2
// source: proto/ballot/ballot.proto

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
import { Choice } from "../choice/choice";

export interface CreateBallotProofRequest {
  oidcId: string;
  voteSlug: string;
  choiceNumber: string;
}

export interface CreateBallotProofResponse {
  proof: Proof | undefined;
}

export interface CreateBallotRequest {
  oidcId: string;
  voteSlug: string;
  proof: Proof | undefined;
}

export interface CreateBallotResponse {
  /** voteId:receipt:encryptedChoiceId */
  ballotKey: string;
}

export interface VerifyBallotRequest {
  oidcId: string;
  ballotKey: string;
}

export interface VerifyBallotResponse {
  isValid: boolean;
  choice?: Choice | undefined;
  timestamp: string;
}

export interface Proof {
  commitment: string;
  blindingFactor: string;
  challenge: string;
  response: string;
  nullifier: string;
  receipt: string;
}

function createBaseCreateBallotProofRequest(): CreateBallotProofRequest {
  return { oidcId: "", voteSlug: "", choiceNumber: "" };
}

export const CreateBallotProofRequest: MessageFns<CreateBallotProofRequest> = {
  encode(message: CreateBallotProofRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.oidcId !== "") {
      writer.uint32(10).string(message.oidcId);
    }
    if (message.voteSlug !== "") {
      writer.uint32(18).string(message.voteSlug);
    }
    if (message.choiceNumber !== "") {
      writer.uint32(26).string(message.choiceNumber);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): CreateBallotProofRequest {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateBallotProofRequest();
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

          message.voteSlug = reader.string();
          continue;
        }
        case 3: {
          if (tag !== 26) {
            break;
          }

          message.choiceNumber = reader.string();
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

  fromJSON(object: any): CreateBallotProofRequest {
    return {
      oidcId: isSet(object.oidcId) ? globalThis.String(object.oidcId) : "",
      voteSlug: isSet(object.voteSlug) ? globalThis.String(object.voteSlug) : "",
      choiceNumber: isSet(object.choiceNumber) ? globalThis.String(object.choiceNumber) : "",
    };
  },

  toJSON(message: CreateBallotProofRequest): unknown {
    const obj: any = {};
    if (message.oidcId !== "") {
      obj.oidcId = message.oidcId;
    }
    if (message.voteSlug !== "") {
      obj.voteSlug = message.voteSlug;
    }
    if (message.choiceNumber !== "") {
      obj.choiceNumber = message.choiceNumber;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<CreateBallotProofRequest>, I>>(base?: I): CreateBallotProofRequest {
    return CreateBallotProofRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<CreateBallotProofRequest>, I>>(object: I): CreateBallotProofRequest {
    const message = createBaseCreateBallotProofRequest();
    message.oidcId = object.oidcId ?? "";
    message.voteSlug = object.voteSlug ?? "";
    message.choiceNumber = object.choiceNumber ?? "";
    return message;
  },
};

function createBaseCreateBallotProofResponse(): CreateBallotProofResponse {
  return { proof: undefined };
}

export const CreateBallotProofResponse: MessageFns<CreateBallotProofResponse> = {
  encode(message: CreateBallotProofResponse, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.proof !== undefined) {
      Proof.encode(message.proof, writer.uint32(10).fork()).join();
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): CreateBallotProofResponse {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateBallotProofResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.proof = Proof.decode(reader, reader.uint32());
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

  fromJSON(object: any): CreateBallotProofResponse {
    return { proof: isSet(object.proof) ? Proof.fromJSON(object.proof) : undefined };
  },

  toJSON(message: CreateBallotProofResponse): unknown {
    const obj: any = {};
    if (message.proof !== undefined) {
      obj.proof = Proof.toJSON(message.proof);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<CreateBallotProofResponse>, I>>(base?: I): CreateBallotProofResponse {
    return CreateBallotProofResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<CreateBallotProofResponse>, I>>(object: I): CreateBallotProofResponse {
    const message = createBaseCreateBallotProofResponse();
    message.proof = (object.proof !== undefined && object.proof !== null) ? Proof.fromPartial(object.proof) : undefined;
    return message;
  },
};

function createBaseCreateBallotRequest(): CreateBallotRequest {
  return { oidcId: "", voteSlug: "", proof: undefined };
}

export const CreateBallotRequest: MessageFns<CreateBallotRequest> = {
  encode(message: CreateBallotRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.oidcId !== "") {
      writer.uint32(10).string(message.oidcId);
    }
    if (message.voteSlug !== "") {
      writer.uint32(18).string(message.voteSlug);
    }
    if (message.proof !== undefined) {
      Proof.encode(message.proof, writer.uint32(26).fork()).join();
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): CreateBallotRequest {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateBallotRequest();
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

          message.voteSlug = reader.string();
          continue;
        }
        case 3: {
          if (tag !== 26) {
            break;
          }

          message.proof = Proof.decode(reader, reader.uint32());
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

  fromJSON(object: any): CreateBallotRequest {
    return {
      oidcId: isSet(object.oidcId) ? globalThis.String(object.oidcId) : "",
      voteSlug: isSet(object.voteSlug) ? globalThis.String(object.voteSlug) : "",
      proof: isSet(object.proof) ? Proof.fromJSON(object.proof) : undefined,
    };
  },

  toJSON(message: CreateBallotRequest): unknown {
    const obj: any = {};
    if (message.oidcId !== "") {
      obj.oidcId = message.oidcId;
    }
    if (message.voteSlug !== "") {
      obj.voteSlug = message.voteSlug;
    }
    if (message.proof !== undefined) {
      obj.proof = Proof.toJSON(message.proof);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<CreateBallotRequest>, I>>(base?: I): CreateBallotRequest {
    return CreateBallotRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<CreateBallotRequest>, I>>(object: I): CreateBallotRequest {
    const message = createBaseCreateBallotRequest();
    message.oidcId = object.oidcId ?? "";
    message.voteSlug = object.voteSlug ?? "";
    message.proof = (object.proof !== undefined && object.proof !== null) ? Proof.fromPartial(object.proof) : undefined;
    return message;
  },
};

function createBaseCreateBallotResponse(): CreateBallotResponse {
  return { ballotKey: "" };
}

export const CreateBallotResponse: MessageFns<CreateBallotResponse> = {
  encode(message: CreateBallotResponse, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.ballotKey !== "") {
      writer.uint32(10).string(message.ballotKey);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): CreateBallotResponse {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateBallotResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.ballotKey = reader.string();
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

  fromJSON(object: any): CreateBallotResponse {
    return { ballotKey: isSet(object.ballotKey) ? globalThis.String(object.ballotKey) : "" };
  },

  toJSON(message: CreateBallotResponse): unknown {
    const obj: any = {};
    if (message.ballotKey !== "") {
      obj.ballotKey = message.ballotKey;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<CreateBallotResponse>, I>>(base?: I): CreateBallotResponse {
    return CreateBallotResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<CreateBallotResponse>, I>>(object: I): CreateBallotResponse {
    const message = createBaseCreateBallotResponse();
    message.ballotKey = object.ballotKey ?? "";
    return message;
  },
};

function createBaseVerifyBallotRequest(): VerifyBallotRequest {
  return { oidcId: "", ballotKey: "" };
}

export const VerifyBallotRequest: MessageFns<VerifyBallotRequest> = {
  encode(message: VerifyBallotRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.oidcId !== "") {
      writer.uint32(10).string(message.oidcId);
    }
    if (message.ballotKey !== "") {
      writer.uint32(18).string(message.ballotKey);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): VerifyBallotRequest {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseVerifyBallotRequest();
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

          message.ballotKey = reader.string();
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

  fromJSON(object: any): VerifyBallotRequest {
    return {
      oidcId: isSet(object.oidcId) ? globalThis.String(object.oidcId) : "",
      ballotKey: isSet(object.ballotKey) ? globalThis.String(object.ballotKey) : "",
    };
  },

  toJSON(message: VerifyBallotRequest): unknown {
    const obj: any = {};
    if (message.oidcId !== "") {
      obj.oidcId = message.oidcId;
    }
    if (message.ballotKey !== "") {
      obj.ballotKey = message.ballotKey;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<VerifyBallotRequest>, I>>(base?: I): VerifyBallotRequest {
    return VerifyBallotRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<VerifyBallotRequest>, I>>(object: I): VerifyBallotRequest {
    const message = createBaseVerifyBallotRequest();
    message.oidcId = object.oidcId ?? "";
    message.ballotKey = object.ballotKey ?? "";
    return message;
  },
};

function createBaseVerifyBallotResponse(): VerifyBallotResponse {
  return { isValid: false, choice: undefined, timestamp: "" };
}

export const VerifyBallotResponse: MessageFns<VerifyBallotResponse> = {
  encode(message: VerifyBallotResponse, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.isValid !== false) {
      writer.uint32(8).bool(message.isValid);
    }
    if (message.choice !== undefined) {
      Choice.encode(message.choice, writer.uint32(18).fork()).join();
    }
    if (message.timestamp !== "") {
      writer.uint32(26).string(message.timestamp);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): VerifyBallotResponse {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseVerifyBallotResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 8) {
            break;
          }

          message.isValid = reader.bool();
          continue;
        }
        case 2: {
          if (tag !== 18) {
            break;
          }

          message.choice = Choice.decode(reader, reader.uint32());
          continue;
        }
        case 3: {
          if (tag !== 26) {
            break;
          }

          message.timestamp = reader.string();
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

  fromJSON(object: any): VerifyBallotResponse {
    return {
      isValid: isSet(object.isValid) ? globalThis.Boolean(object.isValid) : false,
      choice: isSet(object.choice) ? Choice.fromJSON(object.choice) : undefined,
      timestamp: isSet(object.timestamp) ? globalThis.String(object.timestamp) : "",
    };
  },

  toJSON(message: VerifyBallotResponse): unknown {
    const obj: any = {};
    if (message.isValid !== false) {
      obj.isValid = message.isValid;
    }
    if (message.choice !== undefined) {
      obj.choice = Choice.toJSON(message.choice);
    }
    if (message.timestamp !== "") {
      obj.timestamp = message.timestamp;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<VerifyBallotResponse>, I>>(base?: I): VerifyBallotResponse {
    return VerifyBallotResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<VerifyBallotResponse>, I>>(object: I): VerifyBallotResponse {
    const message = createBaseVerifyBallotResponse();
    message.isValid = object.isValid ?? false;
    message.choice = (object.choice !== undefined && object.choice !== null)
      ? Choice.fromPartial(object.choice)
      : undefined;
    message.timestamp = object.timestamp ?? "";
    return message;
  },
};

function createBaseProof(): Proof {
  return { commitment: "", blindingFactor: "", challenge: "", response: "", nullifier: "", receipt: "" };
}

export const Proof: MessageFns<Proof> = {
  encode(message: Proof, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.commitment !== "") {
      writer.uint32(10).string(message.commitment);
    }
    if (message.blindingFactor !== "") {
      writer.uint32(18).string(message.blindingFactor);
    }
    if (message.challenge !== "") {
      writer.uint32(26).string(message.challenge);
    }
    if (message.response !== "") {
      writer.uint32(34).string(message.response);
    }
    if (message.nullifier !== "") {
      writer.uint32(42).string(message.nullifier);
    }
    if (message.receipt !== "") {
      writer.uint32(50).string(message.receipt);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): Proof {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseProof();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.commitment = reader.string();
          continue;
        }
        case 2: {
          if (tag !== 18) {
            break;
          }

          message.blindingFactor = reader.string();
          continue;
        }
        case 3: {
          if (tag !== 26) {
            break;
          }

          message.challenge = reader.string();
          continue;
        }
        case 4: {
          if (tag !== 34) {
            break;
          }

          message.response = reader.string();
          continue;
        }
        case 5: {
          if (tag !== 42) {
            break;
          }

          message.nullifier = reader.string();
          continue;
        }
        case 6: {
          if (tag !== 50) {
            break;
          }

          message.receipt = reader.string();
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

  fromJSON(object: any): Proof {
    return {
      commitment: isSet(object.commitment) ? globalThis.String(object.commitment) : "",
      blindingFactor: isSet(object.blindingFactor) ? globalThis.String(object.blindingFactor) : "",
      challenge: isSet(object.challenge) ? globalThis.String(object.challenge) : "",
      response: isSet(object.response) ? globalThis.String(object.response) : "",
      nullifier: isSet(object.nullifier) ? globalThis.String(object.nullifier) : "",
      receipt: isSet(object.receipt) ? globalThis.String(object.receipt) : "",
    };
  },

  toJSON(message: Proof): unknown {
    const obj: any = {};
    if (message.commitment !== "") {
      obj.commitment = message.commitment;
    }
    if (message.blindingFactor !== "") {
      obj.blindingFactor = message.blindingFactor;
    }
    if (message.challenge !== "") {
      obj.challenge = message.challenge;
    }
    if (message.response !== "") {
      obj.response = message.response;
    }
    if (message.nullifier !== "") {
      obj.nullifier = message.nullifier;
    }
    if (message.receipt !== "") {
      obj.receipt = message.receipt;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<Proof>, I>>(base?: I): Proof {
    return Proof.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<Proof>, I>>(object: I): Proof {
    const message = createBaseProof();
    message.commitment = object.commitment ?? "";
    message.blindingFactor = object.blindingFactor ?? "";
    message.challenge = object.challenge ?? "";
    message.response = object.response ?? "";
    message.nullifier = object.nullifier ?? "";
    message.receipt = object.receipt ?? "";
    return message;
  },
};

export type BallotServiceService = typeof BallotServiceService;
export const BallotServiceService = {
  createBallotProof: {
    path: "/ballot.BallotService/CreateBallotProof",
    requestStream: false,
    responseStream: false,
    requestSerialize: (value: CreateBallotProofRequest) => Buffer.from(CreateBallotProofRequest.encode(value).finish()),
    requestDeserialize: (value: Buffer) => CreateBallotProofRequest.decode(value),
    responseSerialize: (value: CreateBallotProofResponse) =>
      Buffer.from(CreateBallotProofResponse.encode(value).finish()),
    responseDeserialize: (value: Buffer) => CreateBallotProofResponse.decode(value),
  },
  createBallot: {
    path: "/ballot.BallotService/CreateBallot",
    requestStream: false,
    responseStream: false,
    requestSerialize: (value: CreateBallotRequest) => Buffer.from(CreateBallotRequest.encode(value).finish()),
    requestDeserialize: (value: Buffer) => CreateBallotRequest.decode(value),
    responseSerialize: (value: CreateBallotResponse) => Buffer.from(CreateBallotResponse.encode(value).finish()),
    responseDeserialize: (value: Buffer) => CreateBallotResponse.decode(value),
  },
  verifyBallot: {
    path: "/ballot.BallotService/VerifyBallot",
    requestStream: false,
    responseStream: false,
    requestSerialize: (value: VerifyBallotRequest) => Buffer.from(VerifyBallotRequest.encode(value).finish()),
    requestDeserialize: (value: Buffer) => VerifyBallotRequest.decode(value),
    responseSerialize: (value: VerifyBallotResponse) => Buffer.from(VerifyBallotResponse.encode(value).finish()),
    responseDeserialize: (value: Buffer) => VerifyBallotResponse.decode(value),
  },
} as const;

export interface BallotServiceServer extends UntypedServiceImplementation {
  createBallotProof: handleUnaryCall<CreateBallotProofRequest, CreateBallotProofResponse>;
  createBallot: handleUnaryCall<CreateBallotRequest, CreateBallotResponse>;
  verifyBallot: handleUnaryCall<VerifyBallotRequest, VerifyBallotResponse>;
}

export interface BallotServiceClient extends Client {
  createBallotProof(
    request: CreateBallotProofRequest,
    callback: (error: ServiceError | null, response: CreateBallotProofResponse) => void,
  ): ClientUnaryCall;
  createBallotProof(
    request: CreateBallotProofRequest,
    metadata: Metadata,
    callback: (error: ServiceError | null, response: CreateBallotProofResponse) => void,
  ): ClientUnaryCall;
  createBallotProof(
    request: CreateBallotProofRequest,
    metadata: Metadata,
    options: Partial<CallOptions>,
    callback: (error: ServiceError | null, response: CreateBallotProofResponse) => void,
  ): ClientUnaryCall;
  createBallot(
    request: CreateBallotRequest,
    callback: (error: ServiceError | null, response: CreateBallotResponse) => void,
  ): ClientUnaryCall;
  createBallot(
    request: CreateBallotRequest,
    metadata: Metadata,
    callback: (error: ServiceError | null, response: CreateBallotResponse) => void,
  ): ClientUnaryCall;
  createBallot(
    request: CreateBallotRequest,
    metadata: Metadata,
    options: Partial<CallOptions>,
    callback: (error: ServiceError | null, response: CreateBallotResponse) => void,
  ): ClientUnaryCall;
  verifyBallot(
    request: VerifyBallotRequest,
    callback: (error: ServiceError | null, response: VerifyBallotResponse) => void,
  ): ClientUnaryCall;
  verifyBallot(
    request: VerifyBallotRequest,
    metadata: Metadata,
    callback: (error: ServiceError | null, response: VerifyBallotResponse) => void,
  ): ClientUnaryCall;
  verifyBallot(
    request: VerifyBallotRequest,
    metadata: Metadata,
    options: Partial<CallOptions>,
    callback: (error: ServiceError | null, response: VerifyBallotResponse) => void,
  ): ClientUnaryCall;
}

export const BallotServiceClient = makeGenericClientConstructor(
  BallotServiceService,
  "ballot.BallotService",
) as unknown as {
  new (address: string, credentials: ChannelCredentials, options?: Partial<ClientOptions>): BallotServiceClient;
  service: typeof BallotServiceService;
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
