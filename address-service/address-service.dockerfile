FROM alpine:latest

RUN mkdir /app

COPY addressApp /app

CMD [ "/app/addressApp" ]