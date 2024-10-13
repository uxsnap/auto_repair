import { Group, NumberInput, TextInput } from '@mantine/core';
import { useFiltersStore } from './store';

export const Filters = () => {
  const filters = useFiltersStore((state) => state.filters);
  const onChange = useFiltersStore((state) => state.setFilters);

  return (
    <Group gap={12}>
      <TextInput
        onChange={(e) => onChange({ ...filters, name: e.target.value.trim() })}
        label="Имя клиента"
        placeholder="Введите номер склада"
        value={filters.name}
      />

      <TextInput
        onChange={(e) => onChange({ ...filters, employeeName: e.target.value.trim() })}
        label="Имя сотрудника"
        placeholder="Введите имя отвественного сотрудника"
        value={filters.employeeName}
      />

      <NumberInput
        onChange={(v) => onChange({ ...filters, phone: v + '' })}
        label="Телефон"
        placeholder="Введите телефон клиента"
        value={filters.phone}
      />

      <NumberInput
        onChange={(v) => onChange({ ...filters, passport: v + '' })}
        maxLength={10}
        label="Паспорт"
        placeholder="Введите паспорт клиента"
        value={filters.passport}
      />
    </Group>
  );
};
