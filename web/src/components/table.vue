<template>
  <n-data-table
    size="small"
    :columns="columns"
    :data="data"
    :bordered="false"
  />
</template>
<script setup lang="ts">
import { NDataTable } from "naive-ui";
import type { DataTableColumns } from "naive-ui";
import { useData } from "./data";
import { ref } from "vue";

type Panel = {
  status: number;
  method: string;
  protocol: string;
  host: string;
  path: string;
  type: string;
  originBody?: string;
  newBody?: string;
};

const createColumns = (): DataTableColumns<Panel> => {
  return [
    {
      title: "#",
      key: "key",
      render: (_, index) => {
        return `${index + 1}`;
      },
    },
    {
      title: "Status",
      key: "status",
    },
    {
      title: "Method",
      key: "method",
    },
    {
      title: "Protocol",
      key: "protocol",
    },
    {
      title: "Host",
      key: "host",
    },
    {
      title: "Path",
      key: "path",
    },
    {
      title: "Type",
      key: "type",
    },
  ];
};

const data = ref<Panel[]>([]);

const columns = createColumns();

setInterval(async () => {
  const getData = await useData();
  data.value.push(...getData);
}, 1000);
</script>
