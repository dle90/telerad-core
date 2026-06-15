FROM 905418320982.dkr.ecr.ap-southeast-1.amazonaws.com/alpine:3.22
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk*
WORKDIR /app
RUN mkdir -p ./log

# Copy binary đã được build từ ngoài (bởi CodeBuild)
COPY main .
RUN chmod +x ./main

# Copy các file cấu hình khác
COPY .env /app
COPY public_key_pkcs8.pub /app

EXPOSE 8080

ENTRYPOINT ["./main"]