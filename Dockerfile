FROM golang
RUN go get github.com/astaxie/beego && go get github.com/beego/bee

EXPOSE 9090

CMD ["bee", "run"]