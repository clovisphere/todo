FROM golang:1.21-bullseye AS base
WORKDIR /src

# To make argument value available for use in `-ldflags`
ARG RELEASE

COPY go.mod go.sum ./
RUN go mod download -x
RUN go mod verify

COPY . ./

RUN echo "building server using release version: ${RELEASE}"
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-X main.release=${RELEASE}" -o /bin/a.out ./cmd/server


FROM scratch
WORKDIR /app

COPY --from=base /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=base /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=base /etc/passwd /etc/passwd
COPY --from=base /etc/group /etc/group

COPY --from=base /bin/a.out ./

CMD ["./a.out"]
