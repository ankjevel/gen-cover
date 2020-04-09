FROM golang:1.13-stretch AS builder

WORKDIR /build

COPY . .

RUN make release

# # #

FROM golang:1.13-stretch

WORKDIR /app

COPY --from=builder /build/release/gen-cover /app

ENV \
    HOST=0.0.0.0 \
    PORT=8091 \
    ADDR=

EXPOSE 8091/tcp

RUN adduser --disabled-password --gecos '' wejay && \
    chmod -R g+rwX         /app && \
    chgrp -R wejay         /app && \
    chown -R wejay:wejay   /app

USER wejay

ENTRYPOINT [ "/app/bin" ]
CMD [  ]
