#!/bin/bash
sudo docker rm -f web_flask_container
sudo docker build -t web_flask_image ./task
#sudo docker run --name web_flask_container -p <put a public port here>:5000 web_flask_image
sudo docker run --name web_flask_container -p 5010:5000 -m 4g --restart on-failure web_flask_image 
exit
