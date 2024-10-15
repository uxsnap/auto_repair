import { useState } from 'react';
import { Button, Group, Stack } from '@mantine/core';
import { useDisclosure } from '@mantine/hooks';
import { Receipt } from '@/types';
import { ReceiptModal } from './ReceiptModal';
import { Filters } from './Filters';
import { ReceiptTable } from './Table';

export function ReceiptsPage() {
  const [opened, { open, close }] = useDisclosure(false);

  const [curReceipt, setCurReceipt] = useState<Receipt>();

  const handleChange = (Receipt: Receipt) => {
    setCurReceipt(Receipt);
    open();
  };

  return (
    <>
      <ReceiptModal
        onSubmit={() => setCurReceipt(undefined)}
        close={close}
        opened={opened}
        receipt={curReceipt}
        edit={!!curReceipt}
      />

      <Stack mt={20} gap={12}>
        <Group align="flex-end" justify="space-between">
          <Filters />

          <Button onClick={open}>Добавить чек</Button>
        </Group>

        <ReceiptTable onChange={handleChange} />
      </Stack>
    </>
  );
}
