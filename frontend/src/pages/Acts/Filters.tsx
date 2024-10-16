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
        label="Имя клиента"
        placeholder="Введите номер склада"
        value={filters.name}
      />

      <TextInput
        onChange={(e) => onChange({ ...filters, applicationName: e.target.value.trim() })}
        label="Название заявки"
        placeholder="Введите название заявки"
        value={filters.applicationName}
      />

      <TextInput
        onChange={(e) => onChange({ ...filters, serviceName: e.target.value.trim() })}
        label="Название услуги"
        placeholder="Введите название услуги"
        value={filters.serviceName}
      />

      <DateInput
        clearable
        onChange={(v) => onChange({ ...filters, minCreatedAt: v })}
        label="Минимальная дата создания"
        placeholder="Выберите минимальная дату создания"
        value={filters.minCreatedAt}
        valueFormat="DD.MM.YYYY"
      />

      <DateInput
        clearable
        onChange={(v) => onChange({ ...filters, maxCreatedAt: v })}
        label="Максимальная дата создания"
        placeholder="Выберите максимальную дату создания"
        value={filters.maxCreatedAt}
        valueFormat="DD.MM.YYYY"
      />
    </Group>
  );
};
