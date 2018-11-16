FROM registry.docker-cn.com/library/golang:1.8
LABEL maintainer="jonyhy96"
WORKDIR /go/src/k8s
COPY . /go/src/k8s
ENV DB_HOST "192.168.0.22"
ENV KUBERNETES_MASTER http://${KUBERNETES_MASTER}:8080
ENV TZ Asia/Shanghai
EXPOSE 8888
CMD ["./Slyph"]