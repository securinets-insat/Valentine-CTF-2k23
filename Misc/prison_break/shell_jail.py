import threading
import socket
import subprocess
def process_command(command):
    # Prompt the user for a command

    if len(command)>4 :
        return b"Keep Trying ... \n"
    # Filter the input to only allow 4 characters
    command = command[:4]
    if "rm" in command:
        return b"Not Allowed ... \n"
    # Use subprocess to run the command
    output = subprocess.run(command, shell=True, stdout=subprocess.PIPE, stderr=subprocess.PIPE)

    # Print the output
    return output.stdout
 

def handle_client(conn, addr):
    with conn:
        print(f"Accepted connection from {addr}")
        while True:
            data = conn.recv(1024)
            if not data:
                break
            result = process_command(data.decode().strip())
            conn.sendall(result)

with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
    s.bind(("0.0.0.0", 8080))
    s.listen()
    print("Listening for incoming connections")
    while True:
        conn, addr = s.accept()
        t = threading.Thread(target=handle_client, args=(conn, addr))
        t.start()
