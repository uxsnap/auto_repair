import { useEffect } from 'react';
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { AxiosError } from 'axios';
import { Button, Checkbox, Group, Modal, Select, Stack, TextInput } from '@mantine/core';
import { useForm } from '@mantine/form';
import { showNotification } from '@mantine/notifications';
import { addClient } from '@/api/clients/addClient';
import { editClient } from '@/api/clients/editClient';
import { getClients } from '@/api/clients/getClients';
import { getEmployees } from '@/api/employees/getEmployees';
import { Client } from '@/types';

type Props = {
  opened: boolean;
  close: () => void;
  edit?: boolean;
  client?: Client;
  onSubmit: () => void;
};

export const ClientModal = ({ edit = false, opened, close, client, onSubmit }: Props) => {
  const queryClient = useQueryClient();

  const form = useForm({
    mode: 'uncontrolled',
    initialValues: {
      name: client?.name ?? '',
      employeeId: client?.employeeId ?? '',
      phone: client?.phone ?? '',
      passport: client?.passport ?? '',
      hasDocuments: client?.hasDocuments ?? false,
    },
  });

  useEffect(() => {
    if (!client) {
      return;
    }

    form.setValues(client);
  }, [client]);

  const { data: employeesData } = useQuery({
    queryKey: [getEmployees.queryKey],
    queryFn: () => getEmployees(),
    select(data) {
      return data.data.map((employeesData) => ({
        value: employeesData.id,
        label: employeesData.name,
      }));
    },
    staleTime: 5000,
  });

  const addMutation = useMutation({
    mutationFn: addClient,
    onError: (err: AxiosError<{ error: string }>) => {
      showNotification({
        title: 'Ошибка',
        message: err.response?.data.error,
        color: 'red',
      });
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: [getClients.queryKey] });

      showNotification({
        title: 'Клиент',
        message: `Клиент "${form.getValues().name}" был добавлен`,
      });

      form.reset();

      onSubmit();
      close();
    },
  });

  const editMutation = useMutation({
    mutationFn: editClient,
    onError: (err: AxiosError<{ error: string }>) => {
      showNotification({
        title: 'Ошибка',
        message: err.response?.data.error,
        color: 'red',
      });
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: [getClients.queryKey] });

      showNotification({
        title: 'Клиент',
        message: `Клиент "${form.getValues().name}" был обновлен`,
      });

      form.reset();

      onSubmit();
      close();
    },
  });

  const handleSubmit = form.onSubmit((values) => {
    const res = {
      Name: values.name,
      EmployeeId: values.employeeId,
      Phone: values.phone,
      HasDocuments: values.hasDocuments,
      Passport: values.passport,
    };

    if (edit) {
      return editMutation.mutate({ Id: client!.id, ...res });
    }

    addMutation.mutate(res);
  });

  return (
    <Modal opened={opened} onClose={close} title={'Добавить деталь'} centered>
      <form onSubmit={handleSubmit}>
        <Stack gap={12}>
          <TextInput
            withAsterisk
            label="Имя"
            placeholder="Введите имя клиента"
            key={form.key('name')}
            {...form.getInputProps('name')}
          />

          <Select
            withAsterisk
            label="Отвественный сотрудник"
            placeholder="Выберите отвественного сотрудника"
            data={employeesData}
            key={form.key('employeeId')}
            {...form.getInputProps('employeeId')}
          />

          <TextInput
            withAsterisk
            label="Телефон"
            placeholder="Введите телефон клиента"
            key={form.key('phone')}
            {...form.getInputProps('phone')}
          />

          <TextInput
            withAsterisk
            label="Паспорт"
            placeholder="Введите паспорт клиента"
            key={form.key('passport')}
            {...form.getInputProps('passport')}
          />

          <Checkbox
            label="У клиента есть документы на ТС?"
            key={form.key('hasDocuments')}
            {...form.getInputProps('hasDocuments', { type: 'checkbox' })}
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
