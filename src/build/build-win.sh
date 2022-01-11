
#!/bin/bash
path=$(dirname "$0")
cd $path
CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 CGO_ENABLED=1 go build --trimpath -o ../../bin/herbwebdavd.exe ../
