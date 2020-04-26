FROM centos:7

WORKDIR /app

ADD bin/account_service /app/
ADD bin/public/dist /app/public/dist

EXPOSE 8081

RUN chmod +x ./account_service

CMD ./account_service