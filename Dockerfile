FROM golang

ADD testinfluxdb.tar.gz /go/src

#装载执行权限
RUN /bin/sh -c 'cd /go/src/ && chmod a+x testinfluxdb'

EXPOSE 8080

#直接执行可执行文件
CMD /bin/sh -c 'cd /go/src/ && ./testinfluxdb'