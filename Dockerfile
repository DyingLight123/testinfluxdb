FROM golang

ADD testinfluxdb.tar.gz /go/src

#װ��ִ��Ȩ��
RUN /bin/sh -c 'cd /go/src/ && chmod a+x testinfluxdb'

EXPOSE 8080

#ֱ��ִ�п�ִ���ļ�
CMD /bin/sh -c 'cd /go/src/ && ./testinfluxdb'