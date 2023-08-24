FROM golang:1.16

USER golang
WORKDIR /home/go/app
ENV PATH="/go/bin:${PATH}"

CMD ["tail","-f","/dev/null"]