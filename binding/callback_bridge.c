#include "callback_bridge.h"

void callback_bridge(void *callback, char *output) {
  ((Callback*) callback)(output);
};
