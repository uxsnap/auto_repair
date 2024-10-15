import { Group, NumberInput, Select, TextInput } from '@mantine/core';
import { DateInput } from '@mantine/dates';
import { useFiltersStore } from './store';

export const Filters = () => {
  const filters = useFiltersStore((state) => state.filters);
  const onChange = useFiltersStore((state) => state.setFilters);

  return (
    <Group gap={12}>
      <TextInput
        onChange={(e) => onChange({ ...filters, name: e.target.value })}
        label="Имя клиента"
        placeholder="Введите номер склада"
        value={filters.name}
      />

      <NumberInput
        label="Минимальная сумма"
        placeholder="Введите минимальную сумму"
        stepHoldDelay={500}
        stepHoldInterval={100}
        allowNegative={false}
        value={filters.minSum}
        onValueChange={(v) => onChange({ ...filters, minSum: v.floatValue })}
      />

      <NumberInput
        label="Максимальная сумма"
        placeholder="Введите максимальную сумму"
        stepHoldDelay={500}
        stepHoldInterval={100}
        allowNegative={false}
        value={filters.maxSum}
        onValueChange={(v) => onChange({ ...filters, maxSum: v.floatValue })}
      />

      <Select
        onChange={(v) => onChange({ ...filters, status: v + '' })}
        label="Статус"
        placeholder="Выберите статус"
        data={['Новый', 'Действующий', 'Закрыт']}
        value={filters.status}
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
