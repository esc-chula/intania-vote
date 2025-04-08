FROM node:20-alpine AS base

##### DEPENDENCIES

FROM base AS deps
RUN apk add --no-cache libc6-compat openssl

WORKDIR /app

COPY package.json pnpm-lock.yaml\* ./

RUN npm install -g pnpm && pnpm i

##### BUILDER

FROM base AS builder
RUN apk add --no-cache libc6-compat openssl

WORKDIR /app

COPY --from=deps /app/node_modules ./node_modules
COPY . .

ENV NEXT_TELEMETRY_DISABLED=1

RUN npm install -g pnpm && SKIP_ENV_VALIDATION=1 pnpm run build