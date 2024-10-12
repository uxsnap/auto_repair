import { Group, TextInput } from '@mantine/core';
import { useFiltersStore } from './store';

export const Filters = () => {
  const filters = useFiltersStore((state) => state.filters);
  const onChange = useFiltersStore((state) => state.setFilters);

  return (
    <Group gap={12}>
      <TextInput
        onChange={(e) => onChange({ ...filters, storageNum: e.target.value.trim() })}
        label="Номер склада"
        placeholder="Введите номер склада"
        value={filters.storageNum}
      />

      <TextInput
        onChange={(e) => onChange({ ...filters, employeeName: e.target.value.trim() })}
        label="Имя сотрудника"
        placeholder="Введите имя отвественного сотрудника"
        value={filters.employeeName}
      />

      <TextInput
        onChange={(e) => onChange({ ...filters, detailName: e.target.value.trim() })}
        label="Имя детали"
        placeholder="Введите имя детали"
        value={filters.detailName}
      />
    </Group>
  );
};
