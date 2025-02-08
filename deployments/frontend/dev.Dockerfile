FROM node:latest

# Set the working directory
WORKDIR /app

# Copy the package.json and package-lock.json files
COPY frontend/package*.json ./

# Install dependencies
RUN npm install

# Copy everything to the container
COPY ./frontend .

# Expose port
EXPOSE 5173

CMD [ "npm", "run", "dev" ]