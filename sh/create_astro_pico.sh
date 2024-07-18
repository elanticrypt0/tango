#!/bin/bash

mkdir -p ../frontend
cd frontend

echo 'Creating frontend'

npm create astro@latest && npx astro add svelte && npm install @picocss/pico