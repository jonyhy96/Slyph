FROM golang:1.8
LABEL maintainer="jonyhy96"
WORKDIR /go/src/Slyph
COPY . /go/src/Slyph
COPY .kube /.kube
ENV DB_HOST "db"
ENV KUBERNETES_MASTER http://${KUBERNETES_MASTER}:8080
ENV TZ Asia/Shanghai
EXPOSE 8888
CMD ["./Slyph"]