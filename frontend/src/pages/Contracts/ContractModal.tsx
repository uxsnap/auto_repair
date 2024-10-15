import { useEffect } from 'react';
import { useMutation, useQueryClient } from '@tanstack/react-query';
import { AxiosError } from 'axios';
import { Button, Group, Modal, NumberInput, Select, Stack, TextInput } from '@mantine/core';
import { DateInput } from '@mantine/dates';
import { useForm } from '@mantine/form';
import { showNotification } from '@mantine/notifications';
import { addContract } from '@/api/contract/addContract';
import { editContract } from '@/api/contract/editContract';
import { getContracts } from '@/api/contract/getContacts';
import { Contract } from '@/types';
import { fromDateToString } from '@/utils';

type Props = {
  opened: boolean;
  close: () => void;
  edit?: boolean;
  contract?: Contract;
  onSubmit: () => void;
};

const getInitValues = (contract?: Contract) => ({
  name: contract?.name ?? '',
  sum: contract?.sum ?? 0,
  signedAt: contract ? new Date(contract.signedAt.Time) : '',
  status: contract?.status ?? 'Новый',
});

export const ContractModal = ({ edit = false, opened, close, contract, onSubmit }: Props) => {
  const queryContract = useQueryClient();

  console.log(getInitValues(contract));

  const form = useForm({
    mode: 'uncontrolled',
    initialValues: getInitValues(contract),
  });

  useEffect(() => {
    if (!contract) {
      return;
    }

    form.setValues(getInitValues(contract));
  }, [contract]);

  const addMutation = useMutation({
    mutationFn: addContract,
    onError: (err: AxiosError<{ error: string }>) => {
      showNotification({
        title: 'Ошибка',
        message: err.response?.data.error,
        color: 'red',
      });
    },
    onSuccess: () => {
      queryContract.invalidateQueries({ queryKey: [getContracts.queryKey] });

      showNotification({
        title: 'Договор',
        message: `Договор "${form.getValues().name}" был добавлен`,
      });

      form.reset();

      onSubmit();
      close();
    },
  });

  const editMutation = useMutation({
    mutationFn: editContract,
    onError: (err: AxiosError<{ error: string }>) => {
      showNotification({
        title: 'Ошибка',
        message: err.response?.data.error,
        color: 'red',
      });
    },
    onSuccess: () => {
      queryContract.invalidateQueries({ queryKey: [getContracts.queryKey] });

      showNotification({
        title: 'Договор',
        message: `Договор "${form.getValues().name}" был обновлен`,
      });

      form.reset();

      onSubmit();
      close();
    },
  });

  const handleSubmit = form.onSubmit((values) => {
    const res: any = {
      Name: values.name,
      Sum: values.sum,
      Status: values.status,
    };

    if (values.signedAt) {
      res.SignedAt = fromDateToString(new Date(values.signedAt));
    }

    if (edit) {
      return editMutation.mutate({ Id: contract!.id, ...res });
    }

    addMutation.mutate(res);
  });

  return (
    <Modal
      opened={opened}
      onClose={close}
      title={edit ? 'Редактировать договор' : 'Добавить договор'}
      centered
    >
      <form onSubmit={handleSubmit}>
        <Stack gap={12}>
          <TextInput
            withAsterisk
            label="Имя"
            placeholder="Введите имя клиента"
            key={form.key('name')}
            {...form.getInputProps('name')}
          />

          <NumberInput
            withAsterisk
            allowNegative={false}
            label="Cумма"
            placeholder="Введите сумму по договору"
            suffix="₽"
            key={form.key('sum')}
            {...form.getInputProps('sum')}
          />

          <DateInput
            label="Дата подписания"
            placeholder="Выберите дату подписания"
            valueFormat="DD.MM.YYYY"
            lang="ru"
            key={form.key('signedAt')}
            {...form.getInputProps('signedAt')}
          />

          <Select
            withAsterisk
            label="Статус"
            placeholder="Выберите статус"
            data={['Новый', 'Действующий', 'Закрыт']}
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
