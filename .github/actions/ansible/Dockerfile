FROM lexauw/ansible-alpine

WORKDIR /github/workspace
RUN apk add rsync

COPY entrypoint.sh /usr/local/bin/entrypoint
ENTRYPOINT [ "entrypoint" ]