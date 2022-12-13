import React from 'react';
import classnames from 'classnames';

import styles from './styles.module.scss';

interface FooterProps {
    className?: string;
}

export const Footer: React.FC<FooterProps> = ({ className }) => {
    return (
        <footer className={classnames(className, styles.footer)}>
            <div className={styles.footer__title}>MyCode</div>
            <div className={styles.footer__author}>Ivan Antropov, 2022</div>
        </footer>
    );
};
