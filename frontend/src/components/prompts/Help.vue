<template>
  <div class="card floating help">
    <div class="card-title">
      <h2>{{ $t("help.help") }}</h2>
    </div>

    <div class="card-content">
      <ul>
        <li><strong>F1</strong> - {{ $t("help.f1") }}</li> <!-- eslint-disable-line @intlify/vue-i18n/no-raw-text -->
        <li><strong>F2</strong> - {{ $t("help.f2") }}</li> <!-- eslint-disable-line @intlify/vue-i18n/no-raw-text -->
        <li><strong>DEL</strong> - {{ $t("help.del") }}</li> <!-- eslint-disable-line @intlify/vue-i18n/no-raw-text -->
        <li><strong>ESC</strong> - {{ $t("help.esc") }}</li> <!-- eslint-disable-line @intlify/vue-i18n/no-raw-text -->
        <li><strong>CTRL + S</strong> - {{ $t("help.ctrl.s") }}</li> <!-- eslint-disable-line @intlify/vue-i18n/no-raw-text -->
        <li><strong>CTRL + F</strong> - {{ $t("help.ctrl.f") }}</li> <!-- eslint-disable-line @intlify/vue-i18n/no-raw-text -->
        <li><strong>CTRL + Click</strong> - {{ $t("help.ctrl.click") }}</li> <!-- eslint-disable-line @intlify/vue-i18n/no-raw-text -->
        <li><strong>Click</strong> - {{ $t("help.click") }}</li> <!-- eslint-disable-line @intlify/vue-i18n/no-raw-text -->
        <li><strong>Double click</strong> - {{ $t("help.doubleClick") }}</li> <!-- eslint-disable-line @intlify/vue-i18n/no-raw-text -->
      </ul>
    </div>

    <div class="card-action">
      <button
        type="submit"
        @click="closeHovers"
        class="button button--flat"
        :aria-label="$t('buttons.ok')"
        :title="$t('buttons.ok')"
      >
        {{ $t("buttons.ok") }}
      </button>
    </div>
  </div>
</template>

<script>
import { mutations } from "@/store"; // Import the mutations

export default {
  name: "help",
  computed: {
    closeHovers() {
      return mutations.closeHovers; // Return the closeHovers mutation
    },
  },
};
</script>
