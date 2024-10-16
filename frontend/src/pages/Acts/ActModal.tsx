import { useEffect } from 'react';
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { AxiosError } from 'axios';
import { Button, Group, Modal, Select, Stack, TextInput } from '@mantine/core';
import { useForm } from '@mantine/form';
import { showNotification } from '@mantine/notifications';
import { addAct } from '@/api/acts/addAct';
import { getActs } from '@/api/acts/getActs';
import { getEmployees } from '@/api/employees/getEmployees';
import { Act } from '@/types';
import { getServices } from '@/api/services/getServices';
import { getApps } from '@/api/apps/getApps';

type Props = {
  opened: boolean;
  close: () => void;
  edit?: boolean;
  act?: Act;
  onSubmit: () => void;
};

export const ActModal = ({ opened, close, act, onSubmit }: Props) => {
  const queryAct = useQueryClient();

  const form = useForm({
    mode: 'uncontrolled',
    initialValues: {
      name: act?.name ?? '',
      applicationId: act?.applicationId ?? '',
      serviceId: act?.serviceId ?? '',
    },
  });

  useEffect(() => {
    if (!act) {
      return;
    }

    form.setValues(act);
  }, [act]);

  const { data: appsData } = useQuery({
    queryKey: [getApps.queryKey],
    queryFn: () => getApps(),
    select(data) {
      return data.data.map((appsData) => ({
        value: appsData.id,
        label: appsData.name,
      }));
    },
    staleTime: 5000,
  });

  const { data: servicesData } = useQuery({
    queryKey: [getServices.queryKey],
    queryFn: () => getServices(),
    select(data) {
      return data.data.map((servicesData) => ({
        value: servicesData.id,
        label: servicesData.name,
      }));
    },
    staleTime: 5000,
  });

  const addMutation = useMutation({
    mutationFn: addAct,
    onError: (err: AxiosError<{ error: string }>) => {
      showNotification({
        title: 'Ошибка',
        message: err.response?.data.error,
        color: 'red',
      });
    },
    onSuccess: () => {
      queryAct.invalidateQueries({ queryKey: [getActs.queryKey] });

      showNotification({
        title: 'Акт',
        message: `Акт "${form.getValues().name}" был добавлен`,
      });

      form.reset();

      onSubmit();
      close();
    },
  });

  const handleSubmit = form.onSubmit((values) => {
    const res = {
      Name: values.name,
      ApplicationId: values.applicationId,
      ServiceId: values.serviceId,
    };

    addMutation.mutate(res);
  });

  return (
    <Modal opened={opened} onClose={close} title={'Добавить Акт'} centered>
      <form onSubmit={handleSubmit}>
        <Stack gap={12}>
          <TextInput
            withAsterisk
            label="Имя"
            placeholder="Введите название акта"
            key={form.key('name')}
            {...form.getInputProps('name')}
          />

          <Select
            withAsterisk
            label="Заявка"
            placeholder="Выберите связанную заявку"
            data={appsData}
            key={form.key('applicationId')}
            {...form.getInputProps('applicationId')}
          />

          <Select
            withAsterisk
            label="Услуга"
            placeholder="Выберите связанную услугу"
            data={servicesData}
            key={form.key('serviceId')}
            {...form.getInputProps('serviceId')}
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
