#include <stddef.h>


typedef struct person {
	int id;
    int size;
} person_t;

extern float data_sum(person_t* people, size_t size);

