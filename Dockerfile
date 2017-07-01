FROM alpine:3.6
COPY .bin/watworks-public-api /app/watworks-public-api
WORKDIR /app
EXPOSE 80
ENTRYPOINT /app/watworks-public-api
