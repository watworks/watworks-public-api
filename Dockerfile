FROM alpine:3.6

COPY .bin/watworks-public-api /var/app/watworks-public-api

WORKDIR /var/app

CMD ["./watworks-public-api"]

# TODO: somehow generate the config file needed by the binary upon startup (pull from env vars)