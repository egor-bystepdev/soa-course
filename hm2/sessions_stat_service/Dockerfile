FROM golang:1.19

WORKDIR /app

# Copy application source.
COPY . .

RUN mkdir -p bin

# Download dependencies.
RUN go mod download && go mod verify

# Build executable
RUN go build -o bin/graph_stat sessions_stat_service/server.go

CMD ["./bin/graph_stat"]
