import { useEffect } from 'react';
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { AxiosError } from 'axios';
import {
  Button,
  Checkbox,
  Group,
  Modal,
  NumberInput,
  Select,
  Stack,
  TextInput,
} from '@mantine/core';
import { DateInput } from '@mantine/dates';
import { useForm } from '@mantine/form';
import { showNotification } from '@mantine/notifications';
import { getContracts } from '@/api/contract/getContacts';
import { addReceipt } from '@/api/receipt/addReceipt';
import { getReceipts } from '@/api/receipt/getReceipts';
import { Receipt } from '@/types';

type Props = {
  opened: boolean;
  close: () => void;
  edit?: boolean;
  receipt?: Receipt;
  onSubmit: () => void;
};

const getInitValues = (receipt?: Receipt) => ({
  sum: receipt?.sum ?? 0,
  contractId: receipt?.contractId ?? '',
});

export const ReceiptModal = ({ edit = false, opened, close, receipt, onSubmit }: Props) => {
  const queryReceipt = useQueryClient();

  const form = useForm({
    mode: 'uncontrolled',
    initialValues: getInitValues(receipt),
  });

  useEffect(() => {
    if (!receipt) {
      return;
    }

    form.setValues(getInitValues(receipt));
  }, [receipt]);

  const addMutation = useMutation({
    mutationFn: addReceipt,
    onError: (err: AxiosError<{ error: string }>) => {
      showNotification({
        title: 'Ошибка',
        message: err.response?.data.error,
        color: 'red',
      });
    },
    onSuccess: () => {
      queryReceipt.invalidateQueries({ queryKey: [getReceipts.queryKey] });

      showNotification({
        title: 'Чек',
        message: `Новый чек на сумму "${form.getValues().sum}" был добавлен`,
      });

      form.reset();

      onSubmit();
      close();
    },
  });

  const handleSubmit = form.onSubmit((values) => {
    const res: any = {
      Sum: values.sum,
      ContractId: values.contractId,
    };

    addMutation.mutate(res);
  });

  const { data: contractsData } = useQuery({
    queryKey: [getContracts.queryKey],
    queryFn: () => getContracts(),
    select(data) {
      return data.data.map((contractsData) => ({
        value: contractsData.id,
        label: contractsData.name,
      }));
    },
    staleTime: 5000,
  });

  return (
    <Modal
      opened={opened}
      onClose={close}
      title={edit ? 'Редактировать чек' : 'Добавить чек'}
      centered
    >
      <form onSubmit={handleSubmit}>
        <Stack gap={12}>
          <Select
            withAsterisk
            label="Договор"
            placeholder="Выберите связанный договор"
            data={contractsData}
            key={form.key('contractId')}
            {...form.getInputProps('contractId')}
          />

          <NumberInput
            withAsterisk
            allowNegative={false}
            label="Cумма"
            placeholder="Введите сумму"
            suffix="₽"
            key={form.key('sum')}
            {...form.getInputProps('sum')}
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
