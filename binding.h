#include <archive.h>
#include <archive_entry.h>

int64_t seek_cb_binding(struct archive *a, char *client_data, int64_t request, int whence);
int64_t open_cb_binding(struct archive *a, char *client_data);
int64_t close_cb_binding(struct archive *a, char *client_data);
ssize_t read_cb_binding(struct archive *a, char *client_data, const void **block);
int64_t skip_cb_binding(struct archive *a, char *client_data, int64_t request);
int64_t switch_cb_binding(struct archive *a, char *client_data1, char *client_data2);
int64_t read_append_cb_data_binding(struct archive *a, char *client_data);