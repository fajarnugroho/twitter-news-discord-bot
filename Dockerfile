FROM scratch

WORKDIR /app
COPY .docker_build/tndb /app/

CMD ["/app/tndb"]
