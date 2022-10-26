import React, { ReactNode } from 'react';

import { Footer } from '../Footer';
import { Header } from '../Header';

import styles from './styles.module.scss';

export const Layout: React.FC<{ children?: ReactNode }> = ({ children }) => {
    return (
        <div className={styles.layout}>
            <Header className={styles.layout__header} />
            <main className={styles.layout__main}>{children}</main>
            <Footer className={styles.layout__footer} />
        </div>
    );
};
