FROM golang:1.13-alpine

# Adding git, bash and openssh to the image
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o app

EXPOSE 8080

# Run the executable
CMD ["./app"]