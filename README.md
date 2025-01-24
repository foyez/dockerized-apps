# Dockerized Apps

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

A collection of containerized applications using Docker and Docker Compose, featuring:

- React + Vite Application
- Next.js Application
- Node.js Application
- Go Application

## Applications

### React + Vite Application

- React 18 with TypeScript and Vite
- Development server with HMR
- Production build with nginx
- Development port: 4000
- Production port: 80

### Next.js Application

- Production-ready Next.js app with TypeScript
- Uses standalone output for optimal container size
- Development server with hot-reload
- Default port: 3000

### Node.js Application

- Simple Node.js server with TypeScript
- Health check endpoint
- Development mode with ts-node-dev
- Default port: 3000 (configurable via PORT env variable)

### Go Application

- Simple Go HTTP server
- Health check endpoint
- Live reload in development using Air
- Default port: 3000 (configurable via PORT env variable)

## Running the Applications

### Development Mode

Each application can be run in development mode with hot-reload:

```sh
# Navigate to the desired application directory
cd <application-directory>

# Start the development container
docker compose -f docker-compose.yaml -f docker-compose.dev.yaml up -d --build

# Start the development container with custom env file
docker compose -env-file file-name.env \
-f docker-compose.yaml -f docker-compose.dev.yaml up -d --build
```

### Production Mode

To run any application in production mode:

```sh
# Navigate to the desired application directory
cd <application-directory>

# Start the production container
docker compose up -d --build
```

## Stopping the Applications

To stop any running application:

```sh
# Navigate to the application directory
cd <application-directory>

# Stop the containers
docker compose down
```

## Contributing

1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a new Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.
