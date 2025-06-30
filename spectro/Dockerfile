# Use the official nginx image as the base
FROM nginx:latest

# Copy your application's files (e.g., HTML, CSS, JS) into the Nginx web root
COPY ./public_html /usr/share/nginx/html

# Optional: Copy your custom nginx.conf if you need to modify it
# COPY nginx.conf /etc/nginx/nginx.conf

# Expose port 80 (the default HTTP port)
EXPOSE 80

# Optional: If you want to start nginx with a specific command, you can use ENTRYPOINT
# ENTRYPOINT ["/usr/sbin/nginx", "-g", "daemon off;"]
