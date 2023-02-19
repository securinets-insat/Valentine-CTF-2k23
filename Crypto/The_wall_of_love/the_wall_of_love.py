from Crypto.Cipher import AES
from Crypto.Util.number import long_to_bytes,bytes_to_long
from random import choice
key_secret_list=[b"\x00",b"\x01"]


flag=open("flag.txt","rb").read()
# padding the little hole in the relationship to make it works :)
flag=flag+b"F"*(16-len(flag)%16)
# just little
print(flag[:16])

for i in range(2):
    key=b"".join(choice(key_secret_list) for i in range(16))
    cipher=AES.new(key,AES.MODE_ECB)
    flag=cipher.encrypt(flag)
print(flag.hex())
#flag[:16]="git good kiddo S"
#encrypted flag="f9ff02ea6690067ae61e04dedebd1a84805b49c43fa373eddde221921a6c73fee7f6ddbee06806f80372d7a2f592fbb57f1846650e0541942a3db360cd3dabf8"