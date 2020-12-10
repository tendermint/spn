FROM faddat/sos

# fetch spn.  Might want to use copy later, but this should only build with changes to master.
RUN pacman -Syyu --noconfirm go && \
        git clone https://github.com/tendermint/spn && \
        cd spn && \
        go build -mod=mod -o . ./cmd/spnd && \
        mv spnd /usr/bin && \
        pacman -R go && \
        cd .. && \
        rm -rf spn

# Make SPN user, spn hostname so we get spn.local via mdns
RUN useradd --create-home spn && \
        echo "spn" > /etc/hostname && \
        echo "spn:spn" | chpasswd && \
        usermod --append --groups wheel spn
        

