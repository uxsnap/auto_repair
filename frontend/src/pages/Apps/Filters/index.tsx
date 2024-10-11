import { Group, TextInput } from '@mantine/core';

export const Filters = () => {
  return (
    <Group gap={12}>
      <TextInput
        label="Имя"
        placeholder="Введите имя"
      />
    </Group>
  );
};
