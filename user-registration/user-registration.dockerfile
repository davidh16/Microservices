FROM alpine:latest

RUN mkdir /app

COPY userRegistrationApp /app

CMD [ "/app/userRegistrationApp" ]