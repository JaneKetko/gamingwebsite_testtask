FROM alpine
COPY gamingwebsite_testtask .
ENTRYPOINT ["/gamingwebsite_testtask"]