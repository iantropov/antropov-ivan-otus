import React from 'react';

import styles from './styles.module.scss';

interface FooterProps {
    className?: string;
}

export const Footer: React.FC<FooterProps> = ({ className }) => {
    return <footer className={className}>I am footer</footer>;
};
