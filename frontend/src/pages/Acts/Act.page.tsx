import { useState } from 'react';
import { Button, Group, Stack } from '@mantine/core';
import { useDisclosure } from '@mantine/hooks';
import { Act } from '@/types';
import { ActModal } from './ActModal';
import { Filters } from './Filters';
import { ActTable } from './Table';

export function ActsPage() {
  const [opened, { open, close }] = useDisclosure(false);

  const [curAct, setCurAct] = useState<Act>();

  return (
    <>
      <ActModal
        onSubmit={() => setCurAct(undefined)}
        close={close}
        opened={opened}
        act={curAct}
        edit={!!curAct}
      />

      <Stack mt={20} gap={12}>
        <Group align="flex-end" justify="space-between">
          <Filters />

          <Button onClick={open}>Добавить акт</Button>
        </Group>

        <ActTable />
      </Stack>
    </>
  );
}
