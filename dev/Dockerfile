ARG go_version
FROM golang:${go_version}
ARG USERNAME=utrade
ARG USER_UID=1000
ARG USER_GID=$USER_UID
RUN apt-get update && apt-get install -y \
   software-properties-common \
   gnupg \
   ca-certificates \
   lsb-release \
   sudo \
  && groupadd --gid $USER_GID $USERNAME \
  && useradd --uid $USER_UID --gid $USER_GID -m $USERNAME \
  && echo $USERNAME ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/$USERNAME \
  && chmod ug=rw /etc/sudoers.d/$USERNAME
USER $USERNAME
RUN sudo mkdir -p /etc/apt/keyrings \
  && curl -fsSL https://apt.releases.hashicorp.com/gpg | sudo apt-key add - \
  && sudo apt-add-repository "deb [arch=amd64] https://apt.releases.hashicorp.com $(lsb_release -cs) main" \
  && curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg \
  && echo "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/debian \
$(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null \
  && echo "deb [signed-by=/usr/share/keyrings/cloud.google.gpg] https://packages.cloud.google.com/apt cloud-sdk main" | sudo tee -a /etc/apt/sources.list.d/google-cloud-sdk.list \
  && curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key --keyring /usr/share/keyrings/cloud.google.gpg add - \
  && sudo apt-get update && sudo apt-get install -y \
    terraform \
    docker-ce \
    docker-ce-cli \
    containerd.io \
    docker-compose-plugin \
    bash-completion \
    google-cloud-cli \
  && terraform -install-autocomplete \
  && sudo usermod -aG docker $USERNAME \
  && sudo newgrp docker \
  && sudo touch /var/run/docker.sock \
  && sudo chmod o+rw /var/run/docker.sock \
  && GOBIN=/go/bin go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest