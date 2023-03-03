# Netforenotes

Simple notes project for NetForemost interview process. Watch it here [netforenotes.vercel.app](https://netforenotes.vercel.app/ "netforenotes.vercel.app").

## User stories

Time spent in the backend development by user story.

| Usert story           | Golang code | Tests | Total |
| --------------------- | ----------- | ----- | ----- |
| Project setup         | 2           | -     | **2** |
| Create notes          | 2           | 2     | **4** |
| Get and sort notes    | 3           | 2     | **5** |
| Edit and delete notes | 3           | 2     | **5** |
| Search notes          | 2           | 2     | **4** |

## Running tests

```
go test ./...
```

## Running locally

### Go server

It needs a `APP_PORT` env variable, so from the root directory:

```bash
APP_PORT=4000 go run main.go
```

### Svelte front-end

A `VITE_API_URL` env variable that points to the Golang server is needed. Vite allow us to put env variables in a `.env` file, so create one with the command `touch interface/web/.env` and add the following content:

```
VITE_API_URL=http://localhost:4000
```

Navigate to `web`, install the dependencies and run the application:

```bash
cd interface/web
pnpm install
pnpm dev
```
