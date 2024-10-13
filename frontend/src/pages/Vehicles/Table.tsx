import { IconX } from '@tabler/icons-react';
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { Table } from '@mantine/core';
import { useDebouncedValue } from '@mantine/hooks';
import { showNotification } from '@mantine/notifications';
import { deleteVehicle } from '@/api/vehicles/deleteVehicle';
import { getVehicles } from '@/api/vehicles/getVehicles';
import { Container } from '@/components/Container';
import { Vehicle } from '@/types';
import { useFiltersStore } from './store';

type Props = {
  onChange: (vehicle: Vehicle) => void;
};

export const VehicleTable = ({}: Props) => {
  const filters = useFiltersStore((state) => state.filters);
  const [debouncedFilters] = useDebouncedValue(filters, 200);

  const queryVehicle = useQueryClient();

  const { data: vehiclesData, isFetching: isFetchingVehicles } = useQuery({
    queryKey: [getVehicles.queryKey, debouncedFilters],
    queryFn: () => getVehicles(filters),
    staleTime: 5000,
  });

  const deleteMutation = useMutation({
    mutationFn: deleteVehicle,
    onSuccess: () => {
      queryVehicle.invalidateQueries({ queryKey: [getVehicles.queryKey] });
      showNotification({
        title: 'ТС',
        message: `ТС было удалено.`,
      });
      close();
    },
  });

  const onDelete = (id: string) => {
    deleteMutation.mutate({ id });
  };

  const rows = (vehiclesData?.data ?? []).map((element) => (
    <Table.Tr key={element.id}>
      <Table.Td>{element.id}</Table.Td>
      <Table.Td>{element.vehicleNumber}</Table.Td>
      <Table.Td>{element.client.name}</Table.Td>
      <Table.Td>{element.brand}</Table.Td>
      <Table.Td>{element.model}</Table.Td>

      <Table.Td>
        <IconX color="red" style={{ cursor: 'pointer' }} onClick={() => onDelete(element.id)} />
      </Table.Td>
    </Table.Tr>
  ));

  return (
    <Container isFetching={isFetchingVehicles}>
      <Table stickyHeader withColumnBorders highlightOnHover>
        <Table.Thead>
          <Table.Tr>
            <Table.Th>ID записи</Table.Th>
            <Table.Th>Номер машины</Table.Th>
            <Table.Th>Имя клиента</Table.Th>
            <Table.Th>Марка</Table.Th>
            <Table.Th>Модель</Table.Th>
          </Table.Tr>
        </Table.Thead>

        <Table.Tbody>{rows}</Table.Tbody>
      </Table>
    </Container>
  );
};
