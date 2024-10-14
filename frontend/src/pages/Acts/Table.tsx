import { IconEdit, IconX } from '@tabler/icons-react';
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { Table } from '@mantine/core';
import { useDebouncedValue } from '@mantine/hooks';
import { showNotification } from '@mantine/notifications';
import { deleteAct } from '@/api/acts/deleteAct';
import { getActs } from '@/api/acts/getActs';
import { Container } from '@/components/Container';
import { Act, ActWithData } from '@/types';
import { useFiltersStore } from './store';

type Props = {
  onChange: (client: Act) => void;
};

export const ActTable = ({ onChange }: Props) => {
  const filters = useFiltersStore((state) => state.filters);
  const [debouncedFilters] = useDebouncedValue(filters, 200);

  const queryClient = useQueryClient();

  const { data: actsData, isFetching: isFetchingActs } = useQuery({
    queryKey: [getActs.queryKey, debouncedFilters],
    queryFn: () => getActs(filters),
    staleTime: 5000,
  });

  const deleteMutation = useMutation({
    mutationFn: deleteAct,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: [getActs.queryKey] });
      showNotification({
        title: 'Клиент',
        message: `Клиент был удален.`,
      });
      close();
    },
  });

  const onDelete = (id: string) => {
    deleteMutation.mutate({ id });
  };

  const rows = (actsData?.data ?? []).map((element) => (
    <Table.Tr key={element.id}>
      <Table.Td>{element.id}</Table.Td>
      <Table.Td>{element.name}</Table.Td>
      <Table.Td>{element.application.name}</Table.Td>
      <Table.Td>{element.service.name}</Table.Td>
      <Table.Td>{element.createdAt}</Table.Td>

      <Table.Td>
        <IconX color="red" style={{ cursor: 'pointer' }} onClick={() => onDelete(element.id)} />
      </Table.Td>
    </Table.Tr>
  ));

  return (
    <Container isFetching={isFetchingActs}>
      <Table stickyHeader withColumnBorders highlightOnHover>
        <Table.Thead>
          <Table.Tr>
            <Table.Th>ID записи</Table.Th>
            <Table.Th>Имя</Table.Th>
            <Table.Th>Заявка</Table.Th>
            <Table.Th>Услуга</Table.Th>
            <Table.Th>Дата создания</Table.Th>
          </Table.Tr>
        </Table.Thead>

        <Table.Tbody>{rows}</Table.Tbody>
      </Table>
    </Container>
  );
};
