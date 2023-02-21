FROM python:3-slim
WORKDIR /app
RUN pip3 install apprise
COPY junction .
EXPOSE 8025
CMD "/app/junction"