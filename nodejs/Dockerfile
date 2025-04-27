FROM node:18 AS builder

WORKDIR /app

COPY package.json package-lock.json ./

RUN npm install

COPY . .

FROM node:18-slim

WORKDIR /app

COPY --from=builder /app /app

EXPOSE 8080

CMD ["node", "index.js"]

