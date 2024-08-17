#include <archive.h>
#include <archive_entry.h>
#include <stdlib.h>

int64_t seekCb(struct archive *a, char *client_data, int64_t request, int whence);
int64_t openCb(struct archive *a, char *client_data);
int64_t closeCb(struct archive *a, void *client_data);
ssize_t readCb(struct archive *a, void *client_data, const void **block);
int64_t skipCb(struct archive *a, void *client_data, int64_t request);
int64_t switchCb(struct archive *a, void *client_data1, void *client_data2);
