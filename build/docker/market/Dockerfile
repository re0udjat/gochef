########## BUILDING PROCESS ##########
FROM golang:1.23 AS build

# Create working directory inside the container if it does not existed
WORKDIR /app

# Download golang packages into container filesystem
COPY go.sum go.mod ./
RUN go mod download

# Move code into container filesystem for build process
COPY /cmd/market/* ./

# Build services
RUN CGO_ENABLED=0 GOOS=linux go build -o /gochef-market

########## RUNNING PROCESS ##########

# Running distroless base image for:
# 1. Smaller image size to improve download, upload speeds
# 2. Enhanced security because a smaller attack surface makes container less
# vulnerable to exploits.
# 3. Focused and minimalist env
# 4. Performance
# 5. Immutable containers philosophy that everything need to be built into
# the image at build time, ensuring consistency and reliability
FROM gcr.io/distroless/base-debian12 AS release


WORKDIR /

# Move built artifact from previous stage to this stage
COPY --from=build /gochef-market /gochef-market

# Expose port for container
EXPOSE 8080

# Switch user
USER nonroot:nonroot

# Start application
ENTRYPOINT [ "./gochef-market" ]