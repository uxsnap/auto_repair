import { useState } from 'react';
import { Button, Group, Stack } from '@mantine/core';
import { useDisclosure } from '@mantine/hooks';
import { Application } from '@/types';
import { Filters } from './Filters';
import { AppModal } from './AppsModal';
import { AppTable } from './Table';

export function AppsPage() {
  const [opened, { open, close }] = useDisclosure(false);

  const [curApp, setCurApp] = useState<Application>();

  const handleChange = (App: Application) => {
    setCurApp(App);
    open();
  };

  return (
    <>
      <AppModal
        onSubmit={() => setCurApp(undefined)}
        close={close}
        opened={opened}
        app={curApp}
        edit={!!curApp}
      />

      <Stack mt={20} gap={12}>
        <Group align="flex-end" justify="space-between">
          <Filters />

          <Button onClick={open}>Добавить заявку</Button>
        </Group>

        <AppTable onChange={handleChange} />
      </Stack>
    </>
  );
}
