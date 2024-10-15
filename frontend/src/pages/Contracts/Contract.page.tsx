import { useState } from 'react';
import { Button, Group, Stack } from '@mantine/core';
import { useDisclosure } from '@mantine/hooks';
import { Contract } from '@/types';
import { ContractModal } from './ContractModal';
import { Filters } from './Filters';
import { ContractTable } from './Table';

export function ContractsPage() {
  const [opened, { open, close }] = useDisclosure(false);

  const [curContract, setCurContract] = useState<Contract>();

  const handleChange = (Contract: Contract) => {
    setCurContract(Contract);
    open();
  };

  return (
    <>
      <ContractModal
        onSubmit={() => setCurContract(undefined)}
        close={close}
        opened={opened}
        contract={curContract}
        edit={!!curContract}
      />

      <Stack mt={20} gap={12}>
        <Group align="flex-end" justify="space-between">
          <Filters />

          <Button onClick={open}>Добавить договор</Button>
        </Group>

        <ContractTable onChange={handleChange} />
      </Stack>
    </>
  );
}
