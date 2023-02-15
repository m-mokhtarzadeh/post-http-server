FROM gcc:4.9 as build
COPY . /usr/src/
WORKDIR /usr/src/
RUN make server

FROM scratch
WORKDIR /data
COPY --from=build /usr/src/server /server
ENV SERVER_PORT=80
ENTRYPOINT ["/server"]
