FROM tendermintdevelopment/starport AS builder

RUN pacman -S --noconfirm go
ENV PATH=$PATH:/usr/local/go/bin

COPY . /spn
WORKDIR /spn
RUN starport build


FROM faddat/sos

# later, update to tendermint/starport
COPY --from=builder /spn /spn

# Make SPN user, spn hostname so we get spn.local via mdns
RUN useradd --create-home spn && \
        echo "spn" > /etc/hostname && \
        echo "spn:spn" | chpasswd && \
        usermod --append --groups wheel spn

