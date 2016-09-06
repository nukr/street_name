FROM scratch

COPY street_name /street_name
COPY streetName /streetName
COPY country_list /country_list

CMD ["/street_name"]
