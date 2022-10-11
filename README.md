# server on GoLang

## Create env
```
cp .env.example .env
```
### Fill empty keys in .env

## Create docker-compose.override.yml
```
cp docker-compose.example.yml docker-compose.override.yml
```
### Change ports if necessary (if changed app port, change this port in .env file on client)

## Project up with docker-compose
```
docker-compose up -d --build
```

