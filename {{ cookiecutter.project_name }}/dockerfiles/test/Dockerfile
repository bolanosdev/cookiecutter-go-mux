FROM golang:{{ cookiecutter.go_version }}
WORKDIR /opt/{{ cookiecutter.project_name }}
COPY . .
CMD ["make", "integration"]
