FROM golang:1.21-bookworm

ARG KUBECTL_VERSION="v1.29.2"
# aws-cli, azure-cli and kubectl
RUN apt-get update \
 && apt-get --assume-yes upgrade \
 && apt-get --no-install-recommends --assume-yes install ca-certificates curl jq unzip apt-transport-https lsb-release gnupg libc-bin binutils wget \
 && curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"  \
 && unzip awscliv2.zip  \
 && ./aws/install -i /usr/local/aws-cli -b /usr/local/bin \
 && export AZ_DIST=$(lsb_release -cs) \
 && mkdir -p /etc/apt/keyrings \
 && curl -sLS https://packages.microsoft.com/keys/micros.asc | gpg --dearmor | tee /etc/apt/keyrings/microsoft.gpg > /dev/null \
 && chmod go+r /etc/apt/keyrings/microsoft.gpg \
 && echo "deb [arch=`dpkg --print-architecture` signed-by=/etc/apt/keyrings/microsoft.gpg] https://packages.microsoft.com/repos/azure-cli/ $AZ_DIST main" | tee /etc/apt/sources.list.d/azure-cli.list \
 && apt-get update \
 && apt-get --no-install-recommends --assume-yes install azure-cli \
 && wget --quiet "https://dl.k8s.io/release/$KUBECTL_VERSION/bin/linux/amd64/kubectl" -O /usr/local/bin/kubectl \
 && chmod +x /usr/local/bin/kubectl \
 && mkdir /app

COPY . /app

WORKDIR /app

# Build binary
RUN go build -o snd ./cmd/snd-cli/

# Create symlink
RUN ln -s "/app/snd" "/bin/snd"

