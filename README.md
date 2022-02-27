# Example web app

## Run
```
go run main.go
```

## Build
```
cmd /c build.bat
```

## Deploy
```sh
flyctl create silent-glitter
fly volumes create goapp_data --size 1

# to use the fly.io builder
flyctl deploy
```

Alternatively you can build the image locally

```
cmd /c pack.bat
```

and then deploy the image to fly io.
```sh
flyctl deploy -i golang-webapp
```
