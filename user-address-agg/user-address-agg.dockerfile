FROM alpine:latest

RUN mkdir /app

COPY aggregatorApp /app

CMD [ "/app/aggregatorApp" ]