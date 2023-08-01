FROM gcr.io/oss-fuzz-base/base-builder-go:v1
RUN apt-get update && apt-get install -y wget
RUN wget https://go.dev/dl/go1.20.6.linux-amd64.tar.gz \
    && mkdir tmp-go \
    && rm -rf /root/.go/* \
    && tar -C tmp-go/ -xzf go1.20.6.linux-amd64.tar.gz \
    && mv tmp-go/go/* /root/.go/
COPY . $SRC/fuzz
WORKDIR fuzz
COPY .clusterfuzzlite/build.sh $SRC/
