#!/bin/bash

nix_like=('darwin/amd64' 'darwin/arm64' 'linux/amd64' 'linux/arm' 'linux/arm64' 'linux/mips' 'linux/mipsle' )

for osarch in "${nix_like[@]}"; do
    GOOS=${osarch%/*} GOARCH=${osarch#*/} go build -ldflags="-w -s" github.com/xvzc/SpoofDPI/cmd/spoof-dpi &&
        tar -zcvf spoof-dpi-${osarch%/*}-${osarch#*/}.tar.gz ./spoof-dpi &&
        rm -rf ./spoof-dpi
done

for osarch in "${nix_like[@]}"; do
    GOOS=${osarch%/*} GOARCH=${osarch#*/} CGO_ENABLED=0 go build -ldflags="-w -s" github.com/xvzc/SpoofDPI/cmd/spoof-dpi &&
        tar -zcvf spoof-dpi-${osarch%/*}-${osarch#*/}-self-contained.tar.gz ./spoof-dpi &&
        rm -rf ./spoof-dpi
done

for osarch in 'windows/amd64'; do
    GOOS=${osarch%/*} GOARCH=${osarch#*/} go build -o spoof-dpi-${osarch%/*}-${osarch#*/}.exe -ldflags="-w -s" github.com/xvzc/SpoofDPI/cmd/spoof-dpi 
done
