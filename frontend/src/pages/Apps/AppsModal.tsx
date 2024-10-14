import { useEffect } from 'react';
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { AxiosError } from 'axios';
import { Button, Group, Modal, Select, Stack, TextInput } from '@mantine/core';
import { useForm } from '@mantine/form';
import { showNotification } from '@mantine/notifications';
import { addApp } from '@/api/apps/addApp';
import { editApp } from '@/api/apps/editApp';
import { getApps } from '@/api/apps/getApps';
import { getEmployees } from '@/api/employees/getEmployees';
import { Application } from '@/types';
import { getClients } from '@/api/clients/getClients';

type Props = {
  opened: boolean;
  close: () => void;
  edit?: boolean;
  app?: Application;
  onSubmit: () => void;
};

export const AppModal = ({ edit = false, opened, close, app, onSubmit }: Props) => {
  const queryApp = useQueryClient();

  const form = useForm({
    mode: 'uncontrolled',
    initialValues: {
      name: app?.name ?? '',
      employeeId: app?.employeeId ?? '',
      clientId: app?.clientId ?? '',
      createdAt: app?.createdAt ?? '',
      status: app?.status ?? 'Новый',
      contractId: app?.contractId ?? '',
    },
  });

  useEffect(() => {
    if (!app) {
      return;
    }

    form.setValues(app);
  }, [app]);

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

  const { data: clientsData } = useQuery({
    queryKey: [getClients.queryKey],
    queryFn: () => getClients(),
    select(data) {
      return data.data.map((clientsData) => ({
        value: clientsData.id,
        label: clientsData.name,
      }));
    },
    staleTime: 5000,
  });

  const addMutation = useMutation({
    mutationFn: addApp,
    onError: (err: AxiosError<{ error: string }>) => {
      showNotification({
        title: 'Ошибка',
        message: err.response?.data.error,
        color: 'red',
      });
    },
    onSuccess: () => {
      queryApp.invalidateQueries({ queryKey: [getApps.queryKey] });

      showNotification({
        title: 'Контракт',
        message: `Контракт "${form.getValues().name}" был добавлен`,
      });

      form.reset();

      onSubmit();
      close();
    },
  });

  const editMutation = useMutation({
    mutationFn: editApp,
    onError: (err: AxiosError<{ error: string }>) => {
      showNotification({
        title: 'Ошибка',
        message: err.response?.data.error,
        color: 'red',
      });
    },
    onSuccess: () => {
      queryApp.invalidateQueries({ queryKey: [getApps.queryKey] });

      showNotification({
        title: 'Контракт',
        message: `Контракт "${form.getValues().name}" был обновлен`,
      });

      form.reset();

      onSubmit();
      close();
    },
  });

  const handleSubmit = form.onSubmit((values) => {
    const res = {
      EmployeeId: values.employeeId,
      ClientId: values.clientId,
      Name: values.name,
      Status: values.status,
      ContractId: values.contractId,
    };

    if (edit) {
      return editMutation.mutate({ Id: app!.id, ...res });
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

          <Select
            withAsterisk
            label="Клиент"
            placeholder="Выберите клиента"
            data={clientsData}
            key={form.key('clientId')}
            {...form.getInputProps('clientId')}
          />

          {/* <Select
            withAsterisk
            label="Контракт"
            placeholder="Выберите контракт"
            data={contractsData}
            key={form.key('contractId')}
            {...form.getInputProps('contractId')}
          /> */}

          <Select
            withAsterisk
            label="Статус"
            placeholder="Выберите статус"
            data={[
              'Новый',
              'В работе',
              'Закрыт',
            ]}
            key={form.key('status')}
            {...form.getInputProps('status')}
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
