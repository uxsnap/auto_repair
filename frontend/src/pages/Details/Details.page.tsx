import { IconX } from '@tabler/icons-react';
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { Button, Group, Stack, Table } from '@mantine/core';
import { useDebouncedValue, useDisclosure, useSetState } from '@mantine/hooks';
import { showNotification } from '@mantine/notifications';
import { deleteDetail } from '@/api/details/deleteDetail';
import { getDetails } from '@/api/details/getDetails';
import { Container } from '@/components/Container';
import { Detail, FilterValues } from '@/types';
import { DetailsModal } from './DetailsModal';
import { Filters } from './Filters';
import { DetailsTable } from './Table';

export function DetailsPage() {
  const [opened, { open, close }] = useDisclosure(false);

  return (
    <>
      <DetailsModal close={close} opened={opened} />

      <Stack mt={20} gap={12}>
        <Group align="flex-end" justify="space-between">
          <Filters />

          <Button onClick={open}>Добавить деталь</Button>
        </Group>

        <DetailsTable />
      </Stack>
    </>
  );
}
