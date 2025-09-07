#include <stddef.h>

typedef struct data {
	int a;
    int b;
} data_t;


typedef struct person {
	int id;
    int size;
    data_t* data;
} person_t;

extern float data_sum(person_t* people, size_t size);
extern person_t* allocStruct(int num);
extern void convertStructData(person_t* strctArray, int offset, person_t src);
