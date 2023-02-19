from Crypto.Cipher import AES
from Crypto.Hash import SHA256
from Crypto.Util.Padding import pad

flag = b'Securinets{Love_in_the_air!}'
hasher = SHA256.new()
hasher.update(b'ULLULULLLLLLLLLLLLLLLLLLLL')
key = hasher.digest()
iv = b'0123456789012345'
cipher = AES.new(key, AES.MODE_CBC, iv)
b = cipher.encrypt(pad(flag, AES.block_size))
f = open("flag.enc", "wb")
f.write(b)
# print(b)
print("{", end = "")
for elt in b: 
    print(hex(elt), end = ",") 
print("}")


# ct = b"#\x8b\xb6\xb1\x18\x85XV/U\xcc\x93`\x85\xd8\x92\x0e\x90\xbfR\x83\xea\x06\x0b\x0b\xf3aZ,\x0eX\xbe"
# cipher = AES.new(key, AES.MODE_CBC, iv)
# print(cipher.decrypt(ct))
# print(key)