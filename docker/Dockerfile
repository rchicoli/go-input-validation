FROM alpine:3.4

ENV APPLICATION_NAME go-input-validation
ENV PATH /usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/app

RUN apk --update add tzdata \
    && cp /usr/share/zoneinfo/Europe/Berlin /etc/localtime \  
    && addgroup -S app \
    && adduser -s /sbin/nologin -D -S -h /app -G app app

COPY go-input-validation /app/

EXPOSE 8080

USER app
WORKDIR /app

CMD ["go-input-validation"]
