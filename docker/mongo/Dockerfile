FROM mongo:4.2

# Set environment values for configuration
ENV MONGO_INITDB_DATABASE=alerts_db

COPY prepopulate.js /docker-entrypoint-initdb.d/prepopulate.js

# Expose port 27017 for MongoDB communication
EXPOSE 27017