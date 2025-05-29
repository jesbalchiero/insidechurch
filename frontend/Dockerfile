# Build stage
FROM node:20-alpine AS builder

# Add non-root user
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app

# Copy package files
COPY package.json yarn.lock ./

# Install dependencies with cache mount
RUN --mount=type=cache,target=/app/.yarn-cache \
    yarn install --frozen-lockfile --network-timeout 600000

# Copy project files
COPY . .

# Build the application
RUN yarn build

# Development stage
FROM node:20-alpine AS dev

# Add non-root user
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app

# Copy package files
COPY package.json yarn.lock ./

# Install dependencies with cache mount
RUN --mount=type=cache,target=/app/.yarn-cache \
    yarn install --frozen-lockfile --network-timeout 600000

# Copy project files
COPY . .

# Expose port
EXPOSE 3000

# Start development server
CMD ["yarn", "dev"]

# Production stage
FROM node:20-alpine

# Add non-root user
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app

# Copy only necessary files
COPY --from=builder /app/.output ./.output
COPY --from=builder /app/package.json ./
COPY --from=builder /app/yarn.lock ./

# Install production dependencies with cache mount
RUN --mount=type=cache,target=/app/.yarn-cache \
    yarn install --production --frozen-lockfile --network-timeout 600000

# Set proper permissions
RUN chown -R appuser:appgroup /app

# Switch to non-root user
USER appuser

# Expose port
EXPOSE 3000

# Health check
HEALTHCHECK --interval=30s --timeout=30s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:3000/ || exit 1

# Start command
CMD ["node", ".output/server/index.mjs"] 