# base stage
FROM node:22-slim AS base

# switch to node user
USER node

WORKDIR /app

COPY --chown=node:node package*.json .

ENV PORT=3000
EXPOSE ${PORT}

# dev stage
FROM base AS dev
ENV NODE_ENV=development
RUN npm i --only=development
CMD ["npm", "run", "dev"]

# build stage
FROM base AS build
RUN npm ci
COPY --chown=node:node . .
RUN npm run build

# prod stage
FROM base AS prod
ENV NODE_ENV=production
COPY --from=build --chown=node:node /app/public ./public
COPY --from=build --chown=node:node /app/.next/standalone ./
COPY --from=build --chown=node:node /app/.next/static ./.next/static
CMD ["node", "server.js"]