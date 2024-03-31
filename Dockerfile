FROM golang as builder
WORKDIR /src/service
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-X main.Version=`git tag --sort=-version:refname | head -n 1`" -o rms-speech -a -installsuffix cgo rms-speech.go
RUN CGO_ENABLED=0 GOOS=linux go build -o speech-cli -a -installsuffix cgo ./cli/main.go

FROM ubuntu:latest

RUN apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    sudo \
    python3.9 \
    python3-distutils \
    python3-pip \
    ffmpeg

RUN pip install --upgrade pip && pip install -U openai-whisper &&  mkdir /app
WORKDIR /app
COPY --from=builder /src/service/rms-speech .
COPY --from=builder /src/service/configs/rms-speech.json /etc/rms/
CMD ["./rms-speech"]