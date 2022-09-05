FROM alpine:latest

RUN mkdir /app

COPY application /app

CMD [ "/app/application"]