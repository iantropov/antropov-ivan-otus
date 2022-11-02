import React from 'react';
import { NextPage } from 'next';
import { useMutation, useQuery } from '@apollo/client';
import classnames from 'classnames';

import { useUser } from '../../hooks/use-user';
import { UsersResponse } from '../../lib/types';
import { DELETE_USER_MUTATION, GET_USERS_QUERY, WHO_AM_I_QUERY } from '../../lib/graphql';

import styles from './styles.module.scss';

const Users: NextPage = () => {
    const { data, loading } = useQuery<UsersResponse>(GET_USERS_QUERY);
    const [deleteUser] = useMutation(DELETE_USER_MUTATION, {
        refetchQueries: [{ query: WHO_AM_I_QUERY }, GET_USERS_QUERY]
    });

    const [user, isUserLoading] = useUser({ isAdmin: true });

    if (loading || isUserLoading) return <p>Loading...</p>;
    if (!user) return null;

    const onDeleteUserClick = (userId) => {
        deleteUser({ variables: { userId } }).then(
            () => {
                console.log("SUCCESS!");
            },
            error => {
                alert(error);
            }
        );
    };

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
                                <button className="btn btn-sm btn-danger" onClick={() => onDeleteUserClick(user._id)}>
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
