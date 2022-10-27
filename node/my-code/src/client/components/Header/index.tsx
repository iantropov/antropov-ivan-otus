import React from 'react';
import classnames from 'classnames';
import { useRouter } from 'next/router';
import Link from 'next/link';
import { useMutation, useQuery } from '@apollo/client';

import { LOGOUT_USER_MUTATION, WHO_AM_I_QUERY } from '../../lib/graphql-queries';

import styles from './styles.module.scss';

interface HeaderProps {
    className?: string;
}

export const Header: React.FC<HeaderProps> = ({ className }) => {
    const router = useRouter();

    const {
        data: userData,
        loading: userLoading,
        error: userError,
        client
    } = useQuery(WHO_AM_I_QUERY, {
        onError: error => {
            debugger;
            console.log("QUERY", error);
        }
    });
    const [logoutUser] = useMutation(LOGOUT_USER_MUTATION, {
        refetchQueries: [{ query: WHO_AM_I_QUERY }]
    });

    console.log('USER', userData, userLoading, userError);

    const onLogoutClick = () => {
        logoutUser().then(
            () => {
                router.push('/login');
                client.resetStore().then(() => {

                }, (error) => {
                    debugger
                    console.log("RESET", error);
                });
            },
            error => {
                alert(error);
            }
        );
    };

    debugger;

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

            {!userLoading && userData && (
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
            )}

            {!userLoading && !userData && (
                <div className="col-md-3 text-end">
                    <Link href="/login">
                        <a className="btn btn-outline-primary me-2">Login</a>
                    </Link>
                    <Link href="/register">
                        <a className="btn btn-primary">Sign-up</a>
                    </Link>
                </div>
            )}

            {!userLoading && userData && (
                <div className="col-md-3 text-end">
                    <button className="btn btn-outline-primary me-2" onClick={onLogoutClick}>
                        Logout
                    </button>
                </div>
            )}
        </header>
    );
};
