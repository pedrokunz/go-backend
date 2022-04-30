FROM golang:alpine

LABEL mantainer="Pedro Kunz <pedro.kunz@gmail.com>"

RUN apk add --no-cache \
    build-base \
    ca-certificates \
    git

RUN echo 'Installing Go tools'
RUN go install github.com/ramya-rao-a/go-outline@latest \
    && go install github.com/cweill/gotests/gotests@latest \
    && go install github.com/fatih/gomodifytags@latest \
    && go install github.com/josharian/impl@latest \
    && go install github.com/haya14busa/goplay/cmd/goplay@latest \
    && go install github.com/go-delve/delve/cmd/dlv@latest \
    && go install honnef.co/go/tools/cmd/staticcheck@latest \
    && go install golang.org/x/tools/gopls@latest

# Set the working directory to the root of the repository
WORKDIR /app

# Copy the source code
COPY . .

# Download and install dependencies
RUN go mod tidy

# Build the application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o main .

EXPOSE 8080
