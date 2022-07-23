# ============
# Dev
# ============
FROM golang:1.18.3 as Dev
ENV TZ=Asia/Tokyo

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV ROOTPATH=/go/app

WORKDIR ${ROOTPATH}

RUN go install github.com/cosmtrek/air@v1.29.0
COPY go.mod go.sum .air.toml ./
RUN go mod download

COPY src/ ${ROOTPATH}/src/

EXPOSE 80
CMD ["air", "-c", ".air.toml"]

# ============
# Builder
# ============
FROM golang:1.18.3 as Builder
ENV TZ=Asia/Tokyo

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV ROOTPATH=/go/app

WORKDIR ${ROOTPATH}

COPY go.mod go.sum ./
RUN go mod download

COPY src/ ${ROOTPATH}/src/

RUN go build \
    -o ${ROOTPATH}/bin/main \
    -ldflags '-s -w' \
    ./src/main.go


# ============
# Deploy
# ============
FROM gcr.io/distroless/static-debian11
ENV TZ=Asia/Tokyo

COPY --from=Builder /go/app/bin/main /app/main

EXPOSE 80
USER nonroot:nonroot
CMD ["/app/main"]