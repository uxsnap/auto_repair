import { IconEdit, IconX } from '@tabler/icons-react';
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { Table } from '@mantine/core';
import { useDebouncedValue } from '@mantine/hooks';
import { showNotification } from '@mantine/notifications';
import { deleteApp } from '@/api/apps/deleteApp';
import { getApps } from '@/api/apps/getApps';
import { Container } from '@/components/Container';
import { Application, ApplicationWithData } from '@/types';
import { useFiltersStore } from './store';

type Props = {
  onChange: (client: Application) => void;
};

export const AppTable = ({ onChange }: Props) => {
  const filters = useFiltersStore((state) => state.filters);
  const [debouncedFilters] = useDebouncedValue(filters, 200);

  const queryApp = useQueryClient();

  const { data: appsData, isFetching: isFetchingApps } = useQuery({
    queryKey: [getApps.queryKey, debouncedFilters],
    queryFn: () => getApps(filters),
    staleTime: 5000,
  });

  const deleteMutation = useMutation({
    mutationFn: deleteApp,
    onSuccess: () => {
      queryApp.invalidateQueries({ queryKey: [getApps.queryKey] });
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

  const handleChange = (appWithData: ApplicationWithData) => {
    onChange({
      id: appWithData.id,
      name: appWithData.name,
      employeeId: appWithData.employee.id,
      contractId: appWithData.contract.id,
      status: appWithData.status,
      clientId: appWithData.client.id,
      createdAt: appWithData.createdAt,
    });
  };

  const rows = (appsData?.data ?? []).map((element) => (
    <Table.Tr key={element.id}>
      <Table.Td>{element.id}</Table.Td>
      <Table.Td>{element.name}</Table.Td>
      <Table.Td>{element.employee.name}</Table.Td>
      <Table.Td>{element.client.name}</Table.Td>
      <Table.Td>{element.contract.name}</Table.Td>
      <Table.Td>{element.createdAt}</Table.Td>
      <Table.Td>{element.status}</Table.Td>

      <Table.Td>
        <IconEdit style={{ cursor: 'pointer' }} onClick={() => handleChange(element)} />
        <IconX color="red" style={{ cursor: 'pointer' }} onClick={() => onDelete(element.id)} />
      </Table.Td>
    </Table.Tr>
  ));

  return (
    <Container isFetching={isFetchingApps}>
      <Table stickyHeader withColumnBorders highlightOnHover>
        <Table.Thead>
          <Table.Tr>
            <Table.Th>ID записи</Table.Th>
            <Table.Th>Имя</Table.Th>
            <Table.Th>Имя отвественного сотрудника</Table.Th>
            <Table.Th>Имя клиента</Table.Th>
            <Table.Th>Контракт</Table.Th>
            <Table.Th>Дата создания</Table.Th>
            <Table.Th>Статус</Table.Th>
          </Table.Tr>
        </Table.Thead>

        <Table.Tbody>{rows}</Table.Tbody>
      </Table>
    </Container>
  );
};
