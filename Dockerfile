FROM scratch
COPY .bin/watworks-public-api /app/watworks-public-api
WORKDIR /app
EXPOSE 80
CMD ['/app/watworks-public-api']
