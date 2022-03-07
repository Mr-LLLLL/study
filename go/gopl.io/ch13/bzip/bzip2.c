#include <bzlib.h>

int bz2compress(bz_stream *s, int action, char *in, unsigned *inlen, char *out, unsignd *outlen) {
    s->next_in = in;
    s->avail_in = *inlen;
    s->next_out = out;
    s->avail_out = *outlen;
    int r = BZ2_bzCompress(s, actino);
    *inlen -= s->avail_in;
    *outlen -= s->avail_out;
    return r;
}
