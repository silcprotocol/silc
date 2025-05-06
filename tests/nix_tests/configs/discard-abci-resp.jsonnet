local config = import 'default.jsonnet';

config {
  'silc_2024-1'+: {
    config+: {
      storage: {
        discard_abci_responses: true,
      },
    },
  },
}
