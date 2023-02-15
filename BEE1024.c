#include<stdio.h>
#include<string.h>
#include <ctype.h>

#define MAX_INDEX 1000

int main (void)
{
  int n;
  scanf("%d", &n);
  char output[n][MAX_INDEX];

  for(int i = 0; i < n; i++){

    char word[MAX_INDEX];
    getchar();
    scanf("%[^\n]", word);
    int wordlen = strlen(word);

    char encrypted[wordlen+1];

    for (int z = 0, j = wordlen-1; z < wordlen; z++, j--)
    {
      if (isalpha(word[j]))
      {
        word[j] = word[j] + 3;
      }

      if (j < (wordlen+1) / 2)
      {
        word[j] = word[j] - 1;
      }
      encrypted[z] = word[j];
    }
    encrypted[wordlen] = '\0';

    strcpy(output[i], encrypted);
  }
  for(int i = 0; i < n; i++){
    printf("%s\n", output[i]);
  }

  return 0;
}
