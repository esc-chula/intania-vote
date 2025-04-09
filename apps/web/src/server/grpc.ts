import { type ServiceError, credentials } from "@grpc/grpc-js";
import {
  type CreateUserRequest,
  type CreateUserResponse,
  type GetUserByOidcIdRequest,
  type GetUserByOidcIdResponse,
  UserServiceClient,
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

export const grpc = {
  user: {
    CreateUser,
    GetUserByOidcId,
  },
};
