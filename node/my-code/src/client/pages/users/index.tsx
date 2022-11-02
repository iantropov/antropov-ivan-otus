import React from 'react';
import { NextPage } from 'next';
import { useQuery } from '@apollo/client';
import classnames from 'classnames';

import { useUser } from '../../hooks/use-user';
import { ProblemsReponse, UsersResponse } from '../../lib/types';
import { Problems as ProblemsComponent } from '../../components/Problems';
import { GET_USERS_QUERY } from '../../lib/graphql';

import styles from './styles.module.scss';

const Users: NextPage = () => {
    const { data, loading } = useQuery<UsersResponse>(GET_USERS_QUERY);
    const [user, isUserLoading] = useUser({ isAdmin: true });

    if (loading || isUserLoading) return <p>Loading...</p>;
    if (!user) return null;

    const onDeleteUserClick = () => {};

    return (
        <section className={classnames(styles.users)}>
            <h2 className={classnames(styles.users__header)}>Available users:</h2>
            <table className={classnames(styles.users__table, 'table')}>
                <thead>
                    <tr>
                        <th scope="col">#</th>
                        <th scope="col">Name</th>
                        <th scope="col">Email</th>
                        <th scope="col">Admin?</th>
                        <th scope="col">Actions</th>
                    </tr>
                </thead>
                <tbody>
                    {data.users.map(user => (
                        <tr key={user._id}>
                            <th scope="row">{user._id}</th>
                            <td>{user.name}</td>
                            <td>{user.email}</td>
                            <td>{user.isAdmin ? 'Yes' : 'No'}</td>
                            <td>
                                <button className="btn btn-sm btn-danger" onClick={onDeleteUserClick}>
                                    Delete
                                </button>
                            </td>
                        </tr>
                    ))}
                </tbody>
            </table>
        </section>
    );
};

export default Users;
