#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "calc.h"


float data_sum(person_t* people, size_t size) {
	int sum = 0;
	for (int i = 0; i < size; i++) {
        int s = people[i].size;
		sum = sum + s;
	}
	return (float)sum;
}

