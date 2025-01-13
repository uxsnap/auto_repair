import dayjs from 'dayjs';
import { IconEdit, IconX } from '@tabler/icons-react';
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { Table } from '@mantine/core';
import { useDebouncedValue } from '@mantine/hooks';
import { showNotification } from '@mantine/notifications';
import { deleteContract } from '@/api/contract/deleteContract';
import { getContracts } from '@/api/contract/getContacts';
import { Container } from '@/components/Container';
import { Contract } from '@/types';
import { useFiltersStore } from './store';

type Props = {
  onChange: (client: Contract) => void;
};

export const ContractTable = ({ onChange }: Props) => {
  const filters = useFiltersStore((state) => state.filters);
  const [debouncedFilters] = useDebouncedValue(filters, 200);

  const queryContract = useQueryClient();

  const { data: contractsData, isFetching: isFetchingContracts } = useQuery({
    queryKey: [getContracts.queryKey, debouncedFilters],
    queryFn: () => getContracts(filters),
    staleTime: 5000,
  });

  const deleteMutation = useMutation({
    mutationFn: deleteContract,
    onSuccess: () => {
      queryContract.invalidateQueries({ queryKey: [getContracts.queryKey] });
      showNotification({
        title: 'Договор',
        message: `Договор был удален.`,
      });
      close();
    },
  });

  const onDelete = (id: string) => {
    deleteMutation.mutate({ id });
  };

  const rows = (contractsData?.data ?? []).map((element) => (
    <Table.Tr key={element.id}>
      <Table.Td>{element.id}</Table.Td>
      <Table.Td>{element.name}</Table.Td>
      <Table.Td>{element.sum}</Table.Td>
      <Table.Td>{dayjs(element.createdAt.Time).format('DD.MM.YYYY')}</Table.Td>
      <Table.Td>
        {element.signedAt.Status === 1 ? '' : dayjs(element.signedAt.Time).format('DD.MM.YYYY')}
      </Table.Td>
      <Table.Td>{element.status}</Table.Td>

      <Table.Td>
        <IconEdit style={{ cursor: 'pointer' }} onClick={() => onChange(element)} />
        <IconX color="red" style={{ cursor: 'pointer' }} onClick={() => onDelete(element.id)} />
      </Table.Td>
    </Table.Tr>
  ));

  return (
    <Container isFetching={isFetchingContracts}>
      <Table stickyHeader withColumnBorders highlightOnHover>
        <Table.Thead>
          <Table.Tr>
            <Table.Th>ID записи</Table.Th>
            <Table.Th>Имя</Table.Th>
            <Table.Th>Сумма по договору</Table.Th>
            <Table.Th>Дата создания</Table.Th>
            <Table.Th>Дата подписания</Table.Th>
            <Table.Th>Статус</Table.Th>
          </Table.Tr>
        </Table.Thead>

        <Table.Tbody>{rows}</Table.Tbody>
      </Table>
    </Container>
  );
};
