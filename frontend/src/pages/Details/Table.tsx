import { IconX } from '@tabler/icons-react';
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { Table } from '@mantine/core';
import { useDebouncedValue } from '@mantine/hooks';
import { showNotification } from '@mantine/notifications';
import { deleteDetail } from '@/api/details/deleteDetail';
import { getDetails } from '@/api/details/getDetails';
import { useFiltersStore } from './store';
import { Container } from '@/components/Container';

export const DetailsTable = () => {
  const filters = useFiltersStore((state) => state.filters);
  const [debouncedFilters] = useDebouncedValue(filters, 200);

  const queryClient = useQueryClient();

  const { data, isFetching } = useQuery({
    queryKey: [getDetails.queryKey, debouncedFilters],
    queryFn: () => getDetails(filters),
    staleTime: 5000,
  });

  const deleteMutation = useMutation({
    mutationFn: deleteDetail,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: [getDetails.queryKey] });
      showNotification({
        title: 'Детали',
        message: `Деталь была удалена.`,
      });
      close();
    },
  });

  const onDelete = (id: string) => {
    deleteMutation.mutate({ id });
  };

  const rows = (data?.data ?? []).map((element) => (
    <Table.Tr key={element.id}>
      <Table.Td>{element.id}</Table.Td>
      <Table.Td>{element.name}</Table.Td>
      <Table.Td>{element.price}</Table.Td>
      <Table.Td>{element.type}</Table.Td>

      <Table.Td>
        <IconX color="red" style={{ cursor: 'pointer' }} onClick={() => onDelete(element.id)} />
      </Table.Td>
    </Table.Tr>
  ));

  return (
    <Container isFetching={isFetching}>
      <Table stickyHeader withColumnBorders highlightOnHover>
        <Table.Thead>
          <Table.Tr>
            <Table.Th>ID</Table.Th>
            <Table.Th>Имя</Table.Th>
            <Table.Th>Цена</Table.Th>
            <Table.Th>Тип</Table.Th>
            <Table.Th>Действия</Table.Th>
          </Table.Tr>
        </Table.Thead>

        <Table.Tbody>{rows}</Table.Tbody>
      </Table>
    </Container>
  );
};
