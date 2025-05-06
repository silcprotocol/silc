local default = import 'default.jsonnet';

default {
  'silc_2024-1'+: {
    config+: {
      consensus+: {
        timeout_commit: '5s',
      },
    },
  },
}
