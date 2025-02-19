FROM golang:latest

# Set the current working directory inside the container
WORKDIR /app

# Install air
RUN go install github.com/air-verse/air@latest

# Ensure the $GOPATH/bin is in the $PATH
ENV PATH="/go/bin:${PATH}"

# # Copy go.mod and go.sum files to the workspace
# COPY . .

# Download all dependencies
RUN go mod download && go mod verify

# Copy the source from the current directory to the workspace
COPY . .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["air", "-c", ".air.toml"]
