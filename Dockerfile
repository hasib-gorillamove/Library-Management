FROM alpine:latest as certs
RUN apk --update add ca-certificates
FROM scratch
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY golden-infotech-service ./
COPY config.env ./
EXPOSE 2001/tcp
ENTRYPOINT ["/golden-infotech-service"]