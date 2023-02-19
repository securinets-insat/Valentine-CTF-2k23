from scapy.all import * 
from random import sample

packets=rdpcap("merged.pcap")
final=[]
i=0
for pack in packets:
    pack[0][2].type=ord(flag[i])
    i+=1

wrpcap("final1.pcap",packets)

