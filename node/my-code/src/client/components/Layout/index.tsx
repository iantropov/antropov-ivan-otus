import React, { ReactNode, useEffect, useState } from 'react';
import { messageBroker } from '../../lib/message-broker';
import { Message } from '../../lib/types';

import { Footer } from '../Footer';
import { Header } from '../Header';
import { Toasts } from '../Toasts';

import styles from './styles.module.scss';

export const Layout: React.FC<{ children?: ReactNode }> = ({ children }) => {
    const [messages, setMessages] = useState<Message[]>([]);

    useEffect(() => {
        messageBroker.subsribe(messages => {
            debugger
            setMessages(messages);
        });
    }, []);

    return (
        <div className={styles.layout}>
            <Toasts className={styles.layout__toasts} messages={messages} />
            <Header className={styles.layout__header} />
            <div className={styles.layout__main}>{children}</div>
            <Footer className={styles.layout__footer} />
        </div>
    );
};
