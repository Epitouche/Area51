FROM node:20.12.2-alpine AS build

WORKDIR /app

COPY ./package.json ./
COPY ./package-lock.json ./

RUN npm install

FROM build AS run

COPY . .

CMD ["npx", "nuxt", "dev"]