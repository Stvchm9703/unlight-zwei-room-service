WORKDIR /app
COPY ./room_status ./

COPY ./insecure ./insecure

COPY ./config.yaml ./config.yaml

EXPOSE 11000

CMD [ "./roomserver", "start", "-c", "config.yaml" ]



