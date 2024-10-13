import { useEffect } from 'react';
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { AxiosError } from 'axios';
import { Button, Checkbox, Group, Modal, Select, Stack, TextInput } from '@mantine/core';
import { useForm } from '@mantine/form';
import { showNotification } from '@mantine/notifications';
import { getClients } from '@/api/clients/getClients';
import { addVehicle } from '@/api/vehicles/addVehicle';
import { getVehicles } from '@/api/vehicles/getVehicles';
import { Vehicle } from '@/types';

type Props = {
  opened: boolean;
  close: () => void;
  edit?: boolean;
  vehicle?: Vehicle;
  onSubmit: () => void;
};

export const VehicleModal = ({ edit = false, opened, close, vehicle, onSubmit }: Props) => {
  const queryVehicle = useQueryClient();

  const form = useForm({
    mode: 'uncontrolled',
    initialValues: {
      vehicleNumber: vehicle?.vehicleNumber ?? '',
      clientId: vehicle?.clientId ?? '',
      brand: vehicle?.brand ?? '',
      model: vehicle?.model ?? '',
    },
  });

  useEffect(() => {
    if (!vehicle) {
      return;
    }

    form.setValues(vehicle);
  }, [vehicle]);

  const { data: clientsData } = useQuery({
    queryKey: [getClients.queryKey],
    queryFn: () => getClients(),
    select(data) {
      return data.data.map((clientData) => ({
        value: clientData.id,
        label: clientData.name,
      }));
    },
    staleTime: 5000,
  });

  const addMutation = useMutation({
    mutationFn: addVehicle,
    onError: (err: AxiosError<{ error: string }>) => {
      showNotification({
        title: 'Ошибка',
        message: err.response?.data.error,
        color: 'red',
      });
    },
    onSuccess: () => {
      queryVehicle.invalidateQueries({ queryKey: [getVehicles.queryKey] });

      showNotification({
        title: 'ТС',
        message: `ТС "${form.getValues().vehicleNumber}" было добавлено`,
      });

      form.reset();

      onSubmit();
      close();
    },
  });

  // const editMutation = useMutation({
  //   mutationFn: editVehicle,
  //   onError: (err: AxiosError<{ error: string }>) => {
  //     showNotification({
  //       title: 'Ошибка',
  //       message: err.response?.data.error,
  //       color: 'red',
  //     });
  //   },
  //   onSuccess: () => {
  //     queryVehicle.invalidateQueries({ queryKey: [getVehicles.queryKey] });

  //     showNotification({
  //       title: 'ТС',
  //       message: `ТС "${form.getValues().vehicleNumber}" было обновлено`,
  //     });

  //     form.reset();

  //     onSubmit();
  //     close();
  //   },
  // });

  const handleSubmit = form.onSubmit((values) => {
    const res = {
      ClientId: values.clientId,
      VehicleNumber: values.vehicleNumber,
      Brand: values.brand,
      Model: values.model,
    };

    addMutation.mutate(res);
  });

  return (
    <Modal
      opened={opened}
      onClose={close}
      title={edit ? 'Редактировать ТС' : 'Добавить ТС'}
      centered
    >
      <form onSubmit={handleSubmit}>
        <Stack gap={12}>
          <TextInput
            withAsterisk
            label="Номер машины"
            placeholder="Введите номер машины"
            key={form.key('vehicleNumber')}
            {...form.getInputProps('vehicleNumber')}
          />

          <Select
            withAsterisk
            label="Клиент"
            placeholder="Выберите клиента"
            data={clientsData}
            key={form.key('clientId')}
            {...form.getInputProps('clientId')}
          />

          <TextInput
            withAsterisk
            label="Марка"
            placeholder="Введите марку ТС"
            key={form.key('brand')}
            {...form.getInputProps('brand')}
          />

          <TextInput
            withAsterisk
            label="Модель"
            placeholder="Введите модель ТС"
            key={form.key('model')}
            {...form.getInputProps('model')}
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
