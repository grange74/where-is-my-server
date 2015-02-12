FROM scratch

COPY where-is-my-server /

ENTRYPOINT ["/where-is-my-server"]

EXPOSE 8080