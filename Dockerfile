FROM golang:1.18

WORKDIR /auth-service

COPY . .

ENV NODE_VERSION=16.14.2
RUN apt install -y curl
RUN curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.0/install.sh | bash
ENV NVM_DIR=/root/.nvm
RUN . "$NVM_DIR/nvm.sh" && nvm install ${NODE_VERSION}
RUN . "$NVM_DIR/nvm.sh" && nvm use v${NODE_VERSION}
RUN . "$NVM_DIR/nvm.sh" && nvm alias default v${NODE_VERSION}
ENV PATH="/root/.nvm/versions/node/v${NODE_VERSION}/bin/:${PATH}"
RUN npm install -g nodemon

RUN ["go", "mod", "tidy"]
RUN ["go", "mod", "download"]

ENTRYPOINT nodemon --legacy-watch --signal SIGTERM --exec 'go' run ./cmd/api/main.go

