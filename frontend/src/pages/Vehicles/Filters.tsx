import { Group, NumberInput, TextInput } from '@mantine/core';
import { useFiltersStore } from './store';

export const Filters = () => {
  const filters = useFiltersStore((state) => state.filters);
  const onChange = useFiltersStore((state) => state.setFilters);

  return (
    <Group gap={12}>
      <TextInput
        onChange={(e) => onChange({ ...filters, clientName: e.target.value.trim() })}
        label="Имя клиента"
        placeholder="Введите номер склада"
        value={filters.clientName}
      />

      <TextInput
        onChange={(e) => onChange({ ...filters, vehicleNumber: e.target.value.trim() })}
        label="Номер машины"
        placeholder="Введите номер машины"
        value={filters.vehicleNumber}
      />

      <TextInput
        onChange={(e) => onChange({ ...filters, brand: e.target.value.trim() })}
        label="Марка машины"
        placeholder="Введите марку машины"
        value={filters.brand}
      />

      <TextInput
        onChange={(e) => onChange({ ...filters, model: e.target.value.trim() })}
        label="Модель машины"
        placeholder="Введите модель машины"
        value={filters.model}
      />
    </Group>
  );
};
