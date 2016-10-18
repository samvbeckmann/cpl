#include <stdio.h>

int main()
{
        char ch;
        while ((ch = getchar()) != EOF) {
                if (ch == '/') {
                        char next = getchar();
                        if (next == '/') {
                                while ((ch = getchar()) != '\n');
                                putchar(ch);
                        } else if (next == '*') {
                                int end = 0;
                                while (!end) {
                                        if (ch == '*') {
                                                ch = getchar();
                                                if (ch == '/')
                                                        end = 1;
                                        } else {
                                                ch = getchar();
                                        }
                                }
                        }
                } else {
                        putchar(ch);
                }
        }
}
