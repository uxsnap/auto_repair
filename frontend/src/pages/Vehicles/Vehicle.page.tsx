import { useState } from 'react';
import { Button, Group, Stack } from '@mantine/core';
import { useDisclosure } from '@mantine/hooks';
import { Vehicle } from '@/types';
import { Filters } from './Filters';
import { VehicleTable } from './Table';
import { VehicleModal } from './VehicleModal';

export function VehiclesPage() {
  const [opened, { open, close }] = useDisclosure(false);

  const [curVehicle, setCurVehicle] = useState<Vehicle>();

  const handleChange = (Vehicle: Vehicle) => {
    setCurVehicle(Vehicle);
    open();
  };

  return (
    <>
      <VehicleModal
        onSubmit={() => setCurVehicle(undefined)}
        close={close}
        opened={opened}
        vehicle={curVehicle}
        edit={!!curVehicle}
      />

      <Stack mt={20} gap={12}>
        <Group align="flex-end" justify="space-between">
          <Filters />

          <Button onClick={open}>Добавить ТС</Button>
        </Group>

        <VehicleTable onChange={handleChange} />
      </Stack>
    </>
  );
}
