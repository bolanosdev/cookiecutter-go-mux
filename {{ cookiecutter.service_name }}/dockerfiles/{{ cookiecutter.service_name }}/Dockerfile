FROM {{ cookiecutter.docker_image }} 
WORKDIR /app

# COPY /internal/db/migrations/. /app/internal/db/migrations
COPY dist/linux-amd64/. /app

CMD ["/app/{{ cookiecutter.service_name}}"]
