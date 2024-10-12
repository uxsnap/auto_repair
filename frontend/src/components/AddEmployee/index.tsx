import { useEffect } from 'react';
import { useMutation, useQueryClient } from '@tanstack/react-query';
import { AxiosError } from 'axios';
import { Button, Group, Modal, Select, Stack, Text, TextInput } from '@mantine/core';
import { useForm } from '@mantine/form';
import { showNotification } from '@mantine/notifications';
import { addEmployee } from '@/api/employees/addEmployee';
import { editEmployee } from '@/api/employees/editEmployee';
import { getEmployees } from '@/api/employees/getEmployees';
import { Employee } from '@/types';

type Props = {
  opened: boolean;
  close: () => void;
  edit?: boolean;
  employee?: Employee;
};

export const AddEmployee = ({ edit, opened, close, employee }: Props) => {
  const queryClient = useQueryClient();

  const form = useForm({
    mode: 'uncontrolled',
    initialValues: {
      name: employee?.name ?? '',
      position: employee?.position ?? '',
      employeeNum: employee?.employeeNum ?? '',
    },
  });

  useEffect(() => {
    if (!employee) {
      return;
    }

    form.setValues(employee);
  }, [employee]);

  const addMutation = useMutation({
    mutationFn: addEmployee,
    onError: (err: AxiosError<{ error: string }>) => {
      showNotification({
        title: 'Ошибка',
        message: err.response?.data.error,
        color: 'red',
      });
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: [getEmployees.queryKey] });
      showNotification({
        title: 'Сотрудники',
        message: `Cотрудник "${form.getValues().name}" был добавлен`,
      });
      close();
    },
  });

  const editMutation = useMutation({
    mutationFn: editEmployee,
    onError: (err: AxiosError<{ error: string }>) => {
      showNotification({
        title: 'Ошибка',
        message: err.response?.data.error,
        color: 'red',
      });
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: [getEmployees.queryKey] });
      showNotification({
        title: 'Сотрудники',
        message: `Cотрудник "${form.getValues().name}" был обновлен`,
      });
      close();
    },
  });

  const handleSubmit = form.onSubmit((values) => {
    const res = {
      Name: values.name,
      Position: values.position,
      EmployeeNum: values.employeeNum,
    };

    if (!edit) {
      return addMutation.mutate(res);
    }

    editMutation.mutate({ Id: employee!.id, ...res });
  });

  return (
    <Modal
      opened={opened}
      onClose={close}
      title={
        <Text fz={22} fw="bold">
          {edit ? 'Редактировать сотрудника' : 'Добавить сотрудника'}
        </Text>
      }
      centered
    >
      <form onSubmit={handleSubmit}>
        <Stack gap={12}>
          <TextInput
            withAsterisk
            label="Имя"
            placeholder="Введите имя"
            key={form.key('name')}
            {...form.getInputProps('name')}
          />

          <TextInput
            withAsterisk
            label="Номер сотрудника"
            placeholder="Введите номер сотрудника"
            key={form.key('employeeNum')}
            {...form.getInputProps('employeeNum')}
          />
          <Select
            withAsterisk
            label="Должность сотрудника"
            placeholder="Выберите тип"
            data={[
              'Главный менеджер по работе с клиентами',
              'Менеджер по работе с клиентами',
              'Старший механик',
              'Младший механик',
            ]}
            key={form.key('position')}
            {...form.getInputProps('position')}
          />

          <Group wrap="nowrap" mt="md">
            <Button w="100%" type="submit">
              Сохранить
            </Button>
            <Button w="100%" color="red" onClick={close}>
              Отменить
            </Button>
          </Group>
        </Stack>
      </form>
    </Modal>
  );
};
