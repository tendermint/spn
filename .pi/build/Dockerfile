FROM tendermintdevelopment/starport:arm64

RUN pacman -S --noconfirm go
ENV PATH=$PATH:/usr/local/go/bin

COPY . /spn
WORKDIR /spn
RUN starport build
