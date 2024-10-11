import { Group, NumberInput, Select, TextInput } from '@mantine/core';
import { useFiltersStore } from './store';

export const Filters = () => {
  const filters = useFiltersStore((state) => state.filters);
  const onChange = useFiltersStore((state) => state.setFilters);

  return (
    <Group gap={12}>
      <TextInput
        onChange={(e) => onChange({ ...filters, name: e.target.value.trim() })}
        label="Имя"
        placeholder="Введите имя"
        value={filters.name}
      />

      <NumberInput
        label="Минимальная цена"
        placeholder="Введите минимальную цену"
        stepHoldDelay={500}
        stepHoldInterval={100}
        allowNegative={false}
        value={filters.minPrice}
        onValueChange={(v) => onChange({ ...filters, minPrice: v.floatValue })}
      />

      <NumberInput
        label="Максимальная цена"
        placeholder="Введите максимальную цену"
        stepHoldDelay={500}
        stepHoldInterval={100}
        allowNegative={false}
        value={filters.maxPrice}
        onValueChange={(v) => onChange({ ...filters, maxPrice: v.floatValue })}
      />

      <Select
        label="Тип"
        placeholder="Выберите тип"
        data={[
          'Деталь соединительная',
          'Деталь вращательного движения',
          'Деталь обслуживающая передачи',
        ]}
        onChange={(type) => onChange({ ...filters, type: type ?? undefined })}
      />
    </Group>
  );
};
