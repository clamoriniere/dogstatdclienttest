FROM scratch

ADD ddclient-linux /ddclient

ENTRYPOINT [ "/ddclient" ]