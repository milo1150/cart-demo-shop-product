FROM golang:1.24

WORKDIR /app

# Install Air for live reloading
RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./

RUN go mod tidy && go mod download

# 🔍 Debugging Step: Pause and Inspect the Context
RUN ls -lah /app

CMD ["air", "-c", ".air.toml"]

# Copy the rest of the application code
COPY . .

# Another Debugging Step After Copy
RUN ls -lah /app