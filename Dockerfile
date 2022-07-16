FROM alpine:3.16.0

COPY bin/unity-packager-linux-amd64 /bin/unity-packager

WORKDIR /home/src