#include <stdio.h>
#include <math.h>
#include <stdlib.h>

const unsigned int PERMUTATIONS = 1048576;

struct Data {
        int location;
        struct Data *next;
};

static struct Data table[PERMUTATIONS];

static void add_location(int hash, int position)
{
        struct Data *current = &table[hash];

        if (current -> next == NULL) {
                current -> next = malloc(sizeof(struct Data));
                current -> location = position;
                current -> next -> next = NULL;
        } else {
                struct Data *tail = current -> next;
                current -> next = malloc(sizeof(struct Data));
                current -> next -> location = position;
                current -> next -> next = tail;
        }
}

static int get_val(char ch)
{
        switch(ch) {
                case 'a':
                case 'A':
                        return 0;
                case 'c':
                case 'C':
                        return 1;
                case 'g':
                case 'G':
                        return 2;
                case 't':
                case 'T':
                        return 3;
                default:
                        return -1;
        }
}

static char char_from_val(int val)
{
        switch(val) {
                case 0:
                        return 'a';
                case 1:
                        return 'c';
                case 2:
                        return 'g';
                case 3:
                        return 't';
                default:
                        return 'n';
        }
}

static char* unhash(int key)
{
        static char result[11];
        int mutkey = key;

        for (int i = 9; i >= 0; i--) {
                result[i] = char_from_val(mutkey % 4);
                mutkey /= 4;
        }

        result[10] = '\0';
        return result;
}

int main(int argc, char *argv[])
{
        if (argc < 2)
                return -1;

        FILE *fp = fopen(argv[1], "r");
        int hashed = 0;
        char ch;
        int location = 0;

        char buff[10];
        fgets(buff, 10, fp); // remove header

        for (int i = 0; i < 10; i++) {
                do {
                        ch = fgetc(fp);
                } while (ch == '\n'); // ignore newlines

                location++;

                int ch_val = get_val(ch);

                if (ch_val == -1)
                        i--;
                else if (ch_val != 0)
                        hashed += ch_val * pow(4, 9 - i);
        }

        int valid_counter = 0;

        while(ch != EOF) {
                if (valid_counter <= 0)
                        add_location(hashed, location - 10);
                else
                        valid_counter--;

                hashed %= 262144;
                hashed *= 4;

                do {
                        location++;
                        ch = fgetc(fp);

                } while (ch == '\n'); // ignore newlines

                int ch_val = get_val(ch);

                if (ch_val == -1)
                        valid_counter = 10;
                else
                        hashed += ch_val;
        }

        for(int i = 0; i < PERMUTATIONS; i++) {
                struct Data current = table[i];
                if (current.next != NULL) {
                        printf("%s", unhash(i));

                        while(current.next != NULL) {
                                printf("\t%d", current.location);
                                current = *current.next;
                        }
                        printf("\n");
                }
        }
}
