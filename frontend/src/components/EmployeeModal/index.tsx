import { MouseEvent, useState } from 'react';
import { IconEdit, IconX } from '@tabler/icons-react';
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { element } from 'prop-types';
import { Modal, Table, Text } from '@mantine/core';
import { showNotification } from '@mantine/notifications';
import { deleteEmployee } from '@/api/employees/deleteEmployee';
import { getEmployees } from '@/api/employees/getEmployees';
import { Employee } from '@/types';
import { AddEmployee } from '../AddEmployee';
import { Container } from '../Container';

type Props = {
  opened: boolean;
  close: () => void;
};

export const EmployeeModal = ({ opened, close }: Props) => {
  const queryClient = useQueryClient();

  const [curEmployee, setCurEmployee] = useState<Employee>();

  const { data, isFetching } = useQuery({
    queryKey: [getEmployees.queryKey],
    queryFn: getEmployees,
    staleTime: 5000,
  });

  const deleteMutation = useMutation({
    mutationFn: deleteEmployee,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: [getEmployees.queryKey] });
      showNotification({
        title: 'Сотрудники',
        message: `Сотрудник был удален`,
      });
    },
  });

  const onDelete = (e: MouseEvent<SVGSVGElement>, id: string) => {
    e.stopPropagation();
    deleteMutation.mutate({ id });
  };

  const rows = (data?.data ?? []).map((element) => (
    <Table.Tr key={element.id}>
      <Table.Td>{element.name}</Table.Td>
      <Table.Td>{element.position}</Table.Td>
      <Table.Td>{element.employeeNum}</Table.Td>

      <Table.Td>
        <IconEdit style={{ cursor: 'pointer' }} onClick={() => setCurEmployee(element)} />
        <IconX color="red" style={{ cursor: 'pointer' }} onClick={(e) => onDelete(e, element.id)} />
      </Table.Td>
    </Table.Tr>
  ));

  return (
    <>
      <Modal
        size="xl"
        opened={opened}
        onClose={close}
        title={
          <Text fz={22} fw="bold">
            Сотрудники
          </Text>
        }
        centered
      >
        <AddEmployee
          edit
          employee={curEmployee}
          opened={!!curEmployee}
          close={() => setCurEmployee(undefined)}
        />

        <Container isFetching={isFetching}>
          <Table striped highlightOnHover withTableBorder>
            <Table.Thead>
              <Table.Tr>
                <Table.Th>Имя</Table.Th>
                <Table.Th>Должность</Table.Th>
                <Table.Th>Рабочий номер</Table.Th>
                <Table.Th>Действия</Table.Th>
              </Table.Tr>
            </Table.Thead>

            <Table.Tbody>{rows}</Table.Tbody>
          </Table>
        </Container>
      </Modal>
    </>
  );
};
