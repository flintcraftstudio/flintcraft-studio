FROM golang:1.25 AS build

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Build CSS — the marketing site plus each demo vertical's own stylesheet.
# Add a line here when a new demo vertical is added (see tailwind/demos/).
RUN curl -sL https://github.com/tailwindlabs/tailwindcss/releases/download/v3.4.17/tailwindcss-linux-x64 -o /usr/local/bin/tailwindcss && \
    chmod +x /usr/local/bin/tailwindcss && \
    tailwindcss -c tailwind/tailwind.config.js -i tailwind/input.css -o web/static/css/site.css --minify && \
    tailwindcss -c tailwind/demos/chiropractor.config.js -i tailwind/demos/chiropractor.css -o demos/chiropractor/static/css/site.css --minify && \
    tailwindcss -c tailwind/demos/law.config.js -i tailwind/demos/law.css -o demos/law/static/css/site.css --minify

# Generate templ and build Go binary
RUN go install github.com/a-h/templ/cmd/templ@latest && \
    templ generate
RUN CGO_ENABLED=0 go build -o /bin/server ./cmd/server

FROM alpine:3.21
RUN apk add --no-cache ca-certificates
COPY --from=build /bin/server /bin/server
COPY --from=build /src/web /web
# Demo sites' static assets (CSS, fonts, images). Served from
# demos/<vertical>/static relative to CWD (/), so the path must match.
COPY --from=build /src/demos /demos

EXPOSE 8080
ENTRYPOINT ["/bin/server"]
