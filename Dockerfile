FROM golang:1.15 as builder

COPY . /go/src/github.com/zack-jack/pedal-planner-api-v1/
WORKDIR /go/src/github.com/zack-jack/pedal-planner-api-v1/cmd/server

RUN git rev-parse HEAD > /root/commit
RUN git describe --abbrev=0 --tags --always > /root/tag

RUN CGO_ENABLED=0 go build -mod=vendor -ldflags "-s -w -X main.commit=$(cat /root/commit) -X main.gitTag=$(cat /root/tag)" -v

FROM alpine:3.13.0

# create non-root user
RUN apk --no-cache add ca-certificates \
  && apk add shadow \
  && groupadd -r app \
  && useradd -r -g app -s /sbin/nologin -u 1001 -c "Docker image user" app

# run the remaining processes as non-root user
USER app

WORKDIR /server
COPY --from=builder /go/src/github.com/zack-jack/pedal-planner-api-v1/cmd/server/server .
COPY --from=builder /go/src/github.com/zack-jack/pedal-planner-api-v1/docs ./docs

CMD ["./server"]
