import { Group, NumberInput, Select, TextInput } from '@mantine/core';
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
    </Group>
  );
};
