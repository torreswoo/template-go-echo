FROM golang:1.13-stretch
# add source code
COPY . /app/api
# create a working directory
WORKDIR /app/api
# modules
RUN go mod download
# build
RUN go build -o bin/server cmd/server/main.go
# run binary & go run cmd/server/main.go
CMD ["/app/api/bin/server"]
