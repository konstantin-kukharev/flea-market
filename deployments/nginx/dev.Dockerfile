FROM nginx:latest

# Copy Nginx configuration file to the container
COPY deployments/nginx/dev.default.conf /etc/nginx/conf.d/default.conf
# Expose port 80
EXPOSE 80