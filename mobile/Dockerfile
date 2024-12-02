FROM alpine:latest AS build-apk

WORKDIR /app/public/apk

# Do the COPY command once the files and folder are ready

CMD ["echo" "APK has been completed"]

FROM java-install AS run

WORKDIR /app

COPY ./ /app

CMD ["npm", "start"]