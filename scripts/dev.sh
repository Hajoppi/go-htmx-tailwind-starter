#!/bin/bash

# Start Air for live reloading go lang
~/go/bin/air & \

# Start BrowserSync for refreshing the browser for changes
npx browser-sync start \
    --no-open \
    --proxy 'localhost:1337' & \

# Start Tailwind watcher
npx tailwindcss \
    -i ./public/css/index.css \
    -o ./public/css/generated.css \
    --watch
