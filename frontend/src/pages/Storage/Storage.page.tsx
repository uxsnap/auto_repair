import { useEffect, useState } from 'react';
import { Button, Group, Stack } from '@mantine/core';
import { useDisclosure } from '@mantine/hooks';
import { Storage } from '@/types';
import { Filters } from './Filters';
import { StorageModal } from './StorageModal';
import { StorageTable } from './Table';

export function StoragePage() {
  const [opened, { open, close }] = useDisclosure(false);

  const [curStorage, setCurStorage] = useState<Storage>();

  const handleChange = (storage: Storage) => {
    setCurStorage(storage);
    open();
  };

  return (
    <>
      <StorageModal
        onSubmit={() => setCurStorage(undefined)}
        close={close}
        opened={opened}
        storage={curStorage}
        edit={!!curStorage}
      />

      <Stack mt={20} gap={12}>
        <Group align="flex-end" justify="space-between">
          <Filters />

          <Button onClick={open}>Добавить склад</Button>
        </Group>

        <StorageTable onChange={handleChange} />
      </Stack>
    </>
  );
}
