docker build -t shippy-service-consignment ./shippy-service-consignment && \
docker run -p 50051:50051 -e MICRO_SERVER_ADDRESS=:50051 shippy-service-consignment
