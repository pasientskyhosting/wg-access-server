## Motivation
https://www.docker.com/blog/faster-multi-platform-builds-dockerfile-cross-compilation-guide/

For some reason this command has to be repeated after every reboot to install the target emulator in buildx:
```sh
docker run --privileged --rm tonistiigi/binfmt --install arm64
```

## Local machine

### docker buildx emulation

```
=> [website  7/11] RUN npm ci --no-audit --prefer-offline
219.1s

=> [website 11/11] RUN npm run build
337.3s

=> [server 18/18] RUN go build -o wg-access-server
434.3s
```

### Go cross-compile

```
=> [website  7/11] RUN npm ci --no-audit --prefer-offline
27.5s

=> [website 11/11] RUN npm run build
38.0s

=> [server 20/20] RUN CC=`pwd`/${CARCH}-linux-musl-cross/bin/${CARCH}-linux-musl-gcc GOOS=linux GOARCH=arm64 go build -o wg-access-server
44.5s
```

## GitHub CI
<https://github.com/freifunkMUC/wg-access-server/runs/4553130423?check_suite_focus=true>

### docker buildx emulation

```
#76 [linux/amd64 server 18/18] RUN go build -o wg-access-server
#76 DONE 135.3s

#40 [linux/arm/v7 server 18/18] RUN go build -o wg-access-server
#40 DONE 950.2s

#111 [linux/arm64 server 18/18] RUN go build -o wg-access-server
#111 DONE 1044.3s
```
