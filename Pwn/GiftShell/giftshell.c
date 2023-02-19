#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>


void setup(){
    setvbuf(stderr, NULL, _IONBF, 0);
    setvbuf(stdout, NULL, _IONBF, 0);
    setvbuf(stdin, NULL, _IONBF, 0);
}

int main(){

	char input[100];

	setup();
	printf("Pass this test to get to the next step. You can win a discount on one of our products! %p\n", &input);
	puts("Input: ");

	read(0,input,128);

	return 0;


}
