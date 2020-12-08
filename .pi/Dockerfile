FROM faddat/sos

# later, update to tendermint/starport
COPY --from=builder /spn /spn

# Make SPN user, spn hostname so we get spn.local via mdns
RUN useradd --create-home spn && \
        echo "spn" > /etc/hostname && \
        echo "spn:spn" | chpasswd && \
        usermod --append --groups wheel spn
