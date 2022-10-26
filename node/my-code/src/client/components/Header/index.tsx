import React from 'react';

import styles from './styles.module.scss';

interface HeaderProps {
    className?: string;
}

export const Header: React.FC<HeaderProps> = ({ className }) => {
    return <header className={className}>I am Header</header>;
};
