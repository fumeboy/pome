FROM pome:v2
ENV PYTHONUNBUFFERED=1
COPY ./main.py /pome/service.py
RUN apk add --no-cache python3 && ln -sf python3 /usr/bin/python
COPY ./entrypoint.sh /pome/
COPY ./sidecar.json /pome/
WORKDIR /pome
CMD sh ./entrypoint.sh