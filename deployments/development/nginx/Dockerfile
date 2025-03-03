FROM nginx:latest

# Copy Nginx configuration file to the container
COPY default.conf /etc/nginx/conf.d/default.conf

ENTRYPOINT ["nginx-debug"]