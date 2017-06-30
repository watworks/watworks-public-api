FROM scratch

COPY .bin/watworks-public-api /app/watworks-public-api
COPY docker-entrypoint.sh /run.sh

WORKDIR /app

EXPOSE 80

ENTRYPOINT ['/run.sh']
