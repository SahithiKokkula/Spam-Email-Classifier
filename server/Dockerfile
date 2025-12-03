FROM golang:1.21-alpine AS go-builder

WORKDIR /app
COPY server/go.mod ./
COPY server/go.sum ./
RUN go mod download
RUN go mod verify
COPY server/main.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o server .

FROM python:3.10-slim

RUN apt-get update && apt-get install -y --no-install-recommends \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY --from=go-builder /app/server .
COPY server/predict.py .
COPY server/EmailClassifier.keras .
COPY server/tokenizer_email.pickle .
COPY frontend ./frontend

RUN pip install --no-cache-dir tensorflow numpy nltk

RUN mkdir -p /usr/local/nltk_data
RUN python -m nltk.downloader stopwords wordnet -d /usr/local/nltk_data

EXPOSE 8080

CMD ["./server"]

