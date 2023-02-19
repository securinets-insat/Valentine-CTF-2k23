#include <stdio.h>
#include <stdlib.h>
#include <string.h> 
#include <openssl/sha.h>
#include <openssl/conf.h>
#include <openssl/evp.h>
#include <openssl/err.h>

/*
    Author : ironbyte !
    follow me on ironbyte.me
*/

unsigned char iv[] = {'0','1','2','3','4','5','6','7','8','9','0','1','2','3','4','5'};
unsigned char flag[] = {0x13,0xb7,0x7e,0x1a,0x24,0x46,0x1a,0xc6,0x3f,0x68,0x2f,0xb7,0x63,0xf8,0x9e,0x33,0x98,0xa8,0xab,0x93,0x87,0x22,0xc3,0x1,0x45,0xd0,0xb4,0x4,0xb1,0x4,0xed,0xea};

void fail() {
    ERR_print_errors_fp(stderr);
    puts("Patching won't help you!");
    exit(0);
}

int decrypt(unsigned char *ciphertext, int ciphertext_len, unsigned char *key,
            unsigned char *iv, unsigned char *plaintext) {
    EVP_CIPHER_CTX *ctx;
    int len;
    int plaintext_len;
    if(!(ctx = EVP_CIPHER_CTX_new()))
        fail();
    if(1 != EVP_DecryptInit_ex(ctx, EVP_aes_256_cbc(), NULL, key, iv))
        fail();
    if(1 != EVP_DecryptUpdate(ctx, plaintext, &len, ciphertext, ciphertext_len))
        fail();
    plaintext_len = len;
    if(1 != EVP_DecryptFinal_ex(ctx, plaintext + len, &len))
        fail();
    plaintext_len += len;
    EVP_CIPHER_CTX_free(ctx);
    return plaintext_len;
}


int check(char* input) {
    int x = 0, y = 0;
    for(int i = 0; i < strlen(input); i++) {
        if (input[i] == 'U')
            x--; 
        if (input[i] == 'D')
            x++; 
        if (input[i] == 'L')
            y--; 
        if (input[i] == 'R')
            y++;
    }
    if (x == -3 && y == -23)
        return 1; 
    else 
        return 0;
}

int main() {
    char input[27];
    char key[SHA256_DIGEST_LENGTH];
    char plaintext[128];
    int length;

    puts("It's Valentine, mishka is in a maze. He protected his valentine gift using a PASS! Reverse it!"); 
    puts("Password : ");
    fgets(input, sizeof(input), stdin);
 
    if (check(input)) {
        puts("Correct! Printing flag!");
        SHA256(input, 26, key);
        length = decrypt(flag, 32, key, iv, plaintext);
        plaintext[length] = 0;
        printf("Flag = %s\n", plaintext);
    } else {
        puts("Mishka got eaten by a goblin in the maze!"); 
    }
           
    return 0;
}
