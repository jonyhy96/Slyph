FROM registry.docker-cn.com/library/golang:1.8
WORKDIR /go/src/k8s
COPY . /go/src/k8s
ENV KUBERNETES_MASTER http://192.168.34.41:8080
ENV TZ Asia/Shanghai
MAINTAINER haoyun
EXPOSE 8888
CMD ["Slyph"]