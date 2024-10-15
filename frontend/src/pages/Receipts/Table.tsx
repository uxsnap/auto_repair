import dayjs from 'dayjs';
import { useQuery } from '@tanstack/react-query';
import { Table } from '@mantine/core';
import { useDebouncedValue } from '@mantine/hooks';
import { getReceipts } from '@/api/receipt/getReceipts';
import { Container } from '@/components/Container';
import { Receipt } from '@/types';
import { useFiltersStore } from './store';

type Props = {
  onChange: (client: Receipt) => void;
};

export const ReceiptTable = ({}: Props) => {
  const filters = useFiltersStore((state) => state.filters);
  const [debouncedFilters] = useDebouncedValue(filters, 200);

  const { data: receiptsData, isFetching: isFetchingReceipts } = useQuery({
    queryKey: [getReceipts.queryKey, debouncedFilters],
    queryFn: () => getReceipts(filters),
    staleTime: 5000,
  });

  const rows = (receiptsData?.data ?? []).map((element) => (
    <Table.Tr key={element.id}>
      <Table.Td>{element.id}</Table.Td>
      <Table.Td>{element.contract.name}</Table.Td>
      <Table.Td>{element.sum}</Table.Td>
      <Table.Td>{dayjs(element.createdAt.Time).format('DD.MM.YYYY')}</Table.Td>
    </Table.Tr>
  ));

  return (
    <Container isFetching={isFetchingReceipts}>
      <Table stickyHeader withColumnBorders highlightOnHover>
        <Table.Thead>
          <Table.Tr>
            <Table.Th>ID записи</Table.Th>
            <Table.Th>Связанный договор</Table.Th>
            <Table.Th>Сумма</Table.Th>
            <Table.Th>Дата</Table.Th>
          </Table.Tr>
        </Table.Thead>

        <Table.Tbody>{rows}</Table.Tbody>
      </Table>
    </Container>
  );
};
