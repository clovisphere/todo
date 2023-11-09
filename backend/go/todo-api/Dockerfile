# Build application from source
FROM golang:1.21-bullseye AS build-stage
WORKDIR /src

# To make argument value available for use in `-ldflags`
ARG RELEASE

COPY go.mod go.sum ./
RUN go mod download -x
RUN go mod verify

COPY . ./

RUN echo "building server using release version: ${RELEASE}"
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-X main.release=${RELEASE}" -o /bin/a.out ./cmd/server


# Deploy application binary into a lean image
FROM scratch
WORKDIR /

COPY --from=build-stage /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=build-stage /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build-stage /etc/passwd /etc/passwd
COPY --from=build-stage /etc/group /etc/group
COPY --from=build-stage /bin/a.out /a.out

CMD ["/a.out"]
