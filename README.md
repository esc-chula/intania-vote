# Intania Vote

<a alt="Nx logo" href="https://nx.dev" target="_blank" rel="noreferrer"><img src="https://raw.githubusercontent.com/nrwl/nx/master/images/nx-logo.png" width="45"></a>

## Getting Started

### Development

Install NodeJS dependencies

```sh
pnpm i
```

Install Go modules

```sh
go mod download
```

To run the dev server for your app, use:

```sh
pnpm dev # pnpm nx run-many -t serve

# for each app
pnpm dev:web
pnpm dev:api
```

### Add shadcn/ui Components

Completely install the components,

```sh
pnpm ui:add # choose the components you need
pnpm ui:post # to update imports and exports path
```

> For Windows, PS1 scripts is provided in ./scripts
