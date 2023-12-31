FROM archlinux

RUN pacman -Sy go --noconfirm

WORKDIR /app
COPY .. .

RUN go build

CMD ["/sbin/init"]
