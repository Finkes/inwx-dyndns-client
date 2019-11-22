FROM arm32v7/alpine

WORKDIR /app
COPY inwx-dyndns-client /app
RUN chmod +x /app/inwx-dyndns-client

CMD ["./inwx-dyndns-client"]
