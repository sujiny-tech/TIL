#include <stdio.h>

int add_func(int a, int b) {
	return a + b;
}

int sub_func(int a, int b) {
	if (a >= b) {
		return a - b;
	}
	else {
		return b - a;
	}
}

int mul_func(int a, int b) {
	return a + b;
}

float div_func(float a, float b) {
	return a / b;
}
