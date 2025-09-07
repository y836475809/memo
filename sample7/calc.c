#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "calc.h"


float data_sum(person_t* people, size_t size) {
	int sum = 0;
	for (int i = 0; i < size; i++) {
        int s = people[i].size;
        for (int j = 0; j < s; j++) {
		    sum = sum + people[i].data[j].a;
        }
	}
	return (float)sum;
}

person_t* allocStruct(int num) {
    return malloc(num * sizeof(person_t));
}

void convertStructData(person_t* strctArray, int offset, person_t src) {
    memcpy(strctArray + offset, &src, sizeof(person_t));
}
