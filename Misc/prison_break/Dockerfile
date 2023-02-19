FROM python:3.8-alpine
RUN adduser -D -H ctf
RUN mkdir /app
COPY ./shell_jail.py /app
WORKDIR /app/data
USER ctf
CMD ["python", "/app/shell_jail.py"]
expose 8080


