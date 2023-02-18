FROM alpine
COPY ./server /server
EXPOSE 80/tcp
ENTRYPOINT ["/server"]