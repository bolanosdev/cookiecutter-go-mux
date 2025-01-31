FROM golang:{{ cookiecutter.go_version }}

ENV API_ADDRESS http://localhost:9080

WORKDIR /opt/{{ cookiecutter.project_name }}

COPY . .
COPY ./bin/wait-for-it.sh /bin/wait-for-it.sh

CMD ["make integration"]
