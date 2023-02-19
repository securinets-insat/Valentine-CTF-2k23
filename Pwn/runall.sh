#! /bin/bash

cd GiftShell
sudo docker-compose down
sudo docker-compose build
sudo docker-compose up -d
echo "GiftShell up"

cd "../Love Letter"
sudo docker-compose down
sudo docker-compose build
sudo docker-compose up -d
echo "Love letter up"


