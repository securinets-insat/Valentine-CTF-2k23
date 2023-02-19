#!/bin/bash
nohup ./change_port_from_here.sh &
echo "[+] cleaning some files"
sleep 10
sudo docker exec -u root --tty -it web_flask_container /bin/bash -c "rm pin.txt"
sudo docker exec -u root --tty -it web_flask_container /bin/bash -c "rm pin_generator.sh"
sudo docker exec -u root --tty -it web_flask_container /bin/bash -c "rm requirement.txt"
echo "[*] task is up and running"
