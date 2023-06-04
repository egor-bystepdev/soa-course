FROM golang:1.19

WORKDIR /app

# Copy application source.
COPY . .

RUN mkdir -p bin

# Download dependencies.
RUN go mod download && go mod verify

# Build executable
RUN go build -o bin/spammer chat_spam/main.go

