FROM apline
WORKDIR /app
COPY ./room_status ./
COPY ./insecure ./insecure
COPY ./config.test_server.yaml ./config.yaml

EXPOSE 11000
ENTRYPOINT ["room_server"]
CMD [ "start", "-c", "config.yaml" ]



