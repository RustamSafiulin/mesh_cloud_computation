FROM centos:7

WORKDIR /app

ADD bin/job_service /app/

EXPOSE 8082

RUN chmod +x ./job_service

CMD ./job_service