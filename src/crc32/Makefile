CC=gcc
CFLAGS=-Wall -Wextra -Werror -std=c89 -pedantic-errors -Og
PROGNAME=crc32
LIBS=-lz
OBJS=$(PROGNAME).o

.PHONY: all clean clean-all

all: $(PROGNAME)

$(PROGNAME): $(OBJS)
	$(CC) -o '$@' $^ $(LIBS)

$(PROGNAME).o: $(PROGNAME).c
	$(CC) -x c $(CFLAGS) -c -o '$@' '$<'

clean:
	-@rm $(OBJS)
clean-all: clean
	-@rm ./$(PROGNAME)
