# Use an official Node runtime as a parent image
FROM node:14

# Set the working directory in the container
WORKDIR /usr/src/app

# Copy package.json and package-lock.json
COPY package*.json ./

# Install any dependencies
RUN npm install

# Bundle app source
COPY . .

# Build the app for production
RUN npm run build

# Install serve to serve the build files
RUN npm install -g serve

# Command to run the app
CMD ["serve", "-s", "build"]
