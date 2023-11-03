############################
# STEP 1 build executable binary
############################
FROM golang:1.21.3-alpine3.17 AS build

WORKDIR /go/src/app

# Install git + SSL ca certificates.
# Git is required for fetching the dependencies.
# Ca-certificates is required to call HTTPS endpoints.
RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates


# Create appuser
ENV USER=appuser
ENV UID=10001

RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"


# Download dependency    
COPY go.mod ./
RUN go mod download


# Build Project
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/app ./cmd/proxy/main.go


############################
# STEP 2 build a small image
############################
FROM scratch


# Import Certs & SSL
COPY --from=build /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /etc/passwd /etc/passwd
COPY --from=build /etc/group /etc/group


# Copy executable
COPY --from=build /go/bin/app /


# Entrypoint
CMD ["/app"]