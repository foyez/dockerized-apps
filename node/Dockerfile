# base stage
FROM node:22-slim AS base

# create and set the user to node
USER node

WORKDIR /app

# Cache dependencies by copying package*.json
COPY --chown=node:node package*.json .

# default port
ENV PORT=3000
# Expose port (use the default or overridden value)
EXPOSE ${PORT}

# dev stage
# since we bind-mount the source code, we don't need to copy the source code
FROM base AS dev
ENV NODE_ENV=development
RUN npm i --only=development
CMD ["npm", "run", "dev"]

# build stage
FROM base AS build
RUN npm ci
# --chown=node:node is used to set the user and group to node
COPY --chown=node:node . .
RUN npm run build

# prod stage
FROM base AS prod
ENV NODE_ENV=production
COPY --chown=node:node --from=build /app/dist ./dist
CMD ["npm", "start"]