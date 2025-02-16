FROM golang:1.24

WORKDIR /app

# Install Air for live reloading
RUN go install github.com/air-verse/air@latest

# Install Staticcheck
RUN go install honnef.co/go/tools/cmd/staticcheck@latest

COPY go.mod go.sum ./

RUN go mod tidy && go mod download

# Copy the rest of the application code
COPY . .

CMD ["air", "-c", ".air.toml"]