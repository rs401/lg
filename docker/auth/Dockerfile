FROM scratch

# Need to add ENV variables with sane defaults or no default

COPY ./authsvc ./
# COPY ./.env ./

EXPOSE ${AUTHSVC_PORT}

ENTRYPOINT ["/authsvc"]