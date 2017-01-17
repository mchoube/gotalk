FROM alpine:latest

COPY gowebapp /app

ENTRYPOINT /app/gowebapp

EXPOSE 8080
