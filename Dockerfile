FROM scratch

COPY street_name /street_name
COPY streetName /streetName

CMD ["/street_name"]
