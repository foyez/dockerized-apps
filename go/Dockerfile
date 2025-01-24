# base stage
FROM golang:1.23.2-alpine3.19 AS base

# create custom user and group
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
# switch to the custom user
USER appuser:appgroup

WORKDIR /app

# Cache dependencies by copying go.mod and go.sum first
COPY --chown=appuser:appgroup go.mod go.sum ./
RUN go mod download

ENV PORT=3000
EXPOSE ${PORT}

# dev stage
FROM base AS dev
ENV ENVIRONMENT=development
# Install air for live reloading
RUN go install github.com/air-verse/air@latest
# Command to run the application with air for live reloading
CMD ["air", "-c", ".air.toml"]

# build stage
FROM base AS build
COPY --chown=appuser:appgroup . .
# Build the go application
# Using -ldflags="-s -w" removes debugging information from the Go binary, further reducing the size.
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o main main.go

# prod stage
FROM alpine:3.19 AS prod
ENV ENVIRONMENT=production
WORKDIR /app
COPY --from=build /app/main .
COPY app.env .
# create custom user and group
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
# switch to the custom user
USER appuser:appgroup
ENV PORT=3000
EXPOSE ${PORT}
CMD ["/app/main"]