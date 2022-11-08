import React from 'react';
import classnames from 'classnames';

import styles from './styles.module.scss';
import { Message } from '../../lib/types';

interface ToastsProps {
    className?: string;
    messages: Message[];
}

export const Toasts: React.FC<ToastsProps> = ({ className, messages }) => {
    return (
        <div className={classnames(className, styles.toasts)}>
            <div className="toast-container position-absolute top-0 end-0 p-3">
                {messages.map(({ id, type, text }) => (
                    <div
                        key={id}
                        className={`toast show text-white bg-${
                            type === 'error' ? 'danger' : 'success'
                        }`}
                    >
                        <div className="d-flex">
                            <div className="toast-body">{text}</div>
                            <button
                                type="button"
                                className="btn-close btn-close-white me-2 m-auto"
                                data-bs-dismiss="toast"
                                aria-label="Close"
                            ></button>
                        </div>
                    </div>
                ))}
            </div>
        </div>
    );
};
