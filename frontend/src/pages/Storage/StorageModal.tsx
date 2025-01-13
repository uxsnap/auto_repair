import { useEffect } from 'react';
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { AxiosError } from 'axios';
import { Button, Group, Modal, NumberInput, Select, Stack, TextInput } from '@mantine/core';
import { useForm } from '@mantine/form';
import { showNotification } from '@mantine/notifications';
import { getDetails } from '@/api/details/getDetails';
import { getEmployees } from '@/api/employees/getEmployees';
import { addStorage } from '@/api/storages/addStorage';
import { editStorage } from '@/api/storages/editStorage';
import { getStorages } from '@/api/storages/getStorages';
import { Storage } from '@/types';

type Props = {
  opened: boolean;
  close: () => void;
  edit?: boolean;
  storage?: Storage;
  onSubmit: () => void;
};

export const StorageModal = ({ edit = false, opened, close, storage, onSubmit }: Props) => {
  const queryClient = useQueryClient();

  const form = useForm({
    mode: 'uncontrolled',
    initialValues: {
      storageNum: storage?.storageNum ?? '',
      employeeId: storage?.employeeId ?? '',
      detailId: storage?.detailId ?? '',
      detailCount: storage?.detailCount ?? 1,
    },
  });

  useEffect(() => {
    if (!storage) {
      return;
    }

    form.setValues(storage);
  }, [storage]);

  const { data: detailsData } = useQuery({
    queryKey: [getDetails.queryKey],
    queryFn: () => getDetails(),
    select(data) {
      return data.data.map((detail) => ({ value: detail.id, label: detail.name }));
    },
    staleTime: 5000,
  });

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
    mutationFn: addStorage,
    onError: (err: AxiosError<{ error: string }>) => {
      showNotification({
        title: 'Ошибка',
        message: err.response?.data.error,
        color: 'red',
      });
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: [getStorages.queryKey] });

      showNotification({
        title: 'Склад',
        message: `Склад "${form.getValues().storageNum}" был добавлен`,
      });

      form.reset();

      onSubmit();
      close();
    },
  });

  const editMutation = useMutation({
    mutationFn: editStorage,
    onError: (err: AxiosError<{ error: string }>) => {
      showNotification({
        title: 'Ошибка',
        message: err.response?.data.error,
        color: 'red',
      });
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: [getStorages.queryKey] });

      showNotification({
        title: 'Склад',
        message: `Склад "${form.getValues().storageNum}" был обновлен`,
      });

      form.reset();

      onSubmit();
      close();
    },
  });

  const handleSubmit = form.onSubmit((values) => {
    const res = {
      StorageNum: values.storageNum,
      EmployeeId: values.employeeId,
      DetailCount: values.detailCount,
      DetailId: values.detailId,
    };

    if (edit) {
      return editMutation.mutate({ Id: storage!.id, ...res });
    }

    addMutation.mutate(res);
  });

  return (
    <Modal opened={opened} onClose={close} title={'Добавить деталь'} centered>
      <form onSubmit={handleSubmit}>
        <Stack gap={12}>
          <TextInput
            withAsterisk
            label="Номер склада"
            placeholder="Введите номер склада"
            key={form.key('storageNum')}
            {...form.getInputProps('storageNum')}
          />

          <Select
            withAsterisk
            label="Ответственный сотрудник"
            placeholder="Выберите ответственного сотрудника"
            data={employeesData}
            key={form.key('employeeId')}
            {...form.getInputProps('employeeId')}
          />

          <Select
            withAsterisk
            label="Деталь"
            placeholder="Выберите деталь"
            data={detailsData}
            key={form.key('detailId')}
            {...form.getInputProps('detailId')}
          />

          <NumberInput
            withAsterisk
            label="Кол-во деталей"
            placeholder="Введите кол-во деталей"
            key={form.key('detailCount')}
            {...form.getInputProps('detailCount')}
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
