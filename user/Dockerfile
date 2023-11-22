FROM alpine:latest
WORKDIR /user
COPY configs/base.yaml . 
COPY cmd/.env .
COPY main.exe .
EXPOSE 8081
CMD ["./main.exe"]