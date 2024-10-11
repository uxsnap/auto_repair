import { useMutation, useQueryClient } from '@tanstack/react-query';
import { AxiosError } from 'axios';
import { Button, Group, Modal, NumberInput, Select, Stack, TextInput } from '@mantine/core';
import { useForm } from '@mantine/form';
import { showNotification } from '@mantine/notifications';
import { addDetail } from '@/api/details/addDetail';

type Props = {
  opened: boolean;
  close: () => void;
  edit?: boolean;
};

export const DetailsModal = ({ opened, close }: Props) => {
  const queryClient = useQueryClient();

  const form = useForm({
    mode: 'uncontrolled',
    initialValues: {
      name: '',
      price: 0,
      type: '',
    },
  });

  const addMutation = useMutation({
    mutationFn: addDetail,
    onError: (err: AxiosError<{ error: string }>) => {
      showNotification({
        title: 'Ошибка',
        message: err.response?.data.error,
        color: 'red',
      });
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: [addDetail.queryKey] });
      showNotification({
        title: 'Детали',
        message: `Деталь "${form.values.name}" была добавлена`,
      });
      close();
    },
  });

  const handleSubmit = form.onSubmit((values) => {
    addMutation.mutate({
      Name: values.name,
      Price: values.price,
      Type: values.type,
    });
  });

  return (
    <Modal opened={opened} onClose={close} title={'Добавить деталь'} centered>
      <form onSubmit={handleSubmit}>
        <Stack gap={12}>
          <TextInput
            withAsterisk
            label="Имя"
            placeholder="Введите имя"
            key={form.key('name')}
            {...form.getInputProps('name')}
          />

          <NumberInput
            label="Цена"
            placeholder="Введите цену"
            suffix="₽"
            key={form.key('price')}
            {...form.getInputProps('price')}
          />

          <Select
            label="Тип"
            placeholder="Выберите тип"
            data={[
              'Деталь соединительная',
              'Деталь вращательного движения',
              'Деталь обслуживающая передачи',
            ]}
            key={form.key('type')}
            {...form.getInputProps('type')}
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
