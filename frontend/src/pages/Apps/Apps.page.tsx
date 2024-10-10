import { useQuery } from '@tanstack/react-query';
import { Button, Group, Loader, Stack, Table } from '@mantine/core';
import { getApps } from '@/api/apps/getApps';
import { Filters } from '@/components/Filters';
import styles from './Apps.module.css';

export function AppsPage() {
  const { data, isFetching } = useQuery({ queryKey: [getApps.queryKey], queryFn: getApps });

  const rows = (data?.data ?? []).map((element) => (
    <Table.Tr key={element.name}>
      <Table.Td>{element.id}</Table.Td>
      <Table.Td>{element.name}</Table.Td>
      <Table.Td>{element.clientId}</Table.Td>
      <Table.Td>{element.createdAt}</Table.Td>
      <Table.Td>{element.employeeId}</Table.Td>
      <Table.Td>{element.status}</Table.Td>
    </Table.Tr>
  ));

  return (
    <div className={styles.root}>
      {isFetching && <Loader className={styles.loader} />}

      {!isFetching && (
        <Stack gap={20}>
          <Group justify='space-between'>
            <Filters />

            <Button>Создать заявку</Button>
          </Group>

          <Table stickyHeader withColumnBorders highlightOnHover>
            <Table.Thead>
              <Table.Tr>
                <Table.Th>ID</Table.Th>
                <Table.Th>Имя</Table.Th>
                <Table.Th>Имя клиента</Table.Th>
                <Table.Th>Дата создания</Table.Th>
                <Table.Th>Закрепленный сотрудник</Table.Th>
                <Table.Th>Статус</Table.Th>
              </Table.Tr>
            </Table.Thead>

            <Table.Tbody>{rows}</Table.Tbody>
          </Table>
        </Stack>
      )}
    </div>
  );
}
