import { IconEdit, IconX } from '@tabler/icons-react';
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { Table } from '@mantine/core';
import { useDebouncedValue } from '@mantine/hooks';
import { showNotification } from '@mantine/notifications';
import { deleteClient } from '@/api/clients/deleteClient';
import { getClients } from '@/api/clients/getClients';
import { Container } from '@/components/Container';
import { Client, ClientWithData } from '@/types';
import { useFiltersStore } from './store';

type Props = {
  onChange: (client: Client) => void;
};

export const ClientTable = ({ onChange }: Props) => {
  const filters = useFiltersStore((state) => state.filters);
  const [debouncedFilters] = useDebouncedValue(filters, 200);

  const queryClient = useQueryClient();

  const { data: clientsData, isFetching: isFetchingClients } = useQuery({
    queryKey: [getClients.queryKey, debouncedFilters],
    queryFn: () => getClients(filters),
    staleTime: 5000,
  });

  const deleteMutation = useMutation({
    mutationFn: deleteClient,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: [getClients.queryKey] });
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

  const handleChange = (clientWithData: ClientWithData) => {
    onChange({
      id: clientWithData.id,
      name: clientWithData.name,
      employeeId: clientWithData.employee.id,
      phone: clientWithData.phone,
      passport: clientWithData.passport,
      hasDocuments: clientWithData.hasDocuments,
    });
  };

  const rows = (clientsData?.data ?? []).map((element) => (
    <Table.Tr key={element.id}>
      <Table.Td>{element.id}</Table.Td>
      <Table.Td>{element.name}</Table.Td>
      <Table.Td>{element.employee.name}</Table.Td>
      <Table.Td>{element.phone}</Table.Td>
      <Table.Td>{element.passport}</Table.Td>
      <Table.Td>{element.hasDocuments ? 'Да' : 'Нет'}</Table.Td>

      <Table.Td>
        <IconEdit style={{ cursor: 'pointer' }} onClick={() => handleChange(element)} />
        <IconX color="red" style={{ cursor: 'pointer' }} onClick={() => onDelete(element.id)} />
      </Table.Td>
    </Table.Tr>
  ));

  return (
    <Container isFetching={isFetchingClients}>
      <Table stickyHeader withColumnBorders highlightOnHover>
        <Table.Thead>
          <Table.Tr>
            <Table.Th>ID записи</Table.Th>
            <Table.Th>Имя</Table.Th>
            <Table.Th>Имя ответственного сотрудника</Table.Th>
            <Table.Th>Телефон</Table.Th>
            <Table.Th>Паспорт</Table.Th>
            <Table.Th>Наличие документов</Table.Th>
          </Table.Tr>
        </Table.Thead>

        <Table.Tbody>{rows}</Table.Tbody>
      </Table>
    </Container>
  );
};
