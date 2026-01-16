FROM --platform=linux/amd64 debian:stable-slim
RUN apt-get update && apt-get install -y ca-certificates
COPY notely /usr/bin/notely
CMD ["notely"]
