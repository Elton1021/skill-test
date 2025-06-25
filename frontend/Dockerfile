FROM node:18-alpine

WORKDIR /app

# Copy package files
COPY package*.json ./

# Install dependencies
RUN npm ci

# Copy source code
COPY . .

# Set environment variables to prevent auto-open
ENV BROWSER=none
ENV CI=true

# Expose port
EXPOSE 5173

# Start the development server without auto-open
CMD ["npm", "run", "dev", "--", "--host", "0.0.0.0", "--port", "5173"] 