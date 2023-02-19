FROM python:3.8-slim-buster
SHELL ["/bin/bash","-c"]
WORKDIR /TASK
RUN chmod 755 /TASK
COPY . .
RUN /bin/bash
RUN groupadd -r swuser -g 433 && \
    useradd -u 431 -r -g swuser -s /sbin/nologin -c "Docker image user" swuser
RUN pip3 install -r requirement.txt
RUN chmod 777 pin.txt

USER swuser
CMD source ./pin_generator.sh && python3 main.py
 
