FROM alpine:latest
WORKDIR /adventure
COPY configs/base.yaml . 
COPY cmd/.env .
COPY main.exe .
EXPOSE 8082
CMD ["./main.exe"]