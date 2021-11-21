#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include <zlib.h>

unsigned long calculate_crc32(const char *filename);

static unsigned long initial_value;

int main(int argc, char *argv[])
{

	if (argc == 1) {
		exit(EXIT_SUCCESS);
	}

	initial_value = crc32(0, NULL, 0);

	for (++argv; *argv; ++argv)
		if (argc == 2) {
			printf("%08lx\n", calculate_crc32(*argv));
		} else {
			printf("%08lx\t%s\n", calculate_crc32(*argv), *argv);
		}

	exit(EXIT_SUCCESS);
}

unsigned long calculate_crc32(const char *filename)
{
	FILE *f;
	size_t r;
	static unsigned char buf[BUFSIZ];
	unsigned long c;

	f = fopen(filename, "rb");
	if (!f) {
		fprintf(stderr, "crc32: %s: FAILED open or read\n", filename);
		exit(EXIT_FAILURE);
	}

	c = initial_value;
	while ((r = fread(buf, 1, sizeof buf, f)) > 0u) {
		c = crc32(c, buf, r);
	}

	if (ferror(f)) {
		fprintf(stderr, "crc32: %s: Is a directory\n", filename);
		if (fclose(f) == EOF)
			perror("fclose");
		exit(EXIT_FAILURE);
	}

	if (fclose(f) == EOF) {
		perror("fclose");
		exit(EXIT_FAILURE);
	}

	return c;
}
