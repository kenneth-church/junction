FROM python:3-slim
WORKDIR /app
RUN pip3 install apprise==1.7.1
COPY junction .
EXPOSE 8025
CMD "/app/junction"