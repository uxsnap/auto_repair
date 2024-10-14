import { Group, TextInput } from '@mantine/core';
import { DateInput } from '@mantine/dates';
import { useFiltersStore } from './store';

export const Filters = () => {
  const filters = useFiltersStore((state) => state.filters);
  const onChange = useFiltersStore((state) => state.setFilters);

  return (
    <Group gap={12}>
      <TextInput
        onChange={(e) => onChange({ ...filters, name: e.target.value.trim() })}
        label="Имя"
        placeholder="Введите имя заявки"
        value={filters.name}
      />

      <TextInput
        onChange={(e) => onChange({ ...filters, clientName: e.target.value.trim() })}
        label="Имя клиента"
        placeholder="Введите имя клиента"
        value={filters.clientName}
      />

      <TextInput
        onChange={(e) => onChange({ ...filters, employeeName: e.target.value.trim() })}
        label="Имя сотрудника"
        placeholder="Введите имя отвественного сотрудника"
        value={filters.employeeName}
      />

      <TextInput
        onChange={(e) => onChange({ ...filters, contractName: e.target.value.trim() })}
        label="Название контракта"
        placeholder="Введите название контракта"
        value={filters.contractName}
      />

      <TextInput
        onChange={(e) => onChange({ ...filters, status: e.target.value.trim() })}
        label="Статус"
        placeholder="Введите статус"
        value={filters.status}
      />

      <DateInput
        onChange={(v) => onChange({ ...filters, createdAt: v ?? undefined })}
        label="Дата создания"
        placeholder="Выберите дату создания"
        value={filters.createdAt}
      />
    </Group>
  );
};
