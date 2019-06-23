#################
# Builder image #
#################
FROM golang:1.12 as builder

ENV GO111MODULE=on

# Install git + SSL ca certificates.
# Git is required for fetching the dependencies.
# CA-certificates is required to call HTTPS endpoints.
RUN apt-get update && apt-get install git ca-certificates && update-ca-certificates

# Create user for the runtime image
RUN adduser --disabled-password --gecos '' serviceuser

# build binary
WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOFLAGS=-mod=vendor go build  -ldflags="-w -s" -o /bin/user ./cmd/user/main.go


#################
# Runtime image #
#################
FROM scratch

#  service binary and assets
COPY --from=builder /bin/user /user
COPY ./migrations /migrations

# import from builder
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# better not run as root
USER serviceuser

ENTRYPOINT ["/user"]

