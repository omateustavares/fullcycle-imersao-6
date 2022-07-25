#!/bin/sh

if [ ! -f ".env" ]; then
   cp .env.example .env
fi

npm install

npm run dev

{ "emails": [ "mateustm17@gmail.com" ], "subject": "Novos tweets encontrados", "body": "Acesse o link <a href=\"http://localhost:3001/tweets\">Clique aqui</a>" }
