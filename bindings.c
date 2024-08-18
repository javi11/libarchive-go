#include <archive.h>
#include <archive_entry.h>
#include "cgo_export.h"

int64_t seek_cb_binding(struct archive *a, char *client_data, int64_t request, int whence)
{
    return seekCb(a, client_data, request, whence);
}

int64_t close_cb_binding(struct archive *a, char *client_data)
{
    return closeCb(a, (void *)client_data);
}

ssize_t read_cb_binding(struct archive *a, char *client_data, const void **block)
{
    return readCb(a, (void *)client_data, block);
}

int64_t skip_cb_binding(struct archive *a, char *client_data, int64_t request)
{
    return skipCb(a, (void *)client_data, request);
}

int64_t switch_cb_binding(struct archive *a, char *client_data1, char *client_data2)
{
    return switchCb(a, (void *)client_data1, (void *)client_data2);
}

int64_t open_cb_binding(struct archive *a, char *client_data)
{
    return openCb(a, (void *)client_data);
}

int64_t read_append_cb_data_binding(struct archive *a, char *client_data)
{
    return archive_read_append_callback_data(a, (void *)client_data);
}