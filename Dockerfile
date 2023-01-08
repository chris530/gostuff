FROM golang

RUN mkdir /app
COPY . /app/
WORKDIR /app
ENV GOPATH /go

RUN go install github.com/chris530/gostuff/cmd@main
RUN cp /go/bin/cmd /app/program
RUN chmod 755 /app 
RUN chmod a+rx /app/program
ENTRYPOINT ["/app/program"]
