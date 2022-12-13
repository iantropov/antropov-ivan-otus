import React, { ReactNode } from 'react';
import classnames from 'classnames';

import styles from './styles.module.scss';

interface MainProps {
    className?: string;
    align?: 'top' | 'center';
    children: ReactNode;
}

export const Main: React.FC<MainProps> = ({ className, align, children }) => {
    return (
        <main className={classnames(className, styles.main, styles[`main_${align}`])}>
            {children}
        </main>
    );
};
