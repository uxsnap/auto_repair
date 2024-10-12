import { IconEdit, IconX } from '@tabler/icons-react';
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { Table } from '@mantine/core';
import { useDebouncedValue } from '@mantine/hooks';
import { showNotification } from '@mantine/notifications';
import { deleteStorage } from '@/api/storages/deleteStorage';
import { getStorages } from '@/api/storages/getStorages';
import { Container } from '@/components/Container';
import { Storage, StorageWithData } from '@/types';
import { useFiltersStore } from './store';

type Props = {
  onChange: (storage: Storage) => void;
};

export const StorageTable = ({ onChange }: Props) => {
  const filters = useFiltersStore((state) => state.filters);
  const [debouncedFilters] = useDebouncedValue(filters, 200);

  const queryClient = useQueryClient();

  const { data: storagesData, isFetching: isFetchingStorages } = useQuery({
    queryKey: [getStorages.queryKey, debouncedFilters],
    queryFn: () => getStorages(filters),
    staleTime: 5000,
  });

  const deleteMutation = useMutation({
    mutationFn: deleteStorage,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: [getStorages.queryKey] });
      showNotification({
        title: 'Cклад',
        message: `Склад был удален.`,
      });
      close();
    },
  });

  const onDelete = (id: string) => {
    deleteMutation.mutate({ id });
  };

  const handleChange = (storageWithData: StorageWithData) => {
    onChange({
      id: storageWithData.id,
      storageNum: storageWithData.storageNum,
      detailCount: storageWithData.detailCount,
      detailId: storageWithData.detail.id,
      employeeId: storageWithData.employee.id,
    });
  };

  const rows = (storagesData?.data ?? []).map((element) => (
    <Table.Tr key={element.id}>
      <Table.Td>{element.id}</Table.Td>
      <Table.Td>{element.storageNum}</Table.Td>
      <Table.Td>{element.employee.name}</Table.Td>
      <Table.Td>{element.detail.name}</Table.Td>
      <Table.Td>{element.detailCount}</Table.Td>

      <Table.Td>
        <IconEdit style={{ cursor: 'pointer' }} onClick={() => handleChange(element)} />
        <IconX color="red" style={{ cursor: 'pointer' }} onClick={() => onDelete(element.id)} />
      </Table.Td>
    </Table.Tr>
  ));

  return (
    <Container isFetching={isFetchingStorages}>
      <Table stickyHeader withColumnBorders highlightOnHover>
        <Table.Thead>
          <Table.Tr>
            <Table.Th>ID записи</Table.Th>
            <Table.Th>Номер склада</Table.Th>
            <Table.Th>Имя отвественного сотрудника</Table.Th>
            <Table.Th>Наименование запчасти</Table.Th>
            <Table.Th>Кол-во деталей</Table.Th>
            <Table.Th>Действия</Table.Th>
          </Table.Tr>
        </Table.Thead>

        <Table.Tbody>{rows}</Table.Tbody>
      </Table>
    </Container>
  );
};
