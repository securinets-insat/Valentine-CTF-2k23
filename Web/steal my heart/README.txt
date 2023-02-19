## Setup

chmod +x ./change_port_from_here.sh
chmod +x ./start_from_here.sh
./start_from_here.sh
 
*File to attach in /ToGive folder 
*check descriptionAndFlag for the flag and description
*open "./change_port_from_here.sh", comment line 4, uncomment line 3 & put any port u like in the placeholder. (don't change anything if you want the task to be exposed on port 8001)

* if succssful, you can open VM_PUBLIC_IP:8001 with ur browser

## Description

You don't have to be web guru to solve this, your goal is to achieve RCE though.
Try to understand each line in the provided code.
Got stuck ? 
Resource1: https://werkzeug.palletsprojects.com/en/2.2.x/debug/
Resource2: https://ryanflynndev.medium.com/how-to-use-flasks-debug-mode-f42e2566268

Author: n0s

Flag: Securinets{RCE_My_heart_all_yours!!}
