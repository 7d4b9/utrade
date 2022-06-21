FROM golang:1.20
ARG USERNAME=utrade
ARG USER_UID=1000
ARG USER_GID=$USER_UID
RUN mkdir -p /etc/apt/keyrings \
  && groupadd --gid $USER_GID $USERNAME \
  && useradd --uid $USER_UID --gid $USER_GID -m $USERNAME \
  && apt-get update && apt-get install -y \
    software-properties-common \
    lsb-release \
    gnupg \
    ca-certificates \
  && curl -fsSL https://download.docker.com/linux/ubuntu/gpg | gpg --dearmor -o /etc/apt/keyrings/docker.gpg \
  && echo "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/debian \
  $(lsb_release -cs) stable" | tee /etc/apt/sources.list.d/docker.list > /dev/null \
  && echo "deb [signed-by=/usr/share/keyrings/cloud.google.gpg] https://packages.cloud.google.com/apt cloud-sdk main" | tee -a /etc/apt/sources.list.d/google-cloud-sdk.list \
  && curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key --keyring /usr/share/keyrings/cloud.google.gpg add - \
  && apt-get update && apt-get install -y \
    docker-ce \
    docker-ce-cli \
    containerd.io \
    docker-compose-plugin \
    google-cloud-cli \
    unzip \
    sudo \
  # Docker user setup 
  && usermod -aG docker $USERNAME \
  && newgrp docker \
  && touch /var/run/docker.sock \
  && chmod o+rw /var/run/docker.sock \
  # Install Terraform
  && curl -o /tmp/terraform.zip https://releases.hashicorp.com/terraform/1.3.6/terraform_1.3.6_linux_arm.zip \
  && unzip /tmp/terraform.zip -d /usr/local/bin \
  && rm -f /tmp/terraform.zip  \
  # Install Vscode-go extension tools
  && GOPATH=/tmp/go \
    GOBIN=/usr/local/go/bin \
    go run github.com/golang/vscode-go/tools/installtools@latest \
  && ln -sf bash /bin/sh \
  && echo $USERNAME ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/$USERNAME \
  && chmod ug=rw /etc/sudoers.d/$USERNAME

USER ${USERNAME}

RUN echo "export PROMPT_COMMAND='history -a'" >> "/home/$USERNAME/.bashrc"