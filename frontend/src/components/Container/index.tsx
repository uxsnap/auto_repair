import { PropsWithChildren } from 'react';
import { Loader } from '@mantine/core';
import styles from './Container.module.css';

type Props = {
  isFetching: boolean;
};

export const Container = ({ isFetching, children }: PropsWithChildren<Props>) => (
  <div className={styles.root}>
    {isFetching && <Loader className={styles.loader} />}

    {!isFetching && children}
  </div>
);
