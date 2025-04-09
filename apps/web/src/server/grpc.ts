import { type ServiceError, credentials } from "@grpc/grpc-js";
import {
  type CreateUserRequest,
  type CreateUserResponse,
  type CreateVoteRequest,
  type CreateVoteResponse,
  type GetUserByOidcIdRequest,
  type GetUserByOidcIdResponse,
  type GetVoteByIdRequest,
  type GetVoteByIdResponse,
  type GetVoteBySlugRequest,
  type GetVoteBySlugResponse,
  type GetVotesByUserEligibilityRequest,
  type GetVotesByUserEligibilityResponse,
  UserServiceClient,
  VoteServiceClient,
} from "@intania-vote/grpc-ts";

const GRPC_ADDRESS = process.env.GRPC_ADDRESS || "localhost:4000";

export function r<T>(
  resolve: (response: T) => void,
  reject: (reason?: ServiceError) => void,
) {
  return (err: ServiceError | null, response: T) => {
    if (err) {
      reject(err);
      return;
    }
    resolve(response);
  };
}

const userClient = new UserServiceClient(
  GRPC_ADDRESS,
  credentials.createInsecure(),
);

function CreateUser(req: CreateUserRequest): Promise<CreateUserResponse> {
  return new Promise((resolve, reject) => {
    userClient.createUser(req, r(resolve, reject));
  });
}

function GetUserByOidcId(
  req: GetUserByOidcIdRequest,
): Promise<GetUserByOidcIdResponse> {
  return new Promise((resolve, reject) => {
    userClient.getUserByOidcId(req, r(resolve, reject));
  });
}

const voteClient = new VoteServiceClient(
  GRPC_ADDRESS,
  credentials.createInsecure(),
);

function CreateVote(req: CreateVoteRequest): Promise<CreateVoteResponse> {
  return new Promise((resolve, reject) => {
    voteClient.createVote(req, r(resolve, reject));
  });
}

function GetVoteById(req: GetVoteByIdRequest): Promise<GetVoteByIdResponse> {
  return new Promise((resolve, reject) => {
    voteClient.getVoteById(req, r(resolve, reject));
  });
}

function GetVoteBySlug(
  req: GetVoteBySlugRequest,
): Promise<GetVoteBySlugResponse> {
  return new Promise((resolve, reject) => {
    voteClient.getVoteBySlug(req, r(resolve, reject));
  });
}

function GetVotesByUserEligibility(
  req: GetVotesByUserEligibilityRequest,
): Promise<GetVotesByUserEligibilityResponse> {
  return new Promise((resolve, reject) => {
    voteClient.getVotesByUserEligibility(req, r(resolve, reject));
  });
}

export const grpc = {
  user: {
    CreateUser,
    GetUserByOidcId,
  },
  vote: {
    CreateVote,
    GetVoteById,
    GetVoteBySlug,
    GetVotesByUserEligibility,
  },
};
