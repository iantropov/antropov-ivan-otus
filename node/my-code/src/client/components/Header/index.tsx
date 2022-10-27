import React from 'react';
import classnames from 'classnames';

import styles from './styles.module.scss';
import Link from 'next/link';

interface HeaderProps {
    className?: string;
}

export const Header: React.FC<HeaderProps> = ({ className }) => {
    // return <header className={className}>I am Header</header>;
    return (
        <header
            className={classnames(
                styles.header,
                className,
                'd-flex flex-wrap align-items-center justify-content-center justify-content-md-between py-3 mb-4 border-bottom'
            )}
        >
            <a
                href="/"
                className="d-flex align-items-center h3 col-md-3 mb-2 mb-md-0 text-dark text-decoration-none"
            >
                MyCode
            </a>

            <ul className="nav col-12 col-md-auto mb-2 justify-content-center mb-md-0">
                <li>
                    <a href="#" className="nav-link px-2 link-secondary">
                        Home
                    </a>
                </li>
                <li>
                    <a href="#" className="nav-link px-2 link-dark">
                        Features
                    </a>
                </li>
                <li>
                    <a href="#" className="nav-link px-2 link-dark">
                        Pricing
                    </a>
                </li>
                <li>
                    <a href="#" className="nav-link px-2 link-dark">
                        FAQs
                    </a>
                </li>
                <li>
                    <a href="#" className="nav-link px-2 link-dark">
                        About
                    </a>
                </li>
            </ul>

            <div className="col-md-3 text-end">
                <Link href="/login">
                    <a className="btn btn-outline-primary me-2">Login</a>
                </Link>
                <Link href="/register">
                    <a className="btn btn-primary">Sign-up</a>
                </Link>
            </div>
        </header>
    );
};
