import { useEffect, useState } from 'react';
import { Button, Group, Stack } from '@mantine/core';
import { useDisclosure } from '@mantine/hooks';
import { Client } from '@/types';
import { Filters } from './Filters';
import { ClientModal } from './ClientModal';
import { ClientTable } from './Table';

export function ClientsPage() {
  const [opened, { open, close }] = useDisclosure(false);

  const [curClient, setCurClient] = useState<Client>();

  const handleChange = (Client: Client) => {
    setCurClient(Client);
    open();
  };

  return (
    <>
      <ClientModal
        onSubmit={() => setCurClient(undefined)}
        close={close}
        opened={opened}
        client={curClient}
        edit={!!curClient}
      />

      <Stack mt={20} gap={12}>
        <Group align="flex-end" justify="space-between">
          <Filters />

          <Button onClick={open}>Добавить клиента</Button>
        </Group>

        <ClientTable onChange={handleChange} />
      </Stack>
    </>
  );
}
