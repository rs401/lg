FROM scratch

# Need to add ENV variables with sane defaults or no default

COPY ./authapi ./
COPY ./.env ./

EXPOSE ${API_PORT}

ENTRYPOINT ["/authapi"]