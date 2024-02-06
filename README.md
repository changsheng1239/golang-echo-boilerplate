# Golang Echo Boilerplate

This is a boilerplate to start a Golang API project using [echo](https://echo.labstack.com/) web framework.

## Frameworks / Libraries included

- [echo](https://github.com/labstack/echo)
- [charmbracelet/log](https://github.com/charmbracelet/log)
- [cleanenv](https://github.com/ilyakaznacheev/cleanenv)
- [godotenv](https://github.com/joho/godotenv)

## Development Tools

- Makefile
- Devcontainer

## Running

1. Clone this repo.
2. Open the repo using `vscode`.
3. Open the folder in devcontainer (vscode: F1 > Dev Containers: Rebuild and Reopen in Container)
4. Edit `rename-module.sh` and replace `NEW_MODULE_NAME` with your project name.
5. Run `rename-module.sh` to rename the module name.
6. Copy `.env` in place & start using `make`

```shell
cp ./config/.env.example .env
make dev
```

The API server should be running at `http://localhost:8080`. Verify it using `curl`

```shell
curl http://localhost:8080/user/1
# {"ID":1,"CreatedAt":"2024-02-06T10:00:49.4952092+08:00","UpdatedAt":"2024-02-06T10:00:49.4952092+08:00","DeletedAt":"0001-01-01T00:00:00Z","Username":"test","Email":"test@example.com"}
```

## Logging
[charmbracelet/log](https://github.com/charmbracelet/log) was used to enable pretty & coloured printing during development.

Set `ENV` variable as `production` to enable logging in JSON format.

## Environment Variable
|Variable|Description|
|-|-|
|ENV|development/production, default: `development`|
|SERVER_HOST| run on which interface, default: `localhost`|
|SERVER_PORT| run on which port, default: `8080`|