// gcc main.c -o main -no-pie -fno-stack-protector
#include <stdio.h>

#define MAX_BUFFER 6969
#define MAX_READ_LINE 4000

int readInput(char *input, int len){
    printf("> ");
    return read(0, input, len);
}

void setup(){
    setvbuf(stderr, NULL, _IONBF, 0);
    setvbuf(stdout, NULL, _IONBF, 0);
    setvbuf(stdin, NULL, _IONBF, 0);
}

void loveletter(){
    // AS BIG AS POSSIBLE
    // OVERFLOWS ARE BAD
    char buff[MAX_BUFFER + 69];
    int size = 0;
    int tmp;
    char *cursor = buff;

    puts("I'll take take your love letter & deliver it.");
    puts("You can leave it blanc.");
    while(size < MAX_BUFFER){
        if(size+MAX_READ_LINE < MAX_BUFFER){
            tmp = readInput(cursor, MAX_READ_LINE);
        }
        else{
            tmp = readInput(cursor, MAX_BUFFER-size);
        }

        if(tmp <= 0){
            break;
        }

        if(tmp == 1 && cursor[0] == '\n'){
            break;
        }

        size += tmp;
        cursor += tmp;
    }
    cursor[0] = '\0';

    puts("I shall deliver your letter!");
    puts("Good luck!");
}

void doubt(){
    char buff[250];
    int size;

    puts("Well, you seem good.");

    puts("Are you sure you want me to deliver the letter?");
    readInput(buff, 1);

    if(buff[0] != 'Y'){
        puts("Okay you can change this then.");
        scanf("%d", &size);
        if(size >= 250 || size < 0){
            puts("How dare you! You'll be alone forever!");
            exit(0);
        }
    }

    puts("Alright, you should be good now.");
    readInput(buff, size);

    puts("Good luck my friend!");
}

int main(){
    setup();

    loveletter();
    doubt();
    return 0;
}