FROM golang
RUN mkdir /app
ADD logtest /app/logtest
ADD conf /app/conf
WORKDIR /app
EXPOSE 9090
ENTRYPOINT /app/logtest

CMD ["bee", "run"]